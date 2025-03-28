package models

import "gorm.io/gorm"

type Cluster struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	ServerCount int    `gorm:"not null"`
}
