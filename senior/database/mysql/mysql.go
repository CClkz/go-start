package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/cc_db")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	defer Db.Close() // 确保在main函数结束时关闭数据库连接
	// insertData()
	getData()
	// updateData("name100")
	// deleteData()
}
