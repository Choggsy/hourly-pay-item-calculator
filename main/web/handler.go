package web

import (
	"hourly-pay-item-calculator/utils"
	"html/template"
	"net/http"
	"path/filepath"
)

// WebHandler HTTP handlers for form processing
func WebHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(
		"main",
		"web",
		"templates",
		"form.html",
	)))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	hourlyRate, err1 := utils.ValidateAndParseFloat(r.FormValue("hourlyRate"))
	itemPrice, err2 := utils.ValidateAndParseFloat(r.FormValue("itemPrice"))

	if err1 != nil || err2 != nil {
		tmpl.Execute(w, map[string]string{"Error": "Invalid input"})
		return
	}
	tmpl.Execute(w, map[string]float64{"Hours": itemPrice / hourlyRate})
}
