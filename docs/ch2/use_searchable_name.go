package ch2

func main() {
    s := 0.0
    var t []float64
    for j := 0; j < 34; j++ {
        s += (t[j] * 4) / 5
    }

    const WORK_DAYS_PER_WEEK = 5
    const NUMBER_OF_TASKS = 522
    var taskEstimate []int
    var realDaysPerIdealDay = 4
    var sum = 0
    for j := 0; j < NUMBER_OF_TASKS; j++ {
        realTaskDays := taskEstimate[j] * realDaysPerIdealDay
        realTaskWeeks := realTaskDays / WORK_DAYS_PER_WEEK
        sum += realTaskWeeks
    }
}
