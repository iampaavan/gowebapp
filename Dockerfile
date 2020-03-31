FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /app
RUN apk add --no-cache tzdata
RUN go build -o main ./src
FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /app .
EXPOSE 8086
CMD ["main"]