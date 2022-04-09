package pascal

func Triangle(n int) [][]int {
    var result = make([][]int, n)
    for i := 0; i < n; i++ {
        result[i] = getRow(i)
    }
    return result
}

func getRow(n int) []int {
    var result = make([]int, n + 1)
    for i, prev, curr := 0, 1, 1; i <= n; i, prev = i + 1, curr  {
        if i > 0 {
            curr = (prev * (n - i + 1)) / i
        } 
        result[i] = curr
    }
    return result
}
