package biz

type TestCase struct {
	Id                int64  `json:"id"`
	ModuleId          int64  `json:"module_id"`
	ProjectId         int64  `json:"project_id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Maintainer        string `json:"maintainer"`
	Priority          string `json:"priority"`
	Method            string `json:"method"`
	Prerequisite      string `json:"prerequisite"`
	Steps             string `json:"steps"`
	Remark            string `json:"remark"`
	ReviewStatus      int64  `json:"review_status"`
	Status            string `json:"status"`
	StepDescription   string `json:"step_description"`
	ExpectedResult    string `json:"expected_result"`
	CreateUser        string `json:"create_user"`
	OriginalStatus    string `json:"original_status"`
	LastExecuteResult string `json:"last_execute_result"`
	Deleted           int64  `json:"deleted"`
	CreateAt          string `json:"create_at"`
	UpdateAt          string `json:"update_at"`
}
