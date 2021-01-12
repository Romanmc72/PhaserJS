/*
 * This simply serves my content.
 */


package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../phaser/")))
	log.Println(http.ListenAndServe(":8001", nil))
}
