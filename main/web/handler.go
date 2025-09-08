package web

import (
	"hourly-pay-item-calculator/utils"
	"html/template"
	"net/http"
	"path/filepath"
)

// WebHandler HTTP handlers for form processing
func WebHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("main", "web", "templates", "form.html")))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	rawRate := r.FormValue("hourlyRate")
	rawPrice := r.FormValue("itemPrice")

	hourlyRate, err1 := utils.ValidateAndParseFloat(rawRate)
	itemPrice, err2 := utils.ValidateAndParseFloat(rawPrice)

	if err1 != nil || err2 != nil {
		tmpl.Execute(w, map[string]string{"Error": "Invalid input"})
		return
	}

	hours := itemPrice / hourlyRate
	tmpl.Execute(w, map[string]float64{"Hours": hours})
}
