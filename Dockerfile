FROM golang:1.19 as builder
LABEL authors="csp"
WORKDIR /app
COPY backend/main.go app/
COPY backend/go.mod app/
COPY backend/go.sum app/
ENV GOPROXY https://goproxy.cn
RUN go build -o bookManger main.go
ENTRYPOINT ["/app/bookManger"]