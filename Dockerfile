FROM golang:1.21

WORKDIR /app
ADD . /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]