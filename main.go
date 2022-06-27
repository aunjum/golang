package main

import (
	"GOLANG/api"
	"bytes"
	_ "bytes"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	_ "github.com/boltdb/bolt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/thedevsaddam/renderer"
	_ "github.com/thedevsaddam/renderer"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	_ "golang.org/x/crypto/bcrypt"
	_ "io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	_ "net/smtp"
	_ "net/url"
	"os"
	_ "os"
	_ "strconv"
	"sync"
	_ "sync"
	_ "time"
)

// define a renderer object
var rnd *renderer.Render

//initialize what is needed for the renderer:
//which template engine to use,
//where to look for templates,
//and what to call the templates
func init() {
	gin.SetMode(gin.DebugMode)
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}
	rnd = renderer.New(opts)
}

func main() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/login", login)

	go func() {
		err = http.ListenAndServe(":"+port, mux)
		if err != nil {
			return
		}
		wg.Done()
	}()

	api.StartApi()
	// wait until WaitGroup is done
	wg.Wait()
}

//var tpl = template.Must(template.ParseFiles("tpl/index.html"))

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	buf := &bytes.Buffer{}
	//err := tpl.Execute(w, nil) // static tpl site
	err := rnd.HTML(w, http.StatusOK, "index", nil) // dynamic tpl site
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}

func aboutHandler(w http.ResponseWriter, _ *http.Request) {
	buf := &bytes.Buffer{}
	//err := tpl.Execute(w, nil) // static tpl site
	err := rnd.HTML(w, http.StatusOK, "about", nil) // dynamic tpl site
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	fmt.Printf(email)
	password := r.FormValue("password")
	fmt.Printf(password)
}
