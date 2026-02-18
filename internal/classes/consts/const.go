package consts

import (
	"github.com/kafe523/kancolatex/internal/classes/consts/abmode"
	"github.com/kafe523/kancolatex/internal/classes/consts/airstate"
	"github.com/kafe523/kancolatex/internal/classes/consts/celltype"
	"github.com/kafe523/kancolatex/internal/classes/consts/difficulty"
	"github.com/kafe523/kancolatex/internal/classes/consts/formation"
	"github.com/kafe523/kancolatex/internal/classes/consts/shiptype"
	"github.com/kafe523/kancolatex/internal/classes/consts/supporttype"
)

type Formation struct {
	Text       string
	Value      formation.Formation
	Correction float64
}

type AvoidType struct {
	Value int32
	Text  string
	C1    float64
	C2    float64
	C3    float64
	C4    float64
}

// 基地航空隊札一覧
var AIR_STATUS = [...]struct {
	Text  string
	Value airstate.AirState
	Color string
}{
	{"確保", airstate.KAKUHO, ""},
	{"優勢", airstate.YUSEI, "light-green"},
	{"拮抗", airstate.KINKO, "yellow"},
	{"劣勢", airstate.RESSEI, "deep-orange"},
	{"喪失", airstate.SOSHITSU, "red"},
	{"不発", airstate.NONE, "secondary"},
}

// 基地航空隊札一覧
var AB_MODE_ITEM = [...]struct {
	Text  string
	Value abmode.ABMode
}{
	{"出撃", abmode.BATTLE},
	{"防空", abmode.DEFENSE},
	{"待機", abmode.WAIT},
}

// 航空戦に関係する装備カテゴ
var PLANE_TYPES = [...]int32{6, 7, 8, 9, 10, 11, 25, 26, 41, 45, 47, 48, 49, 53, 57}

// 空母のみ搭載可である機体カテゴリ
var CB_PLANE_TYPES = [...]int32{6, 7, 8, 9, 57}

// 水上機
var SP_PLANE_TYPES = [...]int32{10, 11, 41, 45}

// 基地航空隊のみ搭載可である機体カテゴリ
var AB_PLANE_TYPES = [...]int32{47, 48, 49, 53}

// 艦戦カテゴリ
var FIGHTERS = [...]int32{6, 45, 48}

// 攻撃機カテゴリ
var ATTACKERS = [...]int32{7, 8, 11, 47, 53, 57}

// 対潜哨戒機 オートジャイロ
var ASW_PLANES = [...]int32{25, 26}

// 偵察機カテゴリ
var RECONNAISSANCES = [...]int32{9, 10, 41, 49}

// 陸上攻撃機カテゴリ
var AB_ATTACKERS = [...]int32{47, 53}

// 大型陸上機カテゴリ
var AB_ATTACKERS_LARGE = [...]int32{53}

// 狭義の爆雷カテゴリ
var STRICT_DEPTH_CHARGE = [...]int32{226, 227, 378, 439, 488}

// ロケット戦闘機id
var ROCKET = [...]int32{350, 351, 352}

// 爆戦id
var BAKUSEN = [...]int32{60, 154, 219, 447, 487}

// 対地攻撃可能な艦爆id
var ENABLED_LAND_BASE_ATTACK = [...]int32{64, 148, 233, 277, 305, 306, 319, 320, 391, 392, 420, 421, 474, 552}

// 対潜支援可能艦種
var ENABLED_ASW_SUPPORT = [...]int32{7, 8, 10, 11, 45, 41, 25, 26}

// 潜水艦後期型魚雷
var LATE_MODEL_TORPEDO = [...]int32{213, 214, 383, 441, 443, 457, 461, 512}

// 特効情報
var SPECIAL_GROUP = [...]struct {
	Key        string
	Text       string
	Items      []int32
	IsOnlyAB   bool
	IsOnlyShip bool
}{
	//Skip
}

// 海域特効装備
var AIRBASE_MAP_BONUSES = [...]struct {
	Area  int32
	Node  string
	Items []int32
	Bonus float64
	multi bool
}{
	//Skip
}

var ITEM_API_TYPE = [...]struct {
	Id      int32
	Name    string
	SortKey []string
}{
	{6, "艦上戦闘機", []string{"antiAir", "avoid", "accuracy", "scout", "radius", "cost"}},
	{7, "艦上爆撃機", []string{"bomber", "antiAir", "accuracy", "asw", "avoid", "scout", "radius", "avoidId", "cost"}},
	{8, "艦上攻撃機", []string{"torpedo", "antiAir", "accuracy", "asw", "avoid", "scout", "radius", "avoidId", "cost"}},
	{9, "艦上偵察機", []string{"scout", "accuracy", "avoid", "antiAir", "radius", "avoidId", "cost"}},
	{57, "噴式戦闘爆撃機", []string{"bomber", "antiAir", "accuracy", "avoid", "scout", "radius", "avoidId", "cost"}},
	{10, "水上偵察機", []string{"scout", "accuracy", "avoid", "antiAir", "radius", "avoidId", "cost"}},
	{11, "水上爆撃機", []string{"bomber", "antiAir", "accuracy", "avoid", "scout", "radius", "avoidId", "cost"}},
	{45, "水上戦闘機", []string{"antiAir", "avoid", "accuracy", "scout", "radius", "cost"}},
	{41, "大型飛行艇", []string{"scout", "accuracy", "avoid", "antiAir", "radius", "avoidId", "cost"}},
	{47, "陸上攻撃機", []string{"torpedo", "bomber", "antiAir", "asw", "accuracy", "avoid", "scout", "radius", "avoidId", "cost"}},
	{53, "大型陸上機", []string{"torpedo", "bomber", "antiAir", "accuracy", "avoid", "scout", "radius", "avoidId", "cost"}},
	{48, "局地/陸軍戦闘機", []string{"antiAir", "antiBomber", "interception", "radius", "cost", "sortieAntiAir", "defenseAntiAir"}},
	{49, "陸上偵察機", []string{"scout", "accuracy", "avoid", "antiAir", "radius", "avoidId", "cost"}},
	{1, "小口径主砲", []string{"fire", "accuracy", "antiAir", "armor", "range"}},
	{2, "中口径主砲", []string{"fire", "accuracy", "antiAir", "armor", "range"}},
	{3, "大口径主砲", []string{"fire", "accuracy", "antiAir", "armor", "range"}},
	{4, "副砲", []string{"fire", "accuracy", "antiAir", "armor", "range"}},
	{5, "魚雷", []string{"torpedo", "accuracy", "armor"}},
	{12, "小型電探", []string{"antiAir", "accuracy", "scout", "fire", "armor"}},
	{13, "大型電探", []string{"antiAir", "accuracy", "scout", "fire", "armor"}},
	{14, "ソナー", []string{"asw", "accuracy", "scout"}},
	{15, "爆雷", []string{"asw", "accuracy"}},
	{17, "機関部強化", []string{}},
	{18, "対空強化弾", []string{}},
	{19, "対艦強化弾", []string{}},
	{21, "対空機銃", []string{"antiAir", "accuracy", "fire", "armor", "avoid"}},
	{22, "特殊潜航艇", []string{"torpedo", "accuracy", "avoid", "scout"}},
	{23, "応急修理要員", []string{}},
	{24, "上陸用舟艇", []string{"fire", "antiAir", "armor", "avoid", "tp", "tp2"}},
	{25, "回転翼機", []string{}},
	{26, "対潜哨戒機", []string{}},
	{27, "追加装甲(中型)", []string{}},
	{28, "追加装甲(大型)", []string{}},
	{29, "探照灯", []string{}},
	{30, "簡易輸送部材", []string{}},
	{31, "艦艇修理施設", []string{}},
	{32, "潜水艦魚雷", []string{"torpedo", "accuracy", "avoid"}},
	{33, "照明弾", []string{}},
	{34, "司令部施設", []string{}},
	{35, "航空要員", []string{}},
	{36, "高射装置", []string{}},
	{37, "対地装備", []string{}},
	{39, "水上艦要員", []string{}},
	{40, "大型ソナー", []string{}},
	{42, "大型探照灯", []string{}},
	{43, "戦闘糧食", []string{}},
	{44, "補給物資", []string{}},
	{46, "特型内火艇", []string{}},
	{50, "輸送機材", []string{}},
	{51, "潜水艦装備", []string{}},
	{52, "陸戦部隊", []string{}},
	{54, "艦載発煙装置", []string{}},
	{55, "防空気球", []string{}},
}

var SHIP_TYPES_ALT_INFO = [...]struct {
	Id   int32
	Name string
}{
	{6, "金剛型"},
	{26, "扶桑型"},
	{2, "伊勢型"},
	{19, "長門型"},
	{37, "大和型"},
	{14, "赤城型"},
	{3, "加賀型"},
	{17, "蒼龍型"},
	{25, "飛龍型"},
	{33, "翔鶴型"},
	{53, "雲龍型"},
	{43, "大鳳型"},
	{27, "鳳翔型"},
	{32, "龍驤型"},
	{51, "龍鳳型"},
	{11, "祥鳳型"},
	{24, "飛鷹型"},
	{15, "千歳型"},
	{76, "大鷹型"},
	{75, "春日丸級"},
	{72, "神威型"},
	{62, "瑞穂型"},
	{90, "日進型"},
	{59, "秋津洲型"},
	{100, "迅鯨型"},
	{50, "大鯨型"},
	{60, "改風早型"},
	{49, "工作艦"},
	{111, "耐氷型雑用運送艦"},
	{7, "古鷹型"},
	{13, "青葉型"},
	{29, "妙高型"},
	{8, "高雄型"},
	{9, "最上型"},
	{31, "利根型"},
	{21, "天龍型"},
	{4, "球磨型"},
	{20, "長良型"},
	{16, "川内型"},
	{34, "夕張型"},
	{41, "阿賀野型"},
	{52, "大淀型"},
	{56, "香取型"},
	{66, "神風型"},
	{28, "睦月型"},
	{12, "吹雪型"},
	{1, "綾波型"},
	{5, "暁型"},
	{10, "初春型"},
	{23, "白露型"},
	{18, "朝潮型"},
	{30, "陽炎型"},
	{38, "夕雲型"},
	{54, "秋月型"},
	{22, "島風型"},
	{101, "松型"},
	{74, "占守型"},
	{77, "択捉型"},
	{94, "御蔵型"},
	{85, "日振型"},
	{104, "丁型海防艦"},
	{117, "鵜来型"},
	{71, "巡潜甲型改二"},
	{44, "潜特型(伊400型潜水艦)"},
	{35, "海大VI型"},
	{40, "巡潜3型"},
	{39, "巡潜乙型"},
	{103, "巡潜丙型"},
	{36, "巡潜乙型改二"},
	{109, "潜高型"},
	{57, "UボートIXC型"},
	{86, "呂号潜水艦"},
	{46, "三式潜航輸送艇"},
	{97, "陸軍特種船(R1)"},
	{45, "特種船丙型"},
	{115, "特2TL型"},
	{119, "特種船M丙型"},
	{120, "二等輸送艦"},
	{123, "敷島型"},
	{126, "改氷川丸級"},
	{127, "巡潜乙型改一"},
	{130, "大泊型"},
	{132, "特1TL型"},
	{136, "野埼型"},

	{47, "Bismarck級"},
	{55, "Admiral Hipper級"},
	{63, "Graf Zeppelin級"},
	{48, "Z1型"},

	{58, "V.Veneto級"},
	{68, "Aquila級"},
	{64, "Zara級"},
	{92, "L.d.S.D.d.Abruzzi級"},
	{61, "Maestrale級"},
	{113, "Conte di Cavour級"},
	{80, "Guglielmo Marconi級"},
	{124, "Marcello級"},

	{93, "Colorado級"},
	{107, "North Carolina級"},
	{102, "South Dakota級"},
	{65, "Iowa級"},
	{69, "Lexington級"},
	{105, "Yorktown級"},
	{84, "Essex級"},
	{83, "Casablanca級"},
	{95, "Northampton級"},
	{110, "Brooklyn級"},
	{106, "St.Louis級"},
	{99, "Atlanta級"},
	{91, "Fletcher級"},
	{87, "John C.Butler級"},
	{114, "Gato級"},
	{116, "Independence級"},
	{118, "Ranger級"},
	{121, "New Orleans級"},
	{122, "Salmon級"},
	{125, "Nevada級"},

	{67, "Queen Elizabeth級"},
	{88, "Nelson級"},
	{78, "Ark Royal級"},
	{112, "Illustrious級"},
	{108, "Town級"},
	{82, "J級"},
	{134, "Courageous級"},
	{135, "Glorious級"},

	{79, "Richelieu級"},
	{70, "C.Teste級"},
	{128, "La Galissonnière級"},
	{129, "Mogador級"},

	{73, "Гангут級"},
	{81, "Ташкент級"},
	{131, "Киров級"},

	{89, "Gotland級"},

	{98, "De Ruyter級"},

	{96, "Perth級"},

	{133, "Norge級"},

	{137, "Thonburi級"},

	{42, "(霧の艦隊?)"},
}

// 補強増設に搭載可能
// そうそうかわんないけど、定期的に api_mst_equip_exslot チェックして更新
var EXPANDED_ITEM_TYPE = [...]int32{16, 21, 23, 27, 28, 36, 39, 43, 44}

// 特定の艦娘が特定スロットに装備『できない！』やつ 99で増設
var FORBIDDEN_LINK_SHIP_ITEM = [...]struct {
	ShipId   int32
	Index    []int32
	ItemType []int32
	ItemIds  []int32
}{
	// 伊勢改二 3，4，5スロットに 主砲系
	{
		553, []int32{3, 4, 5}, []int32{2, 3}, []int32{0},
	},
	// 日向改二 3，4，5スロットに 主砲系
	{
		554, []int32{3, 4, 5}, []int32{2, 3}, []int32{0},
	},
	// 夕張改二 4スロットに 主砲系 魚雷系 不可
	{
		622, []int32{4}, []int32{1, 2, 5}, []int32{0},
	},
	// 夕張改二特 4スロットに 主砲系 魚雷系 不可
	{
		623, []int32{4}, []int32{1, 2, 5, 22}, []int32{0},
	},
	// 夕張改二丁 4スロットに 主砲系 魚雷系 不可
	{
		624, []int32{4}, []int32{1, 2, 5}, []int32{0},
	},
	// 夕張改二 5スロットに いろいろ装備不可
	{
		622, []int32{5}, []int32{1, 2, 5, 13, 14, 15, 17, 20, 22, 23, 24, 27, 29, 30, 33, 34, 36, 37, 39, 40, 46, 54}, []int32{0},
	},
	// 夕張改二特 5スロットに いろいろ装備不可
	{
		623, []int32{5}, []int32{1, 2, 5, 14, 15, 17, 20, 22, 23, 24, 27, 29, 30, 33, 34, 36, 37, 39, 40, 46, 54}, []int32{0},
	},
	// 夕張改二丁 5スロットに いろいろ装備不可
	{
		624, []int32{5}, []int32{1, 2, 5, 14, 15, 17, 20, 22, 23, 24, 27, 29, 30, 33, 34, 36, 37, 39, 40, 46, 54}, []int32{0},
	},
	// 能代改二 4スロットに 魚雷系 不可
	{
		662, []int32{4}, []int32{5, 22}, []int32{0},
	},
	// 矢矧改二 4スロットに 魚雷系 不可
	{
		663, []int32{4}, []int32{5, 22}, []int32{0},
	},
	// 矢矧改二乙 4スロットに 魚雷系 不可
	{
		668, []int32{4}, []int32{5}, []int32{0},
	},
	// 初月改二 4スロットに 主砲、魚雷、大型電探系 不可
	{
		968, []int32{4}, []int32{1, 5, 13}, []int32{0},
	},
	// 秋月改二 4スロットに 主砲、魚雷、大型電探系 不可
	{
		963, []int32{4}, []int32{1, 5, 13}, []int32{0},
	},
}

// 艦種一覧 省略系と含む艦種
var SHIP_TYPES_ALT = [...]struct {
	Text  string
	Types []int32
}{
	{"空母", []int32{11, 18}},
	{"軽母", []int32{7}},
	{"戦艦", []int32{8, 9}},
	{"航戦", []int32{10}},
	{"重巡", []int32{5}},
	{"航巡", []int32{6}},
	{"軽巡", []int32{3, 4, 21}},
	{"駆逐", []int32{2}},
	{"海防", []int32{1}},
	{"潜水", []int32{13, 14}},
	{"水母", []int32{16}},
	{"その他", []int32{15, 17, 19, 20, 22}},
}

// 艦種一覧 省略系と含む艦種 艦隊分析用
var SHIP_TYPES_ALT2 = [...]struct {
	Text  string
	Types []int32
}{
	// Skip
}

// 艦種一覧 省略形の一覧
var SHIP_TYPES_ALT3 = [...]struct {
	Text string
	Type int32
}{
	{"戦艦", 9},
	{"航戦", 10},
	{"高戦", 8},
	{"空母", 11},
	{"装空", 18},
	{"軽母", 7},
	{"重巡", 5},
	{"航巡", 6},
	{"雷巡", 4},
	{"軽巡", 3},
	{"練巡", 21},
	{"水母", 16},
	{"補給", 15},
	{"補給", 22},
	{"揚陸", 17},
	{"工作", 19},
	{"潜母", 20},
	{"潜水", 13},
	{"潜水", 14},
	{"駆逐", 2},
	{"海防", 1},
}

// 艦種一覧 正式名称の一覧
var SHIP_TYPES_FORMAL = [...]struct {
	Text string
	Type shiptype.ShipType
}{
	{"海防艦", shiptype.DE},
	{"駆逐艦", shiptype.DD},
	{"軽巡", shiptype.CL},
	{"雷巡", shiptype.CLT},
	{"重巡", shiptype.CA},
	{"航巡", shiptype.CAV},
	{"軽空母", shiptype.CVL},
	{"高速戦艦", shiptype.FBB},
	{"戦艦", shiptype.BB},
	{"航空戦艦", shiptype.BBV},
	{"正規空母", shiptype.CV},
	{"潜水艦", shiptype.SS},
	{"潜水空母", shiptype.SSV},
	{"補給", shiptype.AO_2},
	{"水母", shiptype.AV},
	{"揚陸艦", shiptype.LHA},
	{"装甲空母", shiptype.CVB},
	{"工作艦", shiptype.AR},
	{"潜水母艦", shiptype.AS},
	{"練巡", shiptype.CT},
	{"補給", shiptype.AO},
}

var JPN = [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 49, 50, 51, 52, 53, 54, 56, 59, 60, 62, 66, 71, 72, 74, 75, 76, 77, 85, 86, 90, 94, 97, 100, 101, 103, 104, 109, 111, 115, 117, 119, 120, 123, 126, 127, 130, 132, 136}
var USA = [...]int32{65, 69, 83, 84, 87, 91, 93, 95, 99, 102, 105, 106, 107, 110, 114, 116, 118, 121, 122, 125}
var ITA = [...]int32{58, 61, 64, 68, 80, 92, 113, 124}
var GBR = [...]int32{67, 78, 82, 88, 108, 112, 134, 135}
var DEU = [...]int32{47, 48, 55, 57, 63}
var FRA = [...]int32{70, 79, 128, 129}
var RUS = [...]int32{73, 81, 131}
var SWE = [...]int32{89}
var AUS = [...]int32{96}
var NLD = [...]int32{98}
var NOR = [...]int32{133}
var THA = [...]int32{137}

// 陣形一覧
var FORMATIONS = [...]Formation{
	{"単縦陣", formation.LINE_AHEAD, 1.0},
	{"複縦陣", formation.DOUBLE_LINE, 1.2},
	{"輪形陣", formation.DIAMOND, 1.6},
	{"梯形陣", formation.ECHELON, 1.0},
	{"単横陣", formation.LINE_ABREAST, 1.0},
	{"警戒陣", formation.VANGUARD, 1.1},
	{"第一警戒", formation.FORMATION1, 1.1},
	{"第二警戒", formation.FORMATION2, 1.0},
	{"第三警戒", formation.FORMATION3, 1.5},
	{"第四警戒", formation.FORMATION4, 1.0},
}

// 支援一覧
var SUPPORTS = [...]struct {
	Text  string
	Value supporttype.SupportType
}{
	{"支援射撃", supporttype.SHELLING},
	{"航空支援", supporttype.AIRSTRIKE},
	{"対潜支援哨戒", supporttype.ANTI_SUBMARINE},
	{"支援長距離雷撃", supporttype.LONG_RANGE_TORPEDO},
	{"支援不可(要駆逐2)", supporttype.NOT_FOUND_DD},
	{"不発", supporttype.NONE},
}

// 装備一覧 ちょっとまとめたやつと、装備一覧表示のステータス
var ITEM_TYPES_ALT = [...]struct {
	Id         int32
	Text       string
	ViewStatus []string
	Types      []int32
}{
	{
		1, "小口径主砲", []string{"actualFire", "antiAir", "actualAccuracy", "nightBattleFirePower", "antiAirWeight", "antiAirBonus"}, []int32{1},
	},
	{
		2, "中口径主砲", []string{"actualFire", "antiAir", "actualAccuracy", "nightBattleFirePower", "antiAirWeight", "antiAirBonus"}, []int32{2},
	},
	{
		3, "大口径主砲", []string{"actualFire", "antiAir", "actualAccuracy", "nightBattleFirePower", "antiAirBonus", "actualRange"}, []int32{3},
	},
	{
		5, "魚雷", []string{"actualTorpedo", "nightBattleFirePower", "actualFire", "actualAccuracy", "actualAvoid", "actualArmor"}, []int32{5, 22, 32},
	},
	{
		6, "艦戦", []string{"antiAir", "actualAntiAir", "actualAccuracy", "actualAvoid", "radius", "airPower"}, []int32{6},
	},
	{
		7, "艦爆", []string{"actualBomber", "actualAntiAir", "actualAccuracy", "actualAsw", "radius", "avoidId"}, []int32{7},
	},
	{
		8, "艦攻", []string{"actualTorpedo", "actualAntiAir", "actualAccuracy", "actualAsw", "radius", "avoidId"}, []int32{8},
	},
	{
		9, "艦偵", []string{"actualScout", "actualFire", "actualAccuracy", "actualAntiAir", "radius", "cost"}, []int32{9},
	},
	{
		57, "噴式機", []string{"actualBomber", "actualAccuracy", "actualAntiAir", "avoidId", "airPower", "cost"}, []int32{57},
	},
	{
		10, "水偵", []string{"actualScout", "actualFire", "actualAccuracy", "actualAsw", "radius", "cost"}, []int32{10},
	},
	{
		1100, "水爆", []string{"actualBomber", "actualAntiAir", "actualAccuracy", "actualAsw", "radius", "avoidId"}, []int32{11},
	},
	{
		4500, "水戦", []string{"antiAir", "actualAntiAir", "actualScout", "actualAccuracy", "radius", "airPower"}, []int32{45},
	},
	{
		41, "大型飛行艇", []string{"actualScout", "actualAccuracy", "actualAsw", "radius", "cost"}, []int32{41},
	},
	{
		12, "電探", []string{"actualAccuracy", "actualScout", "antiAir", "nightBattleFirePower", "antiAirWeight", "antiAirBonus"}, []int32{12, 13},
	},
	{
		14, "対潜装備", []string{"actualAsw", "actualAccuracy", "actualArmor", "actualAvoid", "actualScout"}, []int32{14, 15, 40},
	},
	{
		4, "副砲", []string{"actualFire", "antiAir", "actualAccuracy", "nightBattleFirePower", "antiAirWeight", "antiAirBonus"}, []int32{4},
	},
	{
		21, "機銃", []string{"antiAir", "actualFire", "actualAccuracy", "actualArmor", "antiAirWeight", "antiAirBonus"}, []int32{21},
	},
	{
		24, "上陸用舟艇", []string{"actualFire", "antiAir", "actualScout", "actualAvoid", "tp", "tp2"}, []int32{24, 30, 46, 52},
	},
	{
		47, "陸攻", []string{"actualTorpedo", "actualBomber", "actualAntiAir", "radius", "cost", "avoidId"}, []int32{47, 53},
	},
	{
		48, "局戦", []string{"actualAntiAir", "actualDefenseAntiAir", "antiBomber", "radius", "airPower", "defenseAirPower"}, []int32{48},
	},
	{
		49, "陸偵", []string{"actualAntiAir", "actualScout", "actualAccuracy", "radius", "actualArmor", "cost"}, []int32{49},
	},
	{
		17, "その他", []string{"radius", "nightBattleFirePower", "actualAccuracy", "actualScout", "actualAsw", "actualArmor"}, []int32{17, 18, 19, 23, 25, 26, 27, 28, 29, 31, 33, 34, 35, 36, 37, 39, 42, 43, 44, 50, 51, 54, 55},
	},
}

// 戦闘マス形式
var CELL_TYPE = [...]struct {
	Text  string
	Value celltype.CellType
}{
	{"通常", celltype.NORMAL},
	{"連合", celltype.GRAND},
	{"空襲", celltype.AIR_RAID},
	{"夜戦", celltype.NIGHT},
	{"重爆", celltype.HIGH_AIR_RAID},
	{"航空戦", celltype.AERIAL_COMBAT},
	{"超重爆", celltype.SUPER_HIGH_AIR_RAID},
	{"対潜空襲", celltype.AIR_SUPPORTED_ASW},
	{"レーダー", celltype.AIR_SUPPORTED_ASW},
}

var DIFFICULTY_LEVELS = [...]struct {
	Text  string
	Value difficulty.Difficulty
}{
	{"甲", difficulty.HARD},
	{"乙", difficulty.MEDIUM},
	{"丙", difficulty.EASY},
	{"丁", difficulty.CASUAL},
}

// 艦載機熟練度ボーダー
var PROF_LEVEL_BORDER = [...]int32{0, 10, 25, 40, 55, 70, 85, 100, 120}

// 対空射撃回避 任意
const MANUAL_AVOID int32 = 99

// 補強増設を識別するindex
const EXPAND_SLOT_INDEX int32 = 99

// 実装最大レベル
const MAX_LEVEL = 185

// 対空射撃回避
var AVOID_TYPE = [...]AvoidType{
	{
		0, "なし", 1, 1, 1, 1,
	},
	{
		1, "弱", 0.6, 1, 1, 1,
	},
	{
		2, "中", 0.6, 0.7, 0.6, 1,
	},
	{
		3, "強", 0.5, 0.7, 0.4, 0.5,
	},
	{
		4, "超", 0.5, 0.5, 0.4, 0.5,
	},
	{
		5, "超+", 0.4, 0.4, 0.4, 0.5,
	},
	{
		MANUAL_AVOID, "任意", 1, 1, 1, 1,
	},
}

// 対空CI優先度配列
var ANTI_AIR_CUT_IN_PRIORITIES = []int32{38, 39, 40, 42, 41, 10, 43, 46, 11, 25, 48, 1, 34, 44, 26, 4, 2, 35, 36, 27, 45, 50, 49, 51, 52, 19, 21, 29, 16, 14, 3, 5, 6, 28, 37, 33, 30, 8, 13, 15, 7, 20, 24, 32, 12, 31, 47, 17, 18, 22, 9, 23}

// 対空CI
var ANTI_AIR_CUTIN = [...]struct {
	Id        int32
	Text      string
	RateBonus float64
	C1        float64
	C2        float64
	Rate      int32
	Remarks   string
}{
	{
		0, "不発", 1, 1, 0, 100, "",
	},
	{
		1, "1種", 1.7, 3, 5, 65, "秋月型",
	},
	{
		2, "2種", 1.7, 3, 4, 55, "秋月型",
	},
	{
		3, "3種", 1.6, 2, 3, 50, "秋月型",
	},
	{
		4, "4種", 1.5, 5, 2, 52, "戦艦",
	},
	{
		5, "5種", 1.5, 2, 3, 50, "汎用",
	},
	{
		6, "6種", 1.45, 4, 1, 40, "戦艦",
	},
	{
		7, "7種", 1.35, 2, 2, 45, "汎用",
	},
	{
		8, "8種", 1.4, 2, 3, 50, "汎用",
	},
	{
		9, "9種", 1.3, 1, 2, 40, "汎用",
	},
	{
		10, "10種", 1.65, 3, 6, 60, "摩耶改二",
	},
	{
		11, "11種", 1.5, 2, 5, 55, "摩耶改二",
	},
	{
		12, "12種", 1.25, 1, 3, 45, "汎用",
	},
	{
		13, "13種", 1.35, 1, 4, 35, "汎用",
	},
	{
		14, "14種", 1.45, 4, 1, 63, "五十鈴改二",
	},
	{
		15, "15種", 1.3, 3, 1, 54, "五十鈴改二",
	},
	{
		16, "16種", 1.4, 4, 1, 62, "霞改二乙 / 夕張改二",
	},
	{
		17, "17種", 1.25, 2, 1, 57, "霞改二乙 / 稲木改二",
	},
	{
		18, "18種", 1.2, 2, 1, 59, "皐月改二 / 稲木改二",
	},
	{
		19, "19種", 1.45, 5, 1, 60, "鬼怒改二",
	},
	{
		20, "20種", 1.25, 3, 1, 65, "鬼怒改二",
	},
	{
		21, "21種", 1.45, 5, 1, 60, "由良改二",
	},
	{
		22, "22種", 1.2, 2, 1, 65, "文月改二",
	},
	{
		23, "23種", 1.05, 1, 1, 80, "UIT-25 / 伊504",
	},
	{
		24, "24種", 1.25, 3, 1, 62, "天龍型改二",
	},
	{
		25, "25種", 1.55, 7, 1, 60, "伊勢型",
	},
	{
		26, "26種", 1.4, 6, 1, 60, "武蔵改二",
	},
	{
		27, "27種", 1.55, 6, 1, 55, "大淀",
	},
	{
		28, "28種", 1.4, 4, 1, 56, "伊勢型 / 武蔵",
	},
	{
		29, "29種", 1.55, 5, 1, 60, "磯風乙改 / 浜風乙改",
	},
	{
		30, "30種", 1.3, 3, 1, 50, "天龍改二",
	},
	{
		31, "31種", 1.25, 2, 1, 50, "天龍改二 / 稲木改二",
	},
	{
		32, "32種", 1.2, 3, 1, 60, "金剛型改二 / 英艦",
	},
	{
		33, "33種", 1.35, 3, 1, 42, "Gotland",
	},
	{
		34, "34種", 1.6, 7, 1, 56, "Fletcher級",
	},
	{
		35, "35種", 1.55, 6, 1, 55, "Fletcher級",
	},
	{
		36, "36種", 1.55, 6, 1, 50, "Fletcher級",
	},
	{
		37, "37種", 1.45, 2, 3, 44, "Fletcher級",
	},
	{
		38, "38種", 1.85, 6, 5, 60, "Atlanta",
	},
	{
		39, "39種", 1.7, 6, 5, 60, "Atlanta",
	},
	{
		40, "40種", 1.7, 6, 5, 60, "Atlanta",
	},
	{
		41, "41種", 1.65, 5, 5, 60, "Atlanta",
	},
	{
		42, "42種", 1.7, 10, 1, 65, "大和型改二",
	},
	{
		43, "43種", 1.6, 8, 1, 60, "大和型改二",
	},
	{
		44, "44種", 1.6, 6, 1, 55, "大和型改二",
	},
	{
		45, "45種", 1.55, 6, 1, 50, "大和型改二",
	},
	{
		46, "46種", 1.55, 8, 1, 50, "榛名改二乙",
	},
	{
		47, "47種", 1.3, 2, 1, 70, "白露型改二",
	},
	{
		48, "48種", 1.75, 8, 1, 65, "秋月型",
	},
	{
		49, "49種", 1.5, 6, 1, 50, "藤波改二 / 吹雪改二 / 白雪改二",
	},
	{
		50, "50種", 1.5, 7, 1, 50, "吹雪改二 / 白雪改二 / 秋月型",
	},
	{
		51, "51種", 1.4, 5, 1, 50, "吹雪改二 / 白雪改二",
	},
	{
		52, "52種", 1.4, 5, 1, 50, "吹雪改二 / 白雪改二",
	},
}

// 経験値最低値ボーダー
var LEVEL_BORDERS = [...]struct {
	Lv  int32
	Req int32
}{
	{185, 16000000},
	{184, 15000000},
	{183, 14200000},
	{182, 13600000},
	{181, 13200000},
	{180, 13000000},
	{179, 12100000},
	{178, 11600000},
	{177, 11300000},
	{176, 11100000},
	{175, 10950000},
	{174, 10266000},
	{173, 9705000},
	{172, 9248000},
	{171, 8875000},
	{170, 8580000},
	{169, 8350000},
	{168, 8172000},
	{167, 8033000},
	{166, 7920000},
	{165, 7820000},
	{164, 7320000},
	{163, 6910000},
	{162, 6580000},
	{161, 6320000},
	{160, 6120000},
	{159, 5970000},
	{158, 5860000},
	{157, 5780000},
	{156, 5720000},
	{155, 5470000},
	{154, 5230000},
	{153, 4999000},
	{152, 4777000},
	{151, 4564000},
	{150, 4360000},
	{149, 4165000},
	{148, 3978000},
	{147, 3799000},
	{146, 3628000},
	{145, 3465000},
	{144, 3310000},
	{143, 3162000},
	{142, 3021000},
	{141, 2887000},
	{140, 2760000},
	{139, 2640000},
	{138, 2525000},
	{137, 2415000},
	{136, 2310000},
	{135, 2210000},
	{134, 2115000},
	{133, 2025000},
	{132, 1940000},
	{131, 1860000},
	{130, 1785000},
	{129, 1714000},
	{128, 1647000},
	{127, 1584000},
	{126, 1525000},
	{125, 1470000},
	{124, 1419000},
	{123, 1372000},
	{122, 1329000},
	{121, 1290000},
	{120, 1255000},
	{119, 1223000},
	{118, 1194000},
	{117, 1168000},
	{116, 1145000},
	{115, 1125000},
	{114, 1107000},
	{113, 1091000},
	{112, 1077000},
	{111, 1065000},
	{110, 1055000},
	{109, 1046000},
	{108, 1038000},
	{107, 1031000},
	{106, 1025000},
	{105, 1020000},
	{104, 1016000},
	{103, 1013000},
	{102, 1011000},
	{101, 1010000},
	{100, 1000000},
	{99, 1000000},
	{98, 851500},
	{97, 761500},
	{96, 701500},
	{95, 661500},
	{94, 631500},
	{93, 606500},
	{92, 584500},
	{91, 564500},
	{90, 545500},
	{89, 527000},
	{88, 509000},
	{87, 491500},
	{86, 474500},
	{85, 458000},
	{84, 442000},
	{83, 426500},
	{82, 411500},
	{81, 397000},
	{80, 383000},
	{79, 369400},
	{78, 356200},
	{77, 343400},
	{76, 331000},
	{75, 319000},
	{74, 307400},
	{73, 296200},
	{72, 285400},
	{71, 275000},
	{70, 265000},
	{69, 255300},
	{68, 245900},
	{67, 236800},
	{66, 228000},
	{65, 219500},
	{64, 211300},
	{63, 203400},
	{62, 195800},
	{61, 188500},
	{60, 181500},
	{59, 174700},
	{58, 168100},
	{57, 161700},
	{56, 155500},
	{55, 149500},
	{54, 143700},
	{53, 138100},
	{52, 132700},
	{51, 127500},
	{50, 122500},
	{49, 117600},
	{48, 112800},
	{47, 108100},
	{46, 103500},
	{45, 99000},
	{44, 94600},
	{43, 90300},
	{42, 86100},
	{41, 82000},
	{40, 78000},
	{39, 74100},
	{38, 70300},
	{37, 66600},
	{36, 63000},
	{35, 59500},
	{34, 56100},
	{33, 52800},
	{32, 49600},
	{31, 46500},
	{30, 43500},
	{29, 40600},
	{28, 37800},
	{27, 35100},
	{26, 32500},
	{25, 30000},
	{24, 27600},
	{23, 25300},
	{22, 23100},
	{21, 21000},
	{20, 19000},
	{19, 17100},
	{18, 15300},
	{17, 13600},
	{16, 12000},
	{15, 10500},
	{14, 9100},
	{13, 7800},
	{12, 6600},
	{11, 5500},
	{10, 4500},
	{9, 3600},
	{8, 2800},
	{7, 2100},
	{6, 1500},
	{5, 1000},
	{4, 600},
	{3, 300},
	{2, 100},
	{1, 0},
}

// ファイル色
var FILE_COLORS = [...]string{
	// SKIP
}

// 戦果関連任務
var RANKING_POINT_QUESTS = [...]struct {
	Id       string
	Name     string
	Requires []struct {
		Area string
		Rank string
	}
	Fuel         int32
	Ammo         int32
	Steel        int32
	Bauxite      int32
	RankingPoint int32
}{
	// SKIP
}

// 戦果関連任務
var EXPEDITIONS = [...]struct {
	World    int32
	Id       string
	Name     string
	Statuses struct {
		Fire    int32
		AntiAir int32
		Asw     int32
		Scout   int32
	}
	Types         []map[shiptype.ShipType]int32
	MinFlagshipLv int32
	TotalLevel    int32
	MinCount      int32
}{
	// SKIP
}
