FROM node:22 as builder

WORKDIR /src
COPY package.json package-lock.json ./
RUN yarn install 
COPY . .
RUN yarn build

FROM node:22
WORKDIR /app
COPY --from=builder /src/.output .
CMD [ "node", "/app/server/index.mjs" ]