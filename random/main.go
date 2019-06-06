package random

// Random Golang
// By Tedezed
//
// Usage:
// fmt.Print(random_string("all", 10))
// fmt.Print(random_string("low_string", 10))
// fmt.Print(random_string("uper_string", 10))
// fmt.Print(random_string("number", 10))
// fmt.Print(random_string("simbols", 10))
// fmt.Print(random_string("string_number", 10))
// fmt.Print(random_int(1, 100))

import (
    //"fmt"
    "math/rand"
    "time"
)

const (
    low_string = "abcdefghijklmnopqrstuvwxyz"
    uper_string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    number_string = "0123456789"
    simbols_string = "~!@#$%^&*()_+-={}|[]:?."
)

func random_string(type_output string, max int) (string) {
    output := ""
    possibles_characters := ""
    if type_output == "low_string" || type_output == "string_number" || type_output == "all" {
        possibles_characters += low_string
    }
    if type_output == "uper_string" || type_output == "string_number" || type_output == "all" {
        possibles_characters += uper_string
    }
    if type_output == "number" || type_output == "string_number" || type_output == "all" {
        possibles_characters += number_string
    }
    if type_output == "simbols" || type_output == "all" {
        possibles_characters += simbols_string
    }

    for i := 0; i < max; i++ {
        output += string([]rune(possibles_characters)[random_int(0,len(possibles_characters))])
    }
    return output
}

func random_int(min int, max int) (int) {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}