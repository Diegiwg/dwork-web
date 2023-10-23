package main

import (
	"fmt"

	dwork_web "github.com/Diegiwg/dwork-web/lib"
	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
)

func main() {
	fmt.Println("...")

	dwork_web.Log()
	dwork_logger.Info("info")
}
