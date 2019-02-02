package main

import "time"

// Directory model...
type Directory struct {
	ID             uint   `gorm:"primary_key"`
	FirstName      string `gorm:"not null"`
	LastName       string
	OfficeNumber   string
	Extension      int
	CompanyCell    string
	PersonalNumber string
	Location       string
	Position       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
