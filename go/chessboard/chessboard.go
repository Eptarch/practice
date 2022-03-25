package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    c := 0
    for _, v := range cb[rank] {
        if v == true {
            c += 1
        }
    }
    return c
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    if !(file > 0 && file <= 8) {
        return 0
    }
    c := 0
    for _, r := range cb {
        if r[file - 1] == true {
            c += 1
        }
    }
    return c
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    c := 0
    for _, v := range cb {
        c += len(v)
    }
    return c
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    c := 0
    for k := range cb {
        c += CountInRank(cb, k)
    }
    return c
}
