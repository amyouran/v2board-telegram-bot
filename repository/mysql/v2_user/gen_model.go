package v2_user

// V2User
//
//go:generate gormgen -structs V2User -input .
type V2User struct {
	Id                int32  //
	InviteUserId      int32  //
	TelegramId        int64  //
	Email             string //
	Password          string //
	PasswordAlgo      string //
	PasswordSalt      string //
	Balance           int32  //
	Discount          int32  //
	CommissionType    int32  // 0: system 1: period 2: onetime
	CommissionRate    int32  //
	CommissionBalance int32  //
	T                 int32  //
	U                 int64  //
	D                 int64  //
	TransferEnable    int64  //
	Banned            int32  //
	IsAdmin           int32  //
	LastLoginAt       int32  //
	IsStaff           int32  //
	LastLoginIp       int32  //
	Uuid              string //
	GroupId           int32  //
	PlanId            int32  //
	SpeedLimit        int32  //
	RemindExpire      int32  //
	RemindTraffic     int32  //
	Token             string //
	ExpiredAt         int64  //
	Remarks           string //
	CreatedAt         int32  //
	UpdatedAt         int32  //
}
