package main

import "fmt"

func selamVer(isim string) string {
    mesaj := fmt.Sprintf("Merhaba, %s!", isim)
    fmt.Println(mesaj)
    return mesaj
}

func main() {
    selamVer("Lonca")
}