
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
	for {
		<-controle
		rand.Seed(time.Now().UnixNano())
		if(rand.Intn(1000)%2==0){
			fmt.Printf("Jogador %s acertou a bola\n.", nome)
			//pointsP1=pointsP1+1
			//fmt.Printf("===*** Messi: %d vs  %d Cristiano ", pointsP1, pointsP2)
			controle <- 1
		}else{

			fmt.Printf("Jogador %s errou a bola\n.", nome)
			if nome == "Messi" {
				pointsP2=pointsP2+1
				fmt.Printf("Ponto para o jogador Cristiano, agora com %d pontos!", pointsP2)
				fmt.Printf("\n \n")
				fmt.Printf("O jogo está: ---- Messi: %d vs %d Cristiano", pointsP1, pointsP2)

			} else {
				pointsP1=pointsP1+1
				fmt.Printf("Ponto para o jogador Messi, totalizando %d pontos!", pointsP2)
				fmt.Printf("\n \n")
				fmt.Printf("O jogo está: ---- Messi: %d vs %d Cristiano", pointsP1, pointsP2)
			}

		}
		if pointsP1 == 5 || pointsP2 == 5 {
			fmt.Printf("O jogo acaboou!! o jogador")
			if (pointsP1 == 5) {
				fmt.Print("O jogador Messi venceu!!")
			} else {
				fmt.Printf("O jogador CR7 venceu!!")
			}
			close(controle)
			return
		} 
		
	}
}
/*
func player2(controle chan int) {
	for {
		<-controle
		rand.Seed(time.Now().UnixNano())
		if(rand.Intn(1000)%2==0){
			controle <- 1
			fmt.Printf("Jogador Neymar acertou a bola.\n")
			
		}else{
			pointsP1=pointsP1+1
			if pointsP1 == 5 {
				fmt.Printf("Jogador Neymar venceu a partida. \n")
				waitGroup.Done()
				//close(controle)
				return
			}
			controle <- 1
		}
		//fmt.Printf("O placar está: %d para o jogador Neymar contra %d do jogador Messi \n.", pointsP1, pointsP2)

	}
}
*/


func main() {
	controle := make(chan int)
	waitGroup.Add(1)
	go player(controle, "Messi")
	go player(controle,"Cristiano")
	controle <- 1
	waitGroup.Wait()
	

}