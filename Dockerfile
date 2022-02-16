FROM golang:1.17.7-alpine3.15 as build
WORKDIR /app
COPY . .
RUN go build main.go

FROM alpine:3.15
COPY --from=build /app/main /
CMD ["/main"]
