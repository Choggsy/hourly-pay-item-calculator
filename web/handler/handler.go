package handler

import (
	"fmt"
	"hourly-pay-item-calculator/utils/model"
	"hourly-pay-item-calculator/utils/validator"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	tmplPath := filepath.Join(wd, "web", "templates", "form.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	data := model.ResultData{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		hourlyRate, err1 := validator.ValidateAndParseFloat(r.FormValue("hourlyRate"))
		itemPrice, err2 := validator.ValidateAndParseFloat(r.FormValue("itemPrice"))

		DataStructureErrorHandler(err1, err2, &data, itemPrice, hourlyRate)
	}

	tmpl.Execute(w, data)
}

func DataStructureErrorHandler(
	err1 error,
	err2 error,
	data *model.ResultData,
	itemPrice float64,
	hourlyRate float64,
) {
	if err1 != nil || err2 != nil {
		data.Error = fmt.Sprintf("Hourly rate error: %v | Item price error: %v", err1, err2)
	} else {
		data.Hours = itemPrice / hourlyRate
	}
}
