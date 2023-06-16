### restful api

- 以用户信息为例，首先设计用户表`t_user`

```go
// file: model/user.go

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

```

- 设计 restful 风格路由

```bash
获取所有用户  GET        /api/users
获取单个用户  GET        /api/users/:id
创建用户     POST       /api/users
更新用户     PUT        /api/users/:id
删除用户     DELETE     /api/users/:id
```

添加对应router
```go
// file: router.go
// 具体方法可见文件夹 /apis/user/

api.GET("/users", apis_user.UsersGet)
api.GET("/users/:id", apis_user.UserGet)
api.POST("/users", apis_user.UserCreate)
api.PUT("/users", apis_user.UserUpdate)
api.DELETE("/users/:id", apis_user.UserDelete)

```
