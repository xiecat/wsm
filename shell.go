package wsm

import (
	"github.com/Go0p/wsm/lib/httpx"
	"github.com/Go0p/wsm/lib/shell"
)

type BaseShell struct {
	// 连接地址
	Url string
	// 连接参数
	Password string
	// shell 类型
	Script shell.ScriptType
	Proxy  string
	// 自定义 header 头
	Headers map[string]string

	Client *httpx.ReqClient
}

func (b BaseShell) Ping(p ...shell.IParams) bool {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) BasicInfo(p ...shell.IParams) shell.IResult {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) CommandExec() shell.IResult {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) OperationFile() shell.IResult {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) OperationDatabase() shell.IResult {
	//TODO implement me
	panic("implement me")
}
