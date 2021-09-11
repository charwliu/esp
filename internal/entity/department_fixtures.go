package entity

type DepartmentMap map[string]Department


func (m DepartmentMap) Get(name string) Department {
    if result, ok := m[name]; ok {
        return result
    }

    return UnknownDepartment
}

func (m DepartmentMap) Pointer(name string) *Department {
    if result, ok := m[name]; ok {
        return &result
    }

    return &UnknownDepartment
}

var DepartmentFixtures = DepartmentMap{
    "100": {
        ID:             100,
        DepartmentName: "中直机关",
        OrderNum:       0,
        ParentID:       1,
    },
    "200": {
        ID:             200,
        DepartmentName: "教育部",
        OrderNum:       0,
        ParentID:       100,
    },
    "201": {
        ID:             201,
        DepartmentName: "办公厅",
        OrderNum:       1,
        ParentID:       200,
    },
    "202": {
        ID:             202,
        DepartmentName: "人事司",
        OrderNum:       2,
        ParentID:       200,
    },
    "203": {
        ID:             203,
        DepartmentName: "财务司",
        OrderNum:       3,
        ParentID:       200,
    },
    "220": {
        ID:             220,
        DepartmentName: "综合改革司",
        OrderNum:       4,
        ParentID:       200,
    },
    "230": {
        ID:             230,
        DepartmentName: "高等教育司",
        OrderNum:       5,
        ParentID:       200,
    },
    "240": {
        ID:             240,
        DepartmentName: "基础教育司",
        OrderNum:       6,
        ParentID:       200,
    },
    "300": {
        ID:             300,
        DepartmentName: "工信部",
        OrderNum:       0,
        ParentID:       100,
    },
    "301": {
        ID:             301,
        DepartmentName: "办公厅",
        OrderNum:       1,
        ParentID:       300,
    },
    "302": {
        ID:             302,
        DepartmentName: "人事司",
        OrderNum:       2,
        ParentID:       300,
    },
    "303": {
        ID:             303,
        DepartmentName: "财务司",
        OrderNum:       3,
        ParentID:       300,
    },
    "320": {
        ID:             320,
        DepartmentName: "综合改革司",
        OrderNum:       4,
        ParentID:       300,
    },
    "400": {
        ID:             400,
        DepartmentName: "科技部",
        OrderNum:       0,
        ParentID:       100,
    },
    "401": {
        ID:             401,
        DepartmentName: "办公厅",
        OrderNum:       1,
        ParentID:       400,
    },
    "402": {
        ID:             402,
        DepartmentName: "人事司",
        OrderNum:       2,
        ParentID:       400,
    },
    "403": {
        ID:             403,
        DepartmentName: "财务司",
        OrderNum:       3,
        ParentID:       400,
    },
    "420": {
        ID:             420,
        DepartmentName: "综合改革司",
        OrderNum:       4,
        ParentID:       400,
    },
}
// CreateDepartmentFixtures inserts known entities into the database for testing.
func CreateDepartmentFixtures() {
    for _, entity := range DepartmentFixtures {
        DB().Create(&entity)
    }
}

