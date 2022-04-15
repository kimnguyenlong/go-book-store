# Introductions

#### This is a backend project written in Go which provides GraphQl apis of a mini book store. It includes some features such as:

- Login
- Get authors, topics, books
- Create, update, remove an author, a topic or a book
- Create a user
- Set cart for a user, get cart of a user
- Update wish list for a user, get wish list of a user
- Get reviews of a books
- Create, update, remove a review

#### Used in this project:

- Gin Framework ([https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin))
- gqlgen ([https://github.com/99designs/gqlgen](https://github.com/99designs/gqlgen))
- MongoDB

# How to use

After cloning this repository, you can run the service via the 2 following ways:

#### 1. With Docker

Run then cmd `docker-compose build` and then `docker-compose up`, then the service will be available on [localhost:8080](localhost:8080)

#### 2. Without Docker

You have to create a `.env` file in the `src` folder, then you need to specify 3 env variables:`MONGODB_CONNECTTION_URI`, `JWT_LIFE_TIME`, `JWT_SECRET`.
After that, you can run the service via cmd `go run main.go` in the`src` folder and the service will be available on [localhost:8080](localhost:8080)
