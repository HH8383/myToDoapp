package test_go_sqlite

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	//	"fmt"
)

type Mydata struct {
	ID        int
	Name      string
	Lat       float64
	Available bool
}

type Mydata_fetch struct {
	ID   int
	Name string
}

func fetch(con *sql.DB) []Mydata_fetch {
	return_fetch := make([]Mydata_fetch, 0)
	q := "select * from test_table"
	re, er := con.Query(q)
	if er != nil {
		panic(er)
	}

	for re.Next() {
		var md Mydata
		er := re.Scan(&md.ID, &md.Name, &md.Lat, &md.Available)
		if er != nil {
			panic(er)
		}
		add := Mydata_fetch{ID: md.ID, Name: md.Name}
		return_fetch = append(return_fetch, add)
	}
	return return_fetch
}

func TestSomething(t *testing.T) {
	con, er := sql.Open("sqlite3", "C:/Program Files/sqlite-tools-win-x64-3460000/data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	actual := fetch(con)
	//	fmt.Println(actual)
	expected := make([]Mydata_fetch, 0)
	add := Mydata_fetch{1, "test"}
	expected = append(expected, add)
	//	fmt.Println(expected)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Test Error!!!\n")
	}
}
