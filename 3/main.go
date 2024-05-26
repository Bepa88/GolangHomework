package main

import (
	"fmt"
)

type Backpack struct {
	flashlight bool
	knife      bool
	matches    bool
}

type Location struct {
	description string
}

type Character struct {
	name             string
	isConscious      bool
	contentsBackpack Backpack
	currentLocation  Location
}

func main() {
	var choice int
	сharacter := Character{
		name:        "Стівен",
		isConscious: true,
		contentsBackpack: Backpack{
			flashlight: true,
			knife:      true,
			matches:    true,
		},
		currentLocation: Location{
			description: "start",
		},
	}

	for сharacter.isConscious {
		switch сharacter.currentLocation.description {
		case "start":
			fmt.Println("Стівен прокинувся біля входу в печеру. Він лише пам'ятає своє ім'я. Поряд з ним рюкзак з сірниками, ліхтариком і ножем.")
			fmt.Println("1. Піти в ліс")
			fmt.Println("2. Піти в печеру")
			fmt.Print("Ваш вибір: ")
			fmt.Scan(&choice)

			if choice == 1 {
				сharacter.currentLocation.description = "forest"
				fmt.Println("Стівен вирішує піти стежкою до лісу.")
				break
			}
			if choice == 2 {
				сharacter.currentLocation.description = "cave"
				fmt.Println("Стівен вирішує піти обстежити печеру.")
				break
			} else {
				fmt.Println("Введіть 1 або 2:")
				fmt.Scan(&choice)
				break
			}

		case "forest":
			fmt.Println("У лісі Стівен натикається на мертве тіло дивної тварини.")
			fmt.Println("1. Оглянути тіло тварини")
			fmt.Println("2. Пройти повз")
			fmt.Print("Ваш вибір: ")
			fmt.Scan(&choice)

			if choice == 1 {
				fmt.Println("Стівен оглядає тіло, але раптом з під тварини вилазить змія яка кусає Стівен, і він падає на землю та вмирає.")
				сharacter.currentLocation.description = "end"
				break
			}
			if choice == 2 {
				fmt.Println("Стівен проходить повз і через деякий час Стівен приходить до безлюдного табору.")
				сharacter.currentLocation.description = "camp"
				break
			} else {
				fmt.Println("Введіть 1 або 2:")
				fmt.Scan(&choice)
				break
			}

		case "cave":
			fmt.Println("У печері Стівен натикається на дивний важіль.")
			fmt.Println("1. Натиснути на важіль")
			fmt.Println("2. Пройти повз")
			fmt.Print("Ваш вибір: ")
			fmt.Scan(&choice)

			if choice == 1 {
				fmt.Println("Стівен натискає важіль але раптом підлога під ним провалюється і Стівен провалюється у прірву.")
				сharacter.currentLocation.description = "end"
				break
			}
			if choice == 2 {
				fmt.Println("Стівен проходить повз але через деякий час заходить у глухий кут. Він вирішує вийти з печери та піти в ліс.")
				сharacter.currentLocation.description = "forest"
				break
			} else {
				fmt.Println("Введіть 1 або 2:")
				fmt.Scan(&choice)
				break
			}

		case "camp":
			fmt.Println("Через деякий час Стівен приходить до безлюдного табору. Стівен вже втомлений і вирішує відпочити. У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.")
			fmt.Println("1. Спробувати відкрити сейф")
			fmt.Println("2. Пропустити сейф і оглянути інший намет")
			fmt.Print("Ваш вибір: ")
			fmt.Scan(&choice)

			if choice == 1 {
				fmt.Println("Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. Стівен непритомніє.")
				сharacter.currentLocation.description = "end"
				break
			}
			if choice == 2 {
				fmt.Println("Стівен проходить повз і раптом до табора приходять люди, які виявилися його друзями, вони розповідають йому хто він є...")
				сharacter.currentLocation.description = "end"
				break
			} else {
				fmt.Println("Введіть 1 або 2:")
				fmt.Scan(&choice)
				break
			}
		case "end":
			fmt.Println("Кінець.")
			fmt.Scan(&choice)
		}
	}
}
