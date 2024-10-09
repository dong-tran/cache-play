-- Drop the table if it exists (to avoid errors if re-running the script)
DROP TABLE IF EXISTS users;

-- Create the table with the necessary columns
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Insert 1,000,000 records using a DO block
DO
$$
DECLARE
    i INTEGER := 0;
BEGIN
    -- Loop to generate and insert random data
    WHILE i < 1000000 LOOP
        INSERT INTO users (email, password)
        VALUES (
            CONCAT('user', FLOOR(RANDOM() * 1000000)::TEXT, '@example.com'),
            CONCAT('password', FLOOR(RANDOM() * 1000000)::TEXT)
        );
        i := i + 1;
    END LOOP;
END
$$;
