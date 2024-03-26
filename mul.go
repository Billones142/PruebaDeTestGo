package mul

type Multiplicador struct{}

func (s *Multiplicador) Multiplicar(numeros []int) int {
	mult := 1
	for _, num := range numeros {
		mult = mult * num
	}
	return mult
}
