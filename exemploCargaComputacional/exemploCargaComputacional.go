package main

import (
	"sync"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// inicializa canais geradores
	chA := gerador("a")
	chB := gerador("b")
	inC := gerador("c")

	// delega o fechamento dos canais à função main
	defer close(chA)
	defer close(chB)
	defer close(inC)

	// inicializa contadores
	a := 0
	b := 0
	c := 0

	// k iterações mestras para cálculos de hashes
	k := 100
	for i := 0; i < k; i++{

		//gera canais de processamento
		hashA := hash(chA, 1)
		hashB := hash(chB, 10)
		hashC := hash(inC, 100)

		// gera canal funil para drenar canais de processamento
		funil := funil(hashA, hashB, hashC)

		// drena o canal funil
		for s := range funil{
			switch s{
			case "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb":
				a++
			case "3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d":
				b++
			case "2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6":
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
	out := make(chan string)
	go func(){
		for {
			out <- s
		}
	}()

	return out
}

// drena o canal in por n iterações
func hash(in chan string, n int) chan string{
	out := make(chan string)

	go func(n int, out chan string){
		defer close(out)
		for i := 0; i < n; i++ {
			s := <- in

			hasher := sha256.New()
			hasher.Write([]byte(s))
			hb := hasher.Sum(nil)
			h := hex.EncodeToString(hb)

			out <- h
		}
	}(n, out)


	return out
}

// funde o fluxo dos canais cs no canal out
func funil(cs ...chan string) chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
