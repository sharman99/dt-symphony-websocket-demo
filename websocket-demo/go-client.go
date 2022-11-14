package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
	"net/url"
	"github.com/gorilla/websocket"
)

var endpoints [2]string = [2]string{"localhost:8080", "localhost:3030"}

func connect(u string, shutdown chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("connecting to %s", u)
	urll := url.URL{Scheme: "ws", Host: u, Path: "/testing_print"}
	c, _, err := websocket.DefaultDialer.Dial(urll.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
			log.Printf("write to server %s: %s", u, t.String())
		case <-shutdown:
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	shutdown := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		log.Println("interrupt")
		close(shutdown)
	}()

	var wg sync.WaitGroup
	for _, u := range endpoints { 
   		wg.Add(1)
    	go connect(u, shutdown, &wg)
	}
	wg.Wait()

}