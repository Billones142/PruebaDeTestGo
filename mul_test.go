package mul_test

import (
	"testing"

	mul "pruebaDeTest.com/multiplicacion"
)

func TestMultiplicador_Multiplicar(t *testing.T) {
	var Multiplicador1 mul.Multiplicador
	numeros := []int{1, 8, 5, 2, 9}
	esperado := 720
	Multiplicador1.Multiplicar(numeros)
	resultado := Multiplicador1.Resultado
	if resultado != esperado {
		t.Errorf("La multiplicacion de %v no dio el resultado esperado, se obtuvo: %d, se esperaba: %d", numeros, resultado, esperado)
	}
}
