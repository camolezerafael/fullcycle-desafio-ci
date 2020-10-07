package main

import "testing"

func TestSoma(t *testing.T)  {
	tabelas := []struct {
		x int
		y int
		n int
	}{
		{0, 2, 2},
		{6, 8, 14},
		{3, 5, 8},
	}

	for _, tabela := range tabelas {
		total := Soma(tabela.x, tabela.y)
		if total != tabela.n {
			t.Errorf("Falha: Soma de (%d+%d) incorreta, obtido: %d, esperado: %d.", tabela.x, tabela.y, total, tabela.n)
		}
	}
}
