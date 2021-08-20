package api

import "github.com/gofiber/fiber/v2"

type GeographicItemType struct {
	Name string `json:"name"`
	Id   string `json:"id""`
}

type GeographicType struct {
	Province GeographicItemType `json:"province"`
	City     GeographicItemType `json:"city"`
}

var provinces = []GeographicItemType{
	{
		Name: "北京市",
		Id:   "110000",
	},
	{
		Name: "天津市",
		Id:   "120000",
	},
	{
		Name: "河北省",
		Id:   "130000",
	},
	{
		Name: "山西省",
		Id:   "140000",
	},
	{
		Name: "内蒙古自治区",
		Id:   "150000",
	},
	{
		Name: "辽宁省",
		Id:   "210000",
	},
	{
		Name: "吉林省",
		Id:   "220000",
	},
	{
		Name: "黑龙江省",
		Id:   "230000",
	},
	{
		Name: "上海市",
		Id:   "310000",
	},
	{
		Name: "江苏省",
		Id:   "320000",
	},
	{
		Name: "浙江省",
		Id:   "330000",
	},
	{
		Name: "安徽省",
		Id:   "340000",
	},
	{
		Name: "福建省",
		Id:   "350000",
	},
	{
		Name: "江西省",
		Id:   "360000",
	},
	{
		Name: "山东省",
		Id:   "370000",
	},
	{
		Name: "河南省",
		Id:   "410000",
	},
	{
		Name: "湖北省",
		Id:   "420000",
	},
	{
		Name: "湖南省",
		Id:   "430000",
	},
	{
		Name: "广东省",
		Id:   "440000",
	},
	{
		Name: "广西壮族自治区",
		Id:   "450000",
	},
	{
		Name: "海南省",
		Id:   "460000",
	},
	{
		Name: "重庆市",
		Id:   "500000",
	},
	{
		Name: "四川省",
		Id:   "510000",
	},
	{
		Name: "贵州省",
		Id:   "520000",
	},
	{
		Name: "云南省",
		Id:   "530000",
	},
	{
		Name: "西藏自治区",
		Id:   "540000",
	},
	{
		Name: "陕西省",
		Id:   "610000",
	},
	{
		Name: "甘肃省",
		Id:   "620000",
	},
	{
		Name: "青海省",
		Id:   "630000",
	},
	{
		Name: "宁夏回族自治区",
		Id:   "640000",
	},
	{
		Name: "新疆维吾尔自治区",
		Id:   "650000",
	},
	{
		Name: "台湾省",
		Id:   "710000",
	},
	{
		Name: "香港特别行政区",
		Id:   "810000",
	},
	{
		Name: "澳门特别行政区",
		Id:   "820000",
	},
}

func GetProvince(router fiber.Router) {
	router.Get("/province", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": provinces,
		})
	})
}

func GetCity(router fiber.Router) {
	router.Get("/city/:province", func(ctx *fiber.Ctx) error {
		province := ctx.Params("province")
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": cities[province],
		})
	})
}

var cities = map[string][]City{
	"110000": {
		{
			Province: "北京市",
			Name:     "市辖区",
			Id:       "110100",
		},
	},
	"120000": {
		{
			Province: "天津市",
			Name:     "市辖区",
			Id:       "120100",
		},
	},
	"130000": {
		{
			Province: "河北省",
			Name:     "石家庄市",
			Id:       "130100",
		},
		{
			Province: "河北省",
			Name:     "唐山市",
			Id:       "130200",
		},
		{
			Province: "河北省",
			Name:     "秦皇岛市",
			Id:       "130300",
		},
		{
			Province: "河北省",
			Name:     "邯郸市",
			Id:       "130400",
		},
		{
			Province: "河北省",
			Name:     "邢台市",
			Id:       "130500",
		},
		{
			Province: "河北省",
			Name:     "保定市",
			Id:       "130600",
		},
		{
			Province: "河北省",
			Name:     "张家口市",
			Id:       "130700",
		},
		{
			Province: "河北省",
			Name:     "承德市",
			Id:       "130800",
		},
		{
			Province: "河北省",
			Name:     "沧州市",
			Id:       "130900",
		},
		{
			Province: "河北省",
			Name:     "廊坊市",
			Id:       "131000",
		},
		{
			Province: "河北省",
			Name:     "衡水市",
			Id:       "131100",
		},
		{
			Province: "河北省",
			Name:     "雄安",
			Id:       "139000",
		},
	},
	"140000": {
		{
			Province: "山西省",
			Name:     "太原市",
			Id:       "140100",
		},
		{
			Province: "山西省",
			Name:     "大同市",
			Id:       "140200",
		},
		{
			Province: "山西省",
			Name:     "阳泉市",
			Id:       "140300",
		},
		{
			Province: "山西省",
			Name:     "长治市",
			Id:       "140400",
		},
		{
			Province: "山西省",
			Name:     "晋城市",
			Id:       "140500",
		},
		{
			Province: "山西省",
			Name:     "朔州市",
			Id:       "140600",
		},
		{
			Province: "山西省",
			Name:     "晋中市",
			Id:       "140700",
		},
		{
			Province: "山西省",
			Name:     "运城市",
			Id:       "140800",
		},
		{
			Province: "山西省",
			Name:     "忻州市",
			Id:       "140900",
		},
		{
			Province: "山西省",
			Name:     "临汾市",
			Id:       "141000",
		},
		{
			Province: "山西省",
			Name:     "吕梁市",
			Id:       "141100",
		},
	},
	"150000": {
		{
			Province: "内蒙古自治区",
			Name:     "呼和浩特市",
			Id:       "150100",
		},
		{
			Province: "内蒙古自治区",
			Name:     "包头市",
			Id:       "150200",
		},
		{
			Province: "内蒙古自治区",
			Name:     "乌海市",
			Id:       "150300",
		},
		{
			Province: "内蒙古自治区",
			Name:     "赤峰市",
			Id:       "150400",
		},
		{
			Province: "内蒙古自治区",
			Name:     "通辽市",
			Id:       "150500",
		},
		{
			Province: "内蒙古自治区",
			Name:     "鄂尔多斯市",
			Id:       "150600",
		},
		{
			Province: "内蒙古自治区",
			Name:     "呼伦贝尔市",
			Id:       "150700",
		},
		{
			Province: "内蒙古自治区",
			Name:     "巴彦淖尔市",
			Id:       "150800",
		},
		{
			Province: "内蒙古自治区",
			Name:     "乌兰察布市",
			Id:       "150900",
		},
		{
			Province: "内蒙古自治区",
			Name:     "兴安盟",
			Id:       "152200",
		},
		{
			Province: "内蒙古自治区",
			Name:     "锡林郭勒盟",
			Id:       "152500",
		},
		{
			Province: "内蒙古自治区",
			Name:     "阿拉善盟",
			Id:       "152900",
		},
	},
	"210000": {
		{
			Province: "辽宁省",
			Name:     "沈阳市",
			Id:       "210100",
		},
		{
			Province: "辽宁省",
			Name:     "大连市",
			Id:       "210200",
		},
		{
			Province: "辽宁省",
			Name:     "鞍山市",
			Id:       "210300",
		},
		{
			Province: "辽宁省",
			Name:     "抚顺市",
			Id:       "210400",
		},
		{
			Province: "辽宁省",
			Name:     "本溪市",
			Id:       "210500",
		},
		{
			Province: "辽宁省",
			Name:     "丹东市",
			Id:       "210600",
		},
		{
			Province: "辽宁省",
			Name:     "锦州市",
			Id:       "210700",
		},
		{
			Province: "辽宁省",
			Name:     "营口市",
			Id:       "210800",
		},
		{
			Province: "辽宁省",
			Name:     "阜新市",
			Id:       "210900",
		},
		{
			Province: "辽宁省",
			Name:     "辽阳市",
			Id:       "211000",
		},
		{
			Province: "辽宁省",
			Name:     "盘锦市",
			Id:       "211100",
		},
		{
			Province: "辽宁省",
			Name:     "铁岭市",
			Id:       "211200",
		},
		{
			Province: "辽宁省",
			Name:     "朝阳市",
			Id:       "211300",
		},
		{
			Province: "辽宁省",
			Name:     "葫芦岛市",
			Id:       "211400",
		},
	},
	"220000": {
		{
			Province: "吉林省",
			Name:     "长春市",
			Id:       "220100",
		},
		{
			Province: "吉林省",
			Name:     "吉林市",
			Id:       "220200",
		},
		{
			Province: "吉林省",
			Name:     "四平市",
			Id:       "220300",
		},
		{
			Province: "吉林省",
			Name:     "辽源市",
			Id:       "220400",
		},
		{
			Province: "吉林省",
			Name:     "通化市",
			Id:       "220500",
		},
		{
			Province: "吉林省",
			Name:     "白山市",
			Id:       "220600",
		},
		{
			Province: "吉林省",
			Name:     "松原市",
			Id:       "220700",
		},
		{
			Province: "吉林省",
			Name:     "白城市",
			Id:       "220800",
		},
		{
			Province: "吉林省",
			Name:     "延边朝鲜族自治州",
			Id:       "222400",
		},
	},
	"230000": {
		{
			Province: "黑龙江省",
			Name:     "哈尔滨市",
			Id:       "230100",
		},
		{
			Province: "黑龙江省",
			Name:     "齐齐哈尔市",
			Id:       "230200",
		},
		{
			Province: "黑龙江省",
			Name:     "鸡西市",
			Id:       "230300",
		},
		{
			Province: "黑龙江省",
			Name:     "鹤岗市",
			Id:       "230400",
		},
		{
			Province: "黑龙江省",
			Name:     "双鸭山市",
			Id:       "230500",
		},
		{
			Province: "黑龙江省",
			Name:     "大庆市",
			Id:       "230600",
		},
		{
			Province: "黑龙江省",
			Name:     "伊春市",
			Id:       "230700",
		},
		{
			Province: "黑龙江省",
			Name:     "佳木斯市",
			Id:       "230800",
		},
		{
			Province: "黑龙江省",
			Name:     "七台河市",
			Id:       "230900",
		},
		{
			Province: "黑龙江省",
			Name:     "牡丹江市",
			Id:       "231000",
		},
		{
			Province: "黑龙江省",
			Name:     "黑河市",
			Id:       "231100",
		},
		{
			Province: "黑龙江省",
			Name:     "绥化市",
			Id:       "231200",
		},
		{
			Province: "黑龙江省",
			Name:     "大兴安岭地区",
			Id:       "232700",
		},
	},
	"310000": {
		{
			Province: "上海市",
			Name:     "市辖区",
			Id:       "310100",
		},
	},
	"320000": {
		{
			Province: "江苏省",
			Name:     "南京市",
			Id:       "320100",
		},
		{
			Province: "江苏省",
			Name:     "无锡市",
			Id:       "320200",
		},
		{
			Province: "江苏省",
			Name:     "徐州市",
			Id:       "320300",
		},
		{
			Province: "江苏省",
			Name:     "常州市",
			Id:       "320400",
		},
		{
			Province: "江苏省",
			Name:     "苏州市",
			Id:       "320500",
		},
		{
			Province: "江苏省",
			Name:     "南通市",
			Id:       "320600",
		},
		{
			Province: "江苏省",
			Name:     "连云港市",
			Id:       "320700",
		},
		{
			Province: "江苏省",
			Name:     "淮安市",
			Id:       "320800",
		},
		{
			Province: "江苏省",
			Name:     "盐城市",
			Id:       "320900",
		},
		{
			Province: "江苏省",
			Name:     "扬州市",
			Id:       "321000",
		},
		{
			Province: "江苏省",
			Name:     "镇江市",
			Id:       "321100",
		},
		{
			Province: "江苏省",
			Name:     "泰州市",
			Id:       "321200",
		},
		{
			Province: "江苏省",
			Name:     "宿迁市",
			Id:       "321300",
		},
	},
	"330000": {
		{
			Province: "浙江省",
			Name:     "杭州市",
			Id:       "330100",
		},
		{
			Province: "浙江省",
			Name:     "宁波市",
			Id:       "330200",
		},
		{
			Province: "浙江省",
			Name:     "温州市",
			Id:       "330300",
		},
		{
			Province: "浙江省",
			Name:     "嘉兴市",
			Id:       "330400",
		},
		{
			Province: "浙江省",
			Name:     "湖州市",
			Id:       "330500",
		},
		{
			Province: "浙江省",
			Name:     "绍兴市",
			Id:       "330600",
		},
		{
			Province: "浙江省",
			Name:     "金华市",
			Id:       "330700",
		},
		{
			Province: "浙江省",
			Name:     "衢州市",
			Id:       "330800",
		},
		{
			Province: "浙江省",
			Name:     "舟山市",
			Id:       "330900",
		},
		{
			Province: "浙江省",
			Name:     "台州市",
			Id:       "331000",
		},
		{
			Province: "浙江省",
			Name:     "丽水市",
			Id:       "331100",
		},
	},
	"340000": {
		{
			Province: "安徽省",
			Name:     "合肥市",
			Id:       "340100",
		},
		{
			Province: "安徽省",
			Name:     "芜湖市",
			Id:       "340200",
		},
		{
			Province: "安徽省",
			Name:     "蚌埠市",
			Id:       "340300",
		},
		{
			Province: "安徽省",
			Name:     "淮南市",
			Id:       "340400",
		},
		{
			Province: "安徽省",
			Name:     "马鞍山市",
			Id:       "340500",
		},
		{
			Province: "安徽省",
			Name:     "淮北市",
			Id:       "340600",
		},
		{
			Province: "安徽省",
			Name:     "铜陵市",
			Id:       "340700",
		},
		{
			Province: "安徽省",
			Name:     "安庆市",
			Id:       "340800",
		},
		{
			Province: "安徽省",
			Name:     "黄山市",
			Id:       "341000",
		},
		{
			Province: "安徽省",
			Name:     "滁州市",
			Id:       "341100",
		},
		{
			Province: "安徽省",
			Name:     "阜阳市",
			Id:       "341200",
		},
		{
			Province: "安徽省",
			Name:     "宿州市",
			Id:       "341300",
		},
		{
			Province: "安徽省",
			Name:     "六安市",
			Id:       "341500",
		},
		{
			Province: "安徽省",
			Name:     "亳州市",
			Id:       "341600",
		},
		{
			Province: "安徽省",
			Name:     "池州市",
			Id:       "341700",
		},
		{
			Province: "安徽省",
			Name:     "宣城市",
			Id:       "341800",
		},
	},
	"350000": {
		{
			Province: "福建省",
			Name:     "福州市",
			Id:       "350100",
		},
		{
			Province: "福建省",
			Name:     "厦门市",
			Id:       "350200",
		},
		{
			Province: "福建省",
			Name:     "莆田市",
			Id:       "350300",
		},
		{
			Province: "福建省",
			Name:     "三明市",
			Id:       "350400",
		},
		{
			Province: "福建省",
			Name:     "泉州市",
			Id:       "350500",
		},
		{
			Province: "福建省",
			Name:     "漳州市",
			Id:       "350600",
		},
		{
			Province: "福建省",
			Name:     "南平市",
			Id:       "350700",
		},
		{
			Province: "福建省",
			Name:     "龙岩市",
			Id:       "350800",
		},
		{
			Province: "福建省",
			Name:     "宁德市",
			Id:       "350900",
		},
	},
	"360000": {
		{
			Province: "江西省",
			Name:     "南昌市",
			Id:       "360100",
		},
		{
			Province: "江西省",
			Name:     "景德镇市",
			Id:       "360200",
		},
		{
			Province: "江西省",
			Name:     "萍乡市",
			Id:       "360300",
		},
		{
			Province: "江西省",
			Name:     "九江市",
			Id:       "360400",
		},
		{
			Province: "江西省",
			Name:     "新余市",
			Id:       "360500",
		},
		{
			Province: "江西省",
			Name:     "鹰潭市",
			Id:       "360600",
		},
		{
			Province: "江西省",
			Name:     "赣州市",
			Id:       "360700",
		},
		{
			Province: "江西省",
			Name:     "吉安市",
			Id:       "360800",
		},
		{
			Province: "江西省",
			Name:     "宜春市",
			Id:       "360900",
		},
		{
			Province: "江西省",
			Name:     "抚州市",
			Id:       "361000",
		},
		{
			Province: "江西省",
			Name:     "上饶市",
			Id:       "361100",
		},
	},
	"370000": {
		{
			Province: "山东省",
			Name:     "济南市",
			Id:       "370100",
		},
		{
			Province: "山东省",
			Name:     "青岛市",
			Id:       "370200",
		},
		{
			Province: "山东省",
			Name:     "淄博市",
			Id:       "370300",
		},
		{
			Province: "山东省",
			Name:     "枣庄市",
			Id:       "370400",
		},
		{
			Province: "山东省",
			Name:     "东营市",
			Id:       "370500",
		},
		{
			Province: "山东省",
			Name:     "烟台市",
			Id:       "370600",
		},
		{
			Province: "山东省",
			Name:     "潍坊市",
			Id:       "370700",
		},
		{
			Province: "山东省",
			Name:     "济宁市",
			Id:       "370800",
		},
		{
			Province: "山东省",
			Name:     "泰安市",
			Id:       "370900",
		},
		{
			Province: "山东省",
			Name:     "威海市",
			Id:       "371000",
		},
		{
			Province: "山东省",
			Name:     "日照市",
			Id:       "371100",
		},
		{
			Province: "山东省",
			Name:     "莱芜市",
			Id:       "371200",
		},
		{
			Province: "山东省",
			Name:     "临沂市",
			Id:       "371300",
		},
		{
			Province: "山东省",
			Name:     "德州市",
			Id:       "371400",
		},
		{
			Province: "山东省",
			Name:     "聊城市",
			Id:       "371500",
		},
		{
			Province: "山东省",
			Name:     "滨州市",
			Id:       "371600",
		},
		{
			Province: "山东省",
			Name:     "菏泽市",
			Id:       "371700",
		},
	},
	"410000": {
		{
			Province: "河南省",
			Name:     "郑州市",
			Id:       "410100",
		},
		{
			Province: "河南省",
			Name:     "开封市",
			Id:       "410200",
		},
		{
			Province: "河南省",
			Name:     "洛阳市",
			Id:       "410300",
		},
		{
			Province: "河南省",
			Name:     "平顶山市",
			Id:       "410400",
		},
		{
			Province: "河南省",
			Name:     "安阳市",
			Id:       "410500",
		},
		{
			Province: "河南省",
			Name:     "鹤壁市",
			Id:       "410600",
		},
		{
			Province: "河南省",
			Name:     "新乡市",
			Id:       "410700",
		},
		{
			Province: "河南省",
			Name:     "焦作市",
			Id:       "410800",
		},
		{
			Province: "河南省",
			Name:     "濮阳市",
			Id:       "410900",
		},
		{
			Province: "河南省",
			Name:     "许昌市",
			Id:       "411000",
		},
		{
			Province: "河南省",
			Name:     "漯河市",
			Id:       "411100",
		},
		{
			Province: "河南省",
			Name:     "三门峡市",
			Id:       "411200",
		},
		{
			Province: "河南省",
			Name:     "南阳市",
			Id:       "411300",
		},
		{
			Province: "河南省",
			Name:     "商丘市",
			Id:       "411400",
		},
		{
			Province: "河南省",
			Name:     "信阳市",
			Id:       "411500",
		},
		{
			Province: "河南省",
			Name:     "周口市",
			Id:       "411600",
		},
		{
			Province: "河南省",
			Name:     "驻马店市",
			Id:       "411700",
		},
		{
			Province: "河南省",
			Name:     "省直辖县级行政区划",
			Id:       "419000",
		},
	},
	"420000": {
		{
			Province: "湖北省",
			Name:     "武汉市",
			Id:       "420100",
		},
		{
			Province: "湖北省",
			Name:     "黄石市",
			Id:       "420200",
		},
		{
			Province: "湖北省",
			Name:     "十堰市",
			Id:       "420300",
		},
		{
			Province: "湖北省",
			Name:     "宜昌市",
			Id:       "420500",
		},
		{
			Province: "湖北省",
			Name:     "襄阳市",
			Id:       "420600",
		},
		{
			Province: "湖北省",
			Name:     "鄂州市",
			Id:       "420700",
		},
		{
			Province: "湖北省",
			Name:     "荆门市",
			Id:       "420800",
		},
		{
			Province: "湖北省",
			Name:     "孝感市",
			Id:       "420900",
		},
		{
			Province: "湖北省",
			Name:     "荆州市",
			Id:       "421000",
		},
		{
			Province: "湖北省",
			Name:     "黄冈市",
			Id:       "421100",
		},
		{
			Province: "湖北省",
			Name:     "咸宁市",
			Id:       "421200",
		},
		{
			Province: "湖北省",
			Name:     "随州市",
			Id:       "421300",
		},
		{
			Province: "湖北省",
			Name:     "恩施土家族苗族自治州",
			Id:       "422800",
		},
		{
			Province: "湖北省",
			Name:     "省直辖县级行政区划",
			Id:       "429000",
		},
	},
	"430000": {
		{
			Province: "湖南省",
			Name:     "长沙市",
			Id:       "430100",
		},
		{
			Province: "湖南省",
			Name:     "株洲市",
			Id:       "430200",
		},
		{
			Province: "湖南省",
			Name:     "湘潭市",
			Id:       "430300",
		},
		{
			Province: "湖南省",
			Name:     "衡阳市",
			Id:       "430400",
		},
		{
			Province: "湖南省",
			Name:     "邵阳市",
			Id:       "430500",
		},
		{
			Province: "湖南省",
			Name:     "岳阳市",
			Id:       "430600",
		},
		{
			Province: "湖南省",
			Name:     "常德市",
			Id:       "430700",
		},
		{
			Province: "湖南省",
			Name:     "张家界市",
			Id:       "430800",
		},
		{
			Province: "湖南省",
			Name:     "益阳市",
			Id:       "430900",
		},
		{
			Province: "湖南省",
			Name:     "郴州市",
			Id:       "431000",
		},
		{
			Province: "湖南省",
			Name:     "永州市",
			Id:       "431100",
		},
		{
			Province: "湖南省",
			Name:     "怀化市",
			Id:       "431200",
		},
		{
			Province: "湖南省",
			Name:     "娄底市",
			Id:       "431300",
		},
		{
			Province: "湖南省",
			Name:     "湘西土家族苗族自治州",
			Id:       "433100",
		},
	},
	"440000": {
		{
			Province: "广东省",
			Name:     "广州市",
			Id:       "440100",
		},
		{
			Province: "广东省",
			Name:     "韶关市",
			Id:       "440200",
		},
		{
			Province: "广东省",
			Name:     "深圳市",
			Id:       "440300",
		},
		{
			Province: "广东省",
			Name:     "珠海市",
			Id:       "440400",
		},
		{
			Province: "广东省",
			Name:     "汕头市",
			Id:       "440500",
		},
		{
			Province: "广东省",
			Name:     "佛山市",
			Id:       "440600",
		},
		{
			Province: "广东省",
			Name:     "江门市",
			Id:       "440700",
		},
		{
			Province: "广东省",
			Name:     "湛江市",
			Id:       "440800",
		},
		{
			Province: "广东省",
			Name:     "茂名市",
			Id:       "440900",
		},
		{
			Province: "广东省",
			Name:     "肇庆市",
			Id:       "441200",
		},
		{
			Province: "广东省",
			Name:     "惠州市",
			Id:       "441300",
		},
		{
			Province: "广东省",
			Name:     "梅州市",
			Id:       "441400",
		},
		{
			Province: "广东省",
			Name:     "汕尾市",
			Id:       "441500",
		},
		{
			Province: "广东省",
			Name:     "河源市",
			Id:       "441600",
		},
		{
			Province: "广东省",
			Name:     "阳江市",
			Id:       "441700",
		},
		{
			Province: "广东省",
			Name:     "清远市",
			Id:       "441800",
		},
		{
			Province: "广东省",
			Name:     "东莞市",
			Id:       "441900",
		},
		{
			Province: "广东省",
			Name:     "中山市",
			Id:       "442000",
		},
		{
			Province: "广东省",
			Name:     "潮州市",
			Id:       "445100",
		},
		{
			Province: "广东省",
			Name:     "揭阳市",
			Id:       "445200",
		},
		{
			Province: "广东省",
			Name:     "云浮市",
			Id:       "445300",
		},
	},
	"450000": {
		{
			Province: "广西壮族自治区",
			Name:     "南宁市",
			Id:       "450100",
		},
		{
			Province: "广西壮族自治区",
			Name:     "柳州市",
			Id:       "450200",
		},
		{
			Province: "广西壮族自治区",
			Name:     "桂林市",
			Id:       "450300",
		},
		{
			Province: "广西壮族自治区",
			Name:     "梧州市",
			Id:       "450400",
		},
		{
			Province: "广西壮族自治区",
			Name:     "北海市",
			Id:       "450500",
		},
		{
			Province: "广西壮族自治区",
			Name:     "防城港市",
			Id:       "450600",
		},
		{
			Province: "广西壮族自治区",
			Name:     "钦州市",
			Id:       "450700",
		},
		{
			Province: "广西壮族自治区",
			Name:     "贵港市",
			Id:       "450800",
		},
		{
			Province: "广西壮族自治区",
			Name:     "玉林市",
			Id:       "450900",
		},
		{
			Province: "广西壮族自治区",
			Name:     "百色市",
			Id:       "451000",
		},
		{
			Province: "广西壮族自治区",
			Name:     "贺州市",
			Id:       "451100",
		},
		{
			Province: "广西壮族自治区",
			Name:     "河池市",
			Id:       "451200",
		},
		{
			Province: "广西壮族自治区",
			Name:     "来宾市",
			Id:       "451300",
		},
		{
			Province: "广西壮族自治区",
			Name:     "崇左市",
			Id:       "451400",
		},
	},
	"460000": {
		{
			Province: "海南省",
			Name:     "海口市",
			Id:       "460100",
		},
		{
			Province: "海南省",
			Name:     "三亚市",
			Id:       "460200",
		},
		{
			Province: "海南省",
			Name:     "三沙市",
			Id:       "460300",
		},
		{
			Province: "海南省",
			Name:     "儋州市",
			Id:       "460400",
		},
		{
			Province: "海南省",
			Name:     "省直辖县级行政区划",
			Id:       "469000",
		},
	},
	"500000": {
		{
			Province: "重庆市",
			Name:     "市辖区",
			Id:       "500100",
		},
		{
			Province: "重庆市",
			Name:     "县",
			Id:       "500200",
		},
	},
	"510000": {
		{
			Province: "四川省",
			Name:     "成都市",
			Id:       "510100",
		},
		{
			Province: "四川省",
			Name:     "自贡市",
			Id:       "510300",
		},
		{
			Province: "四川省",
			Name:     "攀枝花市",
			Id:       "510400",
		},
		{
			Province: "四川省",
			Name:     "泸州市",
			Id:       "510500",
		},
		{
			Province: "四川省",
			Name:     "德阳市",
			Id:       "510600",
		},
		{
			Province: "四川省",
			Name:     "绵阳市",
			Id:       "510700",
		},
		{
			Province: "四川省",
			Name:     "广元市",
			Id:       "510800",
		},
		{
			Province: "四川省",
			Name:     "遂宁市",
			Id:       "510900",
		},
		{
			Province: "四川省",
			Name:     "内江市",
			Id:       "511000",
		},
		{
			Province: "四川省",
			Name:     "乐山市",
			Id:       "511100",
		},
		{
			Province: "四川省",
			Name:     "南充市",
			Id:       "511300",
		},
		{
			Province: "四川省",
			Name:     "眉山市",
			Id:       "511400",
		},
		{
			Province: "四川省",
			Name:     "宜宾市",
			Id:       "511500",
		},
		{
			Province: "四川省",
			Name:     "广安市",
			Id:       "511600",
		},
		{
			Province: "四川省",
			Name:     "达州市",
			Id:       "511700",
		},
		{
			Province: "四川省",
			Name:     "雅安市",
			Id:       "511800",
		},
		{
			Province: "四川省",
			Name:     "巴中市",
			Id:       "511900",
		},
		{
			Province: "四川省",
			Name:     "资阳市",
			Id:       "512000",
		},
		{
			Province: "四川省",
			Name:     "阿坝藏族羌族自治州",
			Id:       "513200",
		},
		{
			Province: "四川省",
			Name:     "甘孜藏族自治州",
			Id:       "513300",
		},
		{
			Province: "四川省",
			Name:     "凉山彝族自治州",
			Id:       "513400",
		},
	},
	"520000": {
		{
			Province: "贵州省",
			Name:     "贵阳市",
			Id:       "520100",
		},
		{
			Province: "贵州省",
			Name:     "六盘水市",
			Id:       "520200",
		},
		{
			Province: "贵州省",
			Name:     "遵义市",
			Id:       "520300",
		},
		{
			Province: "贵州省",
			Name:     "安顺市",
			Id:       "520400",
		},
		{
			Province: "贵州省",
			Name:     "毕节市",
			Id:       "520500",
		},
		{
			Province: "贵州省",
			Name:     "铜仁市",
			Id:       "520600",
		},
		{
			Province: "贵州省",
			Name:     "黔西南布依族苗族自治州",
			Id:       "522300",
		},
		{
			Province: "贵州省",
			Name:     "黔东南苗族侗族自治州",
			Id:       "522600",
		},
		{
			Province: "贵州省",
			Name:     "黔南布依族苗族自治州",
			Id:       "522700",
		},
	},
	"530000": {
		{
			Province: "云南省",
			Name:     "昆明市",
			Id:       "530100",
		},
		{
			Province: "云南省",
			Name:     "曲靖市",
			Id:       "530300",
		},
		{
			Province: "云南省",
			Name:     "玉溪市",
			Id:       "530400",
		},
		{
			Province: "云南省",
			Name:     "保山市",
			Id:       "530500",
		},
		{
			Province: "云南省",
			Name:     "昭通市",
			Id:       "530600",
		},
		{
			Province: "云南省",
			Name:     "丽江市",
			Id:       "530700",
		},
		{
			Province: "云南省",
			Name:     "普洱市",
			Id:       "530800",
		},
		{
			Province: "云南省",
			Name:     "临沧市",
			Id:       "530900",
		},
		{
			Province: "云南省",
			Name:     "楚雄彝族自治州",
			Id:       "532300",
		},
		{
			Province: "云南省",
			Name:     "红河哈尼族彝族自治州",
			Id:       "532500",
		},
		{
			Province: "云南省",
			Name:     "文山壮族苗族自治州",
			Id:       "532600",
		},
		{
			Province: "云南省",
			Name:     "西双版纳傣族自治州",
			Id:       "532800",
		},
		{
			Province: "云南省",
			Name:     "大理白族自治州",
			Id:       "532900",
		},
		{
			Province: "云南省",
			Name:     "德宏傣族景颇族自治州",
			Id:       "533100",
		},
		{
			Province: "云南省",
			Name:     "怒江傈僳族自治州",
			Id:       "533300",
		},
		{
			Province: "云南省",
			Name:     "迪庆藏族自治州",
			Id:       "533400",
		},
	},
	"540000": {
		{
			Province: "西藏自治区",
			Name:     "拉萨市",
			Id:       "540100",
		},
		{
			Province: "西藏自治区",
			Name:     "日喀则市",
			Id:       "540200",
		},
		{
			Province: "西藏自治区",
			Name:     "昌都市",
			Id:       "540300",
		},
		{
			Province: "西藏自治区",
			Name:     "林芝市",
			Id:       "540400",
		},
		{
			Province: "西藏自治区",
			Name:     "山南市",
			Id:       "540500",
		},
		{
			Province: "西藏自治区",
			Name:     "那曲地区",
			Id:       "542400",
		},
		{
			Province: "西藏自治区",
			Name:     "阿里地区",
			Id:       "542500",
		},
	},
	"610000": {
		{
			Province: "陕西省",
			Name:     "西安市",
			Id:       "610100",
		},
		{
			Province: "陕西省",
			Name:     "铜川市",
			Id:       "610200",
		},
		{
			Province: "陕西省",
			Name:     "宝鸡市",
			Id:       "610300",
		},
		{
			Province: "陕西省",
			Name:     "咸阳市",
			Id:       "610400",
		},
		{
			Province: "陕西省",
			Name:     "渭南市",
			Id:       "610500",
		},
		{
			Province: "陕西省",
			Name:     "延安市",
			Id:       "610600",
		},
		{
			Province: "陕西省",
			Name:     "汉中市",
			Id:       "610700",
		},
		{
			Province: "陕西省",
			Name:     "榆林市",
			Id:       "610800",
		},
		{
			Province: "陕西省",
			Name:     "安康市",
			Id:       "610900",
		},
		{
			Province: "陕西省",
			Name:     "商洛市",
			Id:       "611000",
		},
	},
	"620000": {
		{
			Province: "甘肃省",
			Name:     "兰州市",
			Id:       "620100",
		},
		{
			Province: "甘肃省",
			Name:     "嘉峪关市",
			Id:       "620200",
		},
		{
			Province: "甘肃省",
			Name:     "金昌市",
			Id:       "620300",
		},
		{
			Province: "甘肃省",
			Name:     "白银市",
			Id:       "620400",
		},
		{
			Province: "甘肃省",
			Name:     "天水市",
			Id:       "620500",
		},
		{
			Province: "甘肃省",
			Name:     "武威市",
			Id:       "620600",
		},
		{
			Province: "甘肃省",
			Name:     "张掖市",
			Id:       "620700",
		},
		{
			Province: "甘肃省",
			Name:     "平凉市",
			Id:       "620800",
		},
		{
			Province: "甘肃省",
			Name:     "酒泉市",
			Id:       "620900",
		},
		{
			Province: "甘肃省",
			Name:     "庆阳市",
			Id:       "621000",
		},
		{
			Province: "甘肃省",
			Name:     "定西市",
			Id:       "621100",
		},
		{
			Province: "甘肃省",
			Name:     "陇南市",
			Id:       "621200",
		},
		{
			Province: "甘肃省",
			Name:     "临夏回族自治州",
			Id:       "622900",
		},
		{
			Province: "甘肃省",
			Name:     "甘南藏族自治州",
			Id:       "623000",
		},
	},
	"630000": {
		{
			Province: "青海省",
			Name:     "西宁市",
			Id:       "630100",
		},
		{
			Province: "青海省",
			Name:     "海东市",
			Id:       "630200",
		},
		{
			Province: "青海省",
			Name:     "海北藏族自治州",
			Id:       "632200",
		},
		{
			Province: "青海省",
			Name:     "黄南藏族自治州",
			Id:       "632300",
		},
		{
			Province: "青海省",
			Name:     "海南藏族自治州",
			Id:       "632500",
		},
		{
			Province: "青海省",
			Name:     "果洛藏族自治州",
			Id:       "632600",
		},
		{
			Province: "青海省",
			Name:     "玉树藏族自治州",
			Id:       "632700",
		},
		{
			Province: "青海省",
			Name:     "海西蒙古族藏族自治州",
			Id:       "632800",
		},
	},
	"640000": {
		{
			Province: "宁夏回族自治区",
			Name:     "银川市",
			Id:       "640100",
		},
		{
			Province: "宁夏回族自治区",
			Name:     "石嘴山市",
			Id:       "640200",
		},
		{
			Province: "宁夏回族自治区",
			Name:     "吴忠市",
			Id:       "640300",
		},
		{
			Province: "宁夏回族自治区",
			Name:     "固原市",
			Id:       "640400",
		},
		{
			Province: "宁夏回族自治区",
			Name:     "中卫市",
			Id:       "640500",
		},
	},
	"650000": {
		{
			Province: "新疆维吾尔自治区",
			Name:     "乌鲁木齐市",
			Id:       "650100",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "克拉玛依市",
			Id:       "650200",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "吐鲁番市",
			Id:       "650400",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "哈密市",
			Id:       "650500",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "昌吉回族自治州",
			Id:       "652300",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "博尔塔拉蒙古自治州",
			Id:       "652700",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "巴音郭楞蒙古自治州",
			Id:       "652800",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "阿克苏地区",
			Id:       "652900",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "克孜勒苏柯尔克孜自治州",
			Id:       "653000",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "喀什地区",
			Id:       "653100",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "和田地区",
			Id:       "653200",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "伊犁哈萨克自治州",
			Id:       "654000",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "塔城地区",
			Id:       "654200",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "阿勒泰地区",
			Id:       "654300",
		},
		{
			Province: "新疆维吾尔自治区",
			Name:     "自治区直辖县级行政区划",
			Id:       "659000",
		},
	},
}

type City struct {
	Province string `json:"province"`
	Name     string `json:"name"`
	Id       string `json:"id"`
}
