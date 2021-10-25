package main

import "fmt"

func main() {
	/*
		Se debe realizar un ejercicio el cual permita hacer resta sin restar
		sumar sin sumar dividir sin dividir y multiplicar sin multiplicar tambien potenciar.
		Es decir realizar suma sin el signo de sumar implicito, y así susesivamente con cada operación.
	*/
	inicio := "Este programa cálcula 2 números. Resta, Suma, Multiplica, Divide, Potencia y obtiene Raiz "
	inserta_numero := "Por favor ingresa número"
	var num_1, num_2 int
	fmt.Println(inicio)
	fmt.Println(inserta_numero, "uno")
	fmt.Scan(&num_1)
	fmt.Println(inserta_numero, "dos")
	fmt.Scan(&num_2)
	fmt.Println("Número uno", num_1, "Número dos", num_2)

	fmt.Println("Está es la suma", suma(num_1, num_2))
	fmt.Println("Está es la Resta", resta(num_1, num_2))
	fmt.Println("Está es la Multiplicación", multiplicar(num_1, num_2))

}

func suma(a, b int) (result int) {
	c := 0
	if a > b {
		c = -a - b
		return -c
	} else {
		b = -b
		c = b - a
		return -c
	}
}

func resta(a, b int) (result int) {
	if a > b {
		b = -b
		return b + a
	} else {
		a = -a
		return a + b
	}
}
func multiplicar(a, b int) (result int) {
	var cantidad, resultado int
	for i := 1; i <= a; i++ {
		cantidad = suma(b, b)
		resultado = resultado + cantidad
	}
	return resultado / 2
}
