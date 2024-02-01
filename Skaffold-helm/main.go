package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetURL(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/small-test", smallTestHandler)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func smallTestHandler(w http.ResponseWriter, r *http.Request) {
	mypage := `
  <html>
  <head>
  <title>This is html page! Additional text</title>
  </head>
  <body>
	This is a small test! <br>
	Next line...
</body>
</html>
	`
	w.Write([]byte(mypage))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	mypage := `
  <html>
  <head>
  <title>This is html page! Additional text</title>
  </head>
  <body>
  <h1>Hello there! </h1>
  <p>This is a paragraph. More text for this paragraph.</p>
  <br>
  <img src="https://physicsworld.com/wp-content/uploads/2018/04/65-12-163-FEYNMAN.jpg" alt="Feynman" width="635" height="650">
  <br>
  <p> This paragraph is written just to have more text </p>
  <p> I hate lorem ipsum generated text, so I prefer to write this nonsense here </p>
  </body>
  `
	w.Write([]byte(mypage))

	// Calling few external endpoints to check the istio stats
	//GetURL("https://www.gol.bg/")
	//GetURL("https://dir.bg/")
	//GetURL("https://news.ycombinator.com/")
}
