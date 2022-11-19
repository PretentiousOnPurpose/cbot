package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"gopkg.in/zabawaba99/firego.v1"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2"
	
)

var tpl *template.Template
var fb = firego.New("", nil)

func init() {
	d, _ := ioutil.ReadFile("fb.json")

	conf, _ := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")

	fb = firego.New("https://you.firebaseio.com", conf.Client(oauth2.NoContext))

	fmt.Println("Connected to real-time server")
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/voiceCtrl", voiceCtrl)
	http.HandleFunc("/updateTerminal", updateTerminal)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func updateTerminal(w http.ResponseWriter, r *http.Request) {
	
}

func voiceCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form.Get("voiceStr"));
	}
}
