package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		text := ""

		if i%3 == 0 {
			text += "Pin "
		}

		if i%5 == 0 {
			text += "Pan "
		}

		if text == "" {
			text = fmt.Sprintf("%v", i)
		}

		fmt.Println(text)
	}
}
