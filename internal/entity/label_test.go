package entity

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestNewLabel(t *testing.T) {
    t.Run("name Unicorn2000 priority 5", func(t *testing.T) {
        label := NewLabel("Unicorn2000", 5)
        assert.Equal(t, "Unicorn2000", label.LabelName)
        assert.Equal(t, "unicorn2000", label.LabelSlug)
        assert.Equal(t, 5, label.LabelPriority)
    })
    t.Run("name Unknown", func(t *testing.T) {
        label := NewLabel("", -6)
        assert.Equal(t, "Unknown", label.LabelName)
        assert.Equal(t, "unknown", label.LabelSlug)
        assert.Equal(t, -6, label.LabelPriority)
    })
}

func TestLabel_SetName(t *testing.T) {
    t.Run("set name", func(t *testing.T) {
        entity := LabelFixtures["landscape"]

        assert.Equal(t, "Landscape", entity.LabelName)
        assert.Equal(t, "landscape", entity.LabelSlug)
        assert.Equal(t, "landscape", entity.CustomSlug)

        entity.SetName("Landschaft")

        assert.Equal(t, "Landschaft", entity.LabelName)
        assert.Equal(t, "landscape", entity.LabelSlug)
        assert.Equal(t, "landschaft", entity.CustomSlug)
    })

    t.Run("new name empty", func(t *testing.T) {
        entity := LabelFixtures["flower"]

        assert.Equal(t, "Flower", entity.LabelName)
        assert.Equal(t, "flower", entity.LabelSlug)
        assert.Equal(t, "flower", entity.CustomSlug)

        entity.SetName("")

        assert.Equal(t, "Flower", entity.LabelName)
        assert.Equal(t, "flower", entity.LabelSlug)
        assert.Equal(t, "flower", entity.CustomSlug)
    })
}

func TestFirstOrCreateLabel(t *testing.T) {
    label := LabelFixtures.Get("flower")
    result := FirstOrCreateLabel(&label)

    if result == nil {
        t.Fatal("result should not be nil")
    }

    if result.LabelName != label.LabelName {
        t.Errorf("LabelName should be the same: %s %s", result.LabelName, label.LabelName)
    }

    if result.LabelSlug != label.LabelSlug {
        t.Errorf("LabelName should be the same: %s %s", result.LabelSlug, label.LabelSlug)
    }
}

func TestLabel_Save(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        label := NewLabel("Unicorn2000", 5)
        initialDate := label.UpdatedAt
        err := label.Save()

        if err != nil {
            t.Fatal(err)
        }
        afterDate := label.UpdatedAt

        assert.True(t, afterDate.After(initialDate))

    })
}

func TestLabel_Delete(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        label := NewLabel("LabelToBeDeleted", 5)
        err := label.Save()
        assert.False(t, label.Deleted())

        var labels Labels

        if err := DB().Where("label_name = ?", label.LabelName).Find(&labels).Error; err != nil {
            t.Fatal(err)
        }

        assert.Len(t, labels, 1)

        err = label.Delete()
        if err != nil {
            t.Fatal(err)
        }

        if err := DB().Where("label_name = ?", label.LabelName).Find(&labels).Error; err != nil {
            t.Fatal(err)
        }

        assert.Len(t, labels, 0)
    })
}

func TestLabel_Restore(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        var deleteTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

        label := &Label{DeletedAt: &deleteTime, LabelName: "ToBeRestored"}
        err := label.Save()
        if err != nil {
            t.Fatal(err)
        }
        assert.True(t, label.Deleted())

        err = label.Restore()
        if err != nil {
            t.Fatal(err)
        }
        assert.False(t, label.Deleted())
    })
    t.Run("label not deleted", func(t *testing.T) {
        label := &Label{DeletedAt: nil, LabelName: "NotDeleted1234"}
        err := label.Restore()
        if err != nil {
            t.Fatal(err)
        }
        assert.False(t, label.Deleted())
    })
}

func TestFindLabel(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        label := &Label{LabelSlug: "find-me-label", LabelName: "Find Me"}
        err := label.Save()
        if err != nil {
            t.Fatal(err)
        }
        r := FindLabel("find-me-label")
        assert.Equal(t, "Find Me", r.LabelName)
    })
    t.Run("nil", func(t *testing.T) {
        r := FindLabel("XXX")
        assert.Nil(t, r)
    })

}

func TestLabel_Links(t *testing.T) {
    t.Run("1 result", func(t *testing.T) {
        label := LabelFixtures.Get("flower")
        links := label.Links()
        assert.Equal(t, "6jxf3jfn2k", links[0].LinkToken)
    })
}

func TestLabel_Update(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        label := &Label{LabelSlug: "to-be-updated", LabelName: "Update Me Please"}
        err := label.Save()
        if err != nil {
            t.Fatal(err)
        }

        err2 := label.Update("LabelSlug", "my-unique-slug")
        if err2 != nil {
            t.Fatal(err2)
        }
        assert.Equal(t, "my-unique-slug", label.LabelSlug)
    })

}
