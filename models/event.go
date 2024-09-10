package models

import (
	"fmt"
	"time"

	"example.com/main.go/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string  `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int

}
// var events = []Event{}
func(e Event) Save() error {
	//db code 
	query := `INSERT INTO events(name, description, location, dataTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query) //stored in memory and easily reusable
	if err !=nil{
		return err
	}
   defer stmt.Close()
	result, err :=stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
   if err !=nil{
	return err
   }

//    id, err := result.LastInsertId()
//    e.ID = id
fmt.Scan(result)
   
   return err

}

func GetAllEvent() ([]Event,error){
	query := "SELECT * from events"
    rows, err :=	db.DB.Query(query) // easy and evalute for small operation query and exec both r same operation (query-> to retrieve or select the rows in the table)
   if err!=nil{
	return nil,err
    }	
    defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err :=rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
		if err!= nil{
			return nil,err
		}
		events = append(events, event)

	}
    return events,nil
   }