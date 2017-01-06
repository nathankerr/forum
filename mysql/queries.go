package mysql

// import (
// 	"strings"
// )

// func dbSelectAll(table string) {
// 	db.Exec("SELECT * FROM " + table)
// }

// // InsertCategory
// // dbInsert("categories(name, position)", "(?, ?)", []string{"name", "position"})

// // InsertBoard
// // dbInsert("boards(name, category, position)", "(?, ?, ?)", []string{"name", "category", "position"})

// // InsertThread
// // dbInsert("threads(user, title, board)", "(?, ?, ?)", []string{"user", "title", "board"})

// // InsertPost
// // dbInsert("posts(is_sticky, user, title, thread, content, created)", "(?, ?, ?, ?, ?, ?)", []string{"is_sticky", "user", "title", "thread", "content", int32(time.Now().Unix())})

// func dbInsert(table string, placeholder string, values []string) {
// 	db.Exec("INSERT INTO "+table+" VALUES "+placeholder, strings.Join(values, ", "))
// }

// // MoveThread
// // dbUpdate("threads", "threads.board", "threads.id", []string{"1", "1"})

// // MovePost
// // dbUpdate("posts", "posts.thread", "posts.id", []string{"1", "1"})

// func dbUpdate(table string, column string, identifier string, values []string) {
// 	db.Exec("UPDATE "+table+" SET "+column+" = ? WHERE "+identifier+" = ?", strings.Join(values, ", "))
// }

// // DeleteCategory
// // dbDelete("categories", "categories.id", "1")

// // DeleteBoard
// // dbDelete("boards", "boards.id", "1")

// // DeleteThread
// // dbDelete("threads", "threads.id", "1")

// // DeletePost
// // dbDelete("posts", "posts.id", "1")

// func dbDelete(table string, column string, value []string) {
// 	db.Exec("DELETE FROM "+table+" WHERE "+column+" = ?", value)
// }
