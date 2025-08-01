package main
import ("fmt"
		"log"
		"os"
)

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println ("ОШИБКА -", err)
		log.Fatal(err)
	}

	fmt.Println ("Содержимое файла:")
	fmt.Println (string(data))

	err = os.WriteFile("file.txt", []byte("привет, го"), 0644)
		if err != nil {
	fmt.Println ("ОШИБКА -", err)
    log.Fatal(err)
	}
}