FROM golang:1.17
RUN mkdir -p /var/log/
RUN touch /var/log/cert.log
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN mv logger /usr/local/go/src/logger
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
