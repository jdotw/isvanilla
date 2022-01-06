FROM golang AS build

WORKDIR /app

COPY ../../backend . 

RUN CGO_ENABLED=0 go build -o app cmd/scrape/main.go

##
## Deploy
##
FROM alpine:latest  

WORKDIR /root/

COPY --from=build /app/app ./

CMD ["./app"]
