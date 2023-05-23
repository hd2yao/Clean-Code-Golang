package ch2

import (
    "fmt"
    "strconv"
)

func printGuessStatistics(candidate rune, count int) {
    var number string
    var verb string
    var pluralModifier string

    if count == 0 {
        number = "no"
        verb = "are"
        pluralModifier = "s"
    } else if count == 1 {
        number = "1"
        verb = "is"
        pluralModifier = ""
    } else {
        number = strconv.Itoa(count)
        verb = "are"
        pluralModifier = "s"
    }

    guessMessage := fmt.Sprintf("There %s %s %c%s", verb, number, candidate, pluralModifier)
    fmt.Println(guessMessage)
}

type GuessStatusMessage struct {
    number         string
    verb           string
    pluralModifier string
}
