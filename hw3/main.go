package main

import (
	"fmt"
	"math"
)

const (
	applePrice = 5.99
	pearPrice  = 7.
	yourMoney  = 23.
)

func main() {
	appleCount := 9
	pearCount := 8
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити %d яблук та %d груш?\n", appleCount, pearCount)
	applesCost := float64(appleCount) * applePrice
	pearCost := float64(pearCount * pearPrice)
	fmt.Printf("\t- Вони коштують %vгрн.\n", applesCost+pearCost)

	fmt.Println("2. Скільки груш ми можемо купити?")
	pearsCanBuy := math.Floor(yourMoney / pearPrice)
	fmt.Printf("\t- Ми можемо купити %v.\n", pearsCanBuy)

	applesCanBuy := math.Floor(yourMoney / applePrice)
	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Printf("\t- Ми можемо купити %v.\n", applesCanBuy)

	appleCount = 2
	pearCount = 2
	fmt.Printf("4. Чи ми можемо купити %d груші та %d яблука?\n", appleCount, pearCount)
	nApplesCost := float64(appleCount) * applePrice
	nPearsCost := float64(pearCount) * pearPrice
	fmt.Printf("\t- %t\n", nApplesCost+nPearsCost <= yourMoney)

}
