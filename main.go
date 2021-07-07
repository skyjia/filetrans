package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/skyjia/filetrans/internal/translate"
)

// const path = "./testdata/t1"

var (
	appID string
	key   string
	delay int

	path      string
	fileTypes string

	exts []string

	translator *translate.Translator
)

func init() {
	flag.StringVar(&appID, "app_id", "", "百度翻译 App ID")
	flag.StringVar(&key, "key", "", "百度翻译 App Key")
	flag.IntVar(&delay, "delay", 1000, "百度翻译API调用延迟，单位毫秒，默认：1000ms")
	flag.StringVar(&path, "p", ".", "搜索路径")
	flag.StringVar(&fileTypes, "t", "", "搜索文件类型，例如 mp3,m4a 等等")

	flag.Parse()

	exts = strings.Split(fileTypes, ",")
	if len(exts) == 0 {
		exts = DEFAULT_AUDIO_EXTS
	}
}

func main() {
	translator = translate.NewTranslator(appID, key, time.Duration(delay)*time.Millisecond)

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				if m := matchFormat(exts, info.Name()); m {
					oldPath := path
					newPath := genNewPath(path, info)

					fmt.Printf("Renaming %q to %q\r\n", oldPath, newPath)
					err := os.Rename(oldPath, newPath)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

var DEFAULT_AUDIO_EXTS = []string{
	".mp3",
	".m4a",
	".m3u",
	".wav",
}

func matchFormat(exts []string, path string) bool {
	ext := filepath.Ext(path)
	for _, v := range exts {
		if v == ext {
			return true
		}
	}

	return false
}

func genNewPath(path string, info os.FileInfo) string {
	dir := filepath.Dir(path)
	fileName := info.Name()
	ext := filepath.Ext(fileName)
	nameOnly := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	newName := fmt.Sprintf("%s%s", translateName(nameOnly), ext)
	return filepath.Join(dir, newName)
}

func filenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func translateName(name string) string {
	return translator.Translate(name)
}
