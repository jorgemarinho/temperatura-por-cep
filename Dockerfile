FROM golang:1.22.4 as build
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o temperatura-por-cep cmd/temperaturacep/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/temperatura-por-cep .
ENTRYPOINT ["./temperatura-por-cep"]