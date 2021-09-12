package cmd

import (
	"os/exec"
	"strings"
)

// Função para executar um comando no SO
func Cmd(command string) (r string, e error) {
	inputcmd := strings.Split(command, " ")
	cmd := exec.Command(inputcmd[0], inputcmd[1:]...)
	stdout, err := cmd.Output()
	return string(stdout), err
}
