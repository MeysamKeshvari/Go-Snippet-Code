package main 

type Mover interface{
	Run()
	Walk()
}
type Eater interface {
	Eat()
}

type Animal struct{

}

func (a Animal) Run(){

}

func (a Animal) Walk(){

}

func (a animal) Eat(){

}