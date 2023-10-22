package log

import (
	"fmt"
	"os"
)

func Debug(msg string) {
	fmt.Println("\033[34m[DEBUG] \033[0m" + msg) // * Blue
}

func Info(msg string) {
	fmt.Println("\033[32m[INFO] \033[0m" + msg) // * Green
}

func Success(msg string) {
	fmt.Println("\033[32m[SUCCESS] \033[0m" + msg) // * Green
}

func Warning(msg string) {
	fmt.Println("\033[33m[WARNING] \033[0m" + msg) // * Yellow
}

func Error(msg string) {
	fmt.Println("\033[31m[ERROR] \033[0m" + msg) // * Red
}

func Fatal(msg string) {
	fmt.Println("\033[31m[FATAL] \033[0m" + msg) // * Red
	os.Exit(1)
}
