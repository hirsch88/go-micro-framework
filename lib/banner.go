package lib

import "fmt"

func PrintBanner(showBanner bool, port string) {
	if showBanner {
		fmt.Println("")
		fmt.Println("Aloha, your app is ready on http://localhost:" + port)
		fmt.Println("To shut it down, press <CTRL> + C at any time.")
		fmt.Println("")
	}
}
