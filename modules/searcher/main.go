package searcher

import "strings"

var Values = []string{"hola", "fuente", "casa", "cambio", "guacamaya"}

func SearchValue(value string) []string {
	var filteredItems []string
	for _, item := range Values {
		if strings.Contains(strings.ToLower(item), strings.ToLower(value)) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
