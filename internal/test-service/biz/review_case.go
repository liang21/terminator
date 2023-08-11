package biz

import "time"

type ReviewCase struct {
	Id       int64     `json:"id" comment:"测试报告id"`
	ReviewId int64     `json:"review_id" comment:"review_id"`
	CaseId   int64     `json:"name" comment:"测试用例id"`
	Status   string    `json:"status" comment:"状态"`
	Deleted  int64     `json:"deleted"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (c ReviewCase) TableName() string {
	return "test_review_case"
}
