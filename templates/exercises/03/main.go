package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
)

type dateData struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

type header []string

type dates []dateData

type templateData struct {
	Header header
	Dates  dates
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	logError(err)
	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	logError(err)

	header := lines[0]
	d := buildDates(lines[1:])

	td := templateData{
		header,
		d,
	}

	index, err := os.Create("index.html")
	logError(err)
	defer index.Close()

	err = tpl.ExecuteTemplate(index, "tpl.gohtml", td)
	logError(err)
}

func logError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func buildDates(lines [][]string) dates {
	datesData := make(dates, 0)
	for _, line := range lines {
		datesData = append(datesData, dateData{
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
		})
	}
	return datesData
}
