// Filename: internal/models/pigs.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type Room struct {
	RoomID        int64
	Name          string
	NumOfPigsties int64
	Temperature   float64
	Humidity      float64
	CreatedAt     time.Time
}

type RoomModel struct {
	DB *sql.DB
}

func (m *RoomModel) Insert(room string, pigsty string, breed string, age string, dob time.Time, weight string, gender string) (int64, error) {
	var id int64
	// build the query
	statement := `
	             INSERT INTO pigs(room, pigsty, breed, age, dob, weight, gender )
							 VALUES($1, $2, $3, $4, $5, $6, $7)
							 RETURNING pig_id
	             `
	// build a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// write options to the database
	err := m.DB.QueryRowContext(ctx, statement, room, pigsty, breed, age, dob, weight, gender).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *RoomModel) Get() (*Pig, error) {
	var q Pig
	// build the query
	statement := `
				SELECT pig_id, room, pigsty, breed, age, dob, weight, gender
				FROM pigs
				ORDER BY RANDOM()
				LIMIT 1
	             `
	// build a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// write options to the database
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.PigID, &q.Room, &q.Pigsty, &q.Breed, &q.Age, &q.Dob, &q.Weight, &q.Gender) //m is the instance, DB. connectionpool, and we want to send the query row context
	if err != nil {
		return nil, err
	}
	return &q, err
}

func (m *RoomModel) Read() ([]*Pig, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM pigs
		
	`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()

	pigs := []*Pig{} //this will contain the pointer to all quotes

	for rows.Next() {
		q := &Pig{}
		err = rows.Scan(&q.PigID, &q.Room, &q.Pigsty, &q.Breed, &q.Age, &q.Dob, &q.Weight, &q.Gender, &q.CreatedAt)

		if err != nil {
			return nil, err
		}
		pigs = append(pigs, q) //contain first row
	}
	//check to see if there were error generated

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pigs, nil

}

func (m *RoomModel) Delete(pigID int) error {
	// create SQL statement to delete a quote with a given ID
	statement := `
		DELETE FROM pigs
		WHERE pig_id = $1
	`

	// execute the delete statement and check for errors
	_, err := m.DB.Exec(statement, pigID)
	if err != nil {

		return err

	}

	return nil
}

func (m *RoomModel) Update(q *Pig) error {
	// create SQL statement
	statement := `
        UPDATE pigs
        SET room=$1, pigsty=$2, breed=$3, age=$4, dob=$5, weight=$6, gender=$7
        WHERE pig_id=$8
    `
	// execute the SQL statement with the given quote values
	_, err := m.DB.Exec(statement, q.Room, q.Pigsty, q.Breed, q.Age, q.Dob, q.Weight, q.Gender)
	if err != nil {
		return err
	}
	return nil
}
