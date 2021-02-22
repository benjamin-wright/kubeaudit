package server

import (
	"net/http"
	"sync"

	"github.com/benjamin-wright/kubeaudit/web/dist"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Serve serve some files
func Serve(wg *sync.WaitGroup) {
	defer wg.Done()

	gin.SetMode("release")

	r := gin.New()

	f := dist.GetWebContent()
	fs := http.FS(f)

	indexPage, err := f.ReadFile("index.html")
	if err != nil {
		logrus.Errorf("Failed to load index page: %+v", err)
		logrus.Exit(1)
	}

	r.Use(static.Serve("/", &customFileSystem{fs: fs}))

	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(indexPage)
	})

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run("0.0.0.0:3001")
}
