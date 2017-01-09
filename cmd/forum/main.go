package main

import (
	// 	"fmt"

	"github.com/dhenkes/forum/mysql"
	// 	"github.com/dhenkes/forum/mysql/users"
)

func main() {

	// Check if tables exist.
	mysql.CheckTables()

	// // Get user with ID 1.
	// user, err := users.Get(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	// // Get all users.
	// users, err := users.Getall()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(users)

	// // Create user.
	// res, err := users.Create()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(res[:]))
}
