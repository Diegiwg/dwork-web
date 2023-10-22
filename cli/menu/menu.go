package menu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Diegiwg/dwork-web/cli/log"
)

var reader = bufio.NewReader(os.Stdin)

func Question(msg string) string {
	fmt.Println(msg)

	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	return answer
}
