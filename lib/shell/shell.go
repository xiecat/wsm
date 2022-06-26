package shell

type BaseShell struct {
	// 连接地址
	Url string
	// 连接参数
	Password string
	// shell 类型
	Script ScriptType
	Proxy  string
	// 自定义 header 头
	Headers map[string]string
}

func (b BaseShell) Ping(p IParams) bool {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) BasicInfo() Result {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) CommandExec() Result {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) OperationFile() Result {
	//TODO implement me
	panic("implement me")
}

func (b BaseShell) OperationDatabase() Result {
	//TODO implement me
	panic("implement me")
}
