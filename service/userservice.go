package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json {"code": "msg"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"msg": data,
	})

}

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Success 200 {string} json {"code": "msg"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	repassword := c.Query("repassword")
	password := c.Query("password")
	salt := fmt.Sprintf("%06d", rand.Int31())
	if password != repassword {
		c.JSON(-1, gin.H{
			"msg": "两次密码不一致",
		})
		return
	}
	if user.Name == "" || password == "" {
		c.JSON(-1, gin.H{
			"msg": "用户名或密码不能为空",
		})
		return
	}
	data := models.FindUserByName(user.Name)
	if !models.IsEmpty() {
		if data != nil {
			c.JSON(-1, gin.H{
				"msg": "用户名已存在",
			})
			return
		}
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"msg": "创建成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query int true "用户ID"
// @Success 200 {string} json {"code": "msg"}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @Param id query int true "用户ID"
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Param email formData string false "邮箱"
// @Param phone formData string false "手机号"
// @Param repassword formData string false "确认密码"
// @Success 200 {string} json {"code": "msg"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	repassword := c.PostForm("repassword")
	password := c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	if password != repassword {
		c.JSON(-1, gin.H{
			"msg": "两次密码不一致",
		})
	}
	user.Password = password
	if _, err := govalidator.ValidateStruct(user); err != nil {
		fmt.Println(err)
		c.JSON(-1, gin.H{
			"msg": "参数错误",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"msg": "更新成功",
		})
	}
}

// Login
// @Summary 登录
// @Tags 用户模块
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} json {"code": "msg"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user == nil {
		c.JSON(-1, gin.H{
			"msg": "用户名不存在",
		})
		return
	}
	if utils.ValidPassword(password, user.Salt, user.Password) {
		c.JSON(200, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(-1, gin.H{
			"msg": "密码错误",
		})
	}
}
