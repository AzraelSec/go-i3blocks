package files

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func FileWrapper[R int | string](path string, op func(f *os.File) (R, error)) (R, error) {
	f, err := os.Open(path)
	if err != nil {
		return zero[R](), err
	}
	defer f.Close()

	return op(f)
}

func GetIntFileValue(f *os.File) (int, error) {
	buff, err := io.ReadAll(f)
	if err != nil {
		return 0, err
	}

	content, err := strconv.Atoi(strings.Trim(string(buff), "\n"))
	if err != nil {
		return 0, err
	}
	return content, nil
}

func zero[T interface{}]() T {
	return *new(T)
}
