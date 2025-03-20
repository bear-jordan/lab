package iteration

import "strings"

func Repeat(character string, i int) string {
    var result strings.Builder
    for _ = range i {
        result.WriteString(character)
    }

    return result.String()
}
