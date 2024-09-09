package db
import (
   "database/sql"
 _ "github.com/mattn/go-sqlite3"
)


var DB *sql.DB
var err error
func InitDB(){
	DB,err = sql.Open("sqlite3","api.db") //drivername, db name
   if err !=nil{
	panic("could not connect to db")
   }

   DB.SetMaxOpenConns(10) //no of conn can open (controlling ) 
   DB.SetMaxIdleConns(5)

   createTables()

}

func createTables() {
	creatEventTable := `
   CREATE TABLE IF NOT EXISTS events (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL,
   description TEXT NOT NULL,
   location TEXT NOT NULL,
   dataTime DATETIME NOT NULL,
   user_id INTEGER
   )
   `
   _, err := DB.Exec(creatEventTable)
   if err !=nil {
      panic("Could not able to create Table")
   }
}