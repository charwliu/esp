package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDepartByID(t *testing.T) {
	dept := FindDepartmentById(400)
	assert.Equal(t, []string{"科技部"}, dept.Ancestors)
	dept = FindDepartmentById(401)
    assert.Equal(t, []string{"科技部", "办公厅"}, dept.Ancestors)

}
