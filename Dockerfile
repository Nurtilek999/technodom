FROM golang:1.19.4

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o technodom ./main.go

CMD ["./technodom"]