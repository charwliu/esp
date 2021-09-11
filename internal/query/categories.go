package query

import (
    "strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryLabel struct {
    Name  string
    Title string
}

func CategoryLabels(limit, offset int) (results []CategoryLabel) {
    s := DB().Session(&gorm.Session{NewDB: true})

    s = s.Table("categories").
        Select("label_name AS name").
        Joins("JOIN label l ON categories.category_id = l.id").
        Group("label_name").
        Limit(limit).Offset(offset)

    if err := s.Scan(&results).Error; err != nil {
        L().Error("categories", zap.Error(err))
        return results
    }

    for i, l := range results {
        results[i].Title = strings.Title(l.Name)
    }

    return results
}

