package companymodel

import "gorm.io/gorm"

type CompanyHelper struct {
	DB *gorm.DB
}

// SECTION - Create

func (c *CompanyHelper) CreateTable() error {
	tx := c.DB.Begin()
	result := tx.AutoMigrate(&Companymodel{})
	if result != nil {
		tx.Rollback()
		return result
	}
	tx.Commit()
	return nil
}

func (c *CompanyHelper) Insert(cpn []*Companymodel) error {
	tx := c.DB.Begin()
	result := tx.Create(&cpn)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// !Section - Create
