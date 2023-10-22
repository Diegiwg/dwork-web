package main

import (
	"os"

	"github.com/Diegiwg/dwork-web/cli/cmd"
)

// * Ferramenta para criar a estrutura inicial do projeto, além de gerenciar o desenvolvimento com o Framework.
func main() {
	cmd.Commander()
	os.Exit(0)
}
