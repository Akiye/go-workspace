package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*type User struct {
	Name string `json:"name"`
}
*/
func main() {

	fmt.Println("Go MySQL Test")

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/task_group")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	//insert, err := db.Exec("INSERT INTO task_group VALUES(?, ?)")
	/*if err != nil {
		panic(err.Error())
	}
	defer insert.Close()*/

	/*results, err := db.Query("SELECT * FROM task_group")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user user
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
	}*/

	fmt.Println("Successfully Connected to MySQL database")
}
