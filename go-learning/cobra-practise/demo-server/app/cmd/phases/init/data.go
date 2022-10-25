package phases

type InitData interface {
	DryRun() bool
	ConfigDir() string
	ConfigPath() string
	ExternalInitCfg() string
	Version() string
}
