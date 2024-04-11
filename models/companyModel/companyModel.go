package companymodel

import "time"

type Companymodel struct {
	Id        int     `json:"id" gorm:"primaryKey; AutoIncrement; not null" `
	Name      string     `json:"name" gorm:"name"`
	Email     string     `json:"email" gorm:"email"`
	Address   string     `json:"address" gorm:"address"`
	Phone     string     `json:"phone" gorm:"phone"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (Companymodel) TableName() string {
	return "company"
}
