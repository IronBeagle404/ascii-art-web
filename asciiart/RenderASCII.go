package asciiart

import (
	"strings"
)

func RenderASCIIToString(input string, bannerData []string) string {
	runes := []rune(input)
	var final []string

	for _, r := range runes {
		if r == '\n' {
			final = append(final, "")
		} else if int(r) >= 32 && int(r) <= 126 {
			final = append(final, bannerData[int(r)-32])
		}
	}

	var builder strings.Builder
	for i := 0; i < len(final); {
		if final[i] == "" {
			builder.WriteString("\n")
			i++
			continue
		}

		j := i
		for j < len(final) && final[j] != "" {
			j++
		}

		for line := 0; line < 8; line++ {
			var result []string
			for k := i; k < j; k++ {
				lineParts := strings.Split(final[k], "\n")
				if line < len(lineParts) {
					result = append(result, lineParts[line])
				}
			}
			builder.WriteString(strings.Join(result, ""))
			if line != 7 {
				builder.WriteString("\n")
			} else {
				break
			}
		}

		i = j
	}

	return builder.String()
}
