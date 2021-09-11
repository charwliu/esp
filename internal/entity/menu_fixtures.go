package entity

type MenuMap map[string]Menu

func (m MenuMap) Get(name string) Menu {
	if result, ok := m[name]; ok {
		return result
	}

	return UnknownMenu
}

func (m MenuMap) Pointer(name string) *Menu {
	if result, ok := m[name]; ok {
		return &result
	}

	return &UnknownMenu
}

var (
	fileId int64 = 1
	editId int64 = 10
)

var MenuFixtures = MenuMap{
	"File": {
		ID:        fileId,
		Name:      "File",
		OrderNum:  1,
		URL:       "file",
		Target:    "file",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "file",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
	"Edit": {
		ID:        editId,
		Name:      "Edit",
		OrderNum:  1,
		URL:       "edit",
		Target:    "edit",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "edit",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
}

var SubMenuFixtures = MenuMap{
	"Open": {
		ID:        2,
		Name:      "Open",
		ParentID:  &fileId,
		OrderNum:  1,
		URL:       "open",
		Target:    "open",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "open",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
	"Close": {
		ID:        3,
		Name:      "Close",
		ParentID:  &fileId,
		OrderNum:  2,
		URL:       "close",
		Target:    "close",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "close",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
	"Undo": {
		ID:        11,
		Name:      "Undo",
		ParentID:  &editId,
		OrderNum:  1,
		URL:       "undo",
		Target:    "undo",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "undo",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
	"Redo": {
		ID:        12,
		Name:      "Redo",
		ParentID:  &editId,
		OrderNum:  2,
		URL:       "redo",
		Target:    "redo",
		MenuType:  "M",
		Visible:   "0",
		Perms:     "",
		Icon:      "redo",
		CreatedAt: Timestamp(),
		UpdatedAt: Timestamp(),
		Remark:    "",
	},
}

// CreateMenuFixtures inserts known entities into the database for testing.
func CreateMenuFixtures() {
	for _, entity := range MenuFixtures {
		DB().Create(&entity)
	}
	for _, entity := range SubMenuFixtures {
		DB().Create(&entity)
	}
}
