package main

import (
	"github.com/gin-gonic/gin"
)

func aa(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "efgh",
	})
}
func main() {
	r := gin.Default()
	r.GET("/abcd", aa)
	//r.GET("/abcd", check, aa)check是中间件，在执行完check之后或者在check里面执行aa
	r.Run(":1111") //记得加：
} //学习路线：数据库->gin->jwt(token+cookie)/?username=""->/ping(登陆成功)->/JSON{，，，}->crud->log->404->test()->apidoc   into readme.txt
//sshkey
