FROM golang:1.21-bookworm as base

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /parcellab-sre-challenge main.go

FROM gcr.io/distroless/static-debian12

COPY --from=base /parcellab-sre-challenge .

ENV SERVER_PORT 5000

CMD ["./parcellab-sre-challenge"]
