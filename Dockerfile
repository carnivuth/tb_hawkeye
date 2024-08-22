FROM golang:1.23-alpine

COPY . /go/src/tb_hawkeye
WORKDIR /go/src/tb_hawkeye

# get dependencies
RUN go mod tidy

# build
RUN apk add build-base
RUN CGO_ENABLED=1 go build -o /tb_hawkeye
EXPOSE 80
ENV TB_HAWKEYE_DB_PATH=/var/lib/tb_hawkeye
ENV TB_HAWKEYE_PORT=80
ENV TB_HAWKEYE_ADDRESS=0.0.0.0
CMD ["/tb_hawkeye"]
