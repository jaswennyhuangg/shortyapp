package models

import (
	"time"
)

type Shorty struct {
	ID           uint       `gorm:"primary_key"`
	Url	         string     `gorm:"column:url" json:"url"`
	ShortCode    string     `gorm:"column:shortcode" json:"shortcode"`
	RedirectCount int 		`gorm:"column:redirect_count" json:"redirectcount"`
	CreatedAt    *time.Time `sql:"index"`
	UpdatedAt    *time.Time `sql:"index"`
	LastSeen     *time.Time  `sql:"index"`
}

func (Shorty) TableName() string {
	return "rl_shorty"
}
