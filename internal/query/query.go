// Package query contains frequently used database queries for use in commands and API.
package query

import (
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"github.com/jinzhu/inflection"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/pkg/txt"
)

var log = event.Log.Sugar()

// MaxResults Max result limit for queries.
const MaxResults = 10000


type Query struct {
	db *gorm.DB
}

// SearchCount is the total number of search hits.
type SearchCount struct {
	Total int
}

// New returns a new Query type with a given path and db instance.
func New(db *gorm.DB) *Query {
	q := &Query{
		db: db,
	}

	return q
}

// DB returns a database connection instance.
func DB() *gorm.DB {
	return entity.DB()
}

// Unscoped returns an unscoped database connection instance.
func Unscoped() *gorm.DB {
	return entity.DB().Unscoped()
}

// Dialect returns the sql dialect name.
func Dialect() string {
	return DB().Dialector.Name()
}

// LikeAny returns a where condition that matches any keyword in search.
func LikeAny(col, search string) (where string) {
	var wheres []string

	words := txt.UniqueKeywords(search)

	if len(words) == 0 {
		return ""
	}

	for _, w := range words {
		if len(w) > 3 {
			wheres = append(wheres, fmt.Sprintf("%s LIKE '%s%%'", col, w))
		} else {
			wheres = append(wheres, fmt.Sprintf("%s = '%s'", col, w))
		}

		if !txt.ContainsASCIILetters(w) {
			continue
		}

		singular := inflection.Singular(w)

		if singular != w {
			wheres = append(wheres, fmt.Sprintf("%s = '%s'", col, singular))
		}
	}

	return strings.Join(wheres, " OR ")
}

// AnySlug returns a where condition that matches any slug in search.
func AnySlug(col, search, sep string) (where string) {
	if search == "" {
		return ""
	}

	if sep == "" {
		sep = " "
	}

	var wheres []string
	var words []string

	for _, w := range strings.Split(search, sep) {
		w = strings.TrimSpace(w)

		words = append(words, slug.Make(w))

		if !txt.ContainsASCIILetters(w) {
			continue
		}

		singular := inflection.Singular(w)

		if singular != w {
			words = append(words, slug.Make(singular))
		}
	}

	if len(words) == 0 {
		return ""
	}
	for _, w := range words {
		wheres = append(wheres, fmt.Sprintf("%s = '%s'", col, w))
	}

	return strings.Join(wheres, " OR ")
}
