package main

import "github.com/go-sql-driver/mysql"

type Item struct {
	Id          int64  `form:"id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	UserName    string `form:"user_name"`
	Datetime    mysql.NullTime `form:"datetime"`
}
