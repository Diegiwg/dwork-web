package dwlogger

import (
	"fmt"
	"os"
	"strings"
)

func anyToString(msg any) string {
	return fmt.Sprintf("%v", msg)
}

func print(kind string, color string, msg []any) {
	fmt.Println("\033[" + color + "m[" + kind + "]\033[0m " + anyToString(msg[0]))

	identLevel := strings.Repeat(" ", len(kind)+3) // * Kind + 3 spaces: [ and ] and space
	if len(msg) > 1 {
		for i := 1; i < len(msg); i++ {
			fmt.Println(identLevel + anyToString(msg[i])) // * Same ident level
		}
	}
}

func Debug(msg ...any) {
	print("DEBUG", "34", msg)
}

func Info(msg ...any) {
	print("INFO", "32", msg)
}

func Success(msg ...any) {
	print("SUCCESS", "32", msg)
}

func Warning(msg ...any) {
	print("WARNING", "33", msg)
}

func Error(msg ...any) {
	print("ERROR", "31", msg)
}

func Fatal(msg ...any) {
	print("ERROR", "31", msg)
	os.Exit(1)
}
