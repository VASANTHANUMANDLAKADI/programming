package characterfrequency
import "strings"

func FrequencyCount(name string) map[rune]int {

	frequency := make(map[rune]int)

	name = strings.ToLower(name)

	for _, ch := range name {
		frequency[ch]++
	}
	return frequency
}
