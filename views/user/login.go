package user

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	type User struct {
		Name     string `binding:"required"`
		Password string `binding:"required"`
	}

	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
