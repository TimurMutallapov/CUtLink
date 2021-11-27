package main

import (
	//"awesomeProject/Methods"
	"awesomeProject/Migration"
	"awesomeProject/ShortMethods"
	"awesomeProject/console"
	"testing"
	"time"
)



func ShortTest(t *testing.T) {


	result := ShortMethods.Shorting()

	if len(result) != 9 {

		t.Errorf("Incoorect result: %v,%s", time.Now(), result)
	}
}

func MigrationTest(t *testing.T) {

	if Migration.IMDB == nil && Migration.IMAP ==  nil{

		t.Errorf("Incoorect result: %v,%v,%v", time.Now(), Migration.IMDB, Migration.IMAP)
	}
}

func ChoiceTest (t *testing.T)  {


	console.Choice()
	chtest:= console.Choice()

	if  chtest == false {

		t.Errorf("Incoorect result: %v, %v", time.Now(), chtest )
	}

}

