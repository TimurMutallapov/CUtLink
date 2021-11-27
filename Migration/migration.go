package Migration

import (
	"database/sql"
	"fmt"
)
var IMDB =  make(map[string]string)
var IMAP =  make(map[string]string)

type Link struct{
	id int
	Mainlink string
	CuntLink string
}

// Данный метод реалиизует миграцию из pgSQL в Hash таблицу
// нужен для полноценной работы IMDB, после включенния сервера

func Migration (){
	connStr := "user=postgres password=1488 dbname=cutlink sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from lin")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	Links := []Link{}

	for rows.Next(){
		p := Link{}
		err := rows.Scan(&p.id, &p.Mainlink, &p.CuntLink) //получаем значения из основной бд

		if err != nil{
			fmt.Println(err)
			continue
		}
		Links = append(Links, p)
	}
	for _, p := range Links{
		fmt.Println(p.id, p.Mainlink, p.CuntLink)
	//Записываем значения из столбцов
		IMAP[p.Mainlink] = p.CuntLink
		IMDB[p.CuntLink] = p.Mainlink
	}
	fmt.Println("map migration:", IMAP)

}
