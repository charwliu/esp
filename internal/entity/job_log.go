package entity

import (
	"time"
)

// JobLog 定时任务调度日志表
type JobLog struct {
	JobID         int64     `json:"JobId"`                                     // 任务日志ID
	JobName       string    `json:"jobName"`                                   // 任务名称
	JobGroup      string    `json:"jobGroup"`                                  // 任务组名
	InvokeTarget  string    `gorm:"size:500;not null" json:"invokeTarget"`     // 调用目标字符串
	JobMessage    string    `gorm:"size:500" json:"jobMessage"`                // 日志信息
	Status        string    `gorm:"size:1;default:0" json:"status"`            // 执行状态（0正常 1失败）
	ExceptionInfo string    `gorm:"size:2000;default:''" json:"exceptionInfo"` // 异常信息
	CreatedAt     time.Time `json:"createdAt"`                                 // 创建时间
	Job           *Job
}
