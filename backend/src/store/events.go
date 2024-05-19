package store

import (
	"context"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	Location  string    `db:"location"`
}

func (s *Store) CreateEvent(ctx context.Context, event *Event) error {
	stmt := `
		INSERT INTO events (id, name, start_time, end_time, location)
		VALUES ($1, $2, $3, $4, $5);
	`

	_, err := s.db.Exec(stmt, event.ID, event.Name, event.StartTime, event.EndTime, event.Location)
	if err != nil {
		log.Printf("Failed to insert event into the database - values: %+v, err: %v", event, err)
		return err
	}

	return nil
}

func (s *Store) GetEvent(ctx context.Context, eventID string) (*Event, error) {
	stmt := `
		SELECT id, name, start_time, end_time, location
		FROM events
		WHERE id = $1;
	`

	var event []Event
	err := s.db.SelectContext(ctx, &event, stmt, eventID)
	if err != nil {
		log.Printf("Failed to query event table with id %s - err: %v", eventID, err)
		return nil, err
	}
	if len(event) == 0 {
		return nil, nil
	}
	if len(event) > 1 {
		log.Printf("[WARN] There was an unexpected number of events returned from the event table - len: %d", len(event))
	}

	return &event[0], nil
}
