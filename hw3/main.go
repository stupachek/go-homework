package main

import (
	"fmt"
	"math"
)

const (
	priceApple = 5.99
	pricePear  = 7.
	yourMoney  = 23.
)

func main() {
	appleCount := 9
	pearCount := 8
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити %d яблук та %d груш?\n", appleCount, pearCount)
	costApples := float64(appleCount) * priceApple
	costPears := float64(pearCount * pricePear)
	fmt.Printf("\t- Вони коштують %vгрн.\n", costApples+costPears)

	fmt.Println("2. Скільки груш ми можемо купити?")
	pearsCanBuy := math.Floor(yourMoney / pricePear)
	fmt.Printf("\t- Ми можемо купити %v.\n", pearsCanBuy)

	applesCanBuy := math.Floor(yourMoney / priceApple)
	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Printf("\t- Ми можемо купити %v.\n", applesCanBuy)

	nApples := 2
	nPears := 2
	fmt.Printf("4. Чи ми можемо купити %d груші та %d яблука?\n", nApples, nPears)
	costOfnApples := float64(nApples) * priceApple
	costOfnPears := float64(nPears) * priceApple
	fmt.Printf("\t- %t\n", costOfnApples+costOfnPears <= yourMoney)

}
