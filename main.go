package main

import (
	"github.com/kobayashilin1/ginEssential/common"
	"github.com/spf13/viper"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql" //mysql驱动
)

func main() {
	InitConfig()//项目开始时候启动读取配置文件
	db := common.InitDB() //初始化数据库
	defer db.Close()      //延迟关闭db

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}


func InitConfig() {
	workDir, _ := os.Getwd()//获取当前的工作目录
	viper.SetConfigName("application")//设置要读取的文件名
	viper.SetConfigType("yml")//设置要读取的文件类型
	viper.AddConfigPath(workDir + "/config")//设置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}