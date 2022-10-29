package utils

import (
	"math"
	"math/rand"
	"time"
)

func PositiveMod(d, m float64) float64 {
	res := math.Mod(d, m)
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func WeightedRandomCapacity(options []uint, maxValue uint) uint {
	weigtedSlice := []uint{}
	for _, option := range options {
		for i := option; i < maxValue; i++ {
			weigtedSlice = append(weigtedSlice, option)
		}
	}
	return RandomElement(weigtedSlice)
}

func RandomElement[T any](list []T) T {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RemoveElem[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
