package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)



var mrtl Martello
var ccvt Cacciavite
var trpn[2] Trapano
var mutexMrt sync.Mutex
var mutexCvt sync.Mutex
var mutexTrp1 sync.Mutex
var mutexTrp2 sync.Mutex
var wg sync.WaitGroup



type Operaio struct {
	nome string
	mrtBusy bool
	trpBusy bool
}

type Martello struct {
	preso bool
}

type Cacciavite struct {
	preso bool
}

type Trapano struct {
	preso bool
}


func schedLav(lav chan Operaio){
	for i:=0;i<3;i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			lavora(<-lav)
		}()
	}
}


func lavora(op Operaio){

	var t sync.WaitGroup
	t.Add(1)


	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		usaMartello(&op)
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		usaTrapano(&op,&t)
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		usaCacciavite(&op,&t)
	}()

}


func usaMartello(operaio *Operaio){

	for operaio.trpBusy == true {}

	mutexMrt.Lock()
	mrtl.preso = true
	operaio.mrtBusy = true
	fmt.Println(operaio.nome,"sta usando il martello")
	mrtl.preso = false
	fmt.Println(operaio.nome,"non sta piu usando il martello")
	mutexMrt.Unlock()
	operaio.mrtBusy = false

	wg.Done()

}


func usaTrapano(operaio *Operaio,t *sync.WaitGroup){

	for operaio.mrtBusy == true {}

	if trpn[1].preso == false {
		mutexTrp1.Lock()
		trpn[1].preso = true
		operaio.trpBusy = true
		fmt.Println(operaio.nome,"sta usando il trapano 1")
		trpn[1].preso = false
		fmt.Println(operaio.nome,"non sta piu usando il trapano 1")
		mutexTrp1.Unlock()
		operaio.trpBusy = false
		t.Done()
	}else {
		mutexTrp2.Lock()
		trpn[0].preso = true
		operaio.trpBusy = true
		fmt.Println(operaio.nome,"sta usando il trapano 2")
		trpn[0].preso = false
		fmt.Println(operaio.nome,"non sta piu usando il trapano 2")
		mutexTrp2.Unlock()
		operaio.trpBusy = false
		t.Done()
	}

}


func usaCacciavite(operaio *Operaio,t *sync.WaitGroup){

	t.Wait()

	mutexCvt.Lock()
	ccvt.preso = true
	fmt.Println(operaio.nome,"sta usando il cacciavite")
	ccvt.preso = false
	fmt.Println(operaio.nome,"non sta piu usando il cacciavite")
	mutexCvt.Unlock()

	wg.Done()
}




func main(){

	operai := make(chan Operaio,3)
	for i:=1; i<4; i++ {
		operai <- Operaio{nome: fmt.Sprint("Operaio ",i),mrtBusy: false,trpBusy: false}
	}

	wg.Add(6)
	schedLav(operai)
	wg.Wait()
}
