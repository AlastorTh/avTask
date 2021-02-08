FROM golang:1.15-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download

RUN go build -o avtask ./main.go

CMD ["./avtask"]