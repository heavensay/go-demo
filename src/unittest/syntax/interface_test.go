package syntax

import (
	"fmt"
	"testing"
)

type Animal interface{
	say()
}



type Bird struct{

}

func (bird Bird) say(){
	fmt.Println("i am a bird")
}


type Lion struct{

}

func (lion Lion) say(){
	fmt.Println("i am a lion")
}


type Cat struct{}

func (cat *Cat) say(){
	fmt.Println("i am a cat")
}

func Test1(t *testing.T){
	var animal Animal
	animal = new(Bird)
	animal.say()
	fmt.Println(animal)

	animal = Bird{}
	animal.say()
	fmt.Println(animal)

	animal = new (Lion)
	animal.say()
	fmt.Println(animal)

	animal = new(Cat)
	animal.say()
	fmt.Println(animal)

	animal = &Cat{}
	animal.say()
	fmt.Println(animal)

	//animal = Cat{} //编译不通过 Cannot use Cat{} (type Cat) as type Animal in assignment
}