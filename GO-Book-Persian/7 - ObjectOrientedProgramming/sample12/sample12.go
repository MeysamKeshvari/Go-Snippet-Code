package main

import "fmt"

type singerInfo struct {
	name string
	age  int
}

func main() {
	sigers := make([]singerInfo, 0)

	sigers = append(sigers, singerInfo{name: "Ebi"})
	sigers = append(sigers, singerInfo{name: "Googoosh"})
	sigers = append(sigers, singerInfo{name: "Siavash"})

	for _, siger := range sigers {
		fmt.Println(siger.name)
	}

}