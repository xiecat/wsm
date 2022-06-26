package shell

type IManager interface {
	// Ping 验证存活
	Ping(p IParams) bool
	// BasicInfo 获取服务器基本信息
	BasicInfo() Result
	// CommandExec 命令执行
	CommandExec() Result
	OperationFile() Result
	OperationDatabase() Result
}
