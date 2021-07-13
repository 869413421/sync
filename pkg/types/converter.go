package types

import (
	"strconv"
	"sync/pkg/logger"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func UInt64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		logger.Danger(err, "StringToInt Error")
	}

	return num
}

func StringToUInt64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.Danger(err, "StringToInt Error")
	}

	return num
}

func StrMapToString(mp map[string]string) string {
	var value string
	for key, val := range mp {
		value += "&" + key + "=" + val
	}
	newStr := value[1:]
	return newStr
}
