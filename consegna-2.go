/**
Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti in un ristorante. 10 clienti 
ordinano contemporaneamente i loro piatti. In cucina vengono preparati in un massimo di 3 alla volta, 
essendoci solo 3 fornelli. Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi. Dopo che 
un piatto viene preparato, viene portato fuori da un cameriere, che impiega 3 secondi a portarlo fuori. Ci 
sono solamente 2 camerieri nel ristorante.
● Creare la strutture Piatto e Cameriere col relativo campo “nome”.
● Creare le funzioni ordina che aggiunge il piatto a un buffer di piatti da fare; creare la function cucina che 
cucina ogni piatto e lo mette in lista per essere consegnato; creare la function consegna che fa uscire 
un piatto dalla cucina.
● Ogni cameriere può portare solo un piatto alla volta.
● Usate buffered channels per svolgere il compito.
● Attenzione: se per cucinare un piatto lo mandate nel buffer fornello di capienza 3 e lo ritirate dopo 3 
secondi, non è detto che ritiriate lo stesso piatto che avete messo sul fornello. Tenetelo in memoria. 
Ovviamente la vostra soluzione potrebbe differire dalla mia e questo hint potrebbe non servirvi.
*/

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

