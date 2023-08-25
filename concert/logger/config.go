package logger

var (
	LogLevel       int    // log等级,默认设置3级
	NoColor        bool   // 是否开启log输出非颜色版设置
	OutputFileName string // 用于设置log输出名称设置
	NoSave         bool   // not save file // logsync.go 中设置不进行日志写入的设置, 注：在常规的logger中并没有设置该参数
)
