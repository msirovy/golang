package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	type web struct {
		Port    uint
		Domains []string
		Root    string
	}

	var myFuncs = template.FuncMap{
		"velke": func(text string) string {
			return strings.ToUpper(text)
		},
	}

	w1 := web{
		Port:    80,
		Domains: []string{"example1.com", "alias.example1.com"},
		Root:    "/home/example1.com/running/www",
	}
	w2 := web{
		Port:    80,
		Domains: []string{"example2.com", "alias.example2.com", "mail.example2.com"},
		Root:    "/home/example2.com/running/www",
	}

	// Init, hned od zacatku musi byt pridane funkce, jinak je hlaseno, ze nejsou definovane
	tpl := template.Must(template.New("").Funcs(myFuncs).ParseGlob("./tpl/*.tpl"))

	// jeste pridame jeden template
	tpl, err := tpl.ParseFiles("./tpl2/https.tpl")
	if err != nil {
		log.Println("ERR: ", err)
	}

	// vypropaguji moje funkce do sablon
	tpl = tpl.Funcs(myFuncs)

	// definuji ktery template chci renderovat s jakymi daty
	if tpl.ExecuteTemplate(os.Stdout, "https.tpl", w1) != nil {
		log.Println("ERR: Rendering error")
	}

	// definuji ktery template chci renderovat s jakymi daty
	if tpl.ExecuteTemplate(os.Stdout, "static.tpl", w2) != nil {
		log.Println("ERR: Rendering error")
	}

}
