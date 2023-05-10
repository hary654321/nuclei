package cache

// Log task log
type TaskLog struct {
	TaskID         string         `json:"taskid"`         // run taskid
	RunTaskID      string         `json:"runTaskId"`      // run taskid
	StartTime      string         `json:"start_time"`     // ms
	StartTimeStr   string         `json:"start_timestr"`  //
	EndTime        string         `json:"end_time"`       // ms
	EndTimeStr     string         `json:"end_timestr"`    //
	TotalRunTime   int            `json:"total_runtime"`  // ms
	Status         int            `json:"status"`         // 任务运行结果 -1 失败 0进行中 1 成功
	Progress       int            `json:"progress"`       // 任务进度
	RunningThreads map[string]int `json:"runningThreads"` // 任务进度
	Res            string         `json:"res"`            // 任务执行过程日志
	ErrCode        int            `json:"err_code"`       // err code
	ErrMsg         string         `json:"err_msg"`        // 错误原因
	ErrTaskID      string         `json:"err_taskid"`     // task failed id
	ErrTask        string         `json:"err_task"`       // task failed id
}
