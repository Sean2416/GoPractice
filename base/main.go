package main //main為保留字元，build時會產生可執行檔(其餘不會)
import (
	"fmt"
	"net/http"
)

// func main() {
// 	//var card string = "ACE of Spades" //強行別宣告
// 	//cards := newDeck()                  // 動態宣告
// 	//tmp := append(cards, "13 of Heart") //回傳一個新的List，原始List不會變動
// 	// fmt.Println(cards[2]) //錯誤，因為cards不存在第三個元素

// 	// cards := newDeck()
// 	// cards.shuffle()

// 	// _, hand, _ := cards.deal(13)
// 	// hand.print()
// 	// hand.saveTofile("card.txt")

// 	// cards := readFile("Test.txt")
// 	// cards.print()

// 	c := newCard()

// 	c.shuffle()
// 	_, hand, rest := c.deal(5)

// 	p := person{
// 		name:  "Leo",
// 		cards: hand,
// 	}
// 	fmt.Println(p)

// 	_, hand, rest = rest.deal(5)

// 	p = person{
// 		name:  "Sean",
// 		cards: hand,
// 	}

// 	fmt.Println(p)

// 	p.cards[0].Update() //傳直呼叫，因此內容不變
// 	fmt.Println(p)

// 	// pointer := &p.cards[0]
// 	// pointer.UpdateValue()//取記憶體位址
// 	p.cards[0].UpdateValue() //取記憶體位址，接收者定義為取址後Go會直接轉換
// 	fmt.Println(p)

// 	maps := map[string]cards{ // 第一個參數string代表key的型態,cards代表value型態
// 		"Sean": p.cards,
// 		"rest": rest,
// 	}

// 	for key, value := range maps {
// 		fmt.Println(key, value)
// 	}
// 	delete(maps, "rest")
// 	fmt.Println(maps["rest"])
// }

//interface
// func main() {

// 	w := warrior{
// 		roleInfo: info{
// 			name: "狂暴徒",
// 			hp:   999,
// 			mp:   100,
// 		},
// 	}

// 	m := magician{
// 		roleInfo: info{
// 			name: "霓楓",
// 			hp:   100,
// 			mp:   999,
// 		},
// 	}

// 	attack(w, "哥布林")
// 	attack(m, "屋邀王") //傳入的參數要實作Interface內所有的函示才能作為參數

// }

// type express interface {
// 	getAction(s string) string // 類似expression ，定義可共用函式變數(預先定義傳入及回傳內容)
// 	getDefend(s string) string
// }

// func attack(e express, s string) { //傳入的參數要實作Interface內所有的函示才能作為參數
// 	fmt.Println(e.getAction(s))
// }

// func (w warrior) getAction(s string) string {
// 	return w.roleInfo.name + " 使出劍盾凸刺攻擊 " + s
// }

// func (w warrior) getDefend(s string) string {
// 	return w.roleInfo.name + " 使用絕對防禦 " + s
// }

// func (w magician) getAction(s string) string {
// 	return w.roleInfo.name + " 使用暴風雪攻擊 " + s
// }

// routines
func main() {
	links := []string{
		"http://yahoo.com.tw",
		"http://google.com",
		"http://facebook.com",
		"http://golng.org",
	}

	for _, item := range links {
		//透過併發的方式執行函式( multi thread/ 1 CPU - Go預設使用單一CPU，可以調整)
		//同一時間1 CPU只會有執行一個程序，當有程序需要等待結果時執行下一個程序
		go CheckLink(item)
	}
}

func CheckLink(s string) {
	_, err := http.Get(s)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("Success on: ", s)

}
