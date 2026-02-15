package main 

import (
	"fmt"
	"os"
	"io"
)

func main() {

	//createFile()
	//createAndWriteFile()
	//osWriteFile()
	//readFromFile()
	//osReadFile()
	//removeFile()
	//copyFile()
	//createFolder()
	createNestedFolder()
}
func createNestedFolder(){
	err := os.MkdirAll("a/b/c/d",0777)
	if err != nil {
		panic(err)
	}
}

func createFolder(){

	err := os.Mkdir("images",0777)
	if err != nil {
		panic(err)	
	}
}

func copyFile(){
	
	src, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	dst, err := os.Create("file2.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(dst,src)
}

func removeFile(){
	err := os.Remove("myfile.txt")
	if err != nil {
		panic(err)
	}
}

func osReadFile(){
	bs,err := os.ReadFile("file.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(bs))
}


func readFromFile(){
	file,err := os.Open("myfile2.txt")
	if err!=nil {
		panic(err)
	}
	defer file.Close()

	buff,err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buff))
}

func osWriteFile(){
	err := os.WriteFile("myfile2.txt",[]byte("Hello World"),0777)
	if err != nil {
		panic(err)
	}
}

func createAndWriteFile(){
	file,err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("Hello World\n"))
	file.WriteString("A new line\n")
	io.WriteString(file,"Third line\n")
	fmt.Fprint(file,"Fourth line", 1234)
}


func createFile(){
	file,err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}