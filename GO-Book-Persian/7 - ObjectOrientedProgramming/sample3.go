package main 

import "fmt"
 
func main() {
	r := rectangle {
		width:3,
		height:7,
	}
	c:= circle {
		radius:7,
	}
	printRectangleArea(r)
	printCircleArea(c)
}

type rectangle struct { 
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return c.radius * c.radius * 3.14
}

func printRectangleArea(shape rectangle){
	fmt.Println("Area is:", shape.area())
}

func printCircleArea(shape circle){
	fmt.Println("Area is:", shape.area())
}




