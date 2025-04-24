package main

import "fmt"

var nomeEscola = "Escola Técnica SENAI"

func main() {
    nome := "Manu"
    idade := 16

    mensagem := boasVindas(nome)
    fmt.Println(mensagem)

    status := verificaMaioridade(idade)
    fmt.Println(status)
}