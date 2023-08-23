package main

import (
	"github.com/antonmedv/expr"
	"github.com/gin-gonic/gin"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

func main() {
	router := gin.Default()

	router.POST("/calculate", calculate)

	router.Run(":8080")
}
func calculate(c *gin.Context) {
	var req Request
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"result": "wrong of permitted expression", "err": err.Error()})
		return
	}
	if strings.Contains(req.Expression, "*") || strings.Contains(req.Expression, "/") {
		c.AbortWithStatusJSON(400, gin.H{"result": "only subtraction and addition is supported"})
		return
	}
	header := c.Request.Header.Get("User-Access")
	if header == "superuser" {

		result, err := expr.Compile(req.Expression)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"result": "error in expression"})
			return
		}
		c.JSON(200, gin.H{"result": result.Node.String()})
	} else {
		c.AbortWithStatusJSON(401, gin.H{"result": "you are not superuser"})
		return
	}
}
