package entity

import (
	"go.vixal.xyz/esp/pkg/s2"
)

type PlacesMap map[string]Place

func (m PlacesMap) Get(name string) Place {
	if result, ok := m[name]; ok {
		return result
	}

	return UnknownPlace
}

func (m PlacesMap) Pointer(name string) *Place {
	if result, ok := m[name]; ok {
		return &result
	}

	return &UnknownPlace
}

var PlaceFixtures = PlacesMap{
	"mexico": {
		ID:            s2.TokenPrefix + "85d1ea7d3278",
		PlaceLabel:    "Teotihuacán, Mexico, Mexico",
		PlaceCity:     "Teotihuacán",
		PlaceProvince: "State of Mexico",
		PlaceCountry:  "mx",
		PlaceKeywords: "ancient, pyramid",
		PlaceFavorite: false,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"zinkwazi": {
		ID:            s2.TokenPrefix + "1ef744d1e279",
		PlaceLabel:    "KwaDukuza, KwaZulu-Natal, South Africa",
		PlaceCity:     "KwaDukuza",
		PlaceProvince: "KwaZulu-Natal",
		PlaceCountry:  "za",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"holidaypark": {
		ID:            s2.TokenPrefix + "1ef744d1e280",
		PlaceLabel:    "Holiday Park, Amusement",
		PlaceCity:     "",
		PlaceProvince: "Rheinland-Pfalz",
		PlaceCountry:  "de",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"emptyNameLongCity": {
		ID:            s2.TokenPrefix + "1ef744d1e281",
		PlaceLabel:    "labelEmptyNameLongCity",
		PlaceCity:     "longlonglonglonglongcity",
		PlaceProvince: "Rheinland-Pfalz",
		PlaceCountry:  "de",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"emptyNameShortCity": {
		ID:            s2.TokenPrefix + "1ef744d1e282",
		PlaceLabel:    "labelEmptyNameShortCity",
		PlaceCity:     "shortcity",
		PlaceProvince: "Rheinland-Pfalz",
		PlaceCountry:  "de",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"veryLongLocName": {
		ID:            s2.TokenPrefix + "1ef744d1e283",
		PlaceLabel:    "labelVeryLongLocName",
		PlaceCity:     "Mainz",
		PlaceProvince: "Rheinland-Pfalz",
		PlaceCountry:  "de",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
	"mediumLongLocName": {
		ID:            s2.TokenPrefix + "1ef744d1e284",
		PlaceLabel:    "labelMediumLongLocName",
		PlaceCity:     "New york",
		PlaceProvince: "New york",
		PlaceCountry:  "us",
		PlaceKeywords: "",
		PlaceFavorite: true,
		CreatedAt:     Timestamp(),
		UpdatedAt:     Timestamp(),
	},
}

// CreatePlaceFixtures inserts known entities into the database for testing.
func CreatePlaceFixtures() {
	for _, entity := range PlaceFixtures {
		DB().Create(&entity)
	}
}
