package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

type book struct {
	id    int     `db:"id"`
	title string  `db:"title"`
	price float64 `db:"price"`
}

//https://blog.csdn.net/qq_27870421/article/details/118387262

func main() {
	// "用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	db, err := sql.Open("mysql", "huiye:Hy2022.10@tcp(140.210.79.121:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err2 := db.Query("select * from book")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for rows.Next() {
		var b book
		err2 = rows.Scan(&b.id, &b.title, &b.price)
		fmt.Println(b)
	}

	rows.Close()
}
