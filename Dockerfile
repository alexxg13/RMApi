# Используем официальный образ Golang
FROM golang:1.21-alpine AS builder

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Сборка Go приложения
RUN go build -o server .

# Минимальный образ для запуска
FROM alpine:latest

# Копируем скомпилированное приложение и статические файлы из builder
COPY --from=builder /app/server /app/server
COPY --from=builder /app/temp /app/temp

# Указываем рабочую директорию
WORKDIR /app

# Открываем порт 8080
EXPOSE 8080

# Запуск приложения
CMD ["./server"]