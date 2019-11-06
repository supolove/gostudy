package basetest

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

//
//var dd string
//var cc string
//
//func init() {
//	flag.StringVar(&in, "dd", "", "input source dir!")
//	flag.StringVar(&out, "cc", "", "input target dir!")
//}

func main1() {

	//flag.Parse()
	//if len(dd) == 0 {
	//	fmt.Print("请输入-in 文件夹路径")
	//	//return
	//}
	//
	//if len(cc) == 0 {
	//	fmt.Print("请输入-out 文件夹路径")
	//	//return
	//}

	//fmt.Println("----------")
	//
	//str:="eJxtkkFLHDEYhv+KBI9VkskkmcxNULAou9B2KVhkiJPPbdp1ZpiJi4sI9lbRi1XvvfVSKAii9Pd0dfsvmszMLrNQ5vY875sM35cPp8hoyKyxExQjQrGklIRcSo6ZJAK9QkWZf4LUOnuizFAdgWMwdhVHVst8BEkJQ1NZKJ2wxvmYME4IZpgFOGD1CUWF4lO0qooiMdoVed7PyK7YeltOtsVJuLU9dOVaZ/6GGE2/nk8vrl9+XnueflRZBqOmaqGyXdgWnh8uZ7/unp++//1y47WGsUkhSXPt7d5gZ2UvIAHpqPq0A8bDQxA6osAVF4oqmXKsBZORAnFAl/KHyjUGvZ1e/32vK47AOBFxNz6GKWWBZKH3eTX/O5XpMncXNnAMZWXyzHfW8Tr2tBgpm9hJ4cMbvc03/debHrfJxE3Gr6fJzmE7gUa4T3ZlodLPrSFdXuk5DjyuFziETLv1xejl9nEB6/GQgIaMi0gu8NK4Zz++zS7uvaugdDe0pQ5Yiv95Op9e/fb2uFqE//PkGt1W/b4TQkIRRiHn9ULmD27gYu/qF4cCTOQa5mskWiEsZjLGFJ2d7f8DTvfqWw=="
	//
	//decode64Str := base64Decode(str)
	//
	//fmt.Println(zlibDecode(decode64Str))

	// base64编码 压缩
	json := "[{\"identity\":\"b564fe7d83e6a67a3a9c60d7598ae7b3\",\"project\":\"xaigame\",\"event\":\"$startup\",\"time\":1561383553478,\"props\":{\"$first\":true,\"$app_id\":\"6oOn1L7ESryH7x4EHg\",\"$app_name\":\"null\",\"$channel_id\":\"test\",\"$channel_name\":\"测试渠道\",\"$device_code\":\"ZUK Z2121\",\"$device_id\":\"b564fe7d83e6a67a3a9c60d7598ae7b3\",\"$device_idfa\":\"UNKNOWN\",\"$device_imei\":\"86130503352954\",\"$os_name\":\"android\",\"$os_version\":\"8.0.0\",\"$plat_type\":\"ANDROID\",\"$version_app\":\"2.0\",\"$version_channel\":\"1.0.0.10\",\"$version_pack\":\"1.0.1\",\"$version_sdk\":\"1.0.3\"}}]"

	result := zlibEncode(json)

	encodeString := base64Encode(result)

	fmt.Println(encodeString)

}

// 编码
func zlibEncode(data string) string {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write([]byte(data))
	w.Close()
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	return b.String()
}

// 解码
func zlibDecode(data string) string {

	var out bytes.Buffer
	r, err := zlib.NewReader(strings.NewReader(data))
	io.Copy(&out, r)
	r.Close()
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	return out.String()
}

func base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 2)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		if n == 0 || err != nil {
			break
		}
		dst += string(buf[:n])
	}
	return dst
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
