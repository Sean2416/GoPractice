package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string //宣告一個Slice物件

func newDeck() deck {
	card := deck{}

	cardSuits := []string{"黑桃", "紅心", "方塊", "梅花"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, item := range cardSuits {

		for _, value := range cardValues {
			card = append(card, item+value)
		}

		// for i := 1; i <= 13; i++ {
		// 	card = append(card, item+strconv.Itoa(i))
		// }
	}
	return card
}

func (d deck) print() { //Extension涵式 ,只能針對type定義出的物件(類似物件中的屬性概念)

	for _, item := range d[0:] { //從第a[0]開始印到最後
		fmt.Println(item)
	}

	// for _, item := range d[8:10] { //從第8個元素開始印到第10個元素前(A8~A9)
	// 	fmt.Println(item)
	// }

	// for _, item := range d[:10] { //從第0個元素開始印到第10個元素前(A0~A9)
	// 	fmt.Println(item)
	// }

}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())

	for i := range d {
		n := rand.New(s).Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}

func (d deck) deal(size int) (deck, deck, deck) {
	return d, d[:size], d[size:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveTofile(filename string) error {

	s := []byte(d.toString())

	return ioutil.WriteFile(filename, s, 0666)

}

func readFile(filepath string) deck {
	b, err := ioutil.ReadFile(filepath)

	if err != nil {
		println("Error: ", err.Error())
		os.Exit(-1)
	}

	return strings.Split(string(b), ",")
}
