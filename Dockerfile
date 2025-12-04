FROM golang:latest
WORKDIR /app
RUN go mod init go_gin_gorm && \
  go install github.com/air-verse/air@latest && \
  go get -u \
    github.com/gin-gonic/gin \
    gorm.io/gorm \
    gorm.io/driver/mysql \
    golang.org/x/crypto \
    github.com/joho/godotenv \
    github.com/golang-jwt/jwt && \
  go mod tidy && \
  air init
CMD ["air"]
