package shell

type IManager interface {
	// Ping 验证存活
	Ping(p ...IParams) bool
	// BasicInfo 获取服务器基本信息
	BasicInfo(p ...IParams) IResult
	// CommandExec 命令执行
	CommandExec() IResult
	OperationFile() IResult
	OperationDatabase() IResult
}
