package croissant

import (
	"strconv"
)

var ValidTypes map[string]func(string) bool = map[string]func(string) bool{
	"int":  IsInt,
	"bool": IsBool,
}

func IsInt(param string) bool {
	_, err := strconv.Atoi(param)
	return err == nil
}

func IsBool(param string) bool {
	_, err := strconv.ParseBool(param)
	return err == nil
}
