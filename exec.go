// Top K Frequent Elements
// Priority Queue
//  - push
//  - pop
//  - len
package main

import (
	"log"
	"strconv"
)

func main() {
	// nums := []int{4, 5, 3, 1, 2}
	log.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))
	// log.Println(isSolvable([]string{"SE", "ME"}, "GE"))
}

var wordEvalDic111 map[string]int

func isSolvable(words []string, result string) bool {
	charMap := map[byte]uint{}
	charOrder := []byte{}
	firstDigits := map[byte]bool{}
	wordEvalDic111 = map[string]int{}
	makeCharMap(append(words, result), &charMap, &charOrder, &firstDigits)
	log.Println(charMap)
	res := true
	for res {
		if evaluate(words, result, charMap) {
			printMap(charMap)
			return true
		}
		res = plusOne(&charMap, charOrder, firstDigits)
	}
	return res
}

func makeCharMap(arr []string, charMap *map[byte]uint, charOrder *[]byte, fd *map[byte]bool) {
	idx := uint(0)
	zeroCan := false
	for _, w := range arr {
		for i := 0; i < len(w); i++ {
			if _, ok := (*charMap)[w[i]]; ok {
				// do nothing
			} else {
				if i == 0 {
					if idx == 0 {
						idx++
						zeroCan = true
					}
					(*charMap)[w[i]] = idx // no leading zeroes
					(*fd)[w[i]] = true
				} else {
					if zeroCan {
						(*charMap)[w[i]] = 0
					} else {
						(*charMap)[w[i]] = idx
					}

					(*fd)[w[i]] = false
				}
				*charOrder = append(*charOrder, w[i])
			}
		}
	}
}

func plusOne(charMap *map[byte]uint, charOrder []byte, fd map[byte]bool) bool {

	i := len(charOrder) - 1
	for i >= 0 {
		(*charMap)[charOrder[i]]++
		if (*charMap)[charOrder[i]] > 9 { // carry
			if fd[charOrder[i]] {
				(*charMap)[charOrder[i]] = 1
			} else {
				(*charMap)[charOrder[i]] = 0
			}
			i--
		} else {
			return true
		}
	}
	return false
}

func evaluate(arr []string, res string, charMap map[byte]uint) bool {
	sum := 0
	for _, w := range arr {
		wc := ""
		for i := 0; i < len(w); i++ {
			wc = wc + strconv.Itoa(int(charMap[w[i]]))
		}
		intInt, _ := strconv.Atoi(wc)
		sum = sum + intInt
	}
	wc := ""
	for i := 0; i < len(res); i++ {
		wc = wc + strconv.Itoa(int(charMap[res[i]]))
	}
	intInt, _ := strconv.Atoi(wc)
	if sum == intInt {
		return true
	}
	return false
}

func printMap(chm map[byte]uint) {
	for k := range chm {
		log.Println(string(k), chm[k])
	}
}
