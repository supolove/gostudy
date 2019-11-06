package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type AndroidManifest struct {
	XMLName           xml.Name     `xml:"manifest"`
	AndroidPermission []Permission `xml:"uses-permission"`
	Application       Application  `xml:"application"`
	Attrs             []xml.Attr   `xml:",any,attr"`
}

type Application struct {
	XMLName     xml.Name `xml:"application"`
	Application string   `xml:",innerxml"`
}

type Permission struct {
	XMLName xml.Name `xml:"uses-permission"`
	Name    string   `xml:"name,attr"`
}

func main4() {

	//var buffer byte[]
	file, err := os.Open("resources/AndroidManifest.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := AndroidManifest{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v.XMLName)
	fmt.Println(v.Attrs)

	for n, _ := range v.AndroidPermission {
		fmt.Println(v.AndroidPermission[n].Name)
		v.AndroidPermission[n].Name = "11111111"
	}
	fmt.Println("--------")

	//解析文件
	data, _ = xml.Marshal(&v)

	ioutil.WriteFile("resources/AndroidManifestTemp.xml", data, 1)

}
