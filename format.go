package logger

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func getFile(f string) string {
	s := strings.Split(f, "/")
	return s[len(s)-1]
}

func getFunct(f string) string {
	s := strings.Split(f, ".")
	return s[len(s)-1]
}

func formatLog(lvl string, depth int, message string) (string, error) {
	pc, f, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return "", errors.New("[GO-LOGGER] - ERROR - runtime.Caller func has fail and log can't be done")
	}
	file := getFile(f)

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "", errors.New("[GO-LOGGER] - ERROR - runtime.FuncForPC func has fail and log can't be done")
	}
	funct := getFunct(fn.Name())

	s := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s", lvl, "Date", file, funct, strconv.Itoa(line), message)
	return s, nil
}
