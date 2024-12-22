package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/manifoldco/promptui"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide a directory path")
	// 	return
	// }

	asciiArt := `                   
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⠔⠉⠒⠤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠺⠤⠔⠁⠀⠀⢀⣀⠀⠀⠀⠑⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⢀⡰⠋⠉⠈⠉⠉⠳⠤⠔⢦⡈⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⡜⠑⠂⠀⠀⠀⠀⠒⠄⠀⠀⣴⠃⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡸⠀⢰⠄⢀⠆⠀⠰⢦⠀⠀⠀⢽⣳⣰⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠔⠋⢹⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠠⣏⠀⠀⠀⠀⠀⠀⠀⠕⢻⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⢸⠀⠀⠸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⢠⡄⣀⣀⣀⠠⠴⣄⠀⠀⠀⠀⢈⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⢀⣀⣠⣧⣀⠘⠢⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⡌⠉⠒⣦⠖⠒⠉⠉⠀⠀⣠⠖⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⢮⠀⢀⡀⠀⠙⡆⠘⣦⣤⣀⣀⣀⠀⠀⢀⣀⣀⣙⣦⡀⠀⠀⠀⠀⠀⠠⡾⠓⢦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⢸⠉⠉⠉⠉⠚⡅⢀⡟⡄⠀⠈⠉⠉⠉⠉⠀⠀⠀⢸⠉⠳⢶⣤⣀⡤⠖⠁⢀⡞⠀⠈⠓⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠈⡗⠉⠓⠒⠲⡇⢸⢁⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢳⠀⠀⠀⠀⣀⣠⠴⠋⠀⠀⠀⠀⠀⠈⠲⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠙⠶⢖⣒⣺⣵⣣⠊⠀⣀⣀⣀⡠⠤⠔⡿⠀⠀⠀⡞⠀⠀⠀⡞⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠣⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠀⠀⠀⠀⠀⣇⠀⠀⠀⡇⠀⠀⠀⡇⠀⠀⠀⠀⢦⣄⣀⠀⠀⠀⠀⠀⠀⢙⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠃⠀⠀⢸⠇⠀⠀⠀⠀⣿⣿⣿⠛⢲⡄⠀⠀⢀⣼⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⢰⠀⠀⠀⢸⠀⠀⠀⠀⢰⣿⣿⡟⢀⡞⠀⢀⣠⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠦⢄⣈⠀⠀⠀⢸⠀⠀⠀⠀⣼⣿⣿⣇⡎⣠⣴⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⡇⠀⠈⠀⠀⠀⠈⠓⠒⠒⠚⠋⠉⢸⡿⢿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⢄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣸⡇⠀⠉⡹⠻⡅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⠀⠈⠉⠛⠒⠒⠒⠒⠚⣿⣿⣿⣿⣷⣀⣀⣙⣢⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠄⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠛⠛⠉⠉⠀⠀⠀⠀
  _          _          ______ _        __          __        _     __ _               
 | |        | |        |  ____(_)       \ \        / /       | |   / _| |              
 | |     ___| |_ ___   | |__   ___  __   \ \  /\  / /__  _ __| | _| |_| | _____      __
 | |    / _ \ __/ __|  |  __| | \ \/ /    \ \/  \/ / _ \| '__| |/ /  _| |/ _ \ \ /\ / /
 | |___|  __/ |_\__ \  | |    | |>  <      \  /\  / (_) | |  |   <| | | | (_) \ V  V / 
 |______\___|\__|___/  |_|    |_/_/\_\      \/  \/ \___/|_|  |_|\_\_| |_|\___/ \_/\_/  
                                                                                     
`

	fmt.Println(asciiArt)
	fmt.Println()
	// Prompt for import library
	// fmt.Print("Enter the import library to search for (e.g., @untitled-ui/icons-react): ")
	// var importLibrary string
	// fmt.Scanln(&importLibrary)

	prompt := promptui.Select{
		Label: "Select import library",
		Items: []string{
			"@untitled-ui/icons-react",
			"lucide-react",
			"@heroicons/react",
			"react-icons",
			"Custom...",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var importLibrary string
	if result == "Custom..." {
		customPrompt := promptui.Prompt{
			Label: "Enter custom import library",
		}
		importLibrary, err = customPrompt.Run()
		if err != nil {
			fmt.Printf("Custom input failed %v\n", err)
			return
		}
	} else {
		importLibrary = result
	}

	var dirPath string
	if len(os.Args) < 2 || os.Args[1] == "" {
		dirPath = "../workflow"
	} else {
		dirPath = os.Args[1]
	}
	iconCounts := make(map[string]int)
	count := 0

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".tsx") || strings.HasSuffix(path, ".jsx") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			occurrences := strings.Count(string(content), importLibrary)
			count += occurrences
			if occurrences > 0 {
				if idx := strings.Index(string(content), `} from "`+importLibrary+`"`); idx != -1 {
					importLine := string(content[:idx])
					lastImport := strings.LastIndex(importLine, "import {")
					importStrings := string(content[lastImport:idx])

					// Extract content between brackets
					start := strings.Index(importStrings, "{") + 1
					importContent := strings.TrimSpace(importStrings[start:])
					icons := strings.Split(importContent, ",")
					for _, icon := range icons {
						cleanIcon := strings.TrimSpace(icon)
						if cleanIcon != "" {
							iconCounts[cleanIcon]++
						}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Icon\tCount\t")
	fmt.Fprintln(w, "----\t-----\t")

	for icon, iconCount := range iconCounts {
		fmt.Fprintf(w, "%s\t%d\t\n", icon, iconCount)
	}

	w.Flush()

}
