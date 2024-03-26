package mul_test

import (
	"testing"

	structMultiplicacion "pruebaDeTest.com/multiplicacion"
)

func TestMultiplicador_Multiplicar(t *testing.T) {
	Multiplicador := structMultiplicacion.Multiplicador{}
	numeros := []int{1, 8, 5, 2, 9}
	esperado := 720
	resultado := Multiplicador.Multiplicar(numeros)
	if resultado != esperado {
		t.Errorf("La multiplicacion de %v no dio el resultado esperado, se obtuvo: %d, se esperaba: %d.", numeros, resultado, esperado)
	}
}
