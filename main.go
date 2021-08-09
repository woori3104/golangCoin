package main

import (
	"fmt"
	"github/woori3104/golangCoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

type homeData struct {
	// html template에서 쓰려면 대문자로 해야함
	PageTitle string
	Blocks    []*blockchain.Block
}

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	// home.html을 parse
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	// router
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// server open
	log.Fatal(http.ListenAndServe(port, nil))
}
