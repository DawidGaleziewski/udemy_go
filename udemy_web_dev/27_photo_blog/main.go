package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// note on sha1 - its a hashing algorithm, that will take any file and return a string of characters that will always be the same, as long as the file is the same
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := getCookie(w, r)
		pictures := getPictures(c.Value)

		if len(pictures) == 0 {
			c = &http.Cookie{
				Name:  "session",
				Value: updateWithPictures(c.Value),
			}
		}
		http.SetCookie(w, c)
		tpl.ExecuteTemplate(w, "index.gohtml", pictures)
	})

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func getPictures(cookieVal string) []string {
	return strings.Split(cookieVal, "|")[1:]
}

func updateWithPictures(cookieVal string) string {
	var pictures = []string{"cat.jpg", "dog.jpg", "chick.jpg"}
	for _, picture := range pictures {
		cookieVal += "|" + picture
	}
	return cookieVal
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.New().String()
		c = &http.Cookie{
			Name:  "session",
			Value: sID,
		}
	}

	return c
}
