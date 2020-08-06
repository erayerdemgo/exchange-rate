package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gorilla/websocket"
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
		url, err := htmlquery.LoadURL("https://dovizborsa.com/doviz")
		if err != nil {
			log.Fatal(err)
		}
		dolar, err := htmlquery.Query(url, "/html/body/div/div[2]/div/div[1]/div[1]/div/div[1]/div[1]/div[2]/span[1]")
		euro, err := htmlquery.Query(url, "/html/body/div/div[2]/div/div[1]/div[1]/div/div[1]/div[2]/div[2]/span[1]")

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
