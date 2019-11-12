package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
)

type UCert struct{
	Id int `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
	CustomerId string `db:"customerId"`
	Count int `db:"count"`

}

func main() {
	dns := "guomaixinan:b%c3d3F5@tcp(19.15.64.203:15003)/guangdong_public_seal"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("个人用户去重")
	users1 := queryDupUser(db)
	for _, euser := range users1{
		fmt.Println(euser)
		cutDupUser(db, &euser)
	}

	fmt.Println("机构用户去重")

	users2 := queryDupClient(db)
	for _, eclient := range users2{
		fmt.Println(eclient)
		cutDupClient(db, &eclient)
	}


}

/*
*	查重名重号的用户
*/
func queryDupUser(db *sql.DB) []UCert   {

	users :=  []UCert{}
	stmt, err := db.Prepare("select  name, code, count(name) as count from t_user_cert group by name, code having count >1")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	defer rows.Close()

	var i = 0
	for rows.Next(){
		i++
		var user UCert
		err := rows.Scan( &user.Name, &user.Code, &user.Id)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	fmt.Println("查到的记录数据="+ strconv.Itoa(i))

	if(len(users)<1){
		return nil
	}else{
		return users
	}

}

func cutDupUser(db *sql.DB, user0 *UCert){

	name := user0.Name
	code := user0.Code
	fmt.Printf("name=%s, code=%s\n", name, code)

	users :=  []UCert{}
	stmt, err := db.Prepare("select id, name, code, customerId from t_user_cert where name=? and code=? order by id desc")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(name, code)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	defer rows.Close()

	var i = 0
	for rows.Next(){
		var user UCert
		err := rows.Scan( &user.Id, &user.Name, &user.Code, &user.CustomerId)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		fmt.Println(user)
		i++
	}
	fmt.Println("cutDupUser-查到的用户记录数据="+ strconv.Itoa(i))
	if(i<1){
		return
	}
	customerId := users[0].CustomerId
	fmt.Println("索引为0的用户的customerId=" + customerId)

	result, err := db.Exec("update t_user_signature set customer_id= ? where signature_code in (select a.signature_code from (select signature_code from t_user_signature where customer_id in( select customerId from t_user_cert where name=? and code=?) ) a)",
		customerId, name, code)
	if err != nil{
		panic(err)
	}


	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("update parameters=%s, affect rows:%d\n", customerId, rowsAffected)

}

func queryDupClient(db *sql.DB) []UCert   {

	users :=  []UCert{}
	stmt, err := db.Prepare("select  name, code, count(name) as count from t_client_cert group by name, code having count >1")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	defer rows.Close()

	var i = 0
	for rows.Next(){
		i++
		var user UCert
		err := rows.Scan( &user.Name, &user.Code, &user.Id)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	fmt.Println("查到的记录数据="+ strconv.Itoa(i))

	if(len(users)<1){
		return nil
	}else{
		return users
	}

}

func cutDupClient(db *sql.DB, user0 *UCert){

	name := user0.Name
	code := user0.Code
	fmt.Printf("name=%s, code=%s\n", name, code)

	users :=  []UCert{}
	stmt, err := db.Prepare("select id, name, code, customerId from t_client_cert where name=? and code=? order by id desc")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(name, code)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	defer rows.Close()

	var i = 0
	for rows.Next(){
		var user UCert
		err := rows.Scan( &user.Id, &user.Name, &user.Code, &user.CustomerId)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		fmt.Println(user)
		i++
	}
	fmt.Println("cutDupUser-查到的用户记录数据="+ strconv.Itoa(i))
	if(i<1){
		return
	}
	customerId := users[0].CustomerId
	fmt.Println("索引为0的用户的customerId=" + customerId)

	result, err := db.Exec("update t_client_seal set customer_id= ? where seal_code in (select a.seal_code from (select seal_code from t_client_seal where customer_id in( select customerId from t_client_cert where name=? and code=?) ) a)",
		customerId, name, code)
	if err != nil{
		panic(err)
	}


	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("update parameters=%s, affect rows:%d\n", customerId, rowsAffected)

}





