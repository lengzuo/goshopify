package dto

import "time"

type Event struct {
	Arguments   string     `json:"arguments,omitempty"`
	Body        string     `json:"body,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ID          int64      `json:"id,omitempty"`
	Description string     `json:"description,omitempty"`
	Path        string     `json:"path,omitempty"`
	Message     string     `json:"message,omitempty"`
	SubjectID   int64      `json:"subject_id,omitempty"`
	SubjectType string     `json:"subject_type,omitempty"`
	Verb        string     `json:"verb,omitempty"`
}

type EventCollection struct {
	Event Event `json:"event"`
}

type CountEventRequest struct {
	CreatedAtMax *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin *time.Time `url:"created_at_min,omitempty"`
}

type GetEventRequest struct {
	CreatedAtMax *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin *time.Time `url:"created_at_min,omitempty"`
	Fields       string     `url:"fields,omitempty"`
	Filter       string     `url:"filter,omitempty"`
	Limit        string     `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	Verb         string     `url:"verb,omitempty"`
}

type GetEventResponse struct {
	Events []Event `json:"events"`
}
