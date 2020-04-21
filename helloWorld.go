package main

import (
	"fmt"
	"math"
)

var x, y, z int = 1, 2, 3                //inician como 0 si no se define valor
var bulean bool                          //inician como false si no se define valor
var c, python, java = true, false, "no!" //se puden juntar varias si no se define el tipo

const Pi = 3.14 //Las constantes pueden ser caracteres, cadenas de caracteres, Booleanos, o valores numéricos.
const World = "世界"
const Truth = true

const (
	Big   = 1 << 100  // una veces por 2, 100 veces -> 2 elevado a la 100
	Small = Big >> 99 // big veces entre 2, 99 veces -> 2 elevado a la 100 entre 2 elevado a la 99
)

func main() {
	helloWorld()
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(add(32, 34))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(10))

	h := "variable sin asignar dentro de una funcion" //Fuera de las funciones, todas las sentencias empiezan con una palabra reservada y la asignación especial := no existe
	fmt.Println(x, y, z, c, python, java, bulean, h)

	fmt.Println("Hello", World, "Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	bucleFor(10)
	bucleForLight()

}

func helloWorld() {
	fmt.Println("Hello World!")
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func bucleFor(max int) {
	sum := 0
	for i := 0; i < max; i++ {
		sum += i
	}
	fmt.Println("La suma es", sum)
}

func bucleForLight() {
	sum := 1
	for sum < 1000 { //podemos dejar vacias sentencias pre y post
		sum += sum
	}
	fmt.Println("La suma es", sum)
}
