package main

import (
	"context"
	"demo_goent/ent/migrate"
	"log"
	"os"

	"entgo.io/ent/examples/fs/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// dump migration changes to stdout
	if err := client.Schema.WriteTo(ctx, os.Stdout, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Print("ent sample done")
}

// -- DB接続確認
// func main() {
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// 	log.Print("ent sample done")
// }
