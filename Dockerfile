FROM golang:1.16-alpine AS build
WORKDIR /src

COPY go.* main.go ./

RUN go mod download

ENV GOOS=linux
ENV GO_ARCH=amd64
ENV CGO_ENABLED=0

RUN go build -o /app/listener main.go

FROM scratch
LABEL maintainer="Reece Russell <me@reece-russell.co.uk>"
LABEL repository="https://github.com/reecerussell/statsd-listener"

WORKDIR /app
COPY --from=build /app/listener listener

EXPOSE 8125/udp
ENTRYPOINT [ "./listener" ]