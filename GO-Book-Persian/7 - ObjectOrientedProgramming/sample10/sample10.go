package main

import "fmt"

type singerInfo struct {
	name string
}
type album struct {
	snger singerInfo
	name  string
}

func main() {

	myAlbum := album{}
	myAlbum.name = "Jane Javani"
	myAlbum.snger.name = "Ebi"
	fmt.Println(myAlbum)
}