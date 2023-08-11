package biz

import "time"

type TestPlan struct {
	Id             int64     `json:"id"`
	ProjectId      int64     `json:"project_id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Stage          string    `json:"stage"`
	Description    string    `json:"description"`
	ReviewTotal    int64     `json:"review_total"`
	ReviewPassRule int64     `json:"review_pass_rule"`
	PlanedStartAt  time.Time `json:"planed_start_at"`
	PlanedEndAt    time.Time `json:"planed_end_at"`
	ActualStartAt  time.Time `json:"actual_start_at"`
	ActualEndAt    time.Time `json:"actual_end_at"`
	Deleted        int64     `json:"deleted"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

func (p TestPlan) TableName() string {
	return "test_plan"
}
