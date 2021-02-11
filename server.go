package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Compute sum of integers in request body
		n, sum := 0, 0
		scanner := bufio.NewScanner(r.Body)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			text := scanner.Text()
			i, err := strconv.Atoi(text)
			if err != nil {
				log.Printf("Error parsing %q into int: %v", text, err)
			}
			n++
			sum += i
		}
		err := r.Body.Close()
		if err != nil {
			log.Println("Closing request body:", err)
		}

		log.Println("Computed sum of", n, "integers:", sum)
		fmt.Fprintln(w, sum)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}
