package asciiart

import (
	"os"
	"strings"
)

func MainRender(args []string) string {

	input := args[0]
	input = strings.ReplaceAll(input, "\\n", "\n")
	bannerFile := "./ascii-banners/" + args[1] + ".txt"
	banner, err := os.ReadFile(bannerFile)
	if err != nil {
		return err.Error()
	}
	bannerData := strings.Split(string(banner), "\n\n")

	return RenderASCIIToString(input, bannerData)
}

/* func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Print("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n")
		return
	}

	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")

	// security for empty input string
	if input == "" {
		return
	}
	if len(os.Args) == 2 {
		style := "standard.txt"
		banner, err := os.ReadFile(style)
		if err != nil {
			panic(err)
		}

		bannerData := strings.Split(string(banner), "\n\n")

		// security for \n input
		if function.RenderASCIIToString(input, bannerData) == "\n" {
			fmt.Println()
			return
		}

		fmt.Println(function.RenderASCIIToString(input, bannerData))

	} else if len(os.Args) == 3 {
		style := os.Args[2]

		if style == "shadow" {
			style = "shadow.txt"
		} else if style == "standard" {
			style = "standard.txt"
		} else if style == "thinkertoy" {
			style = "thinkertoy.txt"
		} else if style == "graffiti" {
			style = "graffiti.txt"
		} else {
			fmt.Printf("Banner \"%v\" does not exist\n", style)
			return
		}
		banner, err := os.ReadFile(style)
		if err != nil {
			panic(err)
		}
		bannerData := strings.Split(string(banner), "\n\n")
		if function.RenderASCIIToString(input, bannerData) == "\n" {
			fmt.Println()
			return
		}
		fmt.Println(function.RenderASCIIToString(input, bannerData))
	}
}
*/
