package userModel

import (
	"assignment1/helper"
	"log"
	"math"
	"time"

	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

// SECTION - Create
// NOTE - Insert Single values
func (u *UserModelHelper) Insert(user *User) error {
	tx := u.DB.Begin()
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// NOTE - Insert as array
func (u *UserModelHelper) InsertMultiple(users []*User) error {
	result := u.DB.Create(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//!SECTION - Create

// SECTION - Read

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
func (u *UserModelHelper) GetAllUsersPaginated(p *helper.Pagination, f *helper.Filter) ([]*User, error) {
	var users []*User

	// NOTE - เอาข้อมูลทั้งหมด เพื่อหาว่ามีจำนวนกี่ Page
	// allRow := u.DB.Where("deleted_at", nil).Find(&users)
	// p.TotalRows = allRow.RowsAffected
	// p.TotalPages = math.Ceil(float64(allRow.RowsAffected) / float64(p.Row))

	// NOTE - filter add % sting
	// f.Firstname = fmt.Sprint()

	db := u.DB.Debug().Model(&users).Where("deleted_at", nil).Where("firstname like ? and lastname like ?", "%"+f.Firstname+"%", "%"+f.Lastname+"%").Order(p.Sort).Count(&p.TotalRows)
	result := db.Limit(p.Row).Offset((p.Page - 1) * p.Row).Find(&users)
	p.TotalPages = math.Ceil(float64(p.TotalRows) / float64(p.Row))
	log.Println(p.TotalRows)
	log.Print("Pages", p.TotalPages)

	// ANCHOR - ตรวจสอบว่า Page ไม่เกินจำนวน Page ทั้งหมด
	if p.Page >= int(p.TotalPages) {
		p.Page = int(p.TotalPages)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// SECTION - RAW SQL
func (u *UserModelHelper) GetUserProductOrder() ([]UserProductOder, error) {
	result := []UserProductOder{}
	db := u.DB.Raw("SELECT u.id, u.firstname, p.name as productname, po.quantity, o.id as orderid, o.order_date FROM intern.user AS u INNER JOIN intern.order AS o ON u.id = o.user_id left join (intern.product as p left join intern.product_order as po on p.id = po.product_id) on o.id = po.order_id").Scan(&result)
	if db.Error != nil {
		log.Fatalln("Error", db.Error)
		return []UserProductOder{}, db.Error
	}
	return result, nil
}

//!SECTION - Read

// SECTION - Update
// FIXME - ต้องเปลี่ยนเพราะ ทำจริงๆจะต้องเอามาแค่ไอดี แล้วแสดงผลไปเลยว่า fields ไหนเป็นอะไร แล้วเขียนทับเพื่อไม่ให้เป็น Null
// TODO - Toggle ข้อมูลใน database เช่น active to inactive
func (u *UserModelHelper) UpdateUser(user_id string, fields User) (id string, err error) {
	tx := u.DB.Begin()
	result := tx.Where("id =?", user_id).Updates(fields)
	if result.Error != nil {
		tx.Rollback()
		return user_id, result.Error
	}
	tx.Commit()
	return user_id, nil
}

func (u *UserModelHelper) UpdateUserArray(fields []*User) (id []string, err error) {
	tx := u.DB.Begin()

	for _, field := range fields {
		result := tx.Debug().Where("id =?", field.Id).Updates(field)
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	tx.Commit()
	return id, nil
}

// !SECTION - Update

// SECTION - Delete
//NOTE - Erase out of database
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

// NOTE - SoftDelete
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

// NOTE - Soft Delete as Array
func (u *UserModelHelper) SoftArrayDelete(ids []UserId) error {
	now := time.Now()
	listId := []string{}
	for _, v := range ids {
		listId = append(listId, v.Id)
	}
	result := u.DB.Debug().Model(&User{}).Where("id IN (?)", listId).Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// !SECTION - Delete
