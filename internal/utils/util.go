package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	CH      = "ch"
	EM      = "em"
	REM     = "rem"
	PX      = "px"
	VH      = "vh"
	VW      = "vw"
	DVH     = "dvh"
	DVW     = "dvw"
	PERCENT = "%"
)

func ConvertPxToRem(val string) string {

	if !strings.Contains(val, PX) {
		return val
	}

	str := strings.TrimSuffix(val, "px")

	value, err := strconv.Atoi(str)

	if err != nil {
		return val
	}

	res := value / 16

	return fmt.Sprintf("%vrem", res)
}

func CheckPercentage(val float64) string {
	if val < 0 {
		return fmt.Sprintf("%v%%", val*100)
	}
	return fmt.Sprintf("%v%%", val)
}

func ConvertToString[T any](value T) string {
	switch v := any(value).(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	default:
		return fmt.Sprintf("%v", v)
	}
}

func GenerateGridTempalate[T any](value T) string {
	switch v := any(value).(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("repeat(%v, minmax(0, 1fr))", v)
	default:
		return "none"
	}
}
