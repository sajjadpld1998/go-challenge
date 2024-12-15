#---------- builder -------------------------
FROM golang:1.20 AS builder
WORKDIR /

COPY ./app/go.mod ./
RUN go mod download

COPY ./app ./

COPY ./app/configs.yaml.example ./configs.yaml

RUN CGO_ENABLED=0 go build -o main main.go

EXPOSE 8484

CMD ["/main"]
