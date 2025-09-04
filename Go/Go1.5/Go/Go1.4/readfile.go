package main
import ("fmt"
		"os"
)

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println ("ОШИБКА -", err)
	}

	fmt.Println ("Содержимое файла:")
	fmt.Println (string(data))
}