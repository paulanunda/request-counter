FROM golang:alpine AS builder
RUN mkdir $GOPATH/src/app
WORKDIR $GOPATH/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app .

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
