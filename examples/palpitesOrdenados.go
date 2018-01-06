package main

import (
	"fmt"
)

func main() {

	// inicializa os arrays
	var frutas = []string{"banana", "maçã", "uva"}
	var nums = []string{"1", "2", "3", "4"}
	var nomes = []string{"wesley", "rodrigo", "igor", "regina", "marcos"}

	arrays := [][]string{frutas, nums, nomes}

	palpites := palpitesOrdenados(arrays)

	for palpite := range palpites{
		fmt.Println(palpite)
	}
}

// função que retorna canal com palpites ordenados
func palpitesOrdenados(arrays [][]string) chan string{

	// inicializa canal de saida
	saida := make(chan string)

	// lança gorrotina que gera os palpites
	go func(){

		// n máximo a ser processado,
		nMax := 0

		// array com os tamanhos de cada array de strings
		tamanhos := make([]int, 0)

		// calculo de nMax e tamanhos
		for _, array := range arrays{
			nMax += len(array) - 1
			tamanhos = append(tamanhos, len(array))
		}
		nMax++

		// inicializa canal com arrays do produto cartesiano dos indices
		produtoCartesiano := produtoCartesiano(tamanhos)

		// inicializa indices
		indicesOrdenados := make([][][]int, nMax)

		// varre o produto cartesiano dos indices
		for indices := range produtoCartesiano{
			n := 0

			// n é a soma dos indices
			for _, i := range indices{
				n += i
			}

			indicesOrdenados[n] = append(indicesOrdenados[n], indices)
		}

		// varre os indices ordenados, gerando os palpites na ordem correta
		for n := 0; n < nMax; n++{
			l := len(indicesOrdenados[n])

			// varre os elementos do nível n de indicesOrdenados
			for k := 0; k < l; k++{
				indices := indicesOrdenados[n][k]

				// gera palpite do elemento k do nivel n de indicesOrdenados
				palpite := ""
				for i, indice := range indices{
					palpite += arrays[i][indice]
				}

				// envia palpite para a saida
				saida <- palpite
			}
		}

		// fecha o canal de saida
		close(saida)
	}()

	// retorna canal de saida
	return saida
}

// gera canal com arrays do produto cartesiano dos indices
func produtoCartesiano(tamanhos []int) chan []int{

	// inicializa canal de saída
	saida := make(chan []int)

	// lança gorrotina de recursao
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
