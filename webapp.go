package main

import (
	"html/template"
    "log"
	"net/http"
	"os"
)

var templates, assets = "dist/templates/","dist/assets/"
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, files []string, p *Page) {
	ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/index" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        templates + "base.html",
        templates + "index.html",
    }

    renderTemplate(w, files, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        templates + "base.html",
        templates + "signup.html",
    }

    renderTemplate(w, files, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        templates + "base.html",
        templates + "login.html",
    }

    renderTemplate(w, files, nil)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cart" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        templates + "base.html",
        templates + "cart.html",
    }

    renderTemplate(w, files, nil)
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/payment" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        templates + "base.html",
        templates + "payment.html",
    }

    renderTemplate(w, files, nil)
}


func main() {
	http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/signup", signupHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/cart", cartHandler)
    http.HandleFunc("/payment", paymentHandler)

    http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

    http.ListenAndServe(":8080", nil)
}