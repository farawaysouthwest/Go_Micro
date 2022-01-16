
   
FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine

WORKDIR /root/

COPY --from=0 /app/app ./

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["./app"]