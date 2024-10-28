# Этап сборки
FROM golang:1.22 AS builder

# Задаем рабочую директорию для сборки
WORKDIR /app

# Поднимаемся на уровень выше, чтобы скопировать go.mod и go.sum
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Копируем все содержимое проекта
COPY ../../ .

# Сборка приложения, указав путь к main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o alph_pooler ./cmd/pooler

# Этап создания минимального образа
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/alph_pooler .

# Указываем порт для приложения
EXPOSE 8090

# Команда для запуска приложения
CMD ["./alph_pooler"]
