package intro

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Intro() {
	fmt.Printf("I'm alive!\n")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to open db")
	}

	// migrate the schema
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{Code: "D69", Price: 420})

	// read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D69")

	// update
	db.Model(&product).Update("Price", 840)
	db.Model(&product).Updates(Product{Price: 210, Code: "Z96"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 420, "Code": "D69"})

	// delete
	//db.Delete(&product, 1)
}
