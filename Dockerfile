FROM golang:1.12-alpine AS builder

RUN apk add git

EXPOSE 8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

##################

FROM alpine:latest
COPY --from=builder /go/bin/app /bin/app
CMD [ "/bin/app" ]