package logger

import (
	"errors"
	"runtime"
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

func getInfos(depth int) (string, string, int, error) {
	pc, f, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return "", "", 0, errors.New("[GO-LOGGER] - ERROR - runtime.Caller function has fail")
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "", "", 0, errors.New("[GO-LOGGER] - ERROR - runtime.FuncForPC has fail")
	}

	file := getFile(f)
	funct := getFunct(fn.Name())
	return file, funct, line, nil
}
