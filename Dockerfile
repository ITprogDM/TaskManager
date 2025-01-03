# Используем официальный образ Golang
FROM golang:1.22 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o taskmanager ./cmd

# Финальный образ с минимальным размером
FROM alpine:latest

RUN apk add --no-cache libc6-compat

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранный бинарник из builder контейнера
COPY --from=builder /app/taskmanager .

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./taskmanager"]
