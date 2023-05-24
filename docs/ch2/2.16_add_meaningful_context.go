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

type GuessStatisticsMessage struct {
    number         string
    verb           string
    pluralModifier string
}

func (g *GuessStatisticsMessage) make(candidate rune, count int) string {
    g.createPluralDependentMessageParts(count)
    return fmt.Sprintf("There %s %s %c%s", g.verb, g.number, candidate, g.pluralModifier)
}

func (g *GuessStatisticsMessage) createPluralDependentMessageParts(count int) {
    if count == 0 {
        g.thereAreNoLetters()
    } else if count == 1 {
        g.thereIsOneLetter()
    } else {
        g.thereAreManyLetters(count)
    }
}

func (g *GuessStatisticsMessage) thereAreManyLetters(count int) {
    g.number = strconv.Itoa(count)
    g.verb = "are"
    g.pluralModifier = "s"
}

func (g *GuessStatisticsMessage) thereIsOneLetter() {
    g.number = "1"
    g.verb = "is"
    g.pluralModifier = ""
}

func (g *GuessStatisticsMessage) thereAreNoLetters() {
    g.number = "no"
    g.verb = "are"
    g.pluralModifier = "s"
}
