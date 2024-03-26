package mul

// := se utiliza para crear una variable y asignarle el valor sin tener que definirla ni indicar su tipo

type Multiplicador struct {
	Resultado int //el nombre de la variable tiene que comenzar en mayuscula o sino este se vuelve privado
}

func (s *Multiplicador) Multiplicar(numeros []int) {
	var mult *int = &(s.Resultado)
	*mult = 1
	for _, num := range numeros {
		*mult = (*mult) * num
	}
}
