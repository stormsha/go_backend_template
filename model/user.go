package model

type (
	User struct {
		ID           int    `json:"id" gorm:"primaryKey" query:"id"`
		UserAccount  string `json:"user_account" gorm:"column:user_account"`
		UserPassword string `json:"user_password" gorm:"column:user_password"`
		UnionID      string `json:"union_id" gorm:"column:union_id"`
		OpenID       string `json:"open_id" gorm:"column:open_id"`
		UserName     string `json:"user_name" gorm:"column:user_name"`
		UserAvatar   string `json:"user_avatar" gorm:"column:user_avatar"`
		UserProfile  string `json:"user_profile" gorm:"column:user_profile"`
		UserRole     string `json:"user_role" gorm:"-"`
	}

	LoginResp struct {
		ID          int    `json:"id" gorm:"primaryKey"`
		UserAccount string `json:"user_account" gorm:"column:user_account"`
		UserName    string `json:"user_name" gorm:"column:user_name"`
		UserAvatar  string `json:"user_avatar" gorm:"column:user_avatar"`
		Token       string `json:"token"`
	}

	UserRes struct {
		ID           int    `json:"id" gorm:"primaryKey"`
		UserAccount  string `json:"user_account" gorm:"column:user_account"`
		UserPassword string `json:"user_password" gorm:"column:user_password"`
		UnionID      string `json:"union_id" gorm:"column:union_id"`
		OpenID       string `json:"open_id" gorm:"column:open_id"`
		UserName     string `json:"user_name" gorm:"column:user_name"`
		UserAvatar   string `json:"user_avatar" gorm:"column:user_avatar"`
		UserProfile  string `json:"user_profile" gorm:"column:user_profile"`
		UserRole     string `json:"user_role" gorm:"-"`
		Token        string `json:"token"`
	}

	UserRegister struct {
		ID           int    `json:"id" gorm:"primaryKey"`
		UserAccount  string `json:"user_account"`
		UserPassword string `json:"user_password"`
		UserName     string `json:"user_name"`
	}

	UserDetail struct {
		ID          int    `json:"id"`
		UserAccount string `json:"user_account"`
		UnionID     string `json:"union_id"`
		OpenID      string `json:"open_id"`
		UserName    string `json:"user_name"`
		UserAvatar  string `json:"user_avatar"`
		UserProfile string `json:"user_profile"`
		UserRole    string `json:"user_role"`
	}
)

func (User) TableName() string {
	return "users"
}

func (UserRegister) TableName() string {
	return "users"
}
