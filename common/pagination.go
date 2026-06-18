package common

type Pagination struct {
	Page      int   `json:"page" form:"page"`
	PageSize  int   `json:"page_size" form:"page_size"`
	Total     int64 `json:"total" form:"_"`
	TotalPage int   `json:"total_page" form:"_"`
	//Total int64 `json:"total" form:"total"`
	//TotalPage int `json:"total_page" form:"total_page"`
}

func (p *Pagination) SetDefaults() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.PageSize < 1 || p.PageSize > 100 {
		p.PageSize = 10
	}
}

// check here later
//func CalculateTotalPages(total int64, pageSize int) int {
//	if pageSize == 0 {
//		return 0
//	}
//	return int((total + int64(pageSize) - 1) / int64(pageSize))
//}
