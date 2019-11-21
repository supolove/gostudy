package main

import (
	"fmt"
	"github.com/VictorLowther/simplexml/dom"
	"io/ioutil"
	"os"
)

func main123() {

	readFile, readErr := os.Open("resources/AndroidManifest.xml")
	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	dom, domErr := dom.Parse(readFile)

	if domErr != nil {
		fmt.Println(domErr.Error())
	}

	elements := dom.Root().Children()

	for _, element := range elements {

		// 找到application
		if element.Name.Local == "application" {
			for _, child := range element.Children() {
				fixAppid(child)
				fixAppKey(child)
			}
		}
	}

	fmt.Println(dom.String())

	ioutil.WriteFile("resources/myxml.xml", dom.Bytes(), 1)

}

func fixAppid(element *dom.Element) {
	if element.Name.Local == "meta-data" {
		attrs := element.Attributes
		if len(attrs) < 2 {
			return
		}
		if attrs[0].Name.Local == "name" && attrs[0].Value == "xinhun_appid" {
			attrs[1].Value = "123456"
		} else if attrs[1].Name.Local == "name" && attrs[1].Value == "xinhun_appid" {
			attrs[0].Value = "123456"
		}
	}
}

func fixAppKey(element *dom.Element) {
	if element.Name.Local == "meta-data" {
		attrs := element.Attributes
		if len(attrs) < 2 {
			return
		}
		if attrs[0].Name.Local == "name" && attrs[0].Value == "xinhun_appkey" {
			attrs[1].Value = "8888888"
		} else if attrs[1].Name.Local == "name" && attrs[1].Value == "xinhun_appkey" {
			attrs[0].Value = "88888"
		}
	}
}
