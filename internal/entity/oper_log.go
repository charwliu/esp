package entity

import (
	"time"
)

// OperLog 操作日志记录
type OperLog struct {
	OperID        int64     `gorm:"primary_key;not null" json:"operId"`      // 日志主键
	Title         string    `gorm:"size:50;default:''" json:"title"`         // 模块标题
	BusinessType  int       `gorm:"default:0" json:"businessType"`           // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `gorm:"size:100;default:''" json:"method"`       // 方法名称
	RequestMethod string    `gorm:"size:10;default:''" json:"requestMethod"` // 请求方式
	OperatorType  int       `gorm:"default:0" json:"operatorType"`           // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `gorm:"size:50;default:''" json:"operName"`      // 操作人员
	DeptName      string    `gorm:"size:50;default:''" json:"deptName"`      // 部门名称
	OperURL       string    `gorm:"size:255;default:''" json:"operUrl"`      // 请求URL
	OperIP        string    `gorm:"size:50;default:''" json:"operIp"`        // 主机地址
	OperLocation  string    `gorm:"size:255;default:''" json:"operLocation"` // 操作地点
	OperParam     string    `gorm:"size:2000;default:''" json:"operParam"`   // 请求参数
	JSONResult    string    `gorm:"size:2000;default:''" json:"jsonResult"`  // 返回参数
	Status        *int      `gorm:"default:0" json:"status"`                 // 操作状态（0正常 1异常）
	ErrorMsg      string    `gorm:"size:2000;default:''" json:"errorMsg"`    // 错误消息
	OperAt        time.Time `json:"operTime"`                                // 操作时间
}
