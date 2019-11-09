package basetest

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/com"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

func TestHuancun(t *testing.T) {
	systemDir()

	fmt.Println(strings.Repeat("~", 20))
	//saveConfigFile()

	//SaveCacheConfig("Aaaaaa", "11111111111")
	fmt.Println(GetCacheConfig("Bbbbb"))
}

func GetCacheConfig(key string) interface{} {
	cc, e := readConfigContent()
	if e != nil {
		fmt.Println("%V", e)
		return nil
	}
	value, ok := cc[key]
	if ok {
		return value
	}
	return nil
}

// 简单文件缓存
func SaveCacheConfig(key string, value interface{}) {

	if len(key) == 0 {
		log.Println("saveConfigFile key is empty")
		return
	}

	config, err := readConfigContent()
	if err != nil {
		log.Println(err)
		return
	}

	config[key] = value

	writeConfigContent(config)
}

func writeConfigContent(cacheConfig map[string]interface{}) {
	b, e := json.Marshal(cacheConfig)
	if e != nil {
		fmt.Println("error write file %V", e)
	}
	err := ioutil.WriteFile(getCacheFile(), b, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("error write file %V", e)
	}
}

func readConfigContent() (map[string]interface{}, error) {
	var cacheConfig map[string]interface{}
	var content []byte
	var err error
	cacheFile := getCacheFile()
	exist := com.IsExist(cacheFile)
	if exist {
		content, err = ioutil.ReadFile(cacheFile)
		if err != nil {
			log.Println("read configfile content is err: %V", cacheFile)
			return nil, err
		}
		if len(content) > 0 {
			err := json.Unmarshal(content, &cacheConfig)
			if err != nil {
				log.Println("read configfile content is err: %V", cacheFile)
				return nil, err
			}
			return cacheConfig, nil
		}
		return map[string]interface{}{}, nil
	} else {
		_, err := os.Create(getCacheFile())
		if err != nil {
			log.Println("create configfile is err: %V", cacheFile)
			return nil, err
		}
		return map[string]interface{}{}, nil
	}
}

func getCacheFile() string {
	dir := getUserCacheDir()
	fileName := "gofile.conf"
	return path.Join(dir, fileName)
}

func getUserCacheDir() string {
	uc, _ := os.UserCacheDir()
	fmt.Println("usercachedir : " + uc)
	return uc
}

func systemDir() {
	fmt.Println("tempdir : " + os.TempDir())

	uc, _ := os.UserCacheDir()
	fmt.Println("usercachedir : " + uc)

	ucd, _ := os.UserConfigDir()
	fmt.Println("UserConfigDir : " + ucd)

	uhd, _ := os.UserHomeDir()
	fmt.Println("UserHomeDir : " + uhd)
}
