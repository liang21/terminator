package biz

import "time"

type TestModule struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	ProjectId int64     `json:"project_id"`
	Deleted   int64     `json:"deleted"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (m TestModule) TableName() string {
	return "test_module"
}
