package userModel

import (
	"math"
	"time"

	"assignment1/helper"

	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

func (u *UserModelHelper) Insert(user *User) error {
	result := u.DB.Debug().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserModelHelper) FindById(id string) (*User, error) {
	user := User{}
	result := u.DB.First(&user, "id =?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// NOTE -  Fetch user ทั้งหมดแบบไม่เพิ่มเติม
// func (u *UserModelHelper) GetAllUsers() ([]*User, error) {
// 	user := []*User{}
// 	result := u.DB.Debug().Find(&user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return user, nil
// }

// NOTE -  Fetch user and Paginate
func (u *UserModelHelper) GetAllUsersPaginated(p *helper.Pagination) ([]*User, error) {
	var users []*User

	// NOTE - เอาข้อมูลทั้งหมด เพื่อหาว่ามีจำนวนกี่ Page
	allRow := u.DB.Where("deleted_at", nil).Find(&users)
	p.TotalRows = allRow.RowsAffected
	p.TotalPages = math.Ceil(float64(allRow.RowsAffected) / float64(p.Row))

	// ANCHOR - ตรวจสอบว่า Page ไม่เกินจำนวน Page ทั้งหมด
	if p.Page > int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	}

	db := u.DB.Debug().Limit(p.Row).Offset((p.Page-1)*p.Row).Where("deleted_at", nil).Order(p.Sort)
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FIXME - ต้องเปลี่ยนเพราะ ทำจริงๆจะต้องเอามาแค่ไอดี แล้วแสดงผลไปเลยว่า fields ไหนเป็นอะไร แล้วเขียนทับเพื่อไม่ให้เป็น Null
func (u *UserModelHelper) UpdateUser(user *UserUpdate) (*UserUpdate, error) {
	result := u.DB.Debug().Model(&user).Where("id =?", user.Id).Save(user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

// SECTION - Delete function แบบลบออกจากฐานข้อมูล
// func (u *UserModelHelper) FullDelete(id string) (*User, error) {
// 	user := User{}
// 	result := u.DB.Debug().Delete(&user, "id =?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	} else {
// 		log.Print("Deleted user: ", id)
// 		return &user, nil
// 	}
// }

func (u *UserModelHelper) SoftDelete(id string) error {
	now := time.Now()
	result := u.DB.Debug().Model(&User{}).Where("id = ?", id).Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	} else {
		// log.Print("Deleted user: ", id)
		return nil
	}
}

// !SECTION
