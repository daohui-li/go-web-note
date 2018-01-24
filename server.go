package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

const (
	editStr = "/edit/"
	viewStr = "/view/"
	saveStr = "/save/"
)

var (
	// parse template files once and store it in the cache with base files as indicies
	pTemplate *template.Template

	// valid request paths
	validPathRegex     = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9_]+)$")
	validRootPathRegex = regexp.MustCompile("^[/]?$")

	pageMap map[string]Page
)

func serverInit() {
	templateFiles := []string{}
	for _, s := range []string{"view.html", "edit.html", "root.html", "new.html"} {
		templateFiles = append(templateFiles, appConfig.templateDir+"/"+s)
	}
	pTemplate = template.Must(template.ParseFiles(templateFiles...))

	pageMap = make(map[string]Page)

	pages := make(chan []Page)
	go func() {
		pages <- loadPages()
	}()
	for _, page := range <-pages {
		pageMap[page.Title] = page
	}
}
func serverRun(host string, port int) {
	http.HandleFunc(viewStr, genHandler(viewHandler))
	http.HandleFunc(editStr, genHandler(editHandler))
	http.HandleFunc(saveStr, genHandler(saveHandler))
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/", rootHandler)

	url := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println("http.ListenAndServe returns with an error", err)
	}
}

func genHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // closure
		m := validPathRegex.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, ok := pageMap[title]
	if !ok {
		http.Redirect(w, r, editStr+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view.html", &p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// p, err := load(title)
	p, ok := pageMap[title]
	if !ok {
		p = Page{Title: title, Body: []byte{}}
	}
	renderTemplate(w, "edit.html", &p)
}

func renderTemplate(w http.ResponseWriter, fname string, p *Page) {
	err := pTemplate.ExecuteTemplate(w, fname, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	fmt.Printf("saveHandler: body=%s\n", body)
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save() // persistent
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageMap[p.Title] = *p // in memory update
	http.Redirect(w, r, viewStr+p.Title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if !validRootPathRegex.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	pTemplate.ExecuteTemplate(w, "root.html", pageMap)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("newHandler: method=%s\n", r.Method)

	switch {
	case r.Method == "POST":
		title := r.FormValue("title")
		body := r.FormValue("body")
		fmt.Printf("newHandler: body=%s\n", body)
		p, ok := pageMap[title]
		if ok { // entry already exists
			http.Redirect(w, r, editStr+p.Title, http.StatusFound)
		} else {
			saveHandler(w, r, title)
		}
	default:
		p := Page{Title: "", Body: []byte{}}
		renderTemplate(w, "new.html", &p)
	}
}
