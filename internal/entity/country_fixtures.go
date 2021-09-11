package entity

type CountryMap map[string]Country

var CountryFixtures = CountryMap{
	"china": {
		ID:                 "cn",
		CountrySlug:        "china",
		CountryName:        "China",
		CountryDescription: "Country description",
		CountryNotes:       "Country Notes",
		New: false,
	},
}

// CreateCountryFixtures inserts known entities into the database for testing.
func CreateCountryFixtures() {
	for _, entity := range CountryFixtures {
		DB().Create(&entity)
	}
}
