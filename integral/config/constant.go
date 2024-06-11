package config

const (
	// 到期终止 (目前记录于customer_log下的sexnew中,可借此查看类型区别)
	BreakTypeTimeDue = "0"
	// 手动终止
	BreakTypeByHand = "1"
)

const (
	// 关闭用户锁定
	LevelUnLock = "0"
	// 打开用户锁定
	LevelLock = "1"
)
