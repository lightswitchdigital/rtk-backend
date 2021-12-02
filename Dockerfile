FROM golang:latest AS builder
WORKDIR /go/src/github.com/lightswitch/rostelecom-backend/
COPY . .

RUN go get

RUN rm -rf /root/.ssh/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/lightswitch/rostelecom-backend/ .

EXPOSE 8000

CMD ["./app"]