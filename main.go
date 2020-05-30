package main

import (
	"github.com/kobayashilin1/ginEssential/common"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql" //mysql驱动
)

func main() {
	db := common.InitDB() //初始化数据库
	defer db.Close()      //延迟关闭db

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
