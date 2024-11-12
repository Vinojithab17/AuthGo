package models

import (
	"fmt"
	"log"
	"time"

	"example.com/go_app/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	//later: store to database
	query := `INSERT INTO events(name,description,location,datetime,user_id)
				VALUES(?,?,?,?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error Happaned in Preparing the query")
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)
	if err != nil {
		fmt.Println("Error Happaned in Execurting the  query ")

		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	fmt.Println(id)
	if err != nil {
		fmt.Println("Internal Server Error while fetching data")

		return err
	}

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Failed to Query data from Database")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventData(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id  = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `UPDATE events 
	SET 
	name = ?,
	description = ?,
	location = ?,
	datetime = ?
	WHERE id  = ?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.Datetime, event.ID)

	return err
}

func (event Event) DeleteEvent() error {
	query := `DELETE FROM 
	events
	WHERE id  = ?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID)

	return err
}
