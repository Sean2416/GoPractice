package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (c card) Update() {
	c.value = 33
}

func (c *card) UpdateValue() {
	(*c).value = -1
}

func newCard() cards {

	cardSuits := []string{"黑桃", "紅心", "方塊", "梅花"}
	//cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	c := cards{}

	for _, item := range cardSuits {
		for i := 1; i <= 13; i++ {
			c = append(c, card{
				suit:  item,
				value: i,
			})
		}
	}
	return c
}

func (d cards) print() { //Extension涵式 ,只能針對type定義出的物件(類似物件中的屬性概念)

	for _, item := range d[0:] { //從第a[0]開始印到最後
		fmt.Println(item.suit, item.value)
	}

	// for _, item := range d[8:10] { //從第8個元素開始印到第10個元素前(A8~A9)
	// 	fmt.Println(item)
	// }

	// for _, item := range d[:10] { //從第0個元素開始印到第10個元素前(A0~A9)
	// 	fmt.Println(item)
	// }

}

func (d cards) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())

	for i := range d {
		n := rand.New(s).Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}

func (d cards) deal(size int) (cards, cards, cards) {
	return d, d[:size], d[size:]
}

// func (d cards) toString() string {

// 	return strings.Join(d, ",")
// }

// func (d cards) saveTofile(filename string) error {

// 	s := []byte(d.toString())

// 	return ioutil.WriteFile(filename, s, 0666)

// }

// func readFile(filepath string) cards {
// 	b, err := ioutil.ReadFile(filepath)

// 	if err != nil {
// 		println("Error: ", err.Error())
// 		os.Exit(-1)
// 	}

// 	return strings.Split(string(b), ",")
// }
