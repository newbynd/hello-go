package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var x, y, z int = 1, 2, 3                //inician como 0 si no se define valor
var bulean bool                          //inician como false si no se define valor
var c, python, java = true, false, "no!" //se puden juntar varias si no se define el tipo

var (
	ToBe    bool       = false
	MaxInt  uint64     = 1<<64 - 1
	Complex complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14 //Las constantes pueden ser caracteres, cadenas de caracteres, Booleanos, o valores numéricos.
const World = "世界"
const Truth = true

const (
	Big   = 1 << 100  // una veces por 2, 100 veces -> 2 elevado a la 100
	Small = Big >> 99 // big veces entre 2, 99 veces -> 2 elevado a la 100 entre 2 elevado a la 99
)

//struct es una coleccion de campos, y type declara el tipo
type Vertex struct {
	X int
	Y int
}

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

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	basicTypes(ToBe, MaxInt, Complex)

	fmt.Println(Vertex{1, 2})
	fmt.Printf("Tipo: %T\n", Vertex{1, 2})
	ver := Vertex{1, 2}
	ver.X = 4
	fmt.Println(ver)

	ver2 := &ver //cualquier cambio que se haga a ver2, se aplicara a ver, ya que son punteros....
	ver2.X = 1e9
	fmt.Println(ver, ver2)
	ver.X = 2 //.... y viceversa
	fmt.Println(ver, ver2)

	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex
		r = Vertex{X: 1}  // Y:0 is implicit
		s = Vertex{}      // X:0 and Y:0
	)
	fmt.Println(p, q, r, s)

	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	maps()
	literalesMaps()
	moddingMaps()
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

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// las variables declaradas en if solo valen para este condicional, fuera no se pueden usar
	//fmt.Println(v, "hola") undefined si descomentamos
	return lim
}

func basicTypes(ToBe bool, MaxInt uint64, a complex128) {
	const f = "%T(%v)\n"
	// %T substituye por el tipo de variable
	// %v sustiuye por el valor
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, a, a)
}

func maps() {
	type empresas struct {
		string1, string2 string
	}

	//var m map[string]empresas

	/* Un map asocia un valor a cada llave.
	** Antes de usar un map, hay que crearlo con make (en vez de con new); y el map nil por definición está vacío y no se le puede añadir nada.
	 */
	m := make(map[string]empresas)
	m["Taco Bell"] = empresas{
		"Producto", "Tacos",
	}
	m["Mediamarkt"] = empresas{
		"Producto", "Electronica",
	}
	fmt.Println(m)
}

func literalesMaps() {

	// Los literales de map son como los de las estructuras, pero las llaves son necesarias (en el ejemplo, "Bell Labs" y "Google").
	type Vertex struct {
		Lat, Long float64
	}

	m := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func moddingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// Si la llave key está en m, ok será true, Si no, ok será false y elem será el valor cero del tipo de los elementos.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
