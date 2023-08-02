package pagination

type ListMeta struct {
	Name     string `json:"name"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

func GetPageOffset(page, pageSize int64) int64 {
	return (page - 1) * pageSize
}
