package main

import (
	"context"
	"log"

	"entgo.io/ent/examples/fs/ent"
)

func main() {
	client, err := ent.Open("mysql", "myuser:myrootpassword@tcp(<host>:<port>)/db01?parseTIme=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
