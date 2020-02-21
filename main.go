package main

import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err, "xx")
		return
	}
	for {

		time.Sleep(time.Second * 2)
		url, err := htmlquery.LoadURL("https://dovizborsa.com/doviz/dolar")
		if err != nil {
			log.Fatal(err)
		}
		dolar, err := htmlquery.Query(url, "/html/body/div/div[2]/div/div[1]/div[1]/div/div[3]/div/div[2]/div/div/div[1]/div[2]/p[1]")
		euro, err := htmlquery.Query(url, "/html/body/div/div[2]/div/div[1]/div[1]/div/div[3]/div/div[2]/div/div/div[2]/div[2]/p[1]")

		if err != nil {
			log.Fatal(err)
		}

		m := map[string]string{}
		m["dolar"] = dolar.FirstChild.Data
		m["euro"] = euro.FirstChild.Data
		marshal, _ := json.Marshal(m)

		if err := conn.WriteMessage(1, marshal); err != nil {
			log.Println(err)
			return
		}

	}

}

func main() {

	http.HandleFunc("/ws", handler)
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.ListenAndServe(":8080", nil)
}

func handler1(writer http.ResponseWriter, request *http.Request) {

}
