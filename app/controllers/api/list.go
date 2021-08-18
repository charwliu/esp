package api

import (
	"github.com/gofiber/fiber/v2"
	"go.vixal.xyz/esp/pkg/uniuri"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Status string

const (
	Normal    Status = "normal"
	Exception Status = "exception"
	Active    Status = "active"
	Success   Status = "success"
)

type Member struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Id     string `json:"id"`
}

type ListItemDataType struct {
	Id             string      `json:"id"`
	Owner          string      `json:"owner"`
	Title          string      `json:"title"`
	Avatar         string      `json:"avatar"`
	Cover          string      `json:"cover"`
	Status         Status      `json:"status"`
	Percent        int         `json:"percent"`
	Logo           string      `json:"logo"`
	Href           string      `json:"href"`
	Body           interface{} `json:"body"`
	UpdatedAt      time.Time   `json:"updated_at"`
	CreatedAt      time.Time   `json:"created_at"`
	SubDescription string      `json:"subDescription"`
	Description    string      `json:"description"`
	ActiveUser     int         `json:"activeUser"`
	NewUser        int         `json:"newUser"`
	Star           int         `json:"star"`
	Like           int         `json:"like"`
	Message        int         `json:"message"`
	Content        string      `json:"content"`
	Members        []Member    `json:"members"`
}

var user = []string{
	"付小小",
	"曲丽丽",
	"林东东",
	"周星星",
	"吴加好",
	"朱偏右",
	"鱼酱",
	"乐哥",
	"谭小仪",
	"仲尼",
}

var titles = []string{
	"康养结合是未来养老服务服务的趋势",
	"坚持“四合理”养生，身体健康才是王道",
	"智慧养老:以信息化技术创新养老服务",
	"这三种洗澡方法最养生",
	"中医中药",
	"骨质疏松",
	"食品和药物的相互作用",
	"健康睡姿排行榜-错误姿势让你惹上病痛",
}

var avatars = []string{
	"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", // Alipay
	"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png", // Angular
	"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png", // Ant Design
	"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png", // Gosh ESP
	"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png", // Bootstrap
	"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png", // React
	"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png", // Vue
	"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png", // Webpack
}

var covers = []string{
	"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png",
	"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png",
	"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png",
	"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png",
}

var desc = []string{
	"那是一种内在的东西， 他们到达不了，也无法触及的",
	"希望是一个好东西，也许是最好的，好东西是不会消亡的",
	"生命就像一盒巧克力，结果往往出人意料",
	"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆",
	"那时候我只会想自己想要什么，从不想自己拥有什么",
}

func fakeList(count int) []ListItemDataType {
	var list []ListItemDataType
	for i := 0; i < count; i++ {
		list = append(list, ListItemDataType{
			Id:             "fake-list-" + uniuri.NewLen(4) + strconv.Itoa(i),
			Owner:          user[i%len(user)],
			Title:          titles[i%len(titles)],
			Avatar:         avatars[i%len(avatars)],
			Cover:          covers[i%len(covers)],
			Status:         []Status{Active, Exception, Normal}[i%3],
			Percent:        int(math.Ceil(rand.Float64()*50) + 50),
			Logo:           avatars[i%len(avatars)],
			Href:           "http://esp.vixal.xyz",
			Body:           nil,
			UpdatedAt:      time.Time{},
			CreatedAt:      time.Time{},
			SubDescription: desc[i%len(desc)],
			Description:    "在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。",
			ActiveUser:     int(math.Ceil(rand.Float64()*1000000) + 1000000),
			NewUser:        int(math.Ceil(rand.Float64()*10000) + 10000),
			Star:           int(math.Ceil(rand.Float64()*100) + 100),
			Like:           int(math.Ceil(rand.Float64()*100) + 100),
			Message:        int(math.Ceil(rand.Float64()*10) + 10),
			Content:        "大力发展普惠型养老服务，构建居家社区机构相协调、医养康养相结合的养老服务体系，完善社区居家养老服务网络，推进公共设施适老化改造，多措并举扩大养老机构床位供给，提升服务能力和水平。",
			Members: []Member{
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png",
					Name:   "曲丽丽",
					Id:     "member1",
				},
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png",
					Name:   "王昭君",
					Id:     "member2",
				},
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png",
					Name:   "董娜娜",
					Id:     "member3",
				},
			},
		})
	}

	return list
}

func GetFakeList(router fiber.Router) {
	router.Get("/fakelist", func(ctx *fiber.Ctx) error {
		count, err := strconv.Atoi(ctx.Params("count"))
		if err != nil {
			count = 10
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": struct {
				List []ListItemDataType `json:"list"`
			}{
				List: fakeList(count),
			},
		})
	})
}
