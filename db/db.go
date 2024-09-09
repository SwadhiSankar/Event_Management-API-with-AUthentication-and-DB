package db
import ("database/sql"
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
   DB.SetConnMaxIdleTime(5)
}