package entity

type DelegationEvidence struct {
	PolicyIssuer  string `gorm:"primary_key"`
	AccessSubject string `gorm:"primary_key"`
	Polity        string
}

func (DelegationEvidence) TableName() string {
	return "delegation_evidence"
}
