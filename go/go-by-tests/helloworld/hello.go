package helloworld

import "fmt"

const (
    en = "hello"
    es = "hola"
    fr = "bonjour"
)

func greetingPrefix(lang string) string {
    switch lang {
    case "en":
        return en
    case "es":
        return es
    case "fr":
        return fr
    default:
        return ""
    }
}

func Hello(name string, lang string) {
    if name == "" {
        name = "world"
    }

    fmt.Println(greetingPrefix(lang) + " " + name)
}
