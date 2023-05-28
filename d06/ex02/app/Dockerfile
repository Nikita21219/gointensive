FROM golang:alpine

RUN apk add unzip

WORKDIR /app

COPY ./*.go .
COPY pkg pkg
COPY go.mod .
COPY go.sum .
COPY admin_credentials.txt .
COPY ui.zip ui.zip

RUN unzip ui.zip -d . && go build -o blog *.go

CMD ["./blog"]
