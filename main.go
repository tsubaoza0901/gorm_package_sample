package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// --------
// model↓
// --------

// User ...
type User struct {
	ID        uint     `json:"id" gorm:"id"`
	Name      string   `json:"name" gorm:"name"`
	Age       int      `json:"age" gorm:"age"`
	CompanyID uint     `json:"company_id" gorm:"company_id"`
	Company   *Company `json:"company" gorm:""`
	// Role Role   `json:"role" gorm:"role"`
	// Languages []Language `json:"languages" gorm:"many2many:user_language_relations"`
	// UserLanguageRelations []UserLanguageRelation `json:"user_language_relations" gorm:""`
	// CreditCards []CreditCard `json:"credit_cards" gorm:""`
}

// Language ...
type Language struct {
	ID           uint   `json:"id" gorm:"id"`
	LanguageName string `json:"language_name" gorm:"language_name"`
}

// UserLanguageRelation ...
type UserLanguageRelation struct {
	UserID     uint `json:"user_id" gorm:"user_id"`
	LanguageID uint `json:"language_id" gorm:"language_id"`
}

// CreditCard ...
type CreditCard struct {
	ID         uint   `json:"id" gorm:"id"`
	CardNumber string `json:"card_number" gorm:"card_number"`
	UserID     int    `json:"user_id" gorm:"user_id"`
}

// Role ...
type Role struct {
	ID       uint   `json:"id" gorm:"id"`
	RoleName string `json:"role_name" gorm:"role_name"`
	UserID   int    `json:"user_id" gorm:"user_id"`
}

// Company ...
type Company struct {
	ID   uint   `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

// --------
// infrastructure↓
// --------

var db *gorm.DB

// InitDB ...
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/gormsample?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// --------
// router↓
// --------

// InitRouting ...
func InitRouting(e *echo.Echo, u *User) {
	e.POST("user", u.CreateUser)
	e.PUT("user/:id", u.UpdateUser)
	e.DELETE("user/:id", u.DeleteUser)
	e.GET("user/:id", u.GetUser)
	e.GET("users", u.GetUsers)
}

// --------
// handler↓
// --------

// CreateUser ...
func (u *User) CreateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}
	err := db.Debug().Create(&user).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user.ID)
}

// UpdateUser ...
func (u *User) UpdateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	tx := db.Begin()

	// 更新時は対象ユーザーのリレーション情報を一度すべて削除して更新
	err := tx.Debug().Where("user_id = ?", user.ID).Delete(&UserLanguageRelation{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Debug().Save(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return c.JSON(http.StatusOK, "Updated")
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}
	err := db.Debug().Delete(&u).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Deleted")
}

// GetUser ...
func (u *User) GetUser(c echo.Context) error {
	user := User{}

	id := c.Param("id")

	err := db.Debug().Preload("Company").Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers ...
func (u *User) GetUsers(c echo.Context) error {
	users := []*User{}

	err := db.Debug().Preload("Company").Find(&users).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// --------
// main.go↓
// --------

func main() {
	db = InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	e := echo.New()

	u := new(User)
	InitRouting(e, u)

	e.Start(":9002")
}
