package main

import (
	"encoding/json"
	"fmt"
	"github.com/boxie123/BilibiliEmojiDownload/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No json file provided")
		fmt.Scanln()
		return
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var emoji utils.Emoji
	err = json.Unmarshal(data, &emoji)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Output the parsed data for validation
	// fmt.Printf("Read JSON data: %+v\n", emoji)

	utils.DownloadEmoji(emoji)
}
