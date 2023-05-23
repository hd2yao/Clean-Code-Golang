package ch2

func copyChars(a1, a2 []rune) {
    for i := 0; i < len(a1); i++ {
        a2[i] = a1[i]
    }
}
