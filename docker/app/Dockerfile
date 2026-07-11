FROM golang:1.21 AS build

WORKDIR /src

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/ .
RUN CGO_ENABLED=0 go build -o /out/goflix ./cmd/goflix

FROM alpine:3.20

COPY --from=build /out/goflix /goflix

EXPOSE 9000

CMD ["/goflix"]