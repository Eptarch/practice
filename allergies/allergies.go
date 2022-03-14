package allergies

var substances = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(input uint) []string {
	result := make([]string, 0, 8)
	bytes := byte(input)
	for i := 0; bytes > 0; i++ {
		if bytes&1 == 1 {
			result = append(result, substances[i])
		}
		bytes >>= 1
	}
	return result
}

func AllergicTo(score uint, substance string) bool {
	for index, allergen := range substances {
		if substance == allergen {
			return score&(1 << uint(index)) != 0
		}
	}
	return false
}
