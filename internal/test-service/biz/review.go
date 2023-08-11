package biz

import "time"

type CaseReview struct {
	Id             int64     `json:"id" comment:"测试用例评审id"`
	ProjectId      int64     `json:"project_id" comment:"项目id"`
	Name           string    `json:"name" comment:"测试用例评审名称"`
	Status         string    `json:"status" comment:"状态"`
	Description    string    `json:"description" comment:"描述"`
	ReviewTotal    int64     `json:"review_total" comment:"用例数"`
	ReviewPassRule int64     `json:"review_pass_rule" comment:"评审通过率"`
	Deleted        int64     `json:"deleted"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

func (m CaseReview) TableName() string {
	return "test_case_review"
}
