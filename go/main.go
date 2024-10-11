package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dong-tran/cache-play/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Response struct {
	CurrentTime string `json:"currentTime"`
}

type UsersRequest struct {
	Partition int64 `param:"partition"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:58585", "http://localhost:53000"},
	}))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "56379")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dsn := fmt.Sprintf("host=%s user=postgres password=postgres dbname=postgres port=%s sslmode=disable", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "55432"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	dbDateService := services.NewDbDateService(db)
	userService := services.NewUserService(redisClient, db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})
	e.GET("/time", func(c echo.Context) error {
		dbDate := dbDateService.GetCurrent()
		return c.String(http.StatusOK, dbDate.CurrentTimestamp.Format("15:04:05"))
	})
	e.GET("/random", func(c echo.Context) error {
		return c.String(200, randomString(6))
	})
	e.GET("/users/:partition", func(c echo.Context) error {
		req := UsersRequest{}
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		users, err := userService.GetUsers(req.Partition)
		if err != nil {
			return err
		}
		return c.JSON(200, users)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getEnv(key string, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	} else {
		return v
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
