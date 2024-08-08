package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar variaveis de ambiente", err)
	}

	cmd := exec.Command(
		"tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao rodar cmd", err)
	}
}
