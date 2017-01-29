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

func (s *server) UsersGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := users.Getall(s.MySQL)
	if err != nil {
		fmt.Println(err)
	}
	usersJson, _ := json.Marshal(users)
	fmt.Fprint(w, string(usersJson), "\n")
}

func (s *server) UsersGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	u := forum.User{ID: id}
	user, err := users.Get(s.MySQL, u.ID)
	if err != nil {
		fmt.Println(err)
	}
	userJson, _ := json.Marshal(user)
	fmt.Fprint(w, string(userJson), "\n")
}

func (s *server) BoardsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	boards, err := boards.Getall(s.MySQL)
	if err != nil {
		fmt.Println(err)
	}
	boardsJson, _ := json.Marshal(boards)
	fmt.Fprint(w, string(boardsJson), "\n")
}

func (s *server) BoardsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	b := forum.Board{ID: id}
	board, err := boards.Get(s.MySQL, b.ID)
	if err != nil {
		fmt.Println(err)
	}
	boardJson, _ := json.Marshal(board)
	fmt.Fprint(w, string(boardJson), "\n")
}

func (s *server) CategoriesGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories, err := categories.Getall(s.MySQL)
	if err != nil {
		fmt.Println(err)
	}
	categoriesJson, _ := json.Marshal(categories)
	fmt.Fprint(w, string(categoriesJson), "\n")
}

func (s *server) CategoriesGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	c := forum.Category{ID: id}
	category, err := categories.Get(s.MySQL, c.ID)
	if err != nil {
		fmt.Println(err)
	}
	categoryJson, _ := json.Marshal(category)
	fmt.Fprint(w, string(categoryJson), "\n")
}

func (s *server) ThreadsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := threads.Getall(s.MySQL)
	if err != nil {
		fmt.Println(err)
	}
	threadsJson, _ := json.Marshal(threads)
	fmt.Fprint(w, string(threadsJson), "\n")
}

func (s *server) ThreadsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	t := forum.Thread{ID: id}
	thread, err := threads.Get(s.MySQL, t.ID)
	if err != nil {
		fmt.Println(err)
	}
	threadJson, _ := json.Marshal(thread)
	fmt.Fprint(w, string(threadJson), "\n")
}

func (s *server) PostsGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := posts.Getall(s.MySQL)
	if err != nil {
		fmt.Println(err)
	}
	postsJson, _ := json.Marshal(posts)
	fmt.Fprint(w, string(postsJson), "\n")
}

func (s *server) PostsGetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	p := forum.Post{ID: id}
	post, err := posts.Get(s.MySQL, p.ID)
	if err != nil {
		fmt.Println(err)
	}
	postJson, _ := json.Marshal(post)
	fmt.Fprint(w, string(postJson), "\n")
}
