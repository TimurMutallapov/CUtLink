package Methods

import (
	"awesomeProject/ShortMethods"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

// метод для записи данных в PgSQL
func IndexPage (w http.ResponseWriter, r *http.Request ) {

	connStr := "user=postgres password=1488 dbname=cutlink sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	switch r.Method {

	case "GET":
		fmt.Println("pgdb")
		var mainlink string
		blink := r.FormValue("blink")
		//данные получем из основной бд
		rows := db.QueryRow("select mainlink from lin where cutlink = $1 limit 1", blink) //получаем данные через поиск значения в бд
		rows.Scan(&mainlink)
		if err != nil {
			panic(err)
		}
		fmt.Println("расширенная ссылка" + mainlink)
		//fmt.Fprintf(w, "<script>location='%s';</script>", link)

	//ни чем не отличается от IMDB, кроме записи в hash таблицу
	case "POST":

		fmt.Println("post принят, pgdb")
		mainLink := r.FormValue("link")
		lookFor := "."
		contain := strings.Contains(mainLink, lookFor)

		if contain == false {
			fmt.Fprintf(w,"введен не верный формат url")
		} else {
			fmt.Fprintf(w,"основная ссылка = %v\n", mainLink)
			nlink := ShortMethods.Shorting()
			nlink ="http://localhost:8008/" + nlink
			fmt.Fprintf(w, "сокращенная сылка = %s\n", nlink)

			_, err := db.Exec("insert into lin (Mainlink,cutlink) values ($1, $2)",mainLink, nlink)
			if err != nil{
				panic(err)
			}
		}
	}

}