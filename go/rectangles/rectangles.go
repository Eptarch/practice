package rectangles

type Point struct {
    x, y int
}
type Pair struct {
    left, right *Point
}

func Count(diagram []string) (rectangles int) {
    seen := []*Pair{}
    for y, row := range diagram {
        // Store all encountered "+"s
        points := []*Point{}
        for x, column := range row {
            switch column {
                case '+':
                    points = append(points, &Point{x:x, y:y})
                case '-', ' ':
                    i := 0
                    for _, pair := range seen {
                        if pair.left.x != x && pair.right.x != x {
                            seen[i] = pair
                            i++
                        }
                    }
                    seen = seen[:i]
            }
        }
        if len(points) < 2 {
            continue
        }
        // Time to build all the Pairs in this row
        pairs := getCombinedPairs(points)
        // Iterate over pairs to see which of them were already seen, rectangles++ on match
        for _, pair := range pairs {
            for _, occurrence := range seen {
                if pair.left.x == occurrence.left.x && pair.right.x == occurrence.right.x {
                    rectangles++
                }
            }
        }
        seen = append(seen, pairs...)
    }
    return
}

func getCombinedPairs(points []*Point) (result []*Pair) {
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            result = append(result, &Pair{left: points[i], right: points[j]})
        }
    }
    return
}
