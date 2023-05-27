package nomad

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var temp *template.Template

var res struct {
	colon    int
	typeBenz string
	colvo    string
}

func Colon1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		absPath, err := filepath.Abs("front/root.html")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Internal server error")
			return
		}
		temp, err = template.ParseFiles(absPath)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			fmt.Println("Internal server error", err)
			return
		}
	case http.MethodPost:
		typeBenz := r.FormValue("typeBenz")
		colvo := r.FormValue("colvo")
		if colvo != "" {
			var err error
			absPath, err := filepath.Abs("front/pay.html")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Internal server error")
				return
			}
			temp, err = template.ParseFiles(absPath)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}
			err = temp.Execute(w, nil)
			if err != nil {
				fmt.Println("Internal server error", err)
				return
			}
			res.colvo = colvo
			res.colon = 1
			fmt.Println(res)
		} else {
			var err error
			absPath, err := filepath.Abs("front/colvo.html")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Internal server error")
				return
			}
			temp, err = template.ParseFiles(absPath)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}
			err = temp.Execute(w, nil)
			if err != nil {
				fmt.Println("Internal server error", err)
				return
			}
			res.typeBenz = typeBenz
		}
	}
}
