package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	Query(string) string
	Get(string) (interface{}, bool)
	Set(string, interface{})
	AbortWithStatus(code int)
}
