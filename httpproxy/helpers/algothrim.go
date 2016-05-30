package helpers

import (
	"math/rand"
)

func ShuffleStrings(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleInts(slice []int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleUints(slice []uint) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleUint16s(slice []uint16) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ContainsString(slice []string, s string) bool {
	for _, s1 := range slice {
		if s == s1 {
			return true
		}
	}
	return false
}

func UniqueStrings(slice []string) []string {
	ss := make([]string, 0)
	m := map[string]struct{}{}
	for _, s := range slice {
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = struct{}{}
		ss = append(ss, s)
	}
	return ss
}
