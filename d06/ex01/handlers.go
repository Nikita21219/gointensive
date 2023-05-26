package main

import (
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"log"
	"madorsky_go.site/blog/pkg/models"
	"net/http"
	"strconv"
)

type PageData struct {
	Posts    []models.Article
	PrevPage int
	NextPage int
}

func renderTemplate(page string, w http.ResponseWriter, data any) {
	files := []string{
		"./ui/html/" + page + ".page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 3
	offset := (page - 1) * limit

	a := models.ArticleModel{DB: db}
	posts, err := a.GetThreePosts(offset)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	PrevPage := -1
	NextPage := -1
	countPosts, err := a.CountPosts()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if page <= countPosts/limit {
		NextPage = page + 1
	}
	if page > 1 {
		PrevPage = page - 1
	}

	pd := &PageData{
		Posts:    posts,
		PrevPage: PrevPage,
		NextPage: NextPage,
	}

	renderTemplate("index", w, pd)
}

func getCookie(r *http.Request) *http.Cookie {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "",
	}
	for _, c := range r.Cookies() {
		if c.Name == "session" {
			cookie.Value = c.Value
			break
		}
	}
	return cookie
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate("login", w, nil)
	} else if r.Method == "POST" {
		inputLogin := r.FormValue("login")
		password := r.FormValue("password")
		// TODO get access from file
		if inputLogin == cred.Login && password == cred.Password {
			sessionID := uuid.New().String()
			cookie := &http.Cookie{
				Name:  "session",
				Value: sessionID,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
		} else {
			data := map[string]string{"Error": "Логин или пароль неверные"}
			renderTemplate("login", w, data)
		}
	}
}

func permission(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(r)
	if cookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	permission(w, r)
	if r.Method == "GET" {
		renderTemplate("admin", w, nil)
	} else if r.Method == "POST" {
		post := models.Article{Title: r.FormValue("inputTitle"), Text: r.FormValue("inputText")}
		if post.Text == "" || post.Title == "" {
			data := map[string]string{"Error": "Поля формы не могут быть пустыми"}
			renderTemplate("admin", w, data)
			return
		}
		a := models.ArticleModel{DB: db}
		err := a.Insert(post)
		if err != nil {
			data := map[string]string{"Error": "Ошибка при создании статьи"}
			fmt.Println("Error creating aricle:", err)
			renderTemplate("admin", w, data)
		} else {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	}
}
