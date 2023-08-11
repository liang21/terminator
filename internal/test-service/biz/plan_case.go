package biz

import "time"

type PlanCase struct {
	Id       int64     `json:"id" comment:"测试报告id"`
	PlanId   int64     `json:"test_plan_id" comment:"测试计划id"`
	CaseId   int64     `json:"name" comment:"测试用例id"`
	Status   string    `json:"status" comment:"状态"`
	Result   string    `json:"result" comment:"执行结果"`
	Deleted  int64     `json:"deleted"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (c PlanCase) TableName() string {
	return "test_plan_case"
}
