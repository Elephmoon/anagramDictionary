package models

import "time"

type Dictionary struct {
	Word       string    `json:"word"`
	SortedWord string    `json:"sortedWord"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
