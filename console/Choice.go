package console

import (
	"fmt"
	"strings"
)
// выводит на консоль выбор параметра запуска сервера
func Choice () bool {
	var s string
	fmt.Println("для выбора poSQL напишите db, для imdb любой символ")
	fmt.Printf("(db/...) =  ")
	_,err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "db" {
		return true
	}
	return false
}
