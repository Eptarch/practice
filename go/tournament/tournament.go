package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Result struct {
	name            string
	win, draw, loss int
}

type Results map[string]*Result
type SortedResults []*Result

func (sr SortedResults) Len() int {
	return len(sr)
}

func (sr SortedResults) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

// Sort results by descending points, if points are equal - sort by ascending name.
func (sr SortedResults) Less(i, j int) bool {
	if sr[i].Points() == sr[j].Points() {
		return sr[i].name < sr[j].name
	}
	return sr[i].Points() > sr[j].Points()
}

// Defines matches played calc method
func (r *Result) Played() int {
	return r.win + r.loss + r.draw
}

// Defines points calc method
func (r *Result) Points() int {
	return r.win*3 + r.draw
}

func Tally(reader io.Reader, writer io.Writer) error {
	tournament := make(Results)
	// Read to string and split it by newline
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	games := strings.Split(string(b), "\n")
	for _, line := range games {
		// Skip empty or comment lines
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		play := strings.Split(line, ";")
		// We are only interested in valid lines at this point
		if len(play) != 3 {
			return fmt.Errorf("invalid input")
		}
		// Initialize struct in map to be addressable
		if _, ok := tournament[play[0]]; !ok {
			tournament[play[0]] = &Result{name: play[0], win: 0, draw: 0, loss: 0}
		}
		// Just in case the second team is not in map yet
		if _, ok := tournament[play[1]]; !ok {
			tournament[play[1]] = &Result{name: play[1], win: 0, draw: 0, loss: 0}
		}

		switch play[2] {
		default: // no "wi", "lo", "dra" allowed at this point
			return fmt.Errorf("invalid input")
		case "win":
			tournament[play[0]].win++
			tournament[play[1]].loss++
		case "loss":
			tournament[play[0]].loss++
			tournament[play[1]].win++
		case "draw":
			tournament[play[0]].draw++
			tournament[play[1]].draw++
		}
	}

	sr := make(SortedResults, 0, len(tournament))
	for _, val := range tournament {
		sr = append(sr, val)
	}
	sort.Sort(sr)
	// Table headline
	fmt.Fprintln(writer, "Team                           | MP |  W |  D |  L |  P")
	// Witchcraft and Wizardry with padding. "-" is padding to the left.
	for _, d := range sr {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n", d.name, d.Played(), d.win, d.draw, d.loss, d.Points())
	}

	return nil
}
