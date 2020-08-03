package models

import (
	"fmt"
	"log"
	db "test_gin/databases"
)

type User struct {
	Id   int
	Name string
	Age  int // 设置不为空
}

func (user *User) Add() (id int64, err error) {
	insertSql := "insert into user(name,age) values(?,?) "
	fmt.Println("创建sql：", insertSql, user.Name, user.Age)
	rs, err := db.MySql.Exec(insertSql, user.Name, user.Age)

	if err != nil {
		fmt.Println("创建失败")
		log.Fatal(err)
		return
	}
	id, err = rs.LastInsertId()
	fmt.Printf("创建成功「%s」", user.Name)
	return id, err
}

func (user *User) GetUserList() (userList []User, err error) {
	userList = make([]User, 0)

	querySql := "select id,name,age from user"
	rows, err := db.MySql.Query(querySql)
	defer rows.Close() // 类似finally 在整个函数return后执行
	if err != nil {
		return
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Age, &user.Name)
		userList = append(userList, user)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return //FIXME 没有returnuserList，也能自动return吗
}
