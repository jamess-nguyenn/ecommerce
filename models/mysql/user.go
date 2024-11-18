package mysql

type User struct {
	CommonColumn
	CompanyId uint64 `json:"company_id" gorm:"column:company_id"`
	Name      string `json:"name" gorm:"column:name"`
	Email     string `json:"email" gorm:"column:email"`
}

func (User) TableName() string {
	return "users"
}
