package pagination

type ListMeta struct {
	PageSize  int64 `json:"page_size"`
	PageToken int64 `json:"page_token"`
}

func GetPageOffset(pageSize, pageToken int64) int64 {
	return (pageToken - 1) * pageSize
}
