package logger

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getFile(f string) string {
	s := strings.Split(f, "/")
	return s[len(s)-1]
}

func getFunct(f string) string {
	s := strings.Split(f, ".")
	return s[len(s)-1]
}

func getDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
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

	date := getDate()

	s := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s", lvl, date, file, funct, strconv.Itoa(line), message)
	return s, nil
}
