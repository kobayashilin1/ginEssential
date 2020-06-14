package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kobayashilin1/ginEssential/common"
	"github.com/kobayashilin1/ginEssential/model"
	"net/http"
	"strings"
)

func  AuthMiddleware() gin.HandlerFunc  {
	//gin的中间件就是一个函数，返回一个gin.Handlerfunc
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer"){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
			//如果返回的token为空或者不是以Bearer开头，返回错误代码与提示信息，放弃此次操作并返回。
		}
		tokenString = tokenString[7:]
		//否则将返回一个tokenString
		//因为Bearer已经占了七个字符，截取Bearer之后的字符即为token。

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {	//如果解析失败或者解析后的token无效。
			ctx.JSON(http.StatusUnauthorized,gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证成功后获得claim 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在，将user 信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}

}
