package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guygubaby/dytt-server/spiders"
)

func Handler(c *gin.Context, cate int) {
	page := c.DefaultQuery("page", "1")
	spider := spiders.Spider{Name: "movies"}
	url := "http://m.dytt.com/fenlei/" + strconv.Itoa(cate) + "_" + page + ".html"
	list := spider.GetPageData(url)
	c.JSON(http.StatusOK, list)
}

func MovieEndpoint(c *gin.Context) {
	Handler(c, 1)
}

func SoapEndpoint(c *gin.Context) {
	Handler(c, 2)
}

func ZongyiEndpoint(c *gin.Context) {
	Handler(c, 3)
}

func CartoonEndpoint(c *gin.Context) {
	Handler(c, 4)
}

func InitServer() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api/v1")
	{
		router.GET("/movies", MovieEndpoint)
		router.GET("/soap", SoapEndpoint)
		router.GET("/zongyi", ZongyiEndpoint)
		router.GET("/cartoon", CartoonEndpoint)
	}
	return r
}
