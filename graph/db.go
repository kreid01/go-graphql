package graph

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"kreid.com/graphl-go/graph/model"
	"github.com/go-pg/pg/v10/orm"
)

func Connect() *pg.DB {
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	
  //	err = createSchema(db)
//	  if err != nil {
  //      panic(err)
   // }

	var n int
	_, err = db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		panic(fmt.Sprintf("PostgreSQL connection failed: %v", err))
	}

	fmt.Println("PostgreSQL connection established successfully")
	return db
}


func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Message)(nil),
	}

	 for _, model := range models {
        	err := db.Model(model).CreateTable(
		&orm.CreateTableOptions{
        	    Temp: false,
        	})
        if err != nil {
            return err
        }
    }
	
	return nil
}


