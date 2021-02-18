package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gmock/conf"
)

func Run() error {
	c, err := conf.GetConfigure()
	if err != nil {
		return errors.WithStack(err)
	}

	r := gin.Default()
	r.Use(route)

	for _, cf := range c.Conf {
		r.Handle(strings.ToUpper(cf.Method), cf.URI, handler)
	}

	return r.Run()
}

func route(ctx *gin.Context) {
	c, err := conf.GetConfigure()
	if err != nil {
		panic(err)
	}

	for _, cf := range c.Conf {
		ctx.Set(cf.URI, cf)
	}
}

func handler(c *gin.Context) {
	_cf, exist := c.Get(c.Request.RequestURI)
	if !exist {
		c.JSON(600, gin.H{"err": "not found uri"})
		return
	}

	cf := _cf.(conf.Conf)

	//c.Status()
	for key, value := range cf.Header {
		c.Header(key, value)
	}

	switch cf.Body.Type {
	case "string":
		c.String(cf.Status, cf.Body.Data.(string))
	case "json":
		c.JSON(cf.Status, cf.Body.Data)
	}
}
