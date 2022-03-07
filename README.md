# StatsD Listener

An out-of-process StatsD listener. Useful for debugging an application locally, where a StatsD agent is not accessible. There are other listeners available which may be better suited for an application's needs, however, this listener is just used to listen and log the messages.

## Getting started

If you have Docker installed, you can start the StatsD listener by running the command below, which will start the listener on port 8125.

```
> docker run --rm -it -p 8125:8125/udp reecerussell/statsd-listener
```

## Building

This application is written using the Go programming language, so to build it you'll need to ensure you have Go installed - Docker is a nice to have as well.

The process can be run using `go run`:

```
> go run main.go
```

Or, using `go build`:

```
> go build main.go && ./main
```

### Docker

If you'd like to build the listener with Docker, the command is simply:

```
> docker build -t reecerussell/statsd-listener .
```

Which you can then run using the command in the Getting started section.
