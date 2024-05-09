package configs

const CacheTGCheckinKeyFormat = "v2b-tg-ck-%d" // 参数为用户telegram_id

const (
	ErrCacheGet = "系统内部错误: 1001" // 缓存查询出错
	ErrCacheSet = "系统内部错误: 1002" // 缓存设置出错
	ErrCacheDel = "系统内部错误: 1003" // 缓存删除出错
	ErrDbGet    = "系统内部错误: 2001" // 数据库查询出错
	ErrDbSet    = "系统内部错误: 2002" // 数据库创建错误
	ErrStrConv  = "系统内部错误: 3001" // 字符转换错误
	ErrBotApi   = "系统内部错误: 4001" // TGBotApi 请求错误
	ErrRandom   = "系统内部错误: 5001" // 随机错误 请求错误
)
