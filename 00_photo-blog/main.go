package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	// process form submission

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fmt.Println(h.Sum(nil))

		// https://godoc.org/hash#Hash.Sum
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd() // getwd gets the location
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("wd", wd)

		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path) // func Create(name string) (*File, error)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("path", path)

		defer nf.Close()
		// copy
		mf.Seek(0, 0)

		io.Copy(nf, mf)
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// takes in a file name now also
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
