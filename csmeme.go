package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: cs [int]\n")
		return
	}
	http.HandleFunc("/star", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	stars, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(w, "something's up chief\n")
	}
	star(stars, w)
}

func star(x int, w http.ResponseWriter) {
	if x == 1 {
		fmt.Fprintf(w, "*\n")
	} else if x == 0 {
		fmt.Fprintf(w, "\n")
	}else {
		fmt.Fprintf(w, "*")
		for i := x-2; i > 0; i-- {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "*\n")
		for k, _ := strconv.Atoi(os.Args[1]); k >= x; k-=2 {
			fmt.Fprintf(w, " ")
		}

		star(x - 2, w)
		
		for l, _ := strconv.Atoi(os.Args[1]); l > x; l-=2 {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "*")
		for j := x-2; j > 0; j-- {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "*\n")
	}
}