package main
import "testing"

func TestSoma(t *testing.T){
	total := soma(10,10)

	if total != 20{
		t.Errorf("Resultado da soma é inválido: Resultado %d. Valor esperado: %d", total, 20)
	}
}