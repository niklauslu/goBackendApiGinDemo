package model

// TUser 用户
type TUser struct {
	ID         int64  `xorm:"pk autoincr 'id'" json:"id"`
	CreateTime int32  `xorm:"created notnull default 0" json:"createTime"`
	UpdateTime int32  `xorm:"updated notnull default 0" json:"updateTime"`
	Status     int32  `xorm:"TINYINT notnull default 0" json:"status"`
	Username   string `xorm:"varchar(64) notnull default '' comment('用户名') " json:"username"`
	Password   string `xorm:"varchar(64) notnull default '' comment('密码') "  json:"password"`
	Email      string `xorm:"varchar(255) notnull default '' comment('邮箱')" json:"email"`
	Role       string `xorm:"varchar(16) notnull default '' comment('角色')" json:"role"`
	PID        int64  `xorm:"bigint(20) notnull default 0 comment('上级id') 'pid'" json:"pid"`
}
