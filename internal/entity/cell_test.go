package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	t.Run("new label", func(t *testing.T) {
		l := NewCell(1, 1)
		l.CellCategory = "restaurant"
		l.CellName = "LocationName"
		l.Place = PlaceFixtures.Pointer("zinkwazi")

		assert.Equal(t, "restaurant", l.Category())
		assert.Equal(t, false, l.NoCategory())
		assert.Equal(t, false, l.Unknown())
		assert.Equal(t, "LocationName", l.Name())
		assert.Equal(t, false, l.NoName())
		assert.Equal(t, "KwaDukuza, KwaZulu-Natal, South Africa", l.Label())
		assert.Equal(t, "KwaDukuza", l.City())
		assert.Equal(t, false, l.LongCity())
		assert.Equal(t, false, l.CityContains("xxx"))
		assert.Equal(t, false, l.NoCity())
		assert.Equal(t, "KwaZulu-Natal", l.State())
		assert.Equal(t, false, l.NoState())
		assert.Equal(t, "za", l.CountryCode())
		assert.Equal(t, "South Africa", l.CountryName())
	})
}

func TestLocation_Keywords(t *testing.T) {
	t.Run("mexico", func(t *testing.T) {
		m := CellFixtures["mexico"]
		r := m.Keywords()
		assert.Equal(t, []string{"adosada", "ancient", "botanical", "garden", "mexico", "platform", "pyramid", "state-of-mexico", "teotihuacán"}, r)
	})
	t.Run("caravan park", func(t *testing.T) {
		m := CellFixtures["caravan park"]
		r := m.Keywords()
		assert.Equal(t, []string{"camping", "caravan", "kwazulu-natal", "lobotes", "mandeni", "park", "south-africa"}, r)
	})
	t.Run("place id empty", func(t *testing.T) {
		m := &Cell{}
		r := m.Keywords()
		assert.Empty(t, r)
	})
}

func TestLocation_Find(t *testing.T) {
	t.Run("place in db", func(t *testing.T) {
		m := CellFixtures["mexico"]
		r := m.Find("")
		assert.Nil(t, r)
	})
	t.Run("invalid api", func(t *testing.T) {
		l := NewCell(2, 1)
		err := l.Find("")

		if err == nil {
			t.Fatal("error expected")
		}

		assert.Equal(t, "maps: reverse lookup disabled", err.Error())
	})
}

func TestFirstOrCreateLocation(t *testing.T) {
	t.Run("id empty", func(t *testing.T) {
		loc := &Cell{}

		assert.Nil(t, FirstOrCreateCell(loc))
	})
	t.Run("place id empty", func(t *testing.T) {
		loc := &Cell{ID: "1234jhy"}

		assert.Nil(t, FirstOrCreateCell(loc))
	})
	t.Run("success", func(t *testing.T) {
		loc := CellFixtures.Pointer("caravan park")

		result := FirstOrCreateCell(loc)

		if result == nil {
			t.Fatal("result should not be nil")
		}
		assert.NotEmpty(t, result.ID)
	})
}
