package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// tpl = template.Must(template.ParseFiles("composite.gohtml"))
	tpl = template.Must(template.ParseGlob("sage.gohtml"))
}

func main() {
	// sages := map[string]string{
	// 	"India":        "Gandhi",
	// 	"America":      "MLK",
	// 	"Buddhism":     "Buddha",
	// 	"Christianity": "Jesus",
	// 	"Islam":        "Muhamad",
	// }

	b := sage{
		Name:  "Buddha",
		Motto: "The believe of no belief",
	}

	sageStructure(b)
}

func logErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func sageStructure(s sage) {
	err := tpl.Execute(os.Stdout, s)
	logErr(err)
}
func mapStructure(m map[string]string) {
	err := tpl.Execute(os.Stdout, m)
	logErr(err)
}

func compositeStructure(s []string) {
	err := tpl.Execute(os.Stdout, s)
	logErr(err)
}

func withData(s string) {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)
	logErr(err)
}

func parsingGlob() {
	err := tpl.Execute(os.Stdout, nil)
	logErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	logErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "harley.gohtml", nil)
	logErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "trident.gohtml", nil)
	logErr(err)

}

func parsingTemplates() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func basicTemplate(n string) string {
	tpl := `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>` + n + `</h1>	
	</body>
</html>
`

	return tpl
}

func writeToIndexHTML(tpl string) {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
	fmt.Println("Wrote to file 'index.html'.")
}

func writeToFile(tpl string, fn string) {
	nf, err := os.Create(fn)
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
	fmt.Println("Wrote to file '" + fn + "'.")
}
