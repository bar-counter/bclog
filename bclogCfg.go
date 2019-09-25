package log

const (
	DefaultRunAddr   = ":31000"
	DefaultName      = "/template"
	DefaultApiPrefix = "/path"

	RollingPolicySize = "size"
	LogRotateDate     = 1
	LogRotateSize     = 10
	LogBackupCount    = 7
)

func setCfgDefaultOption() *BCLog {
	return &BCLog{
		//Writers:        writers,
		//LoggerLevel:    loggerLevel,
		//LoggerFile:     loggerFile,
		//LogFormatText:  logFormatText,
		//RollingPolicy:  rollingPolicy,
		LogRotateDate:  LogRotateDate,
		LogRotateSize:  LogRotateSize,
		LogBackupCount: LogBackupCount,
	}
}

type BCLog struct {
	Writers        string `yaml:"writers"`
	LoggerLevel    string `yaml:"logger_level"`
	LoggerFile     string `yaml:"logger_file"`
	LogFormatText  bool   `yaml:"log_format_text"`
	RollingPolicy  string `yaml:"rollingPolicy"`
	LogRotateDate  int    `yaml:"log_rotate_date"`
	LogRotateSize  int    `yaml:"log_rotate_size"`
	LogBackupCount int    `yaml:"log_backup_count"`
}

type BCLogFunc func(*BCLog)

var (
	defaultCfgOption = setCfgDefaultOption()
)

func NewConfig(opts ...BCLogFunc) (opt *BCLog) {
	opt = defaultCfgOption
	for _, o := range opts {
		o(opt)
	}
	defaultCfgOption = setCfgDefaultOption()
	return
}
