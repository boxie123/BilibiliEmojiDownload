package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func GetEmojiInfo(itemID int) (*Emoji, error) {
	baseUrl := fmt.Sprintf("https://api.bilibili.com/bapis/main.community.interface.emote.EmoteService/PackageDetail?id=%d", itemID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_1_1 like Mac OS X) AppleWebKit/619.2.8.10.9 (KHTML, like Gecko) Mobile/22B91 BiliApp/83000100 os/ios model/iPhone 13 mobi_app/iphone build/83000100 osVer/18.1.1 network/2 channel/AppStore Buvid/YF4BDFF823E8BA68449892FA07B6F4028355 c_locale/zh-Hans_CN s_locale/zh-Hans_CN sessionID/11bb9479 disable_rcmd/0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	emojiInfoResp := Emoji{}
	err = json.NewDecoder(resp.Body).Decode(&emojiInfoResp)
	if err != nil {
		return nil, err
	}

	return &emojiInfoResp, nil
}

func analyzeEmoji(emoji *Emoji) []DownloadInfo {
	var diList []DownloadInfo
	pkg := emoji.Data.Package
	invalidCharacterRegex := regexp.MustCompile(`[/:*?"<>|]`)
	pkgName := invalidCharacterRegex.ReplaceAllString(fmt.Sprintf("%s_%d", pkg.Text, pkg.ID), "_")
	diList = append(diList, DownloadInfo{
		URL:      pkg.URL,
		PkgName:  pkgName,
		FileName: "cover.png",
	})
	for _, e := range pkg.Emotes {
		if e.GifURL != "" {
			diList = append(diList, DownloadInfo{
				URL:      e.GifURL,
				PkgName:  fmt.Sprintf("%s\\gif", pkgName),
				FileName: invalidCharacterRegex.ReplaceAllString(fmt.Sprintf("%s_%s.gif", pkgName, e.Meta.Alias), "_"),
			})
		}
		diList = append(diList, DownloadInfo{
			URL:      e.URL,
			PkgName:  fmt.Sprintf("%s\\png", pkgName),
			FileName: invalidCharacterRegex.ReplaceAllString(fmt.Sprintf("%s_%s.png", pkgName, e.Meta.Alias), "_"),
		})
	}
	return diList
}

func DownloadEmoji(emoji *Emoji) {
	downloadInfos := analyzeEmoji(emoji)

	var wg sync.WaitGroup
	for _, info := range downloadInfos {
		wg.Add(1)
		go func(info DownloadInfo) {
			defer wg.Done()

			fmt.Println("正在下载：" + info.FileName)
			if err := downloadFile(info); err != nil {
				log.Println("Error downloading file:", err)
				return
			}
		}(info)
	}
	wg.Wait()
	fmt.Println("\n\n下载完成\n按回车键退出")
	fmt.Scanln()
}

// downloadFile
//
//	@Description: 下载文件
//	@param info 需下载文件的信息
//	@return error 错误处理
func downloadFile(info DownloadInfo) error {
	resp, err := http.Get(info.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	dirPath := filepath.Join(".", "data", "suit", info.PkgName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	filePath := filepath.Join(dirPath, info.FileName)
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
