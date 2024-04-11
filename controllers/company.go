package controllers

import (
	"assignment1/database"
	companymodel "assignment1/models/companyModel"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateCompanyTable(c echo.Context) error {
	companyHelper := companymodel.CompanyHelper{DB: database.DBMYSQL}
	companyHelper.CreateTable()
	return c.String(200, "Success")
}

func CreateCompany(c echo.Context) error {
	companyHelper := companymodel.CompanyHelper{DB: database.DBMYSQL}
	now := time.Now()
	// * ANCHOR - การดึงเอาข้อมูลจาก Body มาใส่ตัวแปล
	data := []companymodel.Companymodel{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}

	companies := []*companymodel.Companymodel{}
	for _, item := range data {
		item.CreatedAt = &now
		item.UpdatedAt = &now
		companies = append(companies, &item)
	}
	err = companyHelper.Insert(companies)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"massage": "Invalid request body"})
	}
	//log.Print(data)
	return c.JSON(200, map[string]interface{}{"massage": "Create success"})
}
