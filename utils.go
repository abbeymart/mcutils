// Utility functions: shared library functions for mc-solutions
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"github.com/abbeymart/mcdatago"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
	Hosts        []map[string]string
}

var (
	Config Configuration
	Logger *log.Logger
)

// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("../config/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// Convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
func Session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

// parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
func Info(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

func Danger(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// version
func Version() string {
	return "0.1"
}

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			Danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot read thread")
	} else {
		_, err := Session(writer, request)
		if err != nil {
			GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			Danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			ErrorMessage(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
