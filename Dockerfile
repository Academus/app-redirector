FROM golang:1.12-alpine AS builder

RUN apk add git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

##################

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /go/bin/app /bin/app
CMD [ "/bin/app" ]