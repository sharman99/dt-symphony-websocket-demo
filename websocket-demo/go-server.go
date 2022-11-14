package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var serverAddr = flag.String("addr", "localhost:3030", "http service address")

var upgrader = websocket.Upgrader{} 

func testing_print(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv from client: %s", message)
		err = c.WriteMessage(mt, append([]byte("GO SERVER: "), message...))
		if err != nil {
			log.Println("write:", err)
			break
		}
		log.Printf("write to client: %s", message)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/testing_print", testing_print)
	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}
