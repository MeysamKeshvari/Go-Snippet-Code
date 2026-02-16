package main 

type Animal struct {

}

type Mover interface {
	Run()
	Walk()
} 

type Eater interface {
	Eat()
}

func (a Animal) Run(){
}
func (a Animal) Walk(){
}

func (a Animal) Eat(){
}

func EatSomething(e Eater){
	e.Eat()
}
func Run(m Mover){
	m.Run()
}

func main(){
	lion := Animal{}
	EatSomething(lion)
	Run(lion)
}

