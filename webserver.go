package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
	"os"
	"io"
	"strconv"
)

var PORT = "3333"

func getRoot(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") == "" {
		fmt.Printf("[%s] on root (/) with cURL\n", r.Method)
	} else {
		fmt.Printf("[%s] on root (/) with header %s\n", r.Method, r.Header.Get("Content-Type"))
	}

	if r.Header.Get("Content-Type") == "application/json" {
		var data map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		value, exists := data["credit-card"] // replace "credit-card" with the actual card-card field on your front-end
		if exists {
			
			valueInt, err := strconv.Atoi(value.(string))
			
			if err != nil {
				http.Error(w, "Invalid credit card number", http.StatusBadRequest)
				return
			}
			
			if validator(valueInt) {
				io.WriteString(w, strconv.FormatBool(true))
			} else {
				io.WriteString(w, strconv.FormatBool(false))
			}

		} else {
			http.Error(w, "Missing expected field in JSON", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Expected JSON", http.StatusBadRequest)
	}
}

func validator(cc int) bool {
	count := 0
	sum := 0

	// Iterate over the digits of cc
	for cc > 0 {
		digit := cc % 10 // Extract the last digit
		cc /= 10          // Remove the last digit
		count++

		// Double the digit if count is even
		if count % 2 == 0 {
			digit *= 2
			// If the doubled digit is greater than 9, subtract 9
			if digit > 9 {
				digit -= 9
			}
		}

		// Add the digit to the sum
		sum += digit
	}
	
	return sum % 10 == 0
}

func main() {
	fmt.Printf(":: Webserver started on port %s ::\n", PORT)
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":" + PORT, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf(":: Server closed ::\n")
	} else if err != nil {
		fmt.Printf(":: Error starting server: %s\n ::", err)
		os.Exit(1)
	}

	fmt.Println(":: Webserver started on port 3333 ::")
}