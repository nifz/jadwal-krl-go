package utils
import (
	"os"
	"github.com/joho/godotenv"
)

func Token() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	
	token := os.Getenv("TOKEN")
	return token, nil
}
