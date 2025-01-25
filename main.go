package main

import (
	"fmt"
	"github.com/boxie123/BilibiliEmojiDownload/utils"
)

func main() {
	fmt.Println("想要下载的表情包id(输入错误默认千恋万花动态表情包)：")
	var emojiID int
	_, err := fmt.Scanln(&emojiID)
	if err != nil {
		emojiID = 7563
	}
	emoji, err := utils.GetEmojiInfo(emojiID)
	if err != nil {
		fmt.Errorf("获取表情信息失败：%v", err)
	}
	utils.DownloadEmoji(emoji)
}
