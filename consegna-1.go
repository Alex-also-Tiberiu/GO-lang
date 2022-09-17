package main

import (
	"fmt"
)


type Cliente struct {
	nome string
	presenza string
}

/***
spagna  = 0  / partecipanti = 4
francia = 1  / partecipanti = 2
*/
type Viaggio struct {
	meta string
	letsgo int
	clienti [7]Cliente
}

func stampaPartecipanti(viaggi [2]Viaggio){
	if viaggi[0].letsgo >= 4  {
		for i:=0;i<7;i++ {
			if viaggi[0].clienti[i].presenza == "spagna" {
				fmt.Println("cliente ",viaggi[0].clienti[i].nome ,"ha prenotato per la spagna")
			}
		}
	}else {
		fmt.Println("non ci sono abbastanza iscritti per la spagna")
	}
	if viaggi[1].letsgo >= 2  {
		for i:=0;i<7;i++ {
			if viaggi[0].clienti[i].presenza == "francia" {
				fmt.Println("cliente ",viaggi[1].clienti[i].nome ,"ha prenotato per la francia")
			}
		}
	}else {
		fmt.Println("non ci sono abbastanza iscritti per la francia")
	}
}

func prenota(c Cliente,viaggi [2]Viaggio) {

	for i:=0; i < 7;i++ {
		if c.presenza == "spagna" {
			if viaggi[0].clienti[i] == c {
				viaggi[0].clienti[i].presenza = "spagna"
			}
		} else {
			if viaggi[1].clienti[i] == c {
				viaggi[1].clienti[i].presenza = "francia"
			}
		}
	}

}

func main(){

	var viaggi[2]Viaggio
	var clienti[7]Cliente
	viaggi[0].meta = "spagna"
	viaggi[1].meta = "francia"
	clienti[0].nome = "Giovanni"
	clienti[1].nome = "Michele"
	clienti[2].nome = "Giacomo"
	clienti[3].nome = "Samuele"
	clienti[4].nome = "Silvia"
	clienti[5].nome = "Stefania"
	clienti[6].nome = "Daniela"

	channel1 := make(chan bool)
	channel2 := make(chan bool)

	for i:=0; i<7; i++ {
		go func() {
			channel1 <- true
		}()
		go func() {
			channel2 <- true
		}()
		select {
		case _ = <-channel1:
			viaggi[0].letsgo++
			clienti[i].presenza = "spagna"
			prenota(clienti[i],viaggi)
		case _ = <-channel2:
			viaggi[1].letsgo++
			clienti[i].presenza = "francia"
			prenota(clienti[i],viaggi)
		}
	}

	viaggi[0].clienti = clienti
	viaggi[1].clienti = clienti


	defer stampaPartecipanti(viaggi)

}