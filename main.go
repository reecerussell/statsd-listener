package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {
	log.Println("StatsD Listener!")

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8125})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Listening on 0.0.0.0:8125...")

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Printf("Read error: %v\n", err)
					break
				}
				messages := strings.Split(string(buf[:n]), "\n")
				for _, message := range messages {
					parts := strings.Split(message, "|")
					metricParts := strings.Split(parts[0], ":")
					metric, value := metricParts[0], metricParts[1]
					typ := parts[1]
					log.Printf("Receieved: %s (value=%s,typ=%s)\n", metric, value, typ)
				}
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	go func() {
		<-stop
		log.Println("Forcefully stopping...")
		os.Exit(1)
	}()

	log.Println("Gracefully stopping...")
	cancel()
	wg.Wait()

	log.Println("Stopped!")
}
