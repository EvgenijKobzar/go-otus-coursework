package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const ActionGet = "get"
const ActionAdd = "add"
const ActionList = "list"
const ActionUpdate = "update"
const ActionDelete = "delete"
const ActionRegister = "register"
const ActionLogin = "login"

func GetActonByFullPath(handlerName string) string {

	paths := strings.Split(handlerName, "/")
	items := strings.Split(paths[2], ".")
	if len(items) > 0 && len(items[0]) > 0 {
		return items[len(items)-1]
	}
	return ""
}

func internalizeGet(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New("invalid id")
	} else if id == 0 {
		return errors.New("id is required")
	}
	return nil
}
func internalizeAdd(c *gin.Context) error {
	var err error
	if err = payloadIsEmpty(c); err != nil {
		return err
	}
	return nil
}
func internalizeUpdate(c *gin.Context) error {
	var err error
	if err = internalizeGet(c); err != nil {
		return err
	}
	if err = payloadIsEmpty(c); err != nil {
		return err
	}
	return nil
}
func internalizeList(c *gin.Context) error {
	var err error
	return err
}
func internalizeDelete(c *gin.Context) error {
	return internalizeGet(c)
}
func payloadIsEmpty(c *gin.Context) error {
	req := c.Request
	if req == nil || req.ContentLength == 0 {
		return errors.New("fields is not specified")
	}

	return nil
}
func Internalize(c *gin.Context) (error, bool) {
	var err error
	action := GetActonByFullPath(c.FullPath())
	switch action {
	case ActionGet:
		err = internalizeGet(c)
	case ActionAdd:
		err = internalizeAdd(c)
	case ActionList:
		err = internalizeList(c)
	case ActionUpdate:
		err = internalizeUpdate(c)
	case ActionDelete:
		err = internalizeDelete(c)
	case ActionRegister, ActionLogin:
		err = nil
	default:
		err = errors.New("invalid action")
	}

	return err, err == nil
}
