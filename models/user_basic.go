package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(1[3-9]\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	HeartbeatTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LoginOutTime  time.Time `gorm:"default:CURRENT_TIMESTAMP column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) error {
	now := time.Now()
	if user.LoginTime.IsZero() {
		user.LoginTime = now
	}
	if user.HeartbeatTime.IsZero() {
		user.HeartbeatTime = now
	}
	if user.LoginOutTime.IsZero() {
		user.LoginOutTime = now
	}
	return utils.DB.Create(&user).Error
}

func DeleteUser(user UserBasic) error {
	return utils.DB.Delete(&user).Error
}

func UpdateUser(user UserBasic) error {
	if user.LoginTime.IsZero() {
		user.LoginTime = time.Now()
	}
	if user.HeartbeatTime.IsZero() {
		user.HeartbeatTime = time.Now()
	}
	if user.LoginOutTime.IsZero() {
		user.LoginOutTime = time.Now()
	}
	return utils.DB.Model(&UserBasic{}).Where("id = ?", user.ID).Updates(&user).Error
}

func FindUserByName(name string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return &user
}

func FindUserByPhone(phone string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return &user
}

func FindUserByEmail(email string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return &user
}

func FindUserById(id uint) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return &user
}
