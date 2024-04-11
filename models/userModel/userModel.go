package userModel

import "time"

type User struct {
	Id        string     `json:"id" gorm:"id"`
	Firstname string     `json:"firstname" gorm:"column:firstname"`
	Lastname  string     `json:"lastname" gorm:"column:lastname"`
	Age       int        `json:"age" gorm:"column:age"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type UserUpdate struct {
	Id        string     `json:"id" gorm:"id"`
	Firstname string     `json:"firstname" gorm:"column:firstname"`
	Lastname  string     `json:"lastname" gorm:"column:lastname"`
	Age       int        `json:"age" gorm:"column:age"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type UserProductOder struct {
	Id          string     `json:"id" gorm:"id"`
	Firstname   string     `json:"firstname" gorm:"firstname"`
	Productname string     `json:"productname" gorm:"productname"`
	Quantity    int        `json:"quantity" gorm:"quantity"`
	Orderid     int        `json:"orderid" gorm:"orderid"`
	OderDate    *time.Time `json:"oder_date" gorm:"oder_date"`
}

type UserId struct {
	Id string `json:"id"`
}

func (User) TableName() string {
	return "user"
}

func (UserUpdate) TableName() string {
	return "user"
}
