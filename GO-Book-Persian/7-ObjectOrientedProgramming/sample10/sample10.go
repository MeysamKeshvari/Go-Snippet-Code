package main

import "fmt"

type singerInfo struct {
	name string
}
type album struct {
	singer singerInfo
	name  string
}

func main() {

	myAlbum := album{}
	myAlbum.name = "Jane Javani"
	myAlbum.singer.name = "Ebi"
	fmt.Println(myAlbum)
}