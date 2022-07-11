package godzilla

type IPlugins interface {
	Inject() error
	Use() error
}
