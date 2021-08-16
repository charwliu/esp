package entity

type CountryMap map[string]Country

var CountryFixtures = CountryMap{
	"germany": {
		ID:                 "de",
		CountrySlug:        "germany",
		CountryName:        "Germany",
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
