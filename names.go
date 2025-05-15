package main

import (
    "fmt"
    "os"
)

func main() {
    nomes := []string{"Giovanna", "Mateus", "Miro", "Gilberto"}

    // Adicionando mais um nome
    nomes = append(nomes, "Samira")

    // Abrindo arquivo em modo append (cria se não existir)
    arquivo, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer arquivo.Close()

    // Escrevendo nomes no arquivo, cada nome em uma linha
    for _, nome := range nomes {
        if _, err := arquivo.WriteString(nome + "\n"); err != nil {
            fmt.Println("Erro ao escrever no arquivo:", err)
            return
        }
    }

    fmt.Println("Nomes adicionados com sucesso!")
}


