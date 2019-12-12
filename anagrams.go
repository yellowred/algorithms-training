// Group Anagrams
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/
// 
package main

import (
	"log"
)

func main() {
	log.Println(groupAnagrams([]string{"aaan", "naaa", "naa", "ana", "aan", "aana"}))
}

type AnaGroup struct {
	Checksum []uint8
	Accum    []string
}

func groupAnagrams(strs []string) [][]string {
	anaGroups := []AnaGroup{}
	for _, wd := range strs {
		wdAnd := anagramChecksum(wd)
		found := false
		for i := 0; i < len(anaGroups); i++ {
			if anagramsEqual(anaGroups[i].Checksum, wdAnd) {
				anaGroups[i].Accum = append(anaGroups[i].Accum, wd)
				found = true
			}
		}
		if !found {
			anaGroups = append(anaGroups, AnaGroup{Checksum: wdAnd, Accum: []string{wd}})
		}
	}

	//compose output
	ans := make([][]string, len(anaGroups))
	for k, v := range anaGroups {
		ans[k] = v.Accum
	}
	return ans
}

func anagramChecksum(wd string) []uint8 {
	ana := make([]uint8, 26)
	for i := 0; i < len(wd); i++ {
		ana[rune(wd[i])-97]++
	}
	return ana
}

func anagramsEqual(a, b []uint8) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
