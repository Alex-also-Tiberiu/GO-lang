package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Cameriere struct {
	nome string
}

type Piatto struct {
	nome string
}

/**
	ordina un piatto e lo mette in attesa al buffer
 */
func ordina(ordini chan Piatto,p Piatto){
	fmt.Println(p.nome,"ordinato")
	ordini <- p
}

/**
	cucina ogni piatto in attesa
 */
func cucina(ordini chan Piatto,pronti chan Piatto){
	for {
		p := <-ordini
		time.Sleep(time.Duration(rand.Intn(3)+4) * time.Second)
		fmt.Println(p.nome, "cucinato")
		pronti <- p
	}
}



/**
	fa consegnare un piatto dei due camerieri
 */
func consegna(cameriere Cameriere,pronti chan Piatto,wg *sync.WaitGroup) {
	for {
		p := <-pronti
		time.Sleep(time.Second * 3)
		fmt.Println(cameriere.nome," ha consegnato", p.nome)
		wg.Done()
	}
}

func stampaTempo(){
	t := 0
	time.Sleep(100*time.Millisecond)
	for{
		t = t + 1
		fmt.Println(t)
		time.Sleep(time.Second)

	}
}

func main() {

	var ordini = make(chan Piatto,10)
	var pronti = make(chan Piatto,3)
	cameriere1 := Cameriere{nome : "Stefano"}
	cameriere2 := Cameriere{nome : "Andrea"}
	var piatti[10]Piatto
	var wg sync.WaitGroup
	wg.Add(10)

	piatti[0].nome = "spaghetti"
	piatti[1].nome = "riso"
	piatti[2].nome = "bistecca"
	piatti[3].nome = "salsicce"
	piatti[4].nome = "pancetta"
	piatti[5].nome = "ravioli"
	piatti[6].nome = "tortellini"
	piatti[7].nome = "tiramisu"
	piatti[8].nome = "pollo"
	piatti[9].nome = "caviale"

	for i:=0;i<10;i++ {
		ordina(ordini,piatti[rand.Intn(10)])
	}

	for i:=0;i<3;i++{
		go cucina(ordini,pronti)
	}

	go consegna(cameriere1,pronti,&wg)
	go consegna(cameriere2,pronti,&wg)

	go stampaTempo()



	wg.Wait()
}

