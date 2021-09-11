package query

import (
    "fmt"
    "os"
    "testing"

    "github.com/stretchr/testify/assert"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "go.vixal.xyz/esp/internal/entity"
)

func TestMain(m *testing.M) {
    logger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.ErrorLevel))
    log = logger.Sugar()

    if err := os.Remove(".test.db"); err == nil {
        log.Debug("removed .test.db")
    }

    db := entity.InitTestDb(entity.Postgres,
        fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=UTC",
            "esp", "Pentinum#1", "esp_db", "localhost", 5432))
    defer db.Close()

    code := m.Run()

    if err := os.Remove(".test.db"); err == nil {
        log.Debug("removed .test.db")
    }
    os.Exit(code)
}

func TestLikeAny(t *testing.T) {
    t.Run("table spoon usa img json", func(t *testing.T) {
        where := LikeAny("k.keyword", "table spoon usa img json")
        assert.Equal(t, "k.keyword LIKE 'json%' OR k.keyword LIKE 'spoon%' OR k.keyword LIKE 'table%' OR k.keyword = 'usa'", where)
    })

    t.Run("cat dog", func(t *testing.T) {
        where := LikeAny("k.keyword", "cat dog")
        assert.Equal(t, "k.keyword = 'cat' OR k.keyword = 'dog'", where)
    })

    t.Run("cats dogs", func(t *testing.T) {
        where := LikeAny("k.keyword", "cats dogs")
        assert.Equal(t, "k.keyword LIKE 'cats%' OR k.keyword = 'cat' OR k.keyword LIKE 'dogs%' OR k.keyword = 'dog'", where)
    })

    t.Run("spoon", func(t *testing.T) {
        where := LikeAny("k.keyword", "spoon")
        assert.Equal(t, "k.keyword LIKE 'spoon%'", where)
    })

    t.Run("img", func(t *testing.T) {
        where := LikeAny("k.keyword", "img")
        assert.Equal(t, "", where)
    })

    t.Run("empty", func(t *testing.T) {
        where := LikeAny("k.keyword", "")
        assert.Equal(t, "", where)
    })
}

func TestAnySlug(t *testing.T) {
    t.Run("table spoon usa img json", func(t *testing.T) {
        where := AnySlug("custom_slug", "table spoon usa img json", " ")
        assert.Equal(t, "custom_slug = 'table' OR custom_slug = 'spoon' OR custom_slug = 'usa' OR custom_slug = 'img' OR custom_slug = 'json'", where)
    })

    t.Run("cat dog", func(t *testing.T) {
        where := AnySlug("custom_slug", "cat dog", " ")
        assert.Equal(t, "custom_slug = 'cat' OR custom_slug = 'dog'", where)
    })

    t.Run("cats dogs", func(t *testing.T) {
        where := AnySlug("custom_slug", "cats dogs", " ")
        assert.Equal(t, "custom_slug = 'cats' OR custom_slug = 'cat' OR custom_slug = 'dogs' OR custom_slug = 'dog'", where)
    })

    t.Run("spoon", func(t *testing.T) {
        where := AnySlug("custom_slug", "spoon", " ")
        assert.Equal(t, "custom_slug = 'spoon'", where)
    })

    t.Run("img", func(t *testing.T) {
        where := AnySlug("custom_slug", "img", " ")
        assert.Equal(t, "custom_slug = 'img'", where)
    })

    t.Run("empty", func(t *testing.T) {
        where := AnySlug("custom_slug", "", " ")
        assert.Equal(t, "", where)
    })

    t.Run("comma separated", func(t *testing.T) {
        where := AnySlug("custom_slug", "botanical-garden|landscape|bay", Or)
        assert.Equal(t, "custom_slug = 'botanical-garden' OR custom_slug = 'landscape' OR custom_slug = 'bay'", where)
    })

    t.Run("len = 0", func(t *testing.T) {
        where := AnySlug("custom_slug", " ", "")
        assert.Equal(t, "custom_slug = '' OR custom_slug = ''", where)
    })
}
