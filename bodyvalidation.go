package croissant

import (
	"strconv"
)

var validTypes map[string]func(string) bool = map[string]func(string) bool{
	"int":    IsInt,
	"bool":   IsBool,
	"string": IsStr,
	"float":  IsFloat,
}

/*
Check if a form body parameter can be converted to a string (always true)
*/
func IsStr(_ string) bool {
	return true
}

/*
Check if a form body parameter can be converted to an int
*/
func IsInt(param string) bool {
	_, err := strconv.Atoi(param)
	return err == nil
}

/*
Check if a form body parameter can be converted to a boolean
*/
func IsFloat(param string) bool {
	_, err := strconv.ParseFloat(param, 64)
	return err == nil
}

/*
Check if a form body parameter can be converted to a float
*/
func IsBool(param string) bool {
	_, err := strconv.ParseBool(param)
	return err == nil
}
