package main

import (
	
	"log"
)

func main() {
  list1 := [][]string{}
  list2 := []string{"apple", "orange"}
  log.Println(list1, list2)
  list1 = [][]string{append(list2,  "tomato")}
  log.Println(list1, list2)
//  list2 = list2[0 : len(list2)-1]
  list2 = append(list2, "plum")
  log.Println(list1, list2)
}
