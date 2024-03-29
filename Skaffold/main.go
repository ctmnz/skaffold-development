package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	mypage := `
  <html>
  <head>
  <title> -- Demo Page -- </title>
  </head>
  <body>
  <h1>Hello there!</h1>
  <p>This is a paragraph. More text for this paragraph...</p>
  <p>More text also on the next line. More text for this paragraph.</p>
  <br>
  <img src="https://physicsworld.com/wp-content/uploads/2018/04/65-12-163-FEYNMAN.jpg" alt="Feynman" width="635" height="650">
  <br>
  <p> This paragraph is written just to have more text </p>
  <p> I hate lorem ipsum generated text, so I prefer to write this nonsense here </p>
  </body>
  `
	w.Write([]byte(mypage))
}
