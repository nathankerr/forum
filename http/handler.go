package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql/boards"
	"github.com/dhenkes/forum/mysql/categories"
	"github.com/dhenkes/forum/mysql/posts"
	"github.com/dhenkes/forum/mysql/threads"
	"github.com/dhenkes/forum/mysql/users"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Index\n")
}

func UsersGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := users.Getall()
	if err != nil {
		fmt.Println(err)
	}
	usersJson, _ := json.Marshal(users)
	fmt.Fprint(w, string(usersJson), "\n")
}

func UsersGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	u := forum.User{ID: id}
	user, err := users.Get(u.ID)
	if err != nil {
		fmt.Println(err)
	}
	userJson, _ := json.Marshal(user)
	fmt.Fprint(w, string(userJson), "\n")
}

func BoardsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	boards, err := boards.Getall()
	if err != nil {
		fmt.Println(err)
	}
	boardsJson, _ := json.Marshal(boards)
	fmt.Fprint(w, string(boardsJson), "\n")
}

func BoardsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	b := forum.Board{ID: id}
	board, err := boards.Get(b.ID)
	if err != nil {
		fmt.Println(err)
	}
	boardJson, _ := json.Marshal(board)
	fmt.Fprint(w, string(boardJson), "\n")
}

func CategoriesGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories, err := categories.Getall()
	if err != nil {
		fmt.Println(err)
	}
	categoriesJson, _ := json.Marshal(categories)
	fmt.Fprint(w, string(categoriesJson), "\n")
}

func CategoriesGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	c := forum.Category{ID: id}
	category, err := categories.Get(c.ID)
	if err != nil {
		fmt.Println(err)
	}
	categoryJson, _ := json.Marshal(category)
	fmt.Fprint(w, string(categoryJson), "\n")
}

func ThreadsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := threads.Getall()
	if err != nil {
		fmt.Println(err)
	}
	threadsJson, _ := json.Marshal(threads)
	fmt.Fprint(w, string(threadsJson), "\n")
}

func ThreadsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	t := forum.Thread{ID: id}
	thread, err := threads.Get(t.ID)
	if err != nil {
		fmt.Println(err)
	}
	threadJson, _ := json.Marshal(thread)
	fmt.Fprint(w, string(threadJson), "\n")
}

func PostsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := posts.Getall()
	if err != nil {
		fmt.Println(err)
	}
	postsJson, _ := json.Marshal(posts)
	fmt.Fprint(w, string(postsJson), "\n")
}

func PostsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	p := forum.Post{ID: id}
	post, err := posts.Get(p.ID)
	if err != nil {
		fmt.Println(err)
	}
	postJson, _ := json.Marshal(post)
	fmt.Fprint(w, string(postJson), "\n")
}
