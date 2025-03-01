package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Order represents the structure of the request payload
type Order struct {
	User_id      string  `json:"user_id"`
	Item_ids     string  `json:"item_ids"`
	Total_amount float32 `json:"total_amount"`
}

// GenerateRandomOrder generates a random order with random values
func GenerateRandomOrder() Order {
	// Generate random user_id
	userID := fmt.Sprintf("user-%d", rand.Intn(1000))

	// Generate random item_ids (random number of items between 1 and 5)
	itemCount := rand.Intn(5) + 1
	var itemIDs []string
	for i := 0; i < itemCount; i++ {
		itemIDs = append(itemIDs, fmt.Sprintf("item-%d", rand.Intn(100)))
	}
	itemIDsStr := fmt.Sprintf("%v", itemIDs)

	// Generate random total amount between 10 and 500
	totalAmount := rand.Float32()*490 + 10

	// Return the generated order
	return Order{
		User_id:      userID,
		Item_ids:     itemIDsStr,
		Total_amount: totalAmount,
	}
}

// SendRequest sends the HTTP POST request and measures the response time
func SendRequest(order Order, ch chan<- time.Duration) {
	// Marshal the order struct into JSON
	jsonData, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		ch <- 0
		return
	}

	// Send the POST request
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/order", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		ch <- 0
		return
	}
	defer resp.Body.Close()

	// Calculate the time taken for the request
	duration := time.Since(start)
	ch <- duration
}

func main() {
	// Seed random number generator

	// Number of requests to send
	numRequests := 1000

	// Channel to collect response times
	ch := make(chan time.Duration, numRequests)

	// Start a timer to measure the overall time taken
	startTime := time.Now()

	// Send requests concurrently
	for i := 0; i < numRequests; i++ {
		go SendRequest(GenerateRandomOrder(), ch)
		if i%200 == 0 {
			fmt.Println("sleeping", i)
			time.Sleep(time.Millisecond * 200)
			fmt.Println("awake", i)
		}
	}

	// Collect response times
	var totalDuration time.Duration
	var minDuration, maxDuration time.Duration
	minDuration = time.Hour // Set an initially large value for comparison
	for i := 0; i < numRequests; i++ {
		duration := <-ch
		totalDuration += duration

		// Update min and max durations
		if duration < minDuration {
			minDuration = duration
		}
		if duration > maxDuration {
			maxDuration = duration
		}
	}

	// Calculate the average duration
	avgDuration := totalDuration / time.Duration(numRequests)

	// Measure the total elapsed time
	totalElapsedTime := time.Since(startTime)

	// Print out the results
	fmt.Printf("Total time taken for %d requests: %v\n", numRequests, totalElapsedTime)
	fmt.Printf("Average response time: %v\n", avgDuration)
	fmt.Printf("Min response time: %v\n", minDuration)
	fmt.Printf("Max response time: %v\n", maxDuration)
}
