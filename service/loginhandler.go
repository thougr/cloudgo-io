package service

import (
	"net/http"

	"github.com/unrolled/render"
)
func LoginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		var sex string

		if req.Form["sex"][0] == "1" {
			sex = "Male"
		} else if req.Form["sex"][0] == "2" {
			sex = "Female"
		} else if req.Form["sex"][0] == "3" {
                        sex = "Secret"
    }

		formatter.HTML(w, http.StatusOK, "table", struct {
			Name   string
			Sex    string
      Addr   string
			Phone string
			Email string
		}{Name: req.Form["name"][0], Sex: sex, Addr: req.Form["addr"][0], Phone: req.Form["phone"][0], Email: req.Form["email"][0]})
	}
}
