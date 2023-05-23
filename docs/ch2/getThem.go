package ch2

func getThem(theList [][]int) [][]int {
    var list1 [][]int
    for _, x := range theList {
        if x[0] == 4 {
            list1 = append(list1, x)
        }
    }
    return list1
}
