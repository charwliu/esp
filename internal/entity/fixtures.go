package entity

// CreateTestFixtures inserts all known entities into the database for testing.
func CreateTestFixtures() {
	if err := Admin.SetPassword("esp.design"); err != nil {
		L().Error(err.Error())
	}

	CreateLabelFixtures()
	// CreateCameraFixtures()
	CreateCountryFixtures()
	// CreatePhotoFixtures()
	// CreateAlbumFixtures()
	// CreateAccountFixtures()
	CreateLinkFixtures()
	// CreatePhotoAlbumFixtures()
	// CreateFolderFixtures()
	// CreateFileFixtures()
	// CreateKeywordFixtures()
	// CreatePhotoKeywordFixtures()
	CreateCategoryFixtures()
	CreatePlaceFixtures()
	CreateCellFixtures()
	CreateDepartmentFixtures()
	// CreateFileShareFixtures()
	// CreateFileSyncFixtures()
	// CreateLensFixtures()
	// CreateMarkerFixtures()
	CreateMenuFixtures()
}
