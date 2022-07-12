package godzilla

type IPlugins interface {
	GetPluginName() (string, []byte, error)
	GetParams() (string, *Parameter)
}
