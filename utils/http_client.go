package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func analyzeEmoji(emoji Emoji) []DownloadInfo {
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
		diList = append(diList, DownloadInfo{
			URL:      e.GifURL,
			PkgName:  pkgName,
			FileName: invalidCharacterRegex.ReplaceAllString(fmt.Sprintf("%s_%s.gif", pkgName, e.Meta.Alias), "_"),
		})
	}
	return diList
}

func DownloadEmoji(emoji Emoji) {
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
