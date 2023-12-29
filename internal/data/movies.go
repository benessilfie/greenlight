package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Genres    []string  `json:"genres"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Version   int32     `json:"version,omitempty,string"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"-"`
}
