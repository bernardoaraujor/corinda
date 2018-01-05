package main

import (
	"fmt"
)

func main() {

	// inicializa os arrays
	var frutas = []string{"banana", "maçã", "uva"}
	var nums = []string{"1", "2", "3", "4"}
	var nomes = []string{"wesley", "rodrigo", "igor", "regina", "marcos", "bernardo"}

	arrays := [][]string{frutas, nums, nomes}

	palpites := palpitesOrdenados(arrays)
	for palpite := range palpites{
		fmt.Println(palpite)
	}
}

func palpitesOrdenados(arrays [][]string) chan string{
	saida := make(chan string)

	go func(){
		nMax := 0
		tamanhos := make([]int, 0)
		for _, array := range arrays{
			nMax += len(array) - 1
			tamanhos = append(tamanhos, len(array))
		}
		nMax++

		// inicializa canal com arrays do produto cartesiano dos indices
		produtoCartesiano := produtoCartesiano(tamanhos)

		// inicializa saida
		indicesOrdenados := make([][][]int, nMax)

		for elemento := range produtoCartesiano{
			soma := 0
			for _, i := range elemento{
				soma += i
			}

			indicesOrdenados[soma] = append(indicesOrdenados[soma], elemento)
		}

		for n := 0; n < nMax; n++{
			l := len(indicesOrdenados[n])

			for k := 0; k < l; k++{
				indices := indicesOrdenados[n][k]

				palpite := ""
				for i, indice := range indices{
					palpite += arrays[i][indice]
				}

				saida <- palpite
			}
		}
	}()

	return saida
}

// gera canal com arrays do produto cartesiano dos indices
func produtoCartesiano(tamanhos []int) chan []int{
	// inicializa canal de saída
	saida := make(chan []int)

	// lança gorrotina recursiva
	go recursao(tamanhos, nil, 0, saida)

	// retorna canal de saída
	return saida
}

// lança os valores do produto cartesiano dos arrays no canal de saída
func recursao(tamanhos []int, contadores []int, profundidade int, saida chan []int){
	// profundidade máxima a ser processada recursivamente
	pMax := len(tamanhos)

	// primeiro nível de recursão... inicializa contadores e tamanhos
	if profundidade == 0{

		// inicializa contadores (todos 0)
		contadores = make([]int, pMax)
	}

	// último nível de profundidade na recursão
	if profundidade == pMax {

		//inicializa array resultado
		resultado := make([]int, 0)
		for p := 0; p < pMax; p++{
			i := contadores[p]
			resultado = append(resultado, i)
		}

		// envia elemento no canal de saída
		saida <- resultado

		// qualquer outra profundidade que não seja a última
	}else{
		// varre array da profundidade atual
		for contadores[profundidade] = 0; contadores[profundidade] < tamanhos[profundidade]; contadores[profundidade]++{
			// processa próxima profundidade recursivamente
			recursao(tamanhos, contadores, profundidade+1, saida)
		}
	}

	// hora de fechar canal?
	// analisa contadores de todos arrays... se todos forem iguais aos respectivos tamanhos,
	// então todos elementos do produto cartesiano foram calculados, e o canal pode ser fechado
	fecha := true
	for i := 0; i < pMax; i++{
		if contadores[i] != tamanhos[i]{
			fecha = false
		}
	}

	// fecha canal
	if fecha{
		close(saida)
	}
}