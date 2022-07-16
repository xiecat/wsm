package payloads

import (
	"embed"
	"errors"
	"fmt"
	"github.com/go0p/wsm/lib/encrypt"
	"github.com/go0p/wsm/lib/utils"
	"io/ioutil"
	"os"
	"strings"
)

//go:embed behinder/java/*.class
var behinderClassPayloads embed.FS

//go:embed behinder/php/*.php.txt
var behinderPhpPayloads embed.FS

//go:embed behinder/csharp/*.dll
var behinderCsharpPayloads embed.FS

//go:embed behinder/asp/*.asp.txt
var behinderAspPayloads embed.FS

//go:embed godzilla/java/enpayloadv4.class
var GodzillaClassPayload []byte

//go:embed godzilla/java/plugins/*.class
var godzillaClassPluginsFiles embed.FS

//go:embed godzilla/java/plugins/*.jar
var godzillaJarPluginsFiles embed.FS

//go:embed godzilla/php/enpayloadv4.php.txt
var GodzillaPhpPayload []byte

//go:embed godzilla/php/plugins/*.php.txt
var godzillaPhpPluginsFiles embed.FS

//go:embed godzilla/csharp/enpayload.dll
var GodzillaCsharpPayload []byte

//go:embed godzilla/csharp/plugins/*.dll
var godzillaDllPluginsFiles embed.FS

//go:embed godzilla/asp/enpayload.asp.txt
var GodzillaAspPayload []byte

//go:embed godzilla/asp/plugins/*.asp.txt
var godzillaAspPluginsFiles embed.FS

var bypassKey []byte

func init() {
	bypassKey = utils.SecretKey("wsm-bypass")
	GodzillaAspPayload, _ = encrypt.AESCBCDecrypt(GodzillaAspPayload, bypassKey, bypassKey)
	GodzillaCsharpPayload, _ = encrypt.AESCBCDecrypt(GodzillaCsharpPayload, bypassKey, bypassKey)
	GodzillaClassPayload, _ = encrypt.AESCBCDecrypt(GodzillaClassPayload, bypassKey, bypassKey)
	GodzillaPhpPayload, _ = encrypt.AESCBCDecrypt(GodzillaPhpPayload, bypassKey, bypassKey)
}

func ReadAndDecrypt(payload string) ([]byte, error) {
	payloads := strings.Split(payload, "/")
	t, script := payloads[0], payloads[1]
	var enCode []byte
	var err error
	if t == "behinder" {
		switch script {
		case "asp":
			enCode, err = behinderAspPayloads.ReadFile(payload)
		case "csharp":
			enCode, err = behinderCsharpPayloads.ReadFile(payload)
		case "java":
			enCode, err = behinderClassPayloads.ReadFile(payload)
		case "php":
			enCode, err = behinderPhpPayloads.ReadFile(payload)
		default:
			return nil, errors.New(fmt.Sprintf("script %s is error", script))
		}
	} else if t == "godzilla" {
		switch script {
		case "asp":
			enCode, err = godzillaAspPluginsFiles.ReadFile(payload)
		case "csharp":
			enCode, err = godzillaDllPluginsFiles.ReadFile(payload)
		case "java":
			enCode, err = godzillaClassPluginsFiles.ReadFile(payload)
		case "php":
			enCode, err = godzillaPhpPluginsFiles.ReadFile(payload)
		default:
			return nil, errors.New(fmt.Sprintf("script %s is error", script))
		}
	} else {
		return nil, errors.New("get payload error")
	}
	if err != nil {
		return nil, err
	}
	return encrypt.AESCBCDecrypt(enCode, bypassKey, bypassKey)
}

func en() {
	//root := "F:\\gocode\\wsm\\lib\\payloads\\behinder\\php"
	root := "F:\\gocode\\wsm\\lib\\payloads\\godzilla\\php"
	key := utils.SecretKey("wsm-bypass")
	fmt.Println(key)
	files, _ := ioutil.ReadDir(root)
	for _, f := range files {
		content := read(root + "\\" + f.Name())
		enContent, err := encrypt.AESCBCEncrypt(content, key, key)
		if err != nil {
			panic(err)
		}
		write(enContent, root+"\\en"+f.Name())
	}
}

func read(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return content
}

func write(c []byte, path string) {
	err := ioutil.WriteFile(path, c, 0644)
	if err != nil {
		panic(err)
	}
}
