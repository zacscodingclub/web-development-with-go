package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

var tpl *template.Template

type team struct {
	Name string
	City string
}

var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    firstThree,
	"fdate": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml", "time.gohtml"))
}
func main() {
	b := team{
		Name: "Orioles",
		City: "Baltimore",
	}

	c := team{
		Name: "Cardinals",
		City: "St. Louis",
	}

	cc := team{
		Name: "Cubs",
		City: "Chicago",
	}

	teams := []team{b, c, cc}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", teams)
	logErr(err)
	err = tpl.ExecuteTemplate(os.Stdout, "time.gohtml", time.Now())
	logErr(err)
}

func logErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
