package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DbConnection *gorm.DB

func InitializeDatabase() {
	var err error

	DbConnection, err = gorm.Open("mysql", "employee_dir_api:vkZsJn0Q39x5@tcp(mariadb:3306)/employee_directory?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err)
	}

	// Disable table name's pluralization globally
	DbConnection.SingularTable(true)
}
