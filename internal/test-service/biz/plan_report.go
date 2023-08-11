package biz

import "time"

type PlanReport struct {
	Id         int64     `json:"id" comment:"测试报告id"`
	TestPlanId int64     `json:"test_plan_id" comment:"测试计划id"`
	Name       string    `json:"name" comment:"测试报告名称"`
	Status     string    `json:"status" comment:"状态"`
	Success    int32     `json:"success" comment:"成功数"`
	Deleted    int64     `json:"deleted"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}

func (r PlanReport) TableName() string {
	return "test_plan_report"
}
