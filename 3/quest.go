package main

import (
	"fmt"
	"os"
)

func yesNoChoice(i string) bool {
	if i == "так" {
		return true
	} else {
		return false
	}
}

func printPocket(p []string) {
	for _, item := range p {
		fmt.Println("-", item)
	}
}

func removePocketItem(s string, p *[]string) {
	for i, item := range *p {
		switch item {
		case s:
			*p = append((*p)[:i], (*p)[i+1:]...)
			return
		}
	}

}

func addPocketItem(s string, p *[]string) {
	*p = append(*p, s)
}

func checkPocket(s string, p []string) bool {
	for _, item := range p {
		if item == s {
			return true
		}
	}
	return false
}

func main() {
	pocket := []string{"ніж", "ключі", "запальничка"}

	fmt.Println("Стівен прокинувся на березі моря невідомого острова. Він лише пам’ятає своє ім’я.")
	fmt.Println("У нього в кишенях лежать: ")
	printPocket(pocket)
	fmt.Println("Спереду нього джунглі, а позаду океан. Сонце сильно палить з неба і Стівен хоче швидше заховатись за деревами. ")
	fmt.Println("Підійти до води? ")

	var yesNoAnswer string
	var pocketAnswer string

	fmt.Scan(&yesNoAnswer)

	yesNoResult := yesNoChoice(yesNoAnswer)

	if yesNoResult {
		fmt.Println("Стівен вирішує підійти до океану. Він дивиться в низ і бачить в воді консерву.")
		fmt.Println("Підняти її? ")
		fmt.Scan(&yesNoAnswer)
		yesNoResult := yesNoChoice(yesNoAnswer)
		if yesNoResult {
			fmt.Println("Стівен піднімає консерву, але у нього не вистачає місця в кишенях.")
			fmt.Println("Що викласти? ")
			printPocket(pocket)
			fmt.Scan(&pocketAnswer)
			removePocketItem(pocketAnswer, &pocket)
			addPocketItem("консерва", &pocket)
		}
	}

	fmt.Println("Стівен вирішує піти в джунглі, які здаються нескінченними. Коли сонце вже майже сіло Він вирішує переночувати під деревом. У шлунку пусто. ")
	if checkPocket("консерва", pocket) {
		fmt.Println("Стівен відкриває консерву та з’їдає її.")
	} else {
		fmt.Println("Стівен був голоднішим ніж йому здавалось і на ранок помирає голодною смертю. ")
		os.Exit(1)
	}

	fmt.Println("На наступний ранок Стівен прокидається від нападу дикого звіра. ")
	if checkPocket("ніж", pocket) {
		fmt.Println("Стівен дістає ніж та захищає себе.")
	} else {
		fmt.Println("Стівен немає ножа, тому не може себе захистити і звір вбиває його. ")
		os.Exit(1)
	}

	fmt.Println("Після цього він продовжує свій шлях островом. Погода сильно змінюється. Під кінець дня повітря стало дуже холодним. Стівен хоче розпалити вогонь. ")
	if checkPocket("запальничка", pocket) {
		fmt.Println("У нього є запальничка і він розпалює вогнище, яке допомагає йому дожити до ранку. ")
	} else {
		fmt.Println("У нього немає чим розпалити вогнище. Стівен лягає спати в холоді та замерзає на смерть на ранок. ")
		os.Exit(1)
	}

	fmt.Println("Наступного дня Стівен продовжує свій шлях, нарешті він виходить на берег з іншої сторони. Там причалив корабель і екіпаж забрав Стівена з собою, рятуючи його з невідомого острова. ")
}
