package entity

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/i18n"
	"go.vixal.xyz/esp/pkg/rnd"
	"go.vixal.xyz/esp/pkg/txt"
)

type Users []*User

// User 用户信息表
type User struct {
	ID             int64          `gorm:"primary_key" json:"-"`
	UserUID        string         `gorm:"size:32;not null" json:"UID"`             // 用户ID
	UserName       string         `gorm:"size:32;not null;unique" json:"UserName"` // 登录账号
	UserStatus     string         `gorm:"size:32;" json:"UserStatus" yaml:"UserStatus,omitempty"`
	JobLevel       string         `gorm:"size:32" json:"jobLevel,omitempty"`           // 职位级别
	UserType       string         `gorm:"size:2;default:00" json:"userType,omitempty"` // 用户类型（00系统用户）
	PrimaryEmail   string         `gorm:"size:255;index;" json:"PrimaryEmail" yaml:"PrimaryEmail,omitempty"`
	EmailConfirmed bool           `json:"EmailConfirmed" yaml:"EmailConfirmed,omitempty"`
	PhoneNumber    string         `gorm:"size:16;default:''" json:"phoneNumber,omitempty"` // 手机号码
	Avatar         string         `gorm:"size:100;default:''" json:"avatar,omitempty"`     // 头像路径
	Gender         int16          `gorm:"size:1" json:"gender,omitempty"`                  // 用户性别（0男 1女 2未知）
	Status         int16          `gorm:"size:1;default:0" json:"status,omitempty"`        // 帐号状态（0正常 1停用）
	PersonPhone    string         `gorm:"size:32;" json:"PersonPhone" yaml:"PersonPhone,omitempty"`
	PersonStatus   string         `gorm:"size:32;" json:"PersonStatus" yaml:"PersonStatus,omitempty"`
	PersonAvatar   string         `gorm:"size:255;" json:"PersonAvatar" yaml:"PersonAvatar,omitempty"`
	PersonLocation string         `gorm:"size:128;" json:"PersonLocation" yaml:"PersonLocation,omitempty"`
	PersonBio      string         `gorm:"type:TEXT;" json:"PersonBio" yaml:"PersonBio,omitempty"`
	LoginIP        string         `gorm:"size:128;default:''" json:"loginIp,omitempty"` // 最后登陆IP
	LoginAt        time.Time      `json:"-"`                                            // 最后登陆时间
	LoginAttempts  int16          `json:"-" yaml:"-"`
	Remark         string         `gorm:"size:500" json:"remark,omitempty"` // 备注
	OpenID         string         `gorm:"size:100" json:"openId,omitempty"` // 微信OpenId
	FullName       string         `gorm:"size:128;" json:"FullName" yaml:"FullName,omitempty"`
	NickName       string         `gorm:"size:100" json:"nickName,omitempty"`                 // 昵称
	JobTitle       string         `gorm:"size:64;" json:"JobTitle" yaml:"JobTitle,omitempty"` // 职位
	BirthYear      int            `json:"BirthYear" yaml:"BirthYear,omitempty"`
	BirthMonth     int            `json:"BirthMonth" yaml:"BirthMonth,omitempty"`
	BirthDay       int            `json:"BirthDay" yaml:"BirthDay,omitempty"`
	Roles          []*Role        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:user_role;" json:"roles,omitempty"`
	Address        *Address       `gorm:"association_autoupdate:false;association_autocreate:false;association_save_reference:false;PRELOAD:true;" json:"Address,omitempty" yaml:"Address,omitempty"`
	AddressID      int            `gorm:"default:1" json:"-" yaml:"-"`
	CreatedBy      string         `gorm:"size:64;default:''" json:"-"` // 创建者
	UpdatedBy      string         `gorm:"size:64;default:''" json:"-"` // 更新者
	CreatedAt      time.Time      `json:"-"`                           // 创建时间
	UpdatedAt      time.Time      `json:"-"`                           // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`              // 删除时间
	DepartmentID   int64          `gorm:"default:1" json:"-" yaml:"-"` // 部门ID
	Department     *Department    `json:"Department,omitempty"`
}

// TableName get sql table name.
//  获取数据库表名
func (User) TableName() string {
	return "user"
}

// Admin Default admin user.
var Admin = User{
	ID:           1,
	AddressID:    1,
	UserName:     "admin",
	PrimaryEmail: "Admin@example.com",
	PhoneNumber:  "1329394934",
	DepartmentID:       0,
	Avatar:       "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
}

// CreateDefaultUsers initializes the database with default user accounts.
func CreateDefaultUsers() {
	if user, err := FirstOrCreateUser(&Admin); user != nil || err == nil {
		Admin = *user
	}
}

// Create inserts a new row to the database.
func (m *User) Create() error {
	return DB().Create(m).Error
}

// Save the new row to the database.
func (m *User) Save() error {
	return DB().Save(m).Error
}

// Delete the row from the database.
func (m *User) Delete() error {
	return DB().Delete(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *User) BeforeCreate(*gorm.DB) error {
	if rnd.IsUID(m.UserUID, 'u') {
		return nil
	}
	m.UserUID = rnd.PPID('u')
	return nil
}

// FirstOrCreateUser returns an existing row, inserts a new row or nil in case of errors.
func FirstOrCreateUser(m *User) (*User, error) {
	result := User{}
	var err error
	if err = DB().Preload("Address").
		Where("id = ? OR user_uid = ? OR user_name = ?", m.ID, m.UserUID, m.UserName).
		First(&result).Error; err == nil {
		return &result, errors.New(i18n.Msg(i18n.ErrAlreadyExists, result.UserName))
	} else if err = m.Create(); err != nil {
		L().Error("user:", zap.String("UserUID",m.UserUID), zap.Error(err))
		return nil, err
	}

	return m, nil
}

// FindUserByName returns an existing user or nil if not found.
func FindUserByName(userName string) (*User, error) {
	if userName == "" {
		return nil, errors.New("name is empty")
	}

	result := User{}

	if err := DB().Preload("Address").
		Preload("Department").
		Preload("Roles").
		Joins("LEFT JOIN user_role ur ON ur.user_id = id").
		Where("user_name = ?", userName).First(&result).Error; err == nil {
		return &result, nil
	} else {
		L().Error("user not found", zap.String("name", userName))
		return nil, err
	}
}

// FindUserByUID returns an existing user or nil if not found.
func FindUserByUID(uid string) (*User, error) {
	if uid == "" {
		return nil, errors.New("UID is empty")
	}

	result := User{}

	if err := DB().Preload("Address").
		Preload("Department").
		Preload("Roles").
		Joins("LEFT JOIN user_role ur ON ur.user_id = id").
		Where("user_uid = ?", uid).First(&result).Error; err == nil {
		return &result, nil
	} else {
		L().Error("user not found", zap.String("uid", uid));
		return nil, err
	}
}

// String returns an identifier that can be used in logs.
func (m *User) String() string {
	if m.UserName != "" {
		return m.UserName
	}

	return m.UserUID
}

// Registered User returns true if the user has a username.
func (m *User) Registered() bool {
	return m.UserName != "" && rnd.IsPPID(m.UserUID, 'u')
}

// SetPassword sets a new password stored as hash.
func (m *User) SetPassword(password string) error {
	if !m.Registered() {
		return fmt.Errorf("only registered users can change their password")
	}

	if len(password) < 4 {
		return fmt.Errorf("new password for %s must be at least 4 characters", txt.Quote(m.UserName))
	}

	pw := NewPassword(m.UserUID, password)

	return pw.Save()
}

// InitPassword sets the initial user password stored as hash.
func (m *User) InitPassword(password string) {
	if !m.Registered() {
		L().Warn("only registered users can change their password")
		return
	}

	if password == "" {
		return
	}

	existing := FindPassword(m.UserUID)

	if existing != nil {
		return
	}

	pw := NewPassword(m.UserUID, password)

	if err := pw.Save(); err != nil {
		L().Error(err.Error())
	}
}

// InvalidPassword returns true if the given password does not match the hash.
func (m *User) InvalidPassword(password string) bool {
	if !m.Registered() {
		L().Warn("only registered users can change their password")
		return true
	}

	if password == "" {
		return true
	}

	time.Sleep(time.Second * 5 * time.Duration(m.LoginAttempts))

	pw := FindPassword(m.UserUID)

	if pw == nil {
		return true
	}

	if pw.InvalidPassword(password) {
		if err := DB().Model(m).UpdateColumn("login_attempts", gorm.Expr("login_attempts + ?", 1)).Error; err != nil {

			L().Error("user: update login attempts", zap.Error(err))
		}

		return true
	}

	if err := DB().Model(m).Updates(map[string]interface{}{"login_attempts": 0, "login_at": Timestamp()}).Error; err != nil {
		L().Error("user: update last login", zap.Error(err))
	}

	return false
}

func (m *User) DepartmentName() string {
	return m.Department.DepartmentName
}
