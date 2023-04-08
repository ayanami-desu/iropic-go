package utils

import (
	"math/rand"
	"time"
)

var Rand *rand.Rand

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[Rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func RandomSlice(n int, ls []uint) []uint{
	var s []uint
	length := len(ls)
	if n>length{
		return s
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<n;i++{
		j := r.Intn(length -i) + i
		s = append(s, ls[j])
		ls[i], ls[j] = ls[j], ls[i]
	}
	return s
}

func RangeInt64(left, right int64) int64 {
	return rand.Int63n(left+right) - left
}

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	Rand = rand.New(s)
}
