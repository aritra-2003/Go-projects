package models

import (
	
	"time"

	"example.com/rapi/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserId      int64      `json:"user_id"`
}

// Save event to database
func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, date, location, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
      
	res, err := stmt.Exec(e.Name, e.Description, e.Date, e.Location, e.UserId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

// Retrieve all events
func GetEvents() ([]Event, error) {
	query := "SELECT id, name, description, date, location, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.Location, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// Retrieve a single event
func GetEvent(id int64) (Event, error) {
	query := "SELECT id, name, description, date, location, user_id FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.Location, &event.UserId)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// Update an event
func UpdateEvent(event *Event) error {
	query := `
	UPDATE events
	SET name=?, description=?, date=?, location=?, user_id=?
	WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Date, event.Location, event.UserId, event.ID)

	return err
}
func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}
func (e Event) Register(userId int64)error{
	query:="INSERT INTO registrations(event_id,user_id)  values(?,?)"
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
          return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(e.ID,userId)
	return err



}
func (e Event ) CancelRegistration(userId int64)(error){

    query:="DELETE FROM registrations WHERE event_id=? and user_id=?"

	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
          return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(e.ID,userId)
	return err

}

