package controllers

import (
	"assignment1/database"
	"assignment1/helper"
	"assignment1/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// SECTION - Create - Create
func CreateUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()
	// * ANCHOR - การดึงเอาข้อมูลจาก Body มาใส่ตัวแปล
	user := userModel.User{
		Id:        helper.GenerateUUID(),
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := ctx.Bind(user)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": "Invalid request body"})
	}

	users := userModelHelper.Insert(&user)
	return ctx.JSON(200, map[string]interface{}{"massage": "Create user success", "user": users})
}

//!SECTION

// SECTION - Get

// NOTE - Fetch user ทั้งหมดแบบไม่เพิ่มเติม
// func GetUsers(c echo.Context) error {
// 	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
// 	users, err := userModelHelper.GetAllUsers()
// 	if err != nil {
// 		log.Println("Error get user: ", err)
// 	}
// 	return c.JSON(200, map[string]interface{}{"data": users, "massage": "success"})
// }

// NOTE -  Fetch user ทั้งหมด แล้วทำ Pagination เลย
func GetUsersPaginated(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	pagination := &helper.Pagination{
		Row:  5,
		Page: 1,
	}

	// * ANCHOR การดึงเอา QueryParam มาใส่ตัวแปล
	err := echo.QueryParamsBinder(ctx).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		BindError()
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": err.Error()})
	}

	users, err := userModelHelper.GetAllUsersPaginated(pagination)
	if err != nil {
		log.Println("Error get user: ", err)
		return ctx.JSON(500, map[string]string{"message": "Server Error"})
	}
	return ctx.JSON(200, map[string]interface{}{"data": users, "pagination": pagination, "message": "success"})
}

//!SECTION

// SECTION - Delete
// NOTE - Soft deletes
func DeleteUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	Id := ctx.Param("id")
	users := userModelHelper.SoftDelete(Id)

	if users != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": users})
	}
	return ctx.JSON(200, map[string]string{"massage": "Deleted successfully", "id": Id})
}

//!SECTION
