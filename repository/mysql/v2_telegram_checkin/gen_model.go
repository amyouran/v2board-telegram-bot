package v2_telegram_checkin

import "time"

// V2TelegramCheckin 签到记录表
//
//go:generate gormgen -structs V2TelegramCheckin -input .
type V2TelegramCheckin struct {
	Id        int32     // 主键
	UserTgId  int64     // 用户TGID
	UserId    int32     // 用户id
	Award     int64     // 奖励流量
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
