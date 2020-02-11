FROM golang:alpine as builder

RUN mkdir service

WORKDIR /service/

COPY . .

RUN env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /service/runner



FROM scratch

COPY --from=builder /service/runner /service/runner

EXPOSE 8888

ENTRYPOINT [ "/service/runner" ]