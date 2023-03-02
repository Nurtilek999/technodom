#FROM golang:1.19.4-alpine
#
##WORKDIR /merchant
#ENV GOPATH=/
#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod download
#
#COPY *.go ./
#RUN go build -o merchant ./main.go
#
#CMD ["./merchant"]

FROM golang:1.19.4-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o merchant ./main.go

CMD ["./merchant"]



