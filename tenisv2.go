
package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)
var waitGroup sync.WaitGroup

var pointsP1 int32 = 0
var pointsP2 int32 = 0


func player(controle chan int, nome string) {
	defer waitGroup.Done()
	for {
		valor := <-controle
		if valor != 1 {
			return
		}

		rand.Seed(time.Now().UnixNano())
		if(rand.Intn(1000)%2==0 || rand.Intn(1000)%3==0 || rand.Intn(1000)%5==0){
			fmt.Printf("Jogador %s acertou a bola.\n", nome)

			controle <- 1
		}else{

			fmt.Printf("Jogador %s errou a bola.\n", nome)
			if nome == "Messi" {
				pointsP2=pointsP2+1
				fmt.Printf("Ponto para o jogador Cristiano, agora com %d pontos!", pointsP2)
				fmt.Printf("\n \n")
				fmt.Printf("O jogo está: ---- Messi: %d vs %d Cristiano\n", pointsP1, pointsP2)
				//controle <- 1
				if pointsP2 == 5 {
					fmt.Printf("O jogo acaboou!!\n")
					fmt.Print("O jogador Cristiano venceu!!")
					close(controle)
					return
				}

			} else {
				pointsP1=pointsP1+1
				fmt.Printf("Ponto para o jogador Messi, totalizando %d pontos!", pointsP1)
				fmt.Printf("\n \n")
				fmt.Printf("O jogo está: ---- Messi: %d vs %d Cristiano\n", pointsP1, pointsP2)
				//controle <- 1
				if pointsP1 == 5 {
					fmt.Printf("O jogo acaboou!!\n")
					fmt.Print("O jogador Messi venceu!!")
					close(controle)
					return
				}
			}
			controle <- 1
		}
		
	}
}


func main() {

	controle := make(chan int)
	waitGroup.Add(2)
	go player(controle, "Messi")
	go player(controle,"Cristiano")
	controle <- 1
	waitGroup.Wait()
	

}