package models

import "time"

type Word struct {
	Word       string    `json:"word" gorm:"size:100;not null;unique"`
	SortedWord string    `json:"sortedWord"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
