package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, body())
}

func body() string {
	return `<html>
	<head>
	<style>
		h1 {
			margin-top: 3%;
			text-align: center;
			}
		p {
			text-align: center;
			font-size: x-large;
			font-weight: bold;
			color: #407bab;
			}
	</style>
	</head>
	
	<body> 
	
		<h1>&#129305; &#127796; &#9973;</h1>
		<p>Hello Go!</p>
	
	</body>
	</html>`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
