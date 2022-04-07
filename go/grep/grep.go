package grep

import (
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) (result []string) {
	var n, l, i, v, x bool
	for _, f := range flags {
		switch f {
		case "-n": n = true
		case "-l": l = true
		case "-i": i = true
		case "-v": v = true
		case "-x": x = true
		}
	}
	m := len(files) > 1
    result = []string{} // It's funny that if you delete this, some tests won't pass.
	for _, filename := range files {
		data, _ := os.ReadFile(filename)
        lines := strings.Split(string(data), "\n")
        for j, line := range lines {
            textLine := line
            textPattern := pattern
            prefix := ""
            if i {
                textLine = strings.ToLower(textLine)
                textPattern = strings.ToLower(textPattern)
            }
            if m {
                prefix += filename + ":"
            }
            if n {
                prefix += fmt.Sprintf("%v:", j+1)
            }
            found := strings.Contains(textLine, textPattern)
            if x {
                found = textLine == textPattern
            }
            if l && found {
                result = append(result, filename)
                break
            } else if (v && strings.Trim(textLine, " ") != "" && !found) || (!v && found) {
                result = append(result, prefix+line)
            }
        }
	}
	return
}
