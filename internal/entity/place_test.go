package entity

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "go.vixal.xyz/esp/pkg/s2"
)

func TestCreateUnknownPlace(t *testing.T) {
    r := FirstOrCreatePlace(&UnknownPlace)
    assert.True(t, r.Unknown())
}

func TestFindPlaceByLabel(t *testing.T) {
    t.Run("find by id", func(t *testing.T) {
        r := FindPlace(s2.TokenPrefix+"1ef744d1e280", "")

        if r == nil {
            t.Fatal("result should not be nil")
        }

        assert.Equal(t, "de", r.PlaceCountry)
    })
    t.Run("find by id", func(t *testing.T) {
        r := FindPlace(s2.TokenPrefix+"85d1ea7d3278", "")

        if r == nil {
            t.Fatal("result should not be nil")
        }
        assert.Equal(t, "mx", r.PlaceCountry)
    })
    t.Run("find by label", func(t *testing.T) {
        r := FindPlace("", "KwaDukuza, KwaZulu-Natal, South Africa")

        if r == nil {
            t.Fatal("result should not be nil")
        }

        assert.Equal(t, "za", r.PlaceCountry)
    })
    t.Run("not matching", func(t *testing.T) {
        r := FindPlace("111", "xxx")

        if r != nil {
            t.Fatal("result should be nil")
        }
    })
    t.Run("not matching empty label", func(t *testing.T) {
        r := FindPlace("111", "")

        if r != nil {
            t.Fatal("result should be nil")
        }
    })
}

func TestPlace_Find(t *testing.T) {
    t.Run("record exists", func(t *testing.T) {
        m := PlaceFixtures.Get("mexico")
        if err := m.Find(); err != nil {
            t.Fatal(err)
        }
    })
    t.Run("record does not exist", func(t *testing.T) {
        place := &Place{
            ID:            s2.TokenPrefix + "1110",
            PlaceLabel:    "test",
            PlaceCity:     "testCity",
            PlaceProvince:    "",
            PlaceCountry:  "",
            PlaceKeywords: "",
            PlaceFavorite: false,
            CreatedAt:     Timestamp(),
            UpdatedAt:     Timestamp(),
        }
        err := place.Find()
        assert.EqualError(t, err, "record not found")
    })
}

func TestPlace_LongCity(t *testing.T) {
    t.Run("true", func(t *testing.T) {
        p := Place{PlaceCity: "veryveryveryverylongcity"}
        assert.True(t, p.LongCity())
    })
    t.Run("false", func(t *testing.T) {
        p := Place{PlaceCity: "short"}
        assert.False(t, p.LongCity())
    })
}

func TestPlace_NoCity(t *testing.T) {
    t.Run("true", func(t *testing.T) {
        p := Place{PlaceCity: ""}
        assert.True(t, p.NoCity())
    })
    t.Run("false", func(t *testing.T) {
        p := Place{PlaceCity: "short"}
        assert.False(t, p.NoCity())
    })
}

func TestPlace_CityContains(t *testing.T) {
    t.Run("true", func(t *testing.T) {
        p := Place{PlaceCity: "Beijing"}
        assert.True(t, p.CityContains("Beijing"))
    })
    t.Run("false", func(t *testing.T) {
        p := Place{PlaceCity: "short"}
        assert.False(t, p.CityContains("ing"))
    })
}
