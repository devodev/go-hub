package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	var addr string
	var isClient bool
	flag.StringVar(&addr, "addr", ":8080", "listening address, or connection address if --client is specified")
	flag.BoolVar(&isClient, "client", false, "run a client")
	flag.Parse()

	if isClient {
		if err := startCmdlineClient(addr); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := serveHub(addr); err != nil {
		log.Fatal(err)
	}
}

func serveHub(addr string) error {
	hub := newHub()
	go hub.Run()

	log.Printf("HTTP Server listening on %q", addr)
	return http.ListenAndServe(addr, hub.Handler())
}

func read(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			lines <- scan.Text()
		}
	}()
	return lines
}

func startCmdlineClient(addr string) error {
	ctx := context.Background()

	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}

	log.Printf("Connecting to hub at %q", u.String())
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), http.Header{})
	if err != nil {
		return err
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	stdinChan := read(os.Stdin)
	done := make(chan struct{})

	// Read from websocket connection and write to stdout
	go func() {
		defer close(done)

		conn.SetReadLimit(maxMessageSize)
		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("read error: %v", err)
				}
				log.Printf("clean error: %v", err)
				break
			}
			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			log.Printf("read: %s", message)
		}
	}()

	// Read from stdin and write to websocket connection
Loop:
	for {
		select {
		case <-done:
			break Loop
		case message := <-stdinChan:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				return fmt.Errorf("write error: %v", err)
			}
		case <-interrupt:
			log.Println("interrupt")
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return fmt.Errorf("write close error: %v", err)
			}
			break Loop
		}
	}
	return nil
}
