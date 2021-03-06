/*
Exemplo de como o padrão de concorrência funil, associado à utilização de iterações mestras
com gorrotinas de vida limitada, podem ser utilizados para distribuir a carga computacional
ao processamento de diferentes canais geradores
*/
package main

import (
	"sync"
	"encoding/hex"
	"fmt"
	"crypto/sha1"
)

func main() {
	// inicializa canais geradores
	chA := gerador("a")
	chB := gerador("b")
	chC := gerador("c")

	// delega o fechamento dos canais à função main
	defer close(chA)
	defer close(chB)
	defer close(chC)

	// inicializa contadores
	a := 0
	b := 0
	c := 0

	// gera array de canais geradores
	geradores := []chan string{chA, chB, chC}

	//gera array de indices
	ns := []int{1, 10, 100}

	// inicializa canal de lotes de hashes
	lotes := loteHashes(geradores, ns)

	// k lotes de hashes
	k := 100
	for i := 0; i < k; i++{

		// recebe lote do canal lotes
		lote := <-lotes

		// processa lote
		for _, byte := range lote{
			// converte o hash de hex para string
			s := hex.EncodeToString(byte)

			// incrementa contador correspondente ao valor lido
			switch s{
			case "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8":
				a++
			case "e9d71f5ee7c92d6dc9e92ffdad17b8bd49418f98":
				b++
			case "84a516841ba77a5b4648de2cd0dfcb30ea46dbb4":
				c++
			}
		}
	}

	// imprime os resultados
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// gera canal com fluxo continuo de strings identicas a s
func gerador(s string) chan string{

	// inicializa canal de saída
	out := make(chan string)

	// lança gorrotina que envia cópias de s ao canal de saída indefinidamente, até que este canal seja fechado
	go func(){
		for {
			out <- s
		}
	}()

	// retorna o canal de saída
	return out
}

// drena o canal in por n iterações
func hash(entrada chan string, n int) chan []byte{

	// inicializa canal de saída
	out := make(chan []byte)

	// lança gorrotina que repete por n iterações
	// o fechamento do canal de saída é delegado ao encerramento da gorrotina (defer), o que acontece após a execução das n iterações
	go func(n int, saida chan []byte){
		defer close(saida)
		for i := 0; i < n; i++ {
			// lê o canal de entrada
			s := <-entrada

			// calcula o hash
			hasher := sha1.New()
			hasher.Write([]byte(s))
			hashBytes := hasher.Sum(nil)

			// envia o hash no canal de saída
			saida <- hashBytes
		}
	}(n, out)

	// retorna o canal de saída
	return out
}

// funde o fluxo dos canais cs no canal out
func funil(cs []chan []byte) chan []byte {

	// declara o grupo de espera wg
	var wg sync.WaitGroup

	// inicializa canal de saída
	saida := make(chan []byte)

	// inicializa gorrotina de saída para cada canal de entrada em cs.
	// a gorrotina é responsável por enviar para a saída cópias dos valores drenados de c até que c seja fechado,
	// até por fim chamar wg.Done
	output := func(c <-chan []byte) {
		for n := range c {
			saida <- n
		}
		wg.Done()
	}

	// prepara o grupo de espera para o número de gorrotinas a serem lançadas
	wg.Add(len(cs))

	// lança gorrotinas
	for _, c := range cs {
		go output(c)
	}

	// lança gorrotina para fechar o canal de saída uma vez que todas gorrotinas de saídas estão finalizadas
	go func() {
		wg.Wait()
		close(saida)
	}()

	// retorna canal de saída
	return saida
}

// calcula lote de hashes a partir de um array de canais de entrada
func loteHashes(entradas []chan string, ns []int) chan [][]byte{
	
	// inicializa canal de saida
	saida := make(chan [][]byte)
	
	go func(entradas []chan string, ns []int){

		// processa lotes indefinidamente
		for {
			// gera array de canais de calculo de hashes, a serem utilizados como entrada para o funil
			hashes := make([]chan []byte, 0)
			for i, entrada := range entradas{
				n := ns[i]
				hashes = append(hashes, hash(entrada, n))
			}

			// gera canal funil para drenar canais de processamento
			funil := funil(hashes)

			// inicializa lote de bytes
			lote := make([][]byte, 0)

			// drena canal funil
			for byte := range funil{
				lote = append(lote, byte)
			}

			// envia lote no canal de saída
			saida <- lote
		}
	}(entradas, ns)
	
	return saida
}