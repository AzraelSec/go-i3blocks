package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandColorStr() string {
	r := make([]int, 3)
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixMilli())
		r[i] = rand.Intn(255)
	}

	return fmt.Sprintf("#%x%x%x", r[0], r[1], r[2])
}
