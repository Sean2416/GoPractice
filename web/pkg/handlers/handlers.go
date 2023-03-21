package handlers

import (
	"errors"
	"myapp/pkg/renders"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl")
}

func AddValues(x, y int) (int, error) {
	if x < y {
		return 0, errors.New("Some thing is wrong! ")
	}

	return (x + y), nil
}
