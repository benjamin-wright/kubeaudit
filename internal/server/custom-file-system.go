package server

import (
	"net/http"
	"strings"

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
			logrus.Debugf("Exists check: %+v", err)
			return false
		}

		logrus.Debugf("Exists check: true")
		return true
	}

	logrus.Debugf("Exists check: false")
	return false
}
