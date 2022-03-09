package gin

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	// index
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index",
		})
	})
	// book
	r.GET("/books", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "books",
		})
	})
	// book detail
	r.GET("/book/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "books detial",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


