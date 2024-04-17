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
// @Tags User
// @Summary Create User
// @Description ทดสอบสร้าง User
// @Param Request body userModel.User.Body Param true "JSON Body"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @Router /user/create [post]
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
// @Tags User
// @Summary Create Multiple Users
// @Description ทดสอบสร้าง User แบบหลายคนพร้อมกัน
// @Param Request body []userModel.User.Body Param true "JSON Body"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @Router /users/create [post]
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
// @Tags User
// @Summary Get User
// @Description ทดสอบ get
// @Param page query int false "Page ที่"
// @Param row query int false "Limit ที่"
// @Param sort query string false "Sort ex: firstname desc"
// @Param firstname query string false "ต้องการหาชื่ออะไร"
// @Param lastname query string false "นามสกุลที่ต้องการหา"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @Router /users [get]
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

// @Tags User
// @Summary Get User Order Detail
// @Description ทดสอบ get user order detail
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @Router /users/products/order [get]
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
// @Tags User
// @Summary Update User By Id
// @Description ทดสอบ Update User ด้วย Id
// @Param Request body userModel.UserUpdate.Body Param true "JSON Body"
// @Param id path string true "ID ของผู้ใช้"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @response 4010 {object} helper.UnAuthorizeResponse "HTTP Code 200, ไม่มีการ Authorization "
// @response 4040 {object} helper.NotFoundResponse "HTTP Code 200, ไม่พบข้อมูล"
// @response 5000 {object} helper.InternalServerErrorResponse "HTTP Code 500, ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /user/{id} [put]
func UpdateById(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()
	id := ctx.Param("id")
	fields := userModel.UserUpdate{}
	err := ctx.Bind(&fields.Age)
	fields.UpdatedAt = &now
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	log.Println(fields)
	users, _ := userModelHelper.UpdateUser(id, fields)
	return ctx.JSON(200, map[string]interface{}{"massage": "Update user success", "user": users})
}

// SECTION - Update
// @Tags User
// @Summary Update Multiple User
// @Description ทดสอบ Update User หลายคนพร้อมกัน
// @Param Request body []userModel.User.Body Param true "JSON Body"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @response 4010 {object} helper.UnAuthorizeResponse "HTTP Code 200, ไม่มีการ Authorization "
// @response 4040 {object} helper.NotFoundResponse "HTTP Code 200, ไม่พบข้อมูล"
// @response 5000 {object} helper.InternalServerErrorResponse "HTTP Code 500, ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /users [put]
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
// @Tags User
// @Summary Soft Delete
// @Description ทดสอบ Delete User ด้วย Id
// @Param id path string true "ID ของผู้ใช้"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @response 4010 {object} helper.UnAuthorizeResponse "HTTP Code 200, ไม่มีการ Authorization "
// @response 4040 {object} helper.NotFoundResponse "HTTP Code 200, ไม่พบข้อมูล"
// @response 5000 {object} helper.InternalServerErrorResponse "HTTP Code 500, ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /users/{id} [delete]
func DeleteUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	Id := ctx.Param("id")
	users := userModelHelper.SoftDelete(Id)

	if users != nil {
		return ctx.JSON(400, map[string]interface{}{"massage": users})
	}
	return ctx.JSON(200, map[string]string{"massage": "Deleted successfully", "id": Id})
}

// @Tags User
// @Summary Soft Delete Multiple Users
// @Description ทดสอบ Delete User หลายคนพร้อมกัน
// @Param Request body []userModel.User.Body Param true "JSON Body"
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "HTTP Code 200, สร้างสำเร็จ"
// @response 4010 {object} helper.UnAuthorizeResponse "HTTP Code 200, ไม่มีการ Authorization "
// @response 4040 {object} helper.NotFoundResponse "HTTP Code 200, ไม่พบข้อมูล"
// @response 5000 {object} helper.InternalServerErrorResponse "HTTP Code 500, ข้อผิดพลาดภายในเซิร์ฟเวอร์"
// @Router /users/delete [delete]
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
