package models

import (
	"time"

	"example.com/main.go/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string  `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int64

}
// var events = []Event{}
func(e *Event) Save() error {
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

   id, err := result.LastInsertId()
   e.ID = id
// fmt.Scan(result)
   
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

   func GetEventById(id int64)(*Event,error){
	query := " SELECT * from events WHERE id = ?"
	row := db.DB.QueryRow(query, id) //exactly one row so queryrow
    var event Event
    err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)

   if err !=nil{
	return nil,err
   }
   return &event , nil
   }

   func (event Event) Update()error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location =?, dataTime =?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	
	if err != nil{
       return err
	
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name,event.Description, event.Location, event.DateTime, event.ID)
	return err
   }

   func (event Event)Delete()error{
         query:= "DELETE FROM events where id=?"
		 stmt, err := db.DB.Prepare(query)
		 if err !=nil{
			return err
		 }

		 defer stmt.Close()

		 _ , err = stmt.Exec(event.ID)

		 return err

   }


   func(e Event) Register(userId int64) error{
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID,userId)
	return err
   }

   func (e Event) CancelRegistration(userId int64)error{
     query := "DELETE FROM registrations WHERE event_id =? AND user_id = ?"
	 stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID,userId)
	return err
   }