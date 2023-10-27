package routes

import (
	"fmt"
	"strings"

	"github.com/Diegiwg/dwork-web/lib/logger"
)

func recursiveDump(node Routes, ident int) {

	for key := range node {
		color := "33"
		name := key
		if key == "@" {
			color = "32"
			name = "@:" + node[key].Param
		}

		fmt.Println(strings.Repeat(" ", ident)+"\033["+color+"m["+name+"]:\033[0m", node[key])
		recursiveDump(node[key].Routes, ident+2)
	}

}

func (routes *Routes) Dump() {
	logger.Debug("Dumping Routes")
	recursiveDump(*routes, 0)
	logger.Debug("Routes Dumped")
}
