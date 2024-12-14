#---------- builder -------------------------
FROM golang:buster as builder
WORKDIR /

COPY ./app/go.mod ./
RUN go mod download

COPY ./app ./

RUN CGO_ENABLED=0 go build -o main main.go

EXPOSE 8484

CMD ["/main"]
