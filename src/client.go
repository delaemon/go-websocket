package main
import (
	"net/http"
	"text/template"
)

func view(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./client.html")
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", view)
	http.ListenAndServe(":8080", nil)
}
