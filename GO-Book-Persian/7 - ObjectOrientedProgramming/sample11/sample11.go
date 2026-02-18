package main

import "fmt"

type singerInfo struct {
	name string
	age  int
}
type album struct {
	singerInfo
	name string
}

func main() {
	myAlbum := album{}

	myAlbum.name = "Jane Javani"
	myAlbum.age = 80
	myAlbum.singerInfo.name = "Ebi"

	fmt.Println("Album Name:", myAlbum.name)
	fmt.Println("Singer Name:", myAlbum.singerInfo.name)
	fmt.Println("Album Name:", myAlbum.age)
	

}