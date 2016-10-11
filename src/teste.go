package main

import (
"fmt"
"os"
)

func main(){
        fmt.Println("Ola")

        if len(os.Args) < 3{
            fmt.Println("Nenhum parametro foi recebido")
            os.Exit(1)
        }
}
