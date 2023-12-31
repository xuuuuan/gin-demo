// Author:      xuan
// Date:        2023/6/25
// Description:

package vo

import "time"

type UserVO struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	Username    string    `gorm:"column:username;comment:用户昵称" json:"username"`
	UserAccount string    `gorm:"column:user_account;not null;comment:账号" json:"user_account"`
	UserAvatar  string    `gorm:"column:user_avatar;comment:用户头像" json:"user_avatar"`
	Gender      int32     `gorm:"column:gender;comment:性别 0-保密 1-男 2-女" json:"gender"`
	UserRole    string    `gorm:"column:user_role;not null;default:user;comment:用户角色：user / admin" json:"user_role"`
	UserKey     string    `gorm:"column:user_key;comment:key" json:"user_key"`
	UserSecret  string    `gorm:"column:user_secret;comment:secret" json:"user_secret"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}
