package server

import (
	"net/http"
	"strings"

	"github.com/benjamin-wright/kubeaudit/web/dist"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type customFileSystem struct {
	fs http.FileSystem
}

func (c *customFileSystem) Open(name string) (http.File, error) {
	return c.fs.Open(name)
}

func (c *customFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := c.fs.Open(p); err != nil {
			logrus.Infof("Exists check: %+v", err)
			return false
		}

		logrus.Infof("Exists check: true")
		return true
	}

	logrus.Infof("Exists check: false")
	return false
}

func Serve() {
	r := gin.Default()

	f := dist.GetWebContent()
	fs := http.FS(f)

	entries, err := f.ReadDir(".")
	if err != nil {
		logrus.Errorf("Failed to list dirs: %+v", err)
	} else {
		for index, entry := range entries {
			logrus.Infof("%d - %s", index, entry.Name())
		}
	}

	r.Use(static.Serve("/", &customFileSystem{fs: fs}))

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run("0.0.0.0:3001")
}
