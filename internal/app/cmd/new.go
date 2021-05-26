package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ka1i/wispeeer/internal/pkg/config"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/logeer"
)

// NewArticle ...
func NewArticle(title string) error {
	defer utils.Timer("wispeeer ", time.Now())

	wispeeerConfig, err := config.GetWispeeerConfig()
	if err != nil {
		return err
	}
	logeer.WispeeerLogger("new", "Info", fmt.Sprintf("Location: %s", utils.GetWorkspace()))

	// 检查文章发布文件夹状态 checkAndFixPostDIR
	if !utils.IsExist(path.Join(utils.GetWorkspace(), wispeeerConfig.SourceDir)) {
		os.Mkdir(path.Join(utils.GetWorkspace(), wispeeerConfig.SourceDir), os.ModePerm)
	}
	if !utils.IsExist(path.Join(utils.GetWorkspace(), wispeeerConfig.SourceDir, config.PostsDir)) {
		os.Mkdir(path.Join(utils.GetWorkspace(), wispeeerConfig.SourceDir, config.PostsDir), os.ModePerm)
	}

	// 检查文章存在状态 Check Article status
	title = utils.SafeFormat(title, " ", "", "")
	var safeArticleFileName = utils.SafeFormat(title, "_", "md", ".")
	var filepath = path.Join(utils.GetWorkspace(), wispeeerConfig.SourceDir, config.PostsDir, safeArticleFileName)
	if utils.IsExist(filepath) {
		return fmt.Errorf("article %v is exist", safeArticleFileName)
	}
	// 创建文章文件
	err = createMarkdown(filepath, title)
	if err != nil {
		return fmt.Errorf("create article %s is failed", safeArticleFileName)
	}
	fmt.Printf("title  : %s\n", title)
	fmt.Printf("posted : %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Created: %s\n", safeArticleFileName)
	return nil
}

func createMarkdown(fileName string, title string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	fileWrite := bufio.NewWriter(file)
	//Markdown header
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("title: " + title + "\n")
	fileWrite.WriteString("posted: " + time.Now().In(config.TimeZone()).Format("2006-01-02 15:04:05") + "\n")
	fileWrite.WriteString("tags: void\n")
	fileWrite.WriteString("categories: void\n")
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("\n\n")
	fileWrite.WriteString("# Absract\n")
	fileWrite.WriteString("<!-- more -->\n\n")
	fileWrite.WriteString("# Reference\n\n")

	//Flush buffer
	fileWrite.Flush()
	return nil
}
