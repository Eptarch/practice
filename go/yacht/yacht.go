package yacht

var nominalArg = map[string]int{
    "ones": 1,
    "twos": 2,
    "threes": 3,
    "fours": 4,
    "fives": 5,
    "sixes": 6,
}

func Score(dice []int, category string) (result int) {
    switch category {
        case "ones", "twos", "threes", "fours", "fives", "sixes":
            for _, number := range dice {
                if number == nominalArg[category] {
                    result += number
                }
            }
        case "full house": return fullHouse(dice)
        case "four of a kind": return fourOfAKind(dice)
        case "little straight": return littleStraight(dice)
        case "big straight": return bigStraight(dice)
        case "yacht": return yacht(dice)
        case "choice": return choice(dice)
    }
    return
}


func fullHouse(dice []int) (result int) {
    got2, got3 := false, false
    count := countDice(dice)
    for _, occurrences := range count {
        if occurrences == 2 {
            got2 = true
        } else if occurrences == 3 {
            got3 = true
        }
    }
    if got2 && got3 {
        return choice(dice)
    }
    return 
}

func fourOfAKind(dice []int) (result int) {
    count := countDice(dice)
    for number, occurrences := range count {
        if occurrences >= 4 {
            return 4 * number
        }
    }
    return
}

func yacht(dice []int) (result int) {
    count := countDice(dice)
    for _, occurrences := range count {
        if occurrences == 5 {
            return 50
        }
    }
    return
}

func littleStraight(dice []int) (result int) {
    count := countDice(dice)
    for n, occurrences := range count {
        if n == 6 || occurrences > 1 {
            return
        }

    }
    return 30
}

func bigStraight(dice []int) (result int) {
    count := countDice(dice)
    for n, occurrences := range count {
        if n == 1 || occurrences > 1 {
            return
        }
    }
    return 30
}

func choice(dice []int) (result int) {
    for _, number := range dice {
        result += number
    }
    return 
}

func countDice(dice []int) map[int]int {
    count := make(map[int]int, len(dice))
    for _, n := range dice {
        count[n] += 1
    }
    return count
}
