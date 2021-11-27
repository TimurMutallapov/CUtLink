package Methods

import (
	"awesomeProject/Migration"
	"awesomeProject/ShortMethods"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)


// метод для работы с in memory db
func IndexPageIMDB (w http.ResponseWriter, r *http.Request ) {
	connStr := "user=postgres password=1488 dbname=cutlink sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()



	switch r.Method {

	//2) метод get
	case "GET":

		fmt.Println("IMDB GET принят")
		blink := r.FormValue("blink")

		fmt.Println(blink)
		// получем значение по ключу, полученному с формы
		count, ok := Migration.IMDB[blink]
		if ok != true {
			fmt.Fprintf(w, "данная ссылка не существует")
		}

		//fmt.Println("map:", Migration.IMDB)
		fmt.Fprintf(w, "ссылка развернута = " +count)



	//1) Сначала рассмотрим метод пост, так как он будет вносить данные в IMDB
	// в качестве IMDB используются hash таблици,
	case "POST":

		fmt.Println("IMDB Post принят")
		mainLink := r.FormValue("link")
		lookFor := "."
		contain := strings.Contains(mainLink, lookFor)

		if contain == false {
			fmt.Fprintf(w,"введен не верный формат url")
		} else {

			fmt.Fprintf(w,"основная ссылка = %v\n", mainLink)
			nlink := ShortMethods.Shorting()
			nlink ="http://localhost:8008/" + nlink          //берем наш домен и присоединяем его с кодом для получения URL
			fmt.Fprintf(w, "сокращенная сылка = %s\n", nlink)
			Migration.IMAP[mainLink] = nlink	//добавляем данные в hash Таблицу с где основная ссылка является ключем
			Migration.IMDB[nlink] = mainLink	// главный ключ сгенерировання сылка
			/*
			fmt.Println("map:", Migration.IMAP)
			fmt.Println("map:", Migration.IMDB)
		    */
			//заполняем основную db, для сохранения данных на hdd
			_, err := db.Exec("insert into lin (Mainlink,cutlink) values ($1, $2)",mainLink, nlink)
			if err != nil{
				panic(err)
			}
		}
	}

}
