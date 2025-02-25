package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO Part B: Add the QueryWorkflow() call and log the result.
	// Don't forget to add "context" to your imports.
	response, err := c.QueryWorkflow(context.Background(), "queries", "", "current_state")
	if err != nil {
		log.Fatalln("Error sending the Query", err)
		return
	}
	var result string
	response.Get(&result)
	log.Println("Received Query result. Result: " + result)
}
