package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

/*   ====== This is main ============   */

func main() {

	for {
		fmt.Println("Eng-Ru Dictionary")
		fmt.Println("Menu: ")
		fmt.Println("0 - Create new Base Data(BD)")
		fmt.Println("1 - Show all records from BD")
		fmt.Println("2 - Add record into BD")
		fmt.Println("3 - Edit record in BD")
		fmt.Println("4 - Delete record")
		fmt.Println("5 - Delete all records")
		fmt.Println("6 - Quit")

		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)
		if choice == 6{
			fmt.Println("Good bye!")
			break
		}
		switch choice {
		case 0:
			createBaseData()
			break
		case 1:
			showRecords()
			break
		case 2:
			recordDataBaseData()
			break
		case 3:
			updateEngSlovar()
			break
		case 4:
			deleteEngSlovar()
			break
		case 5:
			deleteAll()
			break
		default:
			fmt.Println("There is no such choice, from!")
			break
		}

	}
}

/* function Create BaseData */
func createBaseData(){
	fmt.Println("createBasedata")
	database, err := sql.Open("sqlite3", "./basedata.db")
	if err != nil{
		panic(err)
	}
	defer database.Close()
	tab, err := database.Exec("CREATE TABLE dictionary (eng text ,rus text)")
	if err != nil{
		panic(err)
	}
	fmt.Println(tab.LastInsertId())
}

/* function Show Records from BaseData */
func showRecords(){
	fmt.Println("---====ZZZ showRecords ZZZ ====---")
	database, err := sql.Open("sqlite3", "./basedata.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	rows, err := database.Query("SELECT * FROM dictionary")
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	var engl string
	var rusl string
	for rows.Next(){
		err := rows.Scan(&engl, &rusl)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println(engl + " - "+ rusl)
		}
		fmt.Println("===================")
	}
/* function Record data into BaseData */
func recordDataBaseData() {
	fmt.Println("recordDataBaseData")
	var eng_L, rus_L string
	fmt.Println("enter English word: ")
	fmt.Fscan(os.Stdin, &eng_L)
	fmt.Println("enter Russian word: ")
	fmt.Fscan(os.Stdin, &rus_L)
	database, err := sql.Open("sqlite3", "./basedata.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	statement, err := database.Prepare("INSERT INTO dictionary (eng, rus) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(eng_L, rus_L)
}

/* function Update records from BaseData */
func updateEngSlovar(){
	fmt.Println("updateEngSlovar")
}
/* function Delete records from BaseData */
func deleteEngSlovar() {
	 fmt.Println("deleteEngSlovar")
}
/* function Delete All Records from  BaseData */
func deleteAll(){
	fmt.Println("deleteAll")
}
