package wsm

import (
	"errors"
	"github.com/xiecat/wsm/lib/httpx"
	"github.com/xiecat/wsm/lib/shell"
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

func (b *BaseShell) Verify() error {
	if len(b.Url) == 0 {
		return errors.New("url is empty")
	}
	if len(b.Password) == 0 {
		return errors.New("password is empty")
	}
	if len(b.Script) == 0 {
		return errors.New("script is empty")
	}
	return nil
}

func (b BaseShell) Ping(p ...shell.IParams) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) BasicInfo(p ...shell.IParams) (shell.IResult, error) {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) CommandExec(p shell.IParams) (shell.IResult, error) {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) FileManagement(p shell.IParams) (shell.IResult, error) {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) DatabaseManagement(p shell.IParams) (shell.IResult, error) {
	//TODO implement me
	panic("implement me")
}
