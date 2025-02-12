package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

const (
	table = "test"
	db    = "projects/labs-169405/instances/alphaus-dev/databases/main"
)

func main() {
	client, err := spanner.NewClient(
		context.Background(),
		db,
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}
	defer client.Close()

	/**

	Query

	**/

	row := client.Single().Query(context.Background(), spanner.Statement{
		SQL: fmt.Sprintf("SELECT * FROM %v", table),
	})

	// We can also query with parameters, try to uncomment the lines below and comment the query above
	// idToQuery := "123"
	// row := client.Single().Query(context.Background(), spanner.Statement{
	// 	SQL: fmt.Sprintf("SELECT * FROM %v WHERE id = @id ", table),
	// 	Params: map[string]interface{}{
	// 		"id": idToQuery,
	// 	},
	// })

	var sample struct {
		Id   spanner.NullString
		Name spanner.NullString
		Age  spanner.NullInt64
	}

	for {
		row, err := row.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		err = row.ToStruct(&sample)
		if err != nil {
			log.Printf("row.ToStruct failed: %v", err)
			break
		}

		log.Println(sample.Id.StringVal, sample.Name.StringVal, sample.Age.Int64)
	}
	/**

	Insert or Update

	**/
	// cols := []string{"id", "name", "age"}
	// vals := []interface{}{"1234", "mynamename", 123}
	// t, err := client.Apply(context.Background(), []*spanner.Mutation{
	// 	spanner.InsertOrUpdate(table, cols, vals),
	// })
	// if err != nil {
	// 	log.Printf("InsertOrUpdate failed: %v", err)
	// 	return
	// }

	// log.Println("time inserted/updated", t)

	/**

	Delete

	**/
	// idToDelete := "1234"
	// t, err = client.Apply(context.Background(), []*spanner.Mutation{
	// 	spanner.Delete(table, spanner.Key{idToDelete}),
	// })
	// if err != nil {
	// 	log.Printf("Delete failed: %v", err)
	// 	return
	// }

	// log.Println("time deleted", t)
}
