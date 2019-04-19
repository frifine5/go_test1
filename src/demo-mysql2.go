package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Age int `db:"age"`
}

// query
func PrepareQuery(db *sql.DB, id int) {
	stmt, err := db.Prepare("select id, name, age from user where id>?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	defer rows.Close()

	for rows.Next(){
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("user: %#v\n", user)
	}
}

// insert
func Insert(db *sql.DB){

	name := "vincent"
	age := 18
	result, err := db.Exec("insert into user(name, age) values(?,?)", name, age)
	if(err != nil){
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Printf("last insert id:%d\naffect rows:%d\n", id, affected)

}

// delete
func Delete(db *sql.DB, age int) {
	result, err := db.Exec("delete from user where age=?", age)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("delect age:%d, affect rows:%d\n", age, rowsAffected)
}

// update
func Update(db *sql.DB) {
	name := "Miles"
	age := 88
	id := 1

	result, err := db.Exec("update user set name=?, age=? where id=?", name, age, id)
	if err != nil {
		panic(err)
	}

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("update id:%d, affect rows:%d\n", id, rowsAffected)

}


func main() {

	dns := "root:123456@tcp(localhost:3306)/tgo1"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	PrepareQuery(db, 0)
	Insert(db)
	//Delete(db, 18)
	//Update(db)


	fmt.Println("--------- end execute-----------")
	PrepareQuery(db, 0)

}
