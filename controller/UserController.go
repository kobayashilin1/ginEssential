package controller

import (
	"github.com/kobayashilin1/ginEssential/common"
	"github.com/kobayashilin1/ginEssential/model"
	"github.com/kobayashilin1/ginEssential/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//电话号码数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号码需为11位"})
		//gin.H: map[string]interface{}
		return
		//停止继续往下执行
	}

	//验证密码
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//如果名称没有，给一个随机的10位字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	//判断手机号是否存在

	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在，请勿重复注册"})
		return
	}
	//用户存在就不允许注册

	//创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),//强制类型转换，将[]byte slice转换为string
	}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"code":"200",
		"message": "注册成功",
	})
}

func Login(ctx *gin.Context){
	DB := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号码需为11位"})
		//gin.H: map[string]interface{}
		return
		//停止继续往下执行
	}

	//验证密码
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//判断手机号码是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断用户密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token给前端
	token := "111"

	//返回结果
	ctx.JSON(200, gin.H{
		"code":"200",
		"date":gin.H{"token":token},
		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
