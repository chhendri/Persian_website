FROM golang:1.13.5-buster

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
COPY . .

WORKDIR /go/golang
ENTRYPOINT [ "/bin/bash", "-c" ]
CMD [ "go run *.go" ]
