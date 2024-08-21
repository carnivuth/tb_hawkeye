FROM golang:1.23-alpine

COPY . /go/src/tb_eagleeye
WORKDIR /go/src/tb_eagleeye

# get dependencies
RUN go mod tidy

# build
RUN apk add build-base
RUN CGO_ENABLED=1 go build -o /tb_eagleeye
EXPOSE 80
ENV TB_EAGLEEYE_DB_PATH=/var/lib/tb_eagleeye
ENV TB_EAGLEEYE_PORT=80
ENV TB_EAGLEEYE_ADDRESS=0.0.0.0
CMD ["/tb_eagleeye"]
