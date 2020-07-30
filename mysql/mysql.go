package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 连接
func Connect() *sql.DB {
	connStr := "root:123456@tcp(127.0.0.1:3306)/go-gin-sample"
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return db
}

// 创建表
func CreateTable(db *sql.DB) {
	var SQL = "CREATE TABLE person (" +
		"id int auto_increment primary key," +
		"name varchar(12) not null," +
		"age int default 1" +
		");"
	_, err := db.Exec(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// 插入
func Insert(db *sql.DB) {
	var SQL = "INSERT INTO person(name, age) " +
		"VALUES(?, ?)"
	_, err := db.Exec(SQL, "Jack", 20)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据插入成功")
	}
}

// 查询
func Select(db *sql.DB) {
	var SQL = "SELECT id, name, age FROM person"
	rows, err := db.Query(SQL)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}
}

type Person struct {
	Id   int
	Name string
	Age  int
}

//---------------------------------------------

// 测试创建表
func TestCreateTable() {
	db := Connect()
	CreateTable(db)
}

// 测试插入
func TestInsert() {
	db := Connect()
	Insert(db)
}

// 测试查询
func TestSelect() {
	db := Connect()
	Select(db)
}
