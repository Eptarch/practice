package binarysearch

func SearchInts(list []int, key int) int {
    low, high := 0, len(list) - 1
    for low <= high {
        middle := (low + high) / 2
        guess := list[middle]
        switch {
            default: return middle
            case guess > key: high = middle - 1
            case guess < key: low = middle + 1
        }
    }
    return -1
}
