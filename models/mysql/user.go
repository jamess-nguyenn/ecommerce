package mysql

type User struct {
	CommonColumn
	CompanyId uint64 `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func (User) TableName() string {
	return "users"
}
