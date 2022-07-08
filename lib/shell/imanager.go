package shell

type IManager interface {
	// Ping 验证存活
	Ping(p ...IParams) (bool, error)
	// BasicInfo 获取服务器基本信息
	BasicInfo(p ...IParams) (IResult, error)
	// CommandExec 命令执行
	CommandExec(p IParams) (IResult, error)
	OperationFile(p IParams) (IResult, error)
	OperationDatabase(p IParams) (IResult, error)
}
