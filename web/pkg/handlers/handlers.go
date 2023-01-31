package handlers

import (
	"errors"
	"fmt"
	"myapp/pkg/renders"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := AddValues(1, 2)

	if err != nil {
		_, _ = fmt.Fprintf(w, fmt.Sprintf("err in : %s", err))
		return
	}

	_, _ = fmt.Fprintf(w, " The value is %d", n)
}

func AddValues(x, y int) (int, error) {
	if x < y {
		return 0, errors.New("Some thing is wrong! ")
	}

	return (x + y), nil
}
