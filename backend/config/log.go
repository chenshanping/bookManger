package config

type Log struct {
	Path       string     `json:"path,omitempty"`
	FilePrefix string     `json:"filePrefix,omitempty"`
	FileFormat string     `json:"fileFormat,omitempty"`
	Level      string     `json:"level,omitempty"`
	WriteWay   string     `json:"writeWay,omitempty"`
	OutFormat  string     `json:"outFormat,omitempty"`
	Zap        Zap        `json:"zap,omitempty"`
	LumberJack LumberJack `json:"lumberJack,omitempty"`
}
type LumberJack struct {
	MaxSize    int  `json:"maxSize,omitempty"`
	MaxBackups int  `json:"maxBackups,omitempty"`
	MaxAge     int  `json:"maxAge,omitempty"`
	Compress   bool `json:"compress,omitempty"`
}

type Zap struct {
	OutFormat string `json:"outFormat,omitempty"`
}
