package entities

type User struct {
	ID int64 `json:"id_user" gorm:"column:id_user"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	IdRole int64 `json:"id_role" gorm:"column:id_role"`
	Role *Role `gorm:"foreignkey:IdRole;references:ID;"`
	Email string `json:"email"`
	NoTelp string `json:"no_telp"`
	Photo string `json:"photo"`
	Status bool `json:"status"`
	Permission *[]Permission `json:"permissions" gorm:"foreignkey:IdUser"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *User) TableName() string {
	return "tb_users"
}