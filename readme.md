# Go server

Web server written in Golang.

# Application environment variables

- `SERVER_URL` - The URL where the HTTP server is available (e.g. `http://127.0.0.1:8080`).
- `SMTP_FROM` - The name of the user who sends email via SMTP (e.g. `John`).
- `SMTP_FROM_EMAIL` - Email address of the user who sends email via SMTP (e.g. `john.doe@gmail.com`).
- `SMTP_PASSWORD` - The password of the user who sends email via SMTP (e.g. `asdf1234`).

# Resources

- [This article describes how to configure the go server to work with docker compose and live reload.](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/)
- [Gin binding in golang, tutorial with examples by LogRocket.](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/)
- [Building REST API with Golang using GIN and GORM.](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/)
