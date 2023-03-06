FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main



FROM scratch

WORKDIR /app

COPY --from=build /app/main ./main

CMD ["/app/main"]
