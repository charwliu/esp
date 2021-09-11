package query

import (
    "go.vixal.xyz/esp/internal/entity"
)

// PurgeOrphanCountries removes countries without any photos.
func PurgeOrphanCountries() error {
    entity.FlushCountryCache()
    switch Dialect() {
    default:
        return Unscoped().Exec(`DELETE FROM countries WHERE country_slug <> ? AND id NOT IN (SELECT photo_country FROM photos)`, entity.UnknownCountry.CountrySlug).Error
    }
}


