package entity

import (
	"time"
)

// Job 定时任务调度表
type Job struct {
	ID             int64     `gorm:"primary_key;not null" json:"jobId"`                            // 任务ID
	Name           string    `gorm:"primary_key;size:64;not null;default:''" json:"jobName"`       // 任务名称
	Group          string    `gorm:"primary_key;size;64;not null;default:DEFAULT" json:"jobGroup"` // 任务组名
	InvokeTarget   string    `gorm:"size:500;not null" json:"invokeTarget"`                        // 调用目标字符串
	CronExpression string    `gorm:"size:255;default:''" json:"cronExpression"`                    // cron执行表达式
	MisfirePolicy  string    `gorm:"size:20;default:3" json:"misfirePolicy"`                       // 计划执行错误策略（1立即>执行 2执行一次 3放弃执行）
	Concurrent     string    `gorm:"size:1;default:1" json:"concurrent"`                           // 是否并发执行（0允许 1禁>止）
	Status         string    `gorm:"size:1;default:0" json:"status"`                               // 状态（0正常 1暂停）
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"`                           // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"`                           // 更新者
	CreatedAt      time.Time `json:"created_at"`                                                   // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`                                                   // 更新时间
	Remark         string    `gorm:"size:500;default:''" json:"remark"`                            // 备注信息
}
