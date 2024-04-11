package controllers

import (
	"assignment1/database"
	"assignment1/helper"
	"assignment1/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// SECTION - Create
// NOTE - Create a single user
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

// NOTE - Create a multiple user
func CreateMultipleUsers(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()
	// * ANCHOR - การดึงเอาข้อมูลจาก Body มาใส่ตัวแปล
	data := []userModel.User{}
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}

	users := []*userModel.User{}
	for _, item := range data {
		item.Id = helper.GenerateUUID()
		item.CreatedAt = &now
		item.UpdatedAt = &now
		users = append(users, &item)
	}
	err = userModelHelper.InsertMultiple(users)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	//log.Print(data)
	return ctx.JSON(200, map[string]interface{}{"massage": "Create user success"})
}

//!SECTION - Create

// SECTION - Read

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
	filter := &helper.Filter{}

	// * ANCHOR การดึงเอา QueryParam มาใส่ตัวแปล
	err := echo.QueryParamsBinder(ctx).
		Int("row", &pagination.Row).
		Int("page", &pagination.Page).
		String("sort", &pagination.Sort).
		String("firstname", &filter.Firstname).
		String("lastname", &filter.Lastname).
		BindError()
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": err.Error()})
	}

	users, err := userModelHelper.GetAllUsersPaginated(pagination, filter)
	if err != nil {
		log.Println("Error get user: ", err)
		return ctx.JSON(500, map[string]string{"message": "Server Error"})
	}
	return ctx.JSON(200, map[string]interface{}{"data": users, "pagination": pagination, "message": "success"})
}

func GetUserProductOrder(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	results, err := userModelHelper.GetUserProductOrder()
	if err != nil {
		log.Println("Error get user: ", err)
		return ctx.JSON(500, map[string]string{"message": "Database Error"})
	}
	return ctx.JSON(200, map[string]interface{}{"data": results, "message": "success"})

}

//!SECTION - Read

// SECTION - Update
func UpdateById(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	id := ctx.Param("id")
	fields := userModel.UserUpdate{}
	err := ctx.Bind(&fields)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	log.Println(fields)
	users, _ := userModelHelper.UpdateUser(id, fields)
	return ctx.JSON(200, map[string]interface{}{"massage": "Update user success", "user": users})
}

func UpdateMultiple(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	data := []userModel.User{}

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	fields := []*userModel.User{}
	for _, item := range data {
		fields = append(fields, &item)
	}
	result, _ := userModelHelper.UpdateUserArray(fields)
	return ctx.JSON(200, map[string]interface{}{"massage": "Request pass", "result": result})
}

//!SECTION - Update

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

func DeleteMultipleUsers(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	ids := []userModel.UserId{}

	err := ctx.Bind(&ids)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	log.Print(ids)
	result := userModelHelper.SoftArrayDelete(ids)
	return ctx.JSON(200, map[string]interface{}{"massage": "Request pass", "result": result})
}

//!SECTION - Delete
