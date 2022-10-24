package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("generator.html")
	err := t.Execute(w, p)
	if err != nil {
		return
	}
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	err := png.Encode(w, qrCode)
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("This is BarCode Project")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator/", viewCodeHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
