package sum

func Sum(numbers []int) int {
    result := 0

    for i := range numbers {
        result += numbers[i]
    }

    return result
}

func SumAll(numbers ...[]int) []int {
    var result []int
    for i := range numbers {
        result = append(result, Sum(numbers[i]))
    }

    return result
}

func SumAllTails(numbers ...[]int) []int {
    var result []int

    for i := range numbers {
        if len(numbers[i]) < 2 {
            result = append(result, 0)
        } else {
            result = append(result, Sum(numbers[i][1:]))
        }
    }

    return result
}
