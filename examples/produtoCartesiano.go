/*
Exemplo de utilização de gorrotina recursiva para gerar fluxo com elementos do produto cartesiano de diversos arrays
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// inicializa as listas de tokens
	var frutas = []string{"banana", "maçã", "uva"}
	var nums = []string{"1", "2", "3", "4"}
	var nomes = []string{"wesley", "rodrigo", "igor", "regina", "marcos"}

	fmt.Println("-------------------frutas x nums-------------------")
	total := 0

	// drena canal com elementos do produto cartesiano dos arrays em arrays1
	for s := range produtoCartesiano(frutas, nums){
		total++
		fmt.Println(s)
	}
	fmt.Println("|frutas x nums| = " + strconv.Itoa(total))


	fmt.Println("---------------frutas x nums x nomes---------------")
	total = 0

	// drena canal com elementos do produto cartesiano dos arrays em arrays2
	for s := range produtoCartesiano(frutas, nums, nomes){
		total++
		fmt.Println(s)
	}
	fmt.Println("|frutas x nums x nomes| = " + strconv.Itoa(total))
}


// gera canal com fluxo de elementos do produto cartesiano dos arrays
func produtoCartesiano(arrays ...[]string) chan string{
	// inicializa canal de saída
	saida := make(chan string)

	// lança gorrotina recursiva
	go recursao(arrays, 0, nil, nil, saida)

	// retorna canal de saída
	return saida
}


// lança os valores do produto cartesiano dos arrays no canal de saída
func recursao(arrays [][]string, profundidade int, contadores []int, tamanhos []int, saida chan string){
	// profundidade máxima a ser processada recursivamente
	n := len(arrays)

	// primeiro nível de recursão... inicializa contadores e tamanhos
	if profundidade == 0{

		// inicializa contadores (todos 0)
		contadores = make([]int, n)

		// inicializa tamanhos
		tamanhos = make([]int, n)
		for i, _ := range arrays{
			tamanhos[i] = len(arrays[i])
		}
	}

	// último nível de profundidade na recursão
	if profundidade == n{
		resultado := ""
		for p := 0; p < n; p++{
			i := contadores[p]
			resultado += arrays[p][i]
		}

		// envia elemento no canal de saída
		saida <- resultado

	// qualquer outra profundidade que não seja a última
	}else{
		// varre array da profundidade atual
		for contadores[profundidade] = 0; contadores[profundidade] < tamanhos[profundidade]; contadores[profundidade]++{
			// processa próxima profundidade recursivamente
			recursao(arrays, profundidade+1, contadores, tamanhos, saida)
		}
	}

	// hora de fechar canal?
	// analisa contadores de todos arrays... se todos forem iguais aos respectivos tamanhos,
	// então todos elementos do produto cartesiano foram calculados, e o canal pode ser fechado
	fecha := true
	for i := 0; i < n; i++{
		if contadores[i] != tamanhos[i]{
			fecha = false
		}
	}

	// fecha canal
	if fecha{
		close(saida)
	}
}
