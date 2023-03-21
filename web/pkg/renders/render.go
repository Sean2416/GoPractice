package renders

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()

	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Some thing is wrong, msg: ", err))
	}

	tmp, inMap := tc[tmpl]

	if !inMap {
		fmt.Fprintf(w, fmt.Sprint("Some thing is wrong"))
	}

	buf := new(bytes.Buffer)

	err = tmp.Execute(buf, nil)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Some thing is wrong, msg: ", err))
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Some thing is wrong, msg: ", err))
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	maches, err := filepath.Glob("./templates/base.layout.tmpl")
	if err != nil {
		return myCache, err
	}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		if len(maches) > 0 {
			ts, err = ts.ParseGlob("./templates/base.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		err = CreateTemplateCache(t)
// 		if err != nil {
// 			fmt.Fprintf(w, fmt.Sprint("Some thing is wrong, msg: ", err))
// 		}
// 	}

// 	err = tc[t].Execute(w, nil)

// 	if err != nil {
// 		fmt.Fprintf(w, fmt.Sprint("Some thing is wrong, msg: ", err))
// 		return
// 	}
// }

// func CreateTemplateCache(t string) error {
// 	arr := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(arr...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }
