package main

import (
	"awesomeProject/Methods"
	"awesomeProject/Migration"
	"awesomeProject/console"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)



func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	choice := console.Choice()
	if choice == true {
		// вслучае если пользователь выбрал db используем метод для PgSQL
		http.HandleFunc("/getpost", Methods.IndexPage)
	} else {
		// для IMDB
		http.HandleFunc("/getpost", Methods.IndexPageIMDB)
		Migration.Migration()
	}


	fmt.Println("Сервер работает...")
	http.ListenAndServe(":8008",nil)
}