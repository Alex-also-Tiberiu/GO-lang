package main

import (
    "fmt"
    "sync"
)
import "time"
import "math/rand"

var wg sync.WaitGroup

type Gruppo struct {
    nome string
    nPalline int
}

type Tunnel struct {
    libero bool
}

func transumanza(g Gruppo, t chan Tunnel, c1 chan int){
    for g.nPalline > 0{
        time.Sleep(time.Duration(rand.Intn(2))*time.Second)
        mandaPersona(&g,t,c1)
    }
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int) {

     tunnel := <-t
     if tunnel.libero == true {
         fmt.Println(g.nome, " ha lanciato")
         time.Sleep(time.Second)
         select {
         case _ = <-c1:
             fmt.Println(g.nome, " è passato")
             g.nPalline--
             fmt.Println("rimangono ", g.nPalline, " nel gruppo ", g.nome)
             tunnel.libero = false
             t <- tunnel

         default:
             fmt.Println("scontro")
             tunnel.libero = true
             t <- tunnel
             c1 <- 1
         }
     }else if g.nPalline == 1 {
         tunnel.libero = true
         t <- tunnel
         c1 <- 1
     }else {
         fmt.Println(g.nome," ha lanciato")
         tunnel.libero = true
         t <- tunnel
     }


     if g.nPalline == 0 {  wg.Done() }
}


func main() {
    wg.Add(2)

    rand.Seed(time.Now().UnixNano())
    gruppo1 := Gruppo{"destra",5}
    gruppo2 := Gruppo{"sinistra",5}

    c1 := make(chan int,1)
    c1<-1

    tunnelChannel := make(chan Tunnel,1)
    tunnelChannel <- Tunnel{libero: true}

    go transumanza(gruppo1,tunnelChannel,c1)
    go transumanza(gruppo2,tunnelChannel,c1)

	wg.Wait()
}

	
	
	
	
	
	
	
	
	
	