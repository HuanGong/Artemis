package weather_cn

type (
	County struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		WeatherCode string `json:"weatherCode"`
	}
	CityInfo struct {
		Id       string   `json:"id"`
		Name     string   `json:"name"`
		Counties []County `json:"county"`
	}
	WeatherItem struct {
		Id   string     `json:"id"`
		Name string     `json:"name"`
		City []CityInfo `json:"city"`
	}
)

const (
	cityCodeCN = `[
    {
       "id": "01",
       "name": "北京",
       "city": [{
          "id": "0101",
          "name": "北京",
          "county": [
             {
                "id": "010101",
                "name": "北京",
                "weatherCode": "101010100"
             },
             {
                "id": "010102",
                "name": "海淀",
                "weatherCode": "101010200"
             },
             {
                "id": "010103",
                "name": "朝阳",
                "weatherCode": "101010300"
             },
             {
                "id": "010104",
                "name": "顺义",
                "weatherCode": "101010400"
             },
             {
                "id": "010105",
                "name": "怀柔",
                "weatherCode": "101010500"
             },
             {
                "id": "010106",
                "name": "通州",
                "weatherCode": "101010600"
             },
             {
                "id": "010107",
                "name": "昌平",
                "weatherCode": "101010700"
             },
             {
                "id": "010108",
                "name": "延庆",
                "weatherCode": "101010800"
             },
             {
                "id": "010109",
                "name": "丰台",
                "weatherCode": "101010900"
             },
             {
                "id": "010110",
                "name": "石景山",
                "weatherCode": "101011000"
             },
             {
                "id": "010111",
                "name": "大兴",
                "weatherCode": "101011100"
             },
             {
                "id": "010112",
                "name": "房山",
                "weatherCode": "101011200"
             },
             {
                "id": "010113",
                "name": "密云",
                "weatherCode": "101011300"
             },
             {
                "id": "010114",
                "name": "门头沟",
                "weatherCode": "101011400"
             },
             {
                "id": "010115",
                "name": "平谷",
                "weatherCode": "101011500"
             }
          ]
       }]
    },
    {
       "id": "02",
       "name": "上海",
       "city": [{
          "id": "0201",
          "name": "上海",
          "county": [
             {
                "id": "020101",
                "name": "上海",
                "weatherCode": "101020100"
             },
             {
                "id": "020102",
                "name": "闵行",
                "weatherCode": "101020200"
             },
             {
                "id": "020103",
                "name": "宝山",
                "weatherCode": "101020300"
             },
             {
                "id": "020104",
                "name": "嘉定",
                "weatherCode": "101020500"
             },
             {
                "id": "020105",
                "name": "浦东南汇",
                "weatherCode": "101020600"
             },
             {
                "id": "020106",
                "name": "金山",
                "weatherCode": "101020700"
             },
             {
                "id": "020107",
                "name": "青浦",
                "weatherCode": "101020800"
             },
             {
                "id": "020108",
                "name": "松江",
                "weatherCode": "101020900"
             },
             {
                "id": "020109",
                "name": "奉贤",
                "weatherCode": "101021000"
             },
             {
                "id": "020110",
                "name": "崇明",
                "weatherCode": "101021100"
             },
             {
                "id": "020111",
                "name": "徐家汇",
                "weatherCode": "101021200"
             },
             {
                "id": "020112",
                "name": "浦东",
                "weatherCode": "101021300"
             }
          ]
       }]
    },
    {
       "id": "03",
       "name": "天津",
       "city": [{
          "id": "0301",
          "name": "天津",
          "county": [
             {
                "id": "030101",
                "name": "天津",
                "weatherCode": "101030100"
             },
             {
                "id": "030102",
                "name": "武清",
                "weatherCode": "101030200"
             },
             {
                "id": "030103",
                "name": "宝坻",
                "weatherCode": "101030300"
             },
             {
                "id": "030104",
                "name": "东丽",
                "weatherCode": "101030400"
             },
             {
                "id": "030105",
                "name": "西青",
                "weatherCode": "101030500"
             },
             {
                "id": "030106",
                "name": "北辰",
                "weatherCode": "101030600"
             },
             {
                "id": "030107",
                "name": "宁河",
                "weatherCode": "101030700"
             },
             {
                "id": "030108",
                "name": "汉沽",
                "weatherCode": "101030800"
             },
             {
                "id": "030109",
                "name": "静海",
                "weatherCode": "101030900"
             },
             {
                "id": "030110",
                "name": "津南",
                "weatherCode": "101031000"
             },
             {
                "id": "030111",
                "name": "塘沽",
                "weatherCode": "101031100"
             },
             {
                "id": "030112",
                "name": "大港",
                "weatherCode": "101031200"
             },
             {
                "id": "030113",
                "name": "蓟县",
                "weatherCode": "101031400"
             }
          ]
       }]
    },
    {
       "id": "04",
       "name": "重庆",
       "city": [{
          "id": "0401",
          "name": "重庆",
          "county": [
             {
                "id": "040101",
                "name": "重庆",
                "weatherCode": "101040100"
             },
             {
                "id": "040102",
                "name": "永川",
                "weatherCode": "101040200"
             },
             {
                "id": "040103",
                "name": "合川",
                "weatherCode": "101040300"
             },
             {
                "id": "040104",
                "name": "南川",
                "weatherCode": "101040400"
             },
             {
                "id": "040105",
                "name": "江津",
                "weatherCode": "101040500"
             },
             {
                "id": "040106",
                "name": "万盛",
                "weatherCode": "101040600"
             },
             {
                "id": "040107",
                "name": "渝北",
                "weatherCode": "101040700"
             },
             {
                "id": "040108",
                "name": "北碚",
                "weatherCode": "101040800"
             },
             {
                "id": "040109",
                "name": "巴南",
                "weatherCode": "101040900"
             },
             {
                "id": "040110",
                "name": "长寿",
                "weatherCode": "101041000"
             },
             {
                "id": "040111",
                "name": "黔江",
                "weatherCode": "101041100"
             },
             {
                "id": "040112",
                "name": "万州",
                "weatherCode": "101041300"
             },
             {
                "id": "040113",
                "name": "涪陵",
                "weatherCode": "101041400"
             },
             {
                "id": "040114",
                "name": "开县",
                "weatherCode": "101041500"
             },
             {
                "id": "040115",
                "name": "城口",
                "weatherCode": "101041600"
             },
             {
                "id": "040116",
                "name": "云阳",
                "weatherCode": "101041700"
             },
             {
                "id": "040117",
                "name": "巫溪",
                "weatherCode": "101041800"
             },
             {
                "id": "040118",
                "name": "奉节",
                "weatherCode": "101041900"
             },
             {
                "id": "040119",
                "name": "巫山",
                "weatherCode": "101042000"
             },
             {
                "id": "040120",
                "name": "潼南",
                "weatherCode": "101042100"
             },
             {
                "id": "040121",
                "name": "垫江",
                "weatherCode": "101042200"
             },
             {
                "id": "040122",
                "name": "梁平",
                "weatherCode": "101042300"
             },
             {
                "id": "040123",
                "name": "忠县",
                "weatherCode": "101042400"
             },
             {
                "id": "040124",
                "name": "石柱",
                "weatherCode": "101042500"
             },
             {
                "id": "040125",
                "name": "大足",
                "weatherCode": "101042600"
             },
             {
                "id": "040126",
                "name": "荣昌",
                "weatherCode": "101042700"
             },
             {
                "id": "040127",
                "name": "铜梁",
                "weatherCode": "101042800"
             },
             {
                "id": "040128",
                "name": "璧山",
                "weatherCode": "101042900"
             },
             {
                "id": "040129",
                "name": "丰都",
                "weatherCode": "101043000"
             },
             {
                "id": "040130",
                "name": "武隆",
                "weatherCode": "101043100"
             },
             {
                "id": "040131",
                "name": "彭水",
                "weatherCode": "101043200"
             },
             {
                "id": "040132",
                "name": "綦江",
                "weatherCode": "101043300"
             },
             {
                "id": "040133",
                "name": "酉阳",
                "weatherCode": "101043400"
             },
             {
                "id": "040134",
                "name": "秀山",
                "weatherCode": "101043600"
             }
          ]
       }]
    },
    {
       "id": "05",
       "name": "黑龙江",
       "city": [
          {
             "id": "0501",
             "name": "哈尔滨",
             "county": [
                {
                   "id": "050101",
                   "name": "哈尔滨",
                   "weatherCode": "101050101"
                },
                {
                   "id": "050102",
                   "name": "双城",
                   "weatherCode": "101050102"
                },
                {
                   "id": "050103",
                   "name": "呼兰",
                   "weatherCode": "101050103"
                },
                {
                   "id": "050104",
                   "name": "阿城",
                   "weatherCode": "101050104"
                },
                {
                   "id": "050105",
                   "name": "宾县",
                   "weatherCode": "101050105"
                },
                {
                   "id": "050106",
                   "name": "依兰",
                   "weatherCode": "101050106"
                },
                {
                   "id": "050107",
                   "name": "巴彦",
                   "weatherCode": "101050107"
                },
                {
                   "id": "050108",
                   "name": "通河",
                   "weatherCode": "101050108"
                },
                {
                   "id": "050109",
                   "name": "方正",
                   "weatherCode": "101050109"
                },
                {
                   "id": "050110",
                   "name": "延寿",
                   "weatherCode": "101050110"
                },
                {
                   "id": "050111",
                   "name": "尚志",
                   "weatherCode": "101050111"
                },
                {
                   "id": "050112",
                   "name": "五常",
                   "weatherCode": "101050112"
                },
                {
                   "id": "050113",
                   "name": "木兰",
                   "weatherCode": "101050113"
                }
             ]
          },
          {
             "id": "0502",
             "name": "齐齐哈尔",
             "county": [
                {
                   "id": "050201",
                   "name": "齐齐哈尔",
                   "weatherCode": "101050201"
                },
                {
                   "id": "050202",
                   "name": "讷河",
                   "weatherCode": "101050202"
                },
                {
                   "id": "050203",
                   "name": "龙江",
                   "weatherCode": "101050203"
                },
                {
                   "id": "050204",
                   "name": "甘南",
                   "weatherCode": "101050204"
                },
                {
                   "id": "050205",
                   "name": "富裕",
                   "weatherCode": "101050205"
                },
                {
                   "id": "050206",
                   "name": "依安",
                   "weatherCode": "101050206"
                },
                {
                   "id": "050207",
                   "name": "拜泉",
                   "weatherCode": "101050207"
                },
                {
                   "id": "050208",
                   "name": "克山",
                   "weatherCode": "101050208"
                },
                {
                   "id": "050209",
                   "name": "克东",
                   "weatherCode": "101050209"
                },
                {
                   "id": "050210",
                   "name": "泰来",
                   "weatherCode": "101050210"
                }
             ]
          },
          {
             "id": "0503",
             "name": "牡丹江",
             "county": [
                {
                   "id": "050301",
                   "name": "牡丹江",
                   "weatherCode": "101050301"
                },
                {
                   "id": "050302",
                   "name": "海林",
                   "weatherCode": "101050302"
                },
                {
                   "id": "050303",
                   "name": "穆棱",
                   "weatherCode": "101050303"
                },
                {
                   "id": "050304",
                   "name": "林口",
                   "weatherCode": "101050304"
                },
                {
                   "id": "050305",
                   "name": "绥芬河",
                   "weatherCode": "101050305"
                },
                {
                   "id": "050306",
                   "name": "宁安",
                   "weatherCode": "101050306"
                },
                {
                   "id": "050307",
                   "name": "东宁",
                   "weatherCode": "101050307"
                }
             ]
          },
          {
             "id": "0504",
             "name": "佳木斯",
             "county": [
                {
                   "id": "050401",
                   "name": "佳木斯",
                   "weatherCode": "101050401"
                },
                {
                   "id": "050402",
                   "name": "汤原",
                   "weatherCode": "101050402"
                },
                {
                   "id": "050403",
                   "name": "抚远",
                   "weatherCode": "101050403"
                },
                {
                   "id": "050404",
                   "name": "桦川",
                   "weatherCode": "101050404"
                },
                {
                   "id": "050405",
                   "name": "桦南",
                   "weatherCode": "101050405"
                },
                {
                   "id": "050406",
                   "name": "同江",
                   "weatherCode": "101050406"
                },
                {
                   "id": "050407",
                   "name": "富锦",
                   "weatherCode": "101050407"
                }
             ]
          },
          {
             "id": "0505",
             "name": "绥化",
             "county": [
                {
                   "id": "050501",
                   "name": "绥化",
                   "weatherCode": "101050501"
                },
                {
                   "id": "050502",
                   "name": "肇东",
                   "weatherCode": "101050502"
                },
                {
                   "id": "050503",
                   "name": "安达",
                   "weatherCode": "101050503"
                },
                {
                   "id": "050504",
                   "name": "海伦",
                   "weatherCode": "101050504"
                },
                {
                   "id": "050505",
                   "name": "明水",
                   "weatherCode": "101050505"
                },
                {
                   "id": "050506",
                   "name": "望奎",
                   "weatherCode": "101050506"
                },
                {
                   "id": "050507",
                   "name": "兰西",
                   "weatherCode": "101050507"
                },
                {
                   "id": "050508",
                   "name": "青冈",
                   "weatherCode": "101050508"
                },
                {
                   "id": "050509",
                   "name": "庆安",
                   "weatherCode": "101050509"
                },
                {
                   "id": "050510",
                   "name": "绥棱",
                   "weatherCode": "101050510"
                }
             ]
          },
          {
             "id": "0506",
             "name": "黑河",
             "county": [
                {
                   "id": "050601",
                   "name": "黑河",
                   "weatherCode": "101050601"
                },
                {
                   "id": "050602",
                   "name": "嫩江",
                   "weatherCode": "101050602"
                },
                {
                   "id": "050603",
                   "name": "孙吴",
                   "weatherCode": "101050603"
                },
                {
                   "id": "050604",
                   "name": "逊克",
                   "weatherCode": "101050604"
                },
                {
                   "id": "050605",
                   "name": "五大连池",
                   "weatherCode": "101050605"
                },
                {
                   "id": "050606",
                   "name": "北安",
                   "weatherCode": "101050606"
                }
             ]
          },
          {
             "id": "0507",
             "name": "大兴安岭",
             "county": [
                {
                   "id": "050701",
                   "name": "大兴安岭",
                   "weatherCode": "101050701"
                },
                {
                   "id": "050702",
                   "name": "塔河",
                   "weatherCode": "101050702"
                },
                {
                   "id": "050703",
                   "name": "漠河",
                   "weatherCode": "101050703"
                },
                {
                   "id": "050704",
                   "name": "呼玛",
                   "weatherCode": "101050704"
                },
                {
                   "id": "050705",
                   "name": "呼中",
                   "weatherCode": "101050705"
                },
                {
                   "id": "050706",
                   "name": "新林",
                   "weatherCode": "101050706"
                },
                {
                   "id": "050707",
                   "name": "加格达奇",
                   "weatherCode": "101050708"
                }
             ]
          },
          {
             "id": "0508",
             "name": "伊春",
             "county": [
                {
                   "id": "050801",
                   "name": "伊春",
                   "weatherCode": "101050801"
                },
                {
                   "id": "050802",
                   "name": "乌伊岭",
                   "weatherCode": "101050802"
                },
                {
                   "id": "050803",
                   "name": "五营",
                   "weatherCode": "101050803"
                },
                {
                   "id": "050804",
                   "name": "铁力",
                   "weatherCode": "101050804"
                },
                {
                   "id": "050805",
                   "name": "嘉荫",
                   "weatherCode": "101050805"
                }
             ]
          },
          {
             "id": "0509",
             "name": "大庆",
             "county": [
                {
                   "id": "050901",
                   "name": "大庆",
                   "weatherCode": "101050901"
                },
                {
                   "id": "050902",
                   "name": "林甸",
                   "weatherCode": "101050902"
                },
                {
                   "id": "050903",
                   "name": "肇州",
                   "weatherCode": "101050903"
                },
                {
                   "id": "050904",
                   "name": "肇源",
                   "weatherCode": "101050904"
                },
                {
                   "id": "050905",
                   "name": "杜尔伯特",
                   "weatherCode": "101050905"
                }
             ]
          },
          {
             "id": "0510",
             "name": "七台河",
             "county": [
                {
                   "id": "051001",
                   "name": "七台河",
                   "weatherCode": "101051002"
                },
                {
                   "id": "051002",
                   "name": "勃利",
                   "weatherCode": "101051003"
                }
             ]
          },
          {
             "id": "0511",
             "name": "鸡西",
             "county": [
                {
                   "id": "051101",
                   "name": "鸡西",
                   "weatherCode": "101051101"
                },
                {
                   "id": "051102",
                   "name": "虎林",
                   "weatherCode": "101051102"
                },
                {
                   "id": "051103",
                   "name": "密山",
                   "weatherCode": "101051103"
                },
                {
                   "id": "051104",
                   "name": "鸡东",
                   "weatherCode": "101051104"
                }
             ]
          },
          {
             "id": "0512",
             "name": "鹤岗",
             "county": [
                {
                   "id": "051201",
                   "name": "鹤岗",
                   "weatherCode": "101051201"
                },
                {
                   "id": "051202",
                   "name": "绥滨",
                   "weatherCode": "101051202"
                },
                {
                   "id": "051203",
                   "name": "萝北",
                   "weatherCode": "101051203"
                }
             ]
          },
          {
             "id": "0513",
             "name": "双鸭山",
             "county": [
                {
                   "id": "051301",
                   "name": "双鸭山",
                   "weatherCode": "101051301"
                },
                {
                   "id": "051302",
                   "name": "集贤",
                   "weatherCode": "101051302"
                },
                {
                   "id": "051303",
                   "name": "宝清",
                   "weatherCode": "101051303"
                },
                {
                   "id": "051304",
                   "name": "饶河",
                   "weatherCode": "101051304"
                },
                {
                   "id": "051305",
                   "name": "友谊",
                   "weatherCode": "101051305"
                }
             ]
          }
       ]
    },
    {
       "id": "06",
       "name": "吉林",
       "city": [
          {
             "id": "0601",
             "name": "长春",
             "county": [
                {
                   "id": "060101",
                   "name": "长春",
                   "weatherCode": "101060101"
                },
                {
                   "id": "060102",
                   "name": "农安",
                   "weatherCode": "101060102"
                },
                {
                   "id": "060103",
                   "name": "德惠",
                   "weatherCode": "101060103"
                },
                {
                   "id": "060104",
                   "name": "九台",
                   "weatherCode": "101060104"
                },
                {
                   "id": "060105",
                   "name": "榆树",
                   "weatherCode": "101060105"
                },
                {
                   "id": "060106",
                   "name": "双阳",
                   "weatherCode": "101060106"
                }
             ]
          },
          {
             "id": "0602",
             "name": "吉林",
             "county": [
                {
                   "id": "060201",
                   "name": "吉林",
                   "weatherCode": "101060201"
                },
                {
                   "id": "060202",
                   "name": "舒兰",
                   "weatherCode": "101060202"
                },
                {
                   "id": "060203",
                   "name": "永吉",
                   "weatherCode": "101060203"
                },
                {
                   "id": "060204",
                   "name": "蛟河",
                   "weatherCode": "101060204"
                },
                {
                   "id": "060205",
                   "name": "磐石",
                   "weatherCode": "101060205"
                },
                {
                   "id": "060206",
                   "name": "桦甸",
                   "weatherCode": "101060206"
                }
             ]
          },
          {
             "id": "0603",
             "name": "延边",
             "county": [
                {
                   "id": "060301",
                   "name": "延吉",
                   "weatherCode": "101060301"
                },
                {
                   "id": "060302",
                   "name": "敦化",
                   "weatherCode": "101060302"
                },
                {
                   "id": "060303",
                   "name": "安图",
                   "weatherCode": "101060303"
                },
                {
                   "id": "060304",
                   "name": "汪清",
                   "weatherCode": "101060304"
                },
                {
                   "id": "060305",
                   "name": "和龙",
                   "weatherCode": "101060305"
                },
                {
                   "id": "060306",
                   "name": "龙井",
                   "weatherCode": "101060307"
                },
                {
                   "id": "060307",
                   "name": "珲春",
                   "weatherCode": "101060308"
                },
                {
                   "id": "060308",
                   "name": "图们",
                   "weatherCode": "101060309"
                }
             ]
          },
          {
             "id": "0604",
             "name": "四平",
             "county": [
                {
                   "id": "060401",
                   "name": "四平",
                   "weatherCode": "101060401"
                },
                {
                   "id": "060402",
                   "name": "双辽",
                   "weatherCode": "101060402"
                },
                {
                   "id": "060403",
                   "name": "梨树",
                   "weatherCode": "101060403"
                },
                {
                   "id": "060404",
                   "name": "公主岭",
                   "weatherCode": "101060404"
                },
                {
                   "id": "060405",
                   "name": "伊通",
                   "weatherCode": "101060405"
                }
             ]
          },
          {
             "id": "0605",
             "name": "通化",
             "county": [
                {
                   "id": "060501",
                   "name": "通化",
                   "weatherCode": "101060501"
                },
                {
                   "id": "060502",
                   "name": "梅河口",
                   "weatherCode": "101060502"
                },
                {
                   "id": "060503",
                   "name": "柳河",
                   "weatherCode": "101060503"
                },
                {
                   "id": "060504",
                   "name": "辉南",
                   "weatherCode": "101060504"
                },
                {
                   "id": "060505",
                   "name": "集安",
                   "weatherCode": "101060505"
                },
                {
                   "id": "060506",
                   "name": "通化县",
                   "weatherCode": "101060506"
                }
             ]
          },
          {
             "id": "0606",
             "name": "白城",
             "county": [
                {
                   "id": "060601",
                   "name": "白城",
                   "weatherCode": "101060601"
                },
                {
                   "id": "060602",
                   "name": "洮南",
                   "weatherCode": "101060602"
                },
                {
                   "id": "060603",
                   "name": "大安",
                   "weatherCode": "101060603"
                },
                {
                   "id": "060604",
                   "name": "镇赉",
                   "weatherCode": "101060604"
                },
                {
                   "id": "060605",
                   "name": "通榆",
                   "weatherCode": "101060605"
                }
             ]
          },
          {
             "id": "0607",
             "name": "辽源",
             "county": [
                {
                   "id": "060701",
                   "name": "辽源",
                   "weatherCode": "101060701"
                },
                {
                   "id": "060702",
                   "name": "东丰",
                   "weatherCode": "101060702"
                },
                {
                   "id": "060703",
                   "name": "东辽",
                   "weatherCode": "101060703"
                }
             ]
          },
          {
             "id": "0608",
             "name": "松原",
             "county": [
                {
                   "id": "060801",
                   "name": "松原",
                   "weatherCode": "101060801"
                },
                {
                   "id": "060802",
                   "name": "乾安",
                   "weatherCode": "101060802"
                },
                {
                   "id": "060803",
                   "name": "前郭",
                   "weatherCode": "101060803"
                },
                {
                   "id": "060804",
                   "name": "长岭",
                   "weatherCode": "101060804"
                },
                {
                   "id": "060805",
                   "name": "扶余",
                   "weatherCode": "101060805"
                }
             ]
          },
          {
             "id": "0609",
             "name": "白山",
             "county": [
                {
                   "id": "060901",
                   "name": "白山",
                   "weatherCode": "101060901"
                },
                {
                   "id": "060902",
                   "name": "靖宇",
                   "weatherCode": "101060902"
                },
                {
                   "id": "060903",
                   "name": "临江",
                   "weatherCode": "101060903"
                },
                {
                   "id": "060904",
                   "name": "东岗",
                   "weatherCode": "101060904"
                },
                {
                   "id": "060905",
                   "name": "长白",
                   "weatherCode": "101060905"
                },
                {
                   "id": "060906",
                   "name": "抚松",
                   "weatherCode": "101060906"
                },
                {
                   "id": "060907",
                   "name": "江源",
                   "weatherCode": "101060907"
                }
             ]
          }
       ]
    },
    {
       "id": "07",
       "name": "辽宁",
       "city": [
          {
             "id": "0701",
             "name": "沈阳",
             "county": [
                {
                   "id": "070101",
                   "name": "沈阳",
                   "weatherCode": "101070101"
                },
                {
                   "id": "070102",
                   "name": "辽中",
                   "weatherCode": "101070103"
                },
                {
                   "id": "070103",
                   "name": "康平",
                   "weatherCode": "101070104"
                },
                {
                   "id": "070104",
                   "name": "法库",
                   "weatherCode": "101070105"
                },
                {
                   "id": "070105",
                   "name": "新民",
                   "weatherCode": "101070106"
                }
             ]
          },
          {
             "id": "0702",
             "name": "大连",
             "county": [
                {
                   "id": "070201",
                   "name": "大连",
                   "weatherCode": "101070201"
                },
                {
                   "id": "070202",
                   "name": "瓦房店",
                   "weatherCode": "101070202"
                },
                {
                   "id": "070203",
                   "name": "金州",
                   "weatherCode": "101070203"
                },
                {
                   "id": "070204",
                   "name": "普兰店",
                   "weatherCode": "101070204"
                },
                {
                   "id": "070205",
                   "name": "旅顺",
                   "weatherCode": "101070205"
                },
                {
                   "id": "070206",
                   "name": "长海",
                   "weatherCode": "101070206"
                },
                {
                   "id": "070207",
                   "name": "庄河",
                   "weatherCode": "101070207"
                }
             ]
          },
          {
             "id": "0703",
             "name": "鞍山",
             "county": [
                {
                   "id": "070301",
                   "name": "鞍山",
                   "weatherCode": "101070301"
                },
                {
                   "id": "070302",
                   "name": "台安",
                   "weatherCode": "101070302"
                },
                {
                   "id": "070303",
                   "name": "岫岩",
                   "weatherCode": "101070303"
                },
                {
                   "id": "070304",
                   "name": "海城",
                   "weatherCode": "101070304"
                }
             ]
          },
          {
             "id": "0704",
             "name": "抚顺",
             "county": [
                {
                   "id": "070401",
                   "name": "抚顺",
                   "weatherCode": "101070401"
                },
                {
                   "id": "070402",
                   "name": "新宾",
                   "weatherCode": "101070402"
                },
                {
                   "id": "070403",
                   "name": "清原",
                   "weatherCode": "101070403"
                },
                {
                   "id": "070404",
                   "name": "章党",
                   "weatherCode": "101070404"
                }
             ]
          },
          {
             "id": "0705",
             "name": "本溪",
             "county": [
                {
                   "id": "070501",
                   "name": "本溪",
                   "weatherCode": "101070501"
                },
                {
                   "id": "070502",
                   "name": "本溪县",
                   "weatherCode": "101070502"
                },
                {
                   "id": "070503",
                   "name": "桓仁",
                   "weatherCode": "101070504"
                }
             ]
          },
          {
             "id": "0706",
             "name": "丹东",
             "county": [
                {
                   "id": "070601",
                   "name": "丹东",
                   "weatherCode": "101070601"
                },
                {
                   "id": "070602",
                   "name": "凤城",
                   "weatherCode": "101070602"
                },
                {
                   "id": "070603",
                   "name": "宽甸",
                   "weatherCode": "101070603"
                },
                {
                   "id": "070604",
                   "name": "东港",
                   "weatherCode": "101070604"
                }
             ]
          },
          {
             "id": "0707",
             "name": "锦州",
             "county": [
                {
                   "id": "070701",
                   "name": "锦州",
                   "weatherCode": "101070701"
                },
                {
                   "id": "070702",
                   "name": "凌海",
                   "weatherCode": "101070702"
                },
                {
                   "id": "070703",
                   "name": "义县",
                   "weatherCode": "101070704"
                },
                {
                   "id": "070704",
                   "name": "黑山",
                   "weatherCode": "101070705"
                },
                {
                   "id": "070705",
                   "name": "北镇",
                   "weatherCode": "101070706"
                }
             ]
          },
          {
             "id": "0708",
             "name": "营口",
             "county": [
                {
                   "id": "070801",
                   "name": "营口",
                   "weatherCode": "101070801"
                },
                {
                   "id": "070802",
                   "name": "大石桥",
                   "weatherCode": "101070802"
                },
                {
                   "id": "070803",
                   "name": "盖州",
                   "weatherCode": "101070803"
                }
             ]
          },
          {
             "id": "0709",
             "name": "阜新",
             "county": [
                {
                   "id": "070901",
                   "name": "阜新",
                   "weatherCode": "101070901"
                },
                {
                   "id": "070902",
                   "name": "彰武",
                   "weatherCode": "101070902"
                }
             ]
          },
          {
             "id": "0710",
             "name": "辽阳",
             "county": [
                {
                   "id": "071001",
                   "name": "辽阳",
                   "weatherCode": "101071001"
                },
                {
                   "id": "071002",
                   "name": "辽阳县",
                   "weatherCode": "101071002"
                },
                {
                   "id": "071003",
                   "name": "灯塔",
                   "weatherCode": "101071003"
                },
                {
                   "id": "071004",
                   "name": "弓长岭",
                   "weatherCode": "101071004"
                }
             ]
          },
          {
             "id": "0711",
             "name": "铁岭",
             "county": [
                {
                   "id": "071101",
                   "name": "铁岭",
                   "weatherCode": "101071101"
                },
                {
                   "id": "071102",
                   "name": "开原",
                   "weatherCode": "101071102"
                },
                {
                   "id": "071103",
                   "name": "昌图",
                   "weatherCode": "101071103"
                },
                {
                   "id": "071104",
                   "name": "西丰",
                   "weatherCode": "101071104"
                },
                {
                   "id": "071105",
                   "name": "调兵山",
                   "weatherCode": "101071105"
                }
             ]
          },
          {
             "id": "0712",
             "name": "朝阳",
             "county": [
                {
                   "id": "071202",
                   "name": "凌源",
                   "weatherCode": "101071203"
                },
                {
                   "id": "071203",
                   "name": "喀左",
                   "weatherCode": "101071204"
                },
                {
                   "id": "071204",
                   "name": "北票",
                   "weatherCode": "101071205"
                },
                {
                   "id": "071205",
                   "name": "建平县",
                   "weatherCode": "101071207"
                }
             ]
          },
          {
             "id": "0713",
             "name": "盘锦",
             "county": [
                {
                   "id": "071301",
                   "name": "盘锦",
                   "weatherCode": "101071301"
                },
                {
                   "id": "071302",
                   "name": "大洼",
                   "weatherCode": "101071302"
                },
                {
                   "id": "071303",
                   "name": "盘山",
                   "weatherCode": "101071303"
                }
             ]
          },
          {
             "id": "0714",
             "name": "葫芦岛",
             "county": [
                {
                   "id": "071401",
                   "name": "葫芦岛",
                   "weatherCode": "101071401"
                },
                {
                   "id": "071402",
                   "name": "建昌",
                   "weatherCode": "101071402"
                },
                {
                   "id": "071403",
                   "name": "绥中",
                   "weatherCode": "101071403"
                },
                {
                   "id": "071404",
                   "name": "兴城",
                   "weatherCode": "101071404"
                }
             ]
          }
       ]
    },
    {
       "id": "08",
       "name": "内蒙古",
       "city": [
          {
             "id": "0801",
             "name": "呼和浩特",
             "county": [
                {
                   "id": "080101",
                   "name": "呼和浩特",
                   "weatherCode": "101080101"
                },
                {
                   "id": "080102",
                   "name": "土左旗",
                   "weatherCode": "101080102"
                },
                {
                   "id": "080103",
                   "name": "托县",
                   "weatherCode": "101080103"
                },
                {
                   "id": "080104",
                   "name": "和林",
                   "weatherCode": "101080104"
                },
                {
                   "id": "080105",
                   "name": "清水河",
                   "weatherCode": "101080105"
                },
                {
                   "id": "080106",
                   "name": "呼市郊区",
                   "weatherCode": "101080106"
                },
                {
                   "id": "080107",
                   "name": "武川",
                   "weatherCode": "101080107"
                }
             ]
          },
          {
             "id": "0802",
             "name": "包头",
             "county": [
                {
                   "id": "080201",
                   "name": "包头",
                   "weatherCode": "101080201"
                },
                {
                   "id": "080202",
                   "name": "白云鄂博",
                   "weatherCode": "101080202"
                },
                {
                   "id": "080203",
                   "name": "满都拉",
                   "weatherCode": "101080203"
                },
                {
                   "id": "080204",
                   "name": "土右旗",
                   "weatherCode": "101080204"
                },
                {
                   "id": "080205",
                   "name": "固阳",
                   "weatherCode": "101080205"
                },
                {
                   "id": "080206",
                   "name": "达茂旗",
                   "weatherCode": "101080206"
                },
                {
                   "id": "080207",
                   "name": "希拉穆仁",
                   "weatherCode": "101080207"
                }
             ]
          },
          {
             "id": "0803",
             "name": "乌海",
             "county": [{
                "id": "080301",
                "name": "乌海",
                "weatherCode": "101080301"
             }]
          },
          {
             "id": "0804",
             "name": "乌兰察布",
             "county": [
                {
                   "id": "080401",
                   "name": "集宁",
                   "weatherCode": "101080401"
                },
                {
                   "id": "080402",
                   "name": "卓资",
                   "weatherCode": "101080402"
                },
                {
                   "id": "080403",
                   "name": "化德",
                   "weatherCode": "101080403"
                },
                {
                   "id": "080404",
                   "name": "商都",
                   "weatherCode": "101080404"
                },
                {
                   "id": "080405",
                   "name": "兴和",
                   "weatherCode": "101080406"
                },
                {
                   "id": "080406",
                   "name": "凉城",
                   "weatherCode": "101080407"
                },
                {
                   "id": "080407",
                   "name": "察右前旗",
                   "weatherCode": "101080408"
                },
                {
                   "id": "080408",
                   "name": "察右中旗",
                   "weatherCode": "101080409"
                },
                {
                   "id": "080409",
                   "name": "察右后旗",
                   "weatherCode": "101080410"
                },
                {
                   "id": "080410",
                   "name": "四子王旗",
                   "weatherCode": "101080411"
                },
                {
                   "id": "080411",
                   "name": "丰镇",
                   "weatherCode": "101080412"
                }
             ]
          },
          {
             "id": "0805",
             "name": "通辽",
             "county": [
                {
                   "id": "080501",
                   "name": "通辽",
                   "weatherCode": "101080501"
                },
                {
                   "id": "080502",
                   "name": "舍伯吐",
                   "weatherCode": "101080502"
                },
                {
                   "id": "080503",
                   "name": "科左中旗",
                   "weatherCode": "101080503"
                },
                {
                   "id": "080504",
                   "name": "科左后旗",
                   "weatherCode": "101080504"
                },
                {
                   "id": "080505",
                   "name": "青龙山",
                   "weatherCode": "101080505"
                },
                {
                   "id": "080506",
                   "name": "开鲁",
                   "weatherCode": "101080506"
                },
                {
                   "id": "080507",
                   "name": "库伦",
                   "weatherCode": "101080507"
                },
                {
                   "id": "080508",
                   "name": "奈曼",
                   "weatherCode": "101080508"
                },
                {
                   "id": "080509",
                   "name": "扎鲁特",
                   "weatherCode": "101080509"
                },
                {
                   "id": "080510",
                   "name": "巴雅尔吐胡硕",
                   "weatherCode": "101080511"
                },
                {
                   "id": "080511",
                   "name": "霍林郭勒",
                   "weatherCode": "101081108"
                }
             ]
          },
          {
             "id": "0806",
             "name": "赤峰",
             "county": [
                {
                   "id": "080601",
                   "name": "赤峰",
                   "weatherCode": "101080601"
                },
                {
                   "id": "080602",
                   "name": "阿鲁旗",
                   "weatherCode": "101080603"
                },
                {
                   "id": "080603",
                   "name": "浩尔吐",
                   "weatherCode": "101080604"
                },
                {
                   "id": "080604",
                   "name": "巴林左旗",
                   "weatherCode": "101080605"
                },
                {
                   "id": "080605",
                   "name": "巴林右旗",
                   "weatherCode": "101080606"
                },
                {
                   "id": "080606",
                   "name": "林西",
                   "weatherCode": "101080607"
                },
                {
                   "id": "080607",
                   "name": "克什克腾",
                   "weatherCode": "101080608"
                },
                {
                   "id": "080608",
                   "name": "翁牛特",
                   "weatherCode": "101080609"
                },
                {
                   "id": "080609",
                   "name": "岗子",
                   "weatherCode": "101080610"
                },
                {
                   "id": "080610",
                   "name": "喀喇沁",
                   "weatherCode": "101080611"
                },
                {
                   "id": "080611",
                   "name": "八里罕",
                   "weatherCode": "101080612"
                },
                {
                   "id": "080612",
                   "name": "宁城",
                   "weatherCode": "101080613"
                },
                {
                   "id": "080613",
                   "name": "敖汉",
                   "weatherCode": "101080614"
                },
                {
                   "id": "080614",
                   "name": "宝国吐",
                   "weatherCode": "101080615"
                }
             ]
          },
          {
             "id": "0807",
             "name": "鄂尔多斯",
             "county": [
                {
                   "id": "080701",
                   "name": "鄂尔多斯",
                   "weatherCode": "101080701"
                },
                {
                   "id": "080702",
                   "name": "达拉特",
                   "weatherCode": "101080703"
                },
                {
                   "id": "080703",
                   "name": "准格尔",
                   "weatherCode": "101080704"
                },
                {
                   "id": "080704",
                   "name": "鄂前旗",
                   "weatherCode": "101080705"
                },
                {
                   "id": "080705",
                   "name": "河南",
                   "weatherCode": "101080706"
                },
                {
                   "id": "080706",
                   "name": "伊克乌素",
                   "weatherCode": "101080707"
                },
                {
                   "id": "080707",
                   "name": "鄂托克",
                   "weatherCode": "101080708"
                },
                {
                   "id": "080708",
                   "name": "杭锦旗",
                   "weatherCode": "101080709"
                },
                {
                   "id": "080709",
                   "name": "乌审旗",
                   "weatherCode": "101080710"
                },
                {
                   "id": "080710",
                   "name": "伊金霍洛",
                   "weatherCode": "101080711"
                },
                {
                   "id": "080711",
                   "name": "乌审召",
                   "weatherCode": "101080712"
                },
                {
                   "id": "080712",
                   "name": "东胜",
                   "weatherCode": "101080713"
                }
             ]
          },
          {
             "id": "0808",
             "name": "巴彦淖尔",
             "county": [
                {
                   "id": "080801",
                   "name": "临河",
                   "weatherCode": "101080801"
                },
                {
                   "id": "080802",
                   "name": "五原",
                   "weatherCode": "101080802"
                },
                {
                   "id": "080803",
                   "name": "磴口",
                   "weatherCode": "101080803"
                },
                {
                   "id": "080804",
                   "name": "乌前旗",
                   "weatherCode": "101080804"
                },
                {
                   "id": "080805",
                   "name": "大佘太",
                   "weatherCode": "101080805"
                },
                {
                   "id": "080806",
                   "name": "乌中旗",
                   "weatherCode": "101080806"
                },
                {
                   "id": "080807",
                   "name": "乌后旗",
                   "weatherCode": "101080807"
                },
                {
                   "id": "080808",
                   "name": "海力素",
                   "weatherCode": "101080808"
                },
                {
                   "id": "080809",
                   "name": "那仁宝力格",
                   "weatherCode": "101080809"
                },
                {
                   "id": "080810",
                   "name": "杭锦后旗",
                   "weatherCode": "101080810"
                }
             ]
          },
          {
             "id": "0809",
             "name": "锡林郭勒",
             "county": [
                {
                   "id": "080901",
                   "name": "锡林浩特",
                   "weatherCode": "101080901"
                },
                {
                   "id": "080902",
                   "name": "二连浩特",
                   "weatherCode": "101080903"
                },
                {
                   "id": "080903",
                   "name": "阿巴嘎",
                   "weatherCode": "101080904"
                },
                {
                   "id": "080904",
                   "name": "苏左旗",
                   "weatherCode": "101080906"
                },
                {
                   "id": "080905",
                   "name": "苏右旗",
                   "weatherCode": "101080907"
                },
                {
                   "id": "080906",
                   "name": "朱日和",
                   "weatherCode": "101080908"
                },
                {
                   "id": "080907",
                   "name": "东乌旗",
                   "weatherCode": "101080909"
                },
                {
                   "id": "080908",
                   "name": "西乌旗",
                   "weatherCode": "101080910"
                },
                {
                   "id": "080909",
                   "name": "太仆寺",
                   "weatherCode": "101080911"
                },
                {
                   "id": "080910",
                   "name": "镶黄旗",
                   "weatherCode": "101080912"
                },
                {
                   "id": "080911",
                   "name": "正镶白旗",
                   "weatherCode": "101080913"
                },
                {
                   "id": "080912",
                   "name": "正蓝旗",
                   "weatherCode": "101080914"
                },
                {
                   "id": "080913",
                   "name": "多伦",
                   "weatherCode": "101080915"
                },
                {
                   "id": "080914",
                   "name": "博克图",
                   "weatherCode": "101080916"
                },
                {
                   "id": "080915",
                   "name": "乌拉盖",
                   "weatherCode": "101080917"
                }
             ]
          },
          {
             "id": "0810",
             "name": "呼伦贝尔",
             "county": [
                {
                   "id": "081001",
                   "name": "海拉尔",
                   "weatherCode": "101081001"
                },
                {
                   "id": "081002",
                   "name": "小二沟",
                   "weatherCode": "101081002"
                },
                {
                   "id": "081003",
                   "name": "阿荣旗",
                   "weatherCode": "101081003"
                },
                {
                   "id": "081004",
                   "name": "莫力达瓦",
                   "weatherCode": "101081004"
                },
                {
                   "id": "081005",
                   "name": "鄂伦春旗",
                   "weatherCode": "101081005"
                },
                {
                   "id": "081006",
                   "name": "鄂温克旗",
                   "weatherCode": "101081006"
                },
                {
                   "id": "081007",
                   "name": "陈旗",
                   "weatherCode": "101081007"
                },
                {
                   "id": "081008",
                   "name": "新左旗",
                   "weatherCode": "101081008"
                },
                {
                   "id": "081009",
                   "name": "新右旗",
                   "weatherCode": "101081009"
                },
                {
                   "id": "081010",
                   "name": "满洲里",
                   "weatherCode": "101081010"
                },
                {
                   "id": "081011",
                   "name": "牙克石",
                   "weatherCode": "101081011"
                },
                {
                   "id": "081012",
                   "name": "扎兰屯",
                   "weatherCode": "101081012"
                },
                {
                   "id": "081013",
                   "name": "额尔古纳",
                   "weatherCode": "101081014"
                },
                {
                   "id": "081014",
                   "name": "根河",
                   "weatherCode": "101081015"
                },
                {
                   "id": "081015",
                   "name": "图里河",
                   "weatherCode": "101081016"
                }
             ]
          },
          {
             "id": "0811",
             "name": "兴安盟",
             "county": [
                {
                   "id": "081101",
                   "name": "高力板",
                   "weatherCode": "101080510"
                },
                {
                   "id": "081102",
                   "name": "乌兰浩特",
                   "weatherCode": "101081101"
                },
                {
                   "id": "081103",
                   "name": "阿尔山",
                   "weatherCode": "101081102"
                },
                {
                   "id": "081104",
                   "name": "科右中旗",
                   "weatherCode": "101081103"
                },
                {
                   "id": "081105",
                   "name": "胡尔勒",
                   "weatherCode": "101081104"
                },
                {
                   "id": "081106",
                   "name": "扎赉特",
                   "weatherCode": "101081105"
                },
                {
                   "id": "081107",
                   "name": "索伦",
                   "weatherCode": "101081106"
                },
                {
                   "id": "081108",
                   "name": "突泉",
                   "weatherCode": "101081107"
                },
                {
                   "id": "081109",
                   "name": "科右前旗",
                   "weatherCode": "101081109"
                }
             ]
          },
          {
             "id": "0812",
             "name": "阿拉善盟",
             "county": [
                {
                   "id": "081201",
                   "name": "阿左旗",
                   "weatherCode": "101081201"
                },
                {
                   "id": "081202",
                   "name": "阿右旗",
                   "weatherCode": "101081202"
                },
                {
                   "id": "081203",
                   "name": "额济纳",
                   "weatherCode": "101081203"
                },
                {
                   "id": "081204",
                   "name": "拐子湖",
                   "weatherCode": "101081204"
                },
                {
                   "id": "081205",
                   "name": "吉兰太",
                   "weatherCode": "101081205"
                },
                {
                   "id": "081206",
                   "name": "锡林高勒",
                   "weatherCode": "101081206"
                },
                {
                   "id": "081207",
                   "name": "头道湖",
                   "weatherCode": "101081207"
                },
                {
                   "id": "081208",
                   "name": "中泉子",
                   "weatherCode": "101081208"
                },
                {
                   "id": "081209",
                   "name": "诺尔公",
                   "weatherCode": "101081209"
                },
                {
                   "id": "081210",
                   "name": "雅布赖",
                   "weatherCode": "101081210"
                },
                {
                   "id": "081211",
                   "name": "乌斯泰",
                   "weatherCode": "101081211"
                },
                {
                   "id": "081212",
                   "name": "孪井滩",
                   "weatherCode": "101081212"
                }
             ]
          }
       ]
    },
    {
       "id": "09",
       "name": "河北",
       "city": [
          {
             "id": "0901",
             "name": "石家庄",
             "county": [
                {
                   "id": "090101",
                   "name": "石家庄",
                   "weatherCode": "101090101"
                },
                {
                   "id": "090102",
                   "name": "井陉",
                   "weatherCode": "101090102"
                },
                {
                   "id": "090103",
                   "name": "正定",
                   "weatherCode": "101090103"
                },
                {
                   "id": "090104",
                   "name": "栾城",
                   "weatherCode": "101090104"
                },
                {
                   "id": "090105",
                   "name": "行唐",
                   "weatherCode": "101090105"
                },
                {
                   "id": "090106",
                   "name": "灵寿",
                   "weatherCode": "101090106"
                },
                {
                   "id": "090107",
                   "name": "高邑",
                   "weatherCode": "101090107"
                },
                {
                   "id": "090108",
                   "name": "深泽",
                   "weatherCode": "101090108"
                },
                {
                   "id": "090109",
                   "name": "赞皇",
                   "weatherCode": "101090109"
                },
                {
                   "id": "090110",
                   "name": "无极",
                   "weatherCode": "101090110"
                },
                {
                   "id": "090111",
                   "name": "平山",
                   "weatherCode": "101090111"
                },
                {
                   "id": "090112",
                   "name": "元氏",
                   "weatherCode": "101090112"
                },
                {
                   "id": "090113",
                   "name": "赵县",
                   "weatherCode": "101090113"
                },
                {
                   "id": "090114",
                   "name": "辛集",
                   "weatherCode": "101090114"
                },
                {
                   "id": "090115",
                   "name": "藁城",
                   "weatherCode": "101090115"
                },
                {
                   "id": "090116",
                   "name": "晋州",
                   "weatherCode": "101090116"
                },
                {
                   "id": "090117",
                   "name": "新乐",
                   "weatherCode": "101090117"
                },
                {
                   "id": "090118",
                   "name": "鹿泉",
                   "weatherCode": "101090118"
                }
             ]
          },
          {
             "id": "0902",
             "name": "保定",
             "county": [
                {
                   "id": "090201",
                   "name": "保定",
                   "weatherCode": "101090201"
                },
                {
                   "id": "090202",
                   "name": "满城",
                   "weatherCode": "101090202"
                },
                {
                   "id": "090203",
                   "name": "阜平",
                   "weatherCode": "101090203"
                },
                {
                   "id": "090204",
                   "name": "徐水",
                   "weatherCode": "101090204"
                },
                {
                   "id": "090205",
                   "name": "唐县",
                   "weatherCode": "101090205"
                },
                {
                   "id": "090206",
                   "name": "高阳",
                   "weatherCode": "101090206"
                },
                {
                   "id": "090207",
                   "name": "容城",
                   "weatherCode": "101090207"
                },
                {
                   "id": "090208",
                   "name": "涞源",
                   "weatherCode": "101090209"
                },
                {
                   "id": "090209",
                   "name": "望都",
                   "weatherCode": "101090210"
                },
                {
                   "id": "090210",
                   "name": "安新",
                   "weatherCode": "101090211"
                },
                {
                   "id": "090211",
                   "name": "易县",
                   "weatherCode": "101090212"
                },
                {
                   "id": "090212",
                   "name": "曲阳",
                   "weatherCode": "101090214"
                },
                {
                   "id": "090213",
                   "name": "蠡县",
                   "weatherCode": "101090215"
                },
                {
                   "id": "090214",
                   "name": "顺平",
                   "weatherCode": "101090216"
                },
                {
                   "id": "090215",
                   "name": "雄县",
                   "weatherCode": "101090217"
                },
                {
                   "id": "090216",
                   "name": "涿州",
                   "weatherCode": "101090218"
                },
                {
                   "id": "090217",
                   "name": "定州",
                   "weatherCode": "101090219"
                },
                {
                   "id": "090218",
                   "name": "安国",
                   "weatherCode": "101090220"
                },
                {
                   "id": "090219",
                   "name": "高碑店",
                   "weatherCode": "101090221"
                },
                {
                   "id": "090220",
                   "name": "涞水",
                   "weatherCode": "101090222"
                },
                {
                   "id": "090221",
                   "name": "定兴",
                   "weatherCode": "101090223"
                },
                {
                   "id": "090222",
                   "name": "清苑",
                   "weatherCode": "101090224"
                },
                {
                   "id": "090223",
                   "name": "博野",
                   "weatherCode": "101090225"
                }
             ]
          },
          {
             "id": "0903",
             "name": "张家口",
             "county": [
                {
                   "id": "090301",
                   "name": "张家口",
                   "weatherCode": "101090301"
                },
                {
                   "id": "090302",
                   "name": "宣化",
                   "weatherCode": "101090302"
                },
                {
                   "id": "090303",
                   "name": "张北",
                   "weatherCode": "101090303"
                },
                {
                   "id": "090304",
                   "name": "康保",
                   "weatherCode": "101090304"
                },
                {
                   "id": "090305",
                   "name": "沽源",
                   "weatherCode": "101090305"
                },
                {
                   "id": "090306",
                   "name": "尚义",
                   "weatherCode": "101090306"
                },
                {
                   "id": "090307",
                   "name": "蔚县",
                   "weatherCode": "101090307"
                },
                {
                   "id": "090308",
                   "name": "阳原",
                   "weatherCode": "101090308"
                },
                {
                   "id": "090309",
                   "name": "怀安",
                   "weatherCode": "101090309"
                },
                {
                   "id": "090310",
                   "name": "万全",
                   "weatherCode": "101090310"
                },
                {
                   "id": "090311",
                   "name": "怀来",
                   "weatherCode": "101090311"
                },
                {
                   "id": "090312",
                   "name": "涿鹿",
                   "weatherCode": "101090312"
                },
                {
                   "id": "090313",
                   "name": "赤城",
                   "weatherCode": "101090313"
                },
                {
                   "id": "090314",
                   "name": "崇礼",
                   "weatherCode": "101090314"
                }
             ]
          },
          {
             "id": "0904",
             "name": "承德",
             "county": [
                {
                   "id": "090401",
                   "name": "承德",
                   "weatherCode": "101090402"
                },
                {
                   "id": "090402",
                   "name": "承德县",
                   "weatherCode": "101090403"
                },
                {
                   "id": "090403",
                   "name": "兴隆",
                   "weatherCode": "101090404"
                },
                {
                   "id": "090404",
                   "name": "平泉",
                   "weatherCode": "101090405"
                },
                {
                   "id": "090405",
                   "name": "滦平",
                   "weatherCode": "101090406"
                },
                {
                   "id": "090406",
                   "name": "隆化",
                   "weatherCode": "101090407"
                },
                {
                   "id": "090407",
                   "name": "丰宁",
                   "weatherCode": "101090408"
                },
                {
                   "id": "090408",
                   "name": "宽城",
                   "weatherCode": "101090409"
                },
                {
                   "id": "090409",
                   "name": "围场",
                   "weatherCode": "101090410"
                }
             ]
          },
          {
             "id": "0905",
             "name": "唐山",
             "county": [
                {
                   "id": "090501",
                   "name": "唐山",
                   "weatherCode": "101090501"
                },
                {
                   "id": "090502",
                   "name": "丰南",
                   "weatherCode": "101090502"
                },
                {
                   "id": "090503",
                   "name": "丰润",
                   "weatherCode": "101090503"
                },
                {
                   "id": "090504",
                   "name": "滦县",
                   "weatherCode": "101090504"
                },
                {
                   "id": "090505",
                   "name": "滦南",
                   "weatherCode": "101090505"
                },
                {
                   "id": "090506",
                   "name": "乐亭",
                   "weatherCode": "101090506"
                },
                {
                   "id": "090507",
                   "name": "迁西",
                   "weatherCode": "101090507"
                },
                {
                   "id": "090508",
                   "name": "玉田",
                   "weatherCode": "101090508"
                },
                {
                   "id": "090509",
                   "name": "唐海",
                   "weatherCode": "101090509"
                },
                {
                   "id": "090510",
                   "name": "遵化",
                   "weatherCode": "101090510"
                },
                {
                   "id": "090511",
                   "name": "迁安",
                   "weatherCode": "101090511"
                },
                {
                   "id": "090512",
                   "name": "曹妃甸",
                   "weatherCode": "101090512"
                }
             ]
          },
          {
             "id": "0906",
             "name": "廊坊",
             "county": [
                {
                   "id": "090601",
                   "name": "廊坊",
                   "weatherCode": "101090601"
                },
                {
                   "id": "090602",
                   "name": "固安",
                   "weatherCode": "101090602"
                },
                {
                   "id": "090603",
                   "name": "永清",
                   "weatherCode": "101090603"
                },
                {
                   "id": "090604",
                   "name": "香河",
                   "weatherCode": "101090604"
                },
                {
                   "id": "090605",
                   "name": "大城",
                   "weatherCode": "101090605"
                },
                {
                   "id": "090606",
                   "name": "文安",
                   "weatherCode": "101090606"
                },
                {
                   "id": "090607",
                   "name": "大厂",
                   "weatherCode": "101090607"
                },
                {
                   "id": "090608",
                   "name": "霸州",
                   "weatherCode": "101090608"
                },
                {
                   "id": "090609",
                   "name": "三河",
                   "weatherCode": "101090609"
                }
             ]
          },
          {
             "id": "0907",
             "name": "沧州",
             "county": [
                {
                   "id": "090701",
                   "name": "沧州",
                   "weatherCode": "101090701"
                },
                {
                   "id": "090702",
                   "name": "青县",
                   "weatherCode": "101090702"
                },
                {
                   "id": "090703",
                   "name": "东光",
                   "weatherCode": "101090703"
                },
                {
                   "id": "090704",
                   "name": "海兴",
                   "weatherCode": "101090704"
                },
                {
                   "id": "090705",
                   "name": "盐山",
                   "weatherCode": "101090705"
                },
                {
                   "id": "090706",
                   "name": "肃宁",
                   "weatherCode": "101090706"
                },
                {
                   "id": "090707",
                   "name": "南皮",
                   "weatherCode": "101090707"
                },
                {
                   "id": "090708",
                   "name": "吴桥",
                   "weatherCode": "101090708"
                },
                {
                   "id": "090709",
                   "name": "献县",
                   "weatherCode": "101090709"
                },
                {
                   "id": "090710",
                   "name": "孟村",
                   "weatherCode": "101090710"
                },
                {
                   "id": "090711",
                   "name": "泊头",
                   "weatherCode": "101090711"
                },
                {
                   "id": "090712",
                   "name": "任丘",
                   "weatherCode": "101090712"
                },
                {
                   "id": "090713",
                   "name": "黄骅",
                   "weatherCode": "101090713"
                },
                {
                   "id": "090714",
                   "name": "河间",
                   "weatherCode": "101090714"
                },
                {
                   "id": "090715",
                   "name": "沧县",
                   "weatherCode": "101090716"
                }
             ]
          },
          {
             "id": "0908",
             "name": "衡水",
             "county": [
                {
                   "id": "090801",
                   "name": "衡水",
                   "weatherCode": "101090801"
                },
                {
                   "id": "090802",
                   "name": "枣强",
                   "weatherCode": "101090802"
                },
                {
                   "id": "090803",
                   "name": "武邑",
                   "weatherCode": "101090803"
                },
                {
                   "id": "090804",
                   "name": "武强",
                   "weatherCode": "101090804"
                },
                {
                   "id": "090805",
                   "name": "饶阳",
                   "weatherCode": "101090805"
                },
                {
                   "id": "090806",
                   "name": "安平",
                   "weatherCode": "101090806"
                },
                {
                   "id": "090807",
                   "name": "故城",
                   "weatherCode": "101090807"
                },
                {
                   "id": "090808",
                   "name": "景县",
                   "weatherCode": "101090808"
                },
                {
                   "id": "090809",
                   "name": "阜城",
                   "weatherCode": "101090809"
                },
                {
                   "id": "090810",
                   "name": "冀州",
                   "weatherCode": "101090810"
                },
                {
                   "id": "090811",
                   "name": "深州",
                   "weatherCode": "101090811"
                }
             ]
          },
          {
             "id": "0909",
             "name": "邢台",
             "county": [
                {
                   "id": "090901",
                   "name": "邢台",
                   "weatherCode": "101090901"
                },
                {
                   "id": "090902",
                   "name": "临城",
                   "weatherCode": "101090902"
                },
                {
                   "id": "090903",
                   "name": "内丘",
                   "weatherCode": "101090904"
                },
                {
                   "id": "090904",
                   "name": "柏乡",
                   "weatherCode": "101090905"
                },
                {
                   "id": "090905",
                   "name": "隆尧",
                   "weatherCode": "101090906"
                },
                {
                   "id": "090906",
                   "name": "南和",
                   "weatherCode": "101090907"
                },
                {
                   "id": "090907",
                   "name": "宁晋",
                   "weatherCode": "101090908"
                },
                {
                   "id": "090908",
                   "name": "巨鹿",
                   "weatherCode": "101090909"
                },
                {
                   "id": "090909",
                   "name": "新河",
                   "weatherCode": "101090910"
                },
                {
                   "id": "090910",
                   "name": "广宗",
                   "weatherCode": "101090911"
                },
                {
                   "id": "090911",
                   "name": "平乡",
                   "weatherCode": "101090912"
                },
                {
                   "id": "090912",
                   "name": "威县",
                   "weatherCode": "101090913"
                },
                {
                   "id": "090913",
                   "name": "清河",
                   "weatherCode": "101090914"
                },
                {
                   "id": "090914",
                   "name": "临西",
                   "weatherCode": "101090915"
                },
                {
                   "id": "090915",
                   "name": "南宫",
                   "weatherCode": "101090916"
                },
                {
                   "id": "090916",
                   "name": "沙河",
                   "weatherCode": "101090917"
                },
                {
                   "id": "090917",
                   "name": "任县",
                   "weatherCode": "101090918"
                }
             ]
          },
          {
             "id": "0910",
             "name": "邯郸",
             "county": [
                {
                   "id": "091001",
                   "name": "邯郸",
                   "weatherCode": "101091001"
                },
                {
                   "id": "091002",
                   "name": "峰峰",
                   "weatherCode": "101091002"
                },
                {
                   "id": "091003",
                   "name": "临漳",
                   "weatherCode": "101091003"
                },
                {
                   "id": "091004",
                   "name": "成安",
                   "weatherCode": "101091004"
                },
                {
                   "id": "091005",
                   "name": "大名",
                   "weatherCode": "101091005"
                },
                {
                   "id": "091006",
                   "name": "涉县",
                   "weatherCode": "101091006"
                },
                {
                   "id": "091007",
                   "name": "磁县",
                   "weatherCode": "101091007"
                },
                {
                   "id": "091008",
                   "name": "肥乡",
                   "weatherCode": "101091008"
                },
                {
                   "id": "091009",
                   "name": "永年",
                   "weatherCode": "101091009"
                },
                {
                   "id": "091010",
                   "name": "邱县",
                   "weatherCode": "101091010"
                },
                {
                   "id": "091011",
                   "name": "鸡泽",
                   "weatherCode": "101091011"
                },
                {
                   "id": "091012",
                   "name": "广平",
                   "weatherCode": "101091012"
                },
                {
                   "id": "091013",
                   "name": "馆陶",
                   "weatherCode": "101091013"
                },
                {
                   "id": "091014",
                   "name": "魏县",
                   "weatherCode": "101091014"
                },
                {
                   "id": "091015",
                   "name": "曲周",
                   "weatherCode": "101091015"
                },
                {
                   "id": "091016",
                   "name": "武安",
                   "weatherCode": "101091016"
                }
             ]
          },
          {
             "id": "0911",
             "name": "秦皇岛",
             "county": [
                {
                   "id": "091101",
                   "name": "秦皇岛",
                   "weatherCode": "101091101"
                },
                {
                   "id": "091102",
                   "name": "青龙",
                   "weatherCode": "101091102"
                },
                {
                   "id": "091103",
                   "name": "昌黎",
                   "weatherCode": "101091103"
                },
                {
                   "id": "091104",
                   "name": "抚宁",
                   "weatherCode": "101091104"
                },
                {
                   "id": "091105",
                   "name": "卢龙",
                   "weatherCode": "101091105"
                },
                {
                   "id": "091106",
                   "name": "北戴河",
                   "weatherCode": "101091106"
                }
             ]
          }
       ]
    },
    {
       "id": "10",
       "name": "山西",
       "city": [
          {
             "id": "1001",
             "name": "太原",
             "county": [
                {
                   "id": "100101",
                   "name": "太原",
                   "weatherCode": "101100101"
                },
                {
                   "id": "100102",
                   "name": "清徐",
                   "weatherCode": "101100102"
                },
                {
                   "id": "100103",
                   "name": "阳曲",
                   "weatherCode": "101100103"
                },
                {
                   "id": "100104",
                   "name": "娄烦",
                   "weatherCode": "101100104"
                },
                {
                   "id": "100105",
                   "name": "古交",
                   "weatherCode": "101100105"
                },
                {
                   "id": "100106",
                   "name": "尖草坪区",
                   "weatherCode": "101100106"
                },
                {
                   "id": "100107",
                   "name": "小店区",
                   "weatherCode": "101100107"
                }
             ]
          },
          {
             "id": "1002",
             "name": "大同",
             "county": [
                {
                   "id": "100201",
                   "name": "大同",
                   "weatherCode": "101100201"
                },
                {
                   "id": "100202",
                   "name": "阳高",
                   "weatherCode": "101100202"
                },
                {
                   "id": "100203",
                   "name": "大同县",
                   "weatherCode": "101100203"
                },
                {
                   "id": "100204",
                   "name": "天镇",
                   "weatherCode": "101100204"
                },
                {
                   "id": "100205",
                   "name": "广灵",
                   "weatherCode": "101100205"
                },
                {
                   "id": "100206",
                   "name": "灵丘",
                   "weatherCode": "101100206"
                },
                {
                   "id": "100207",
                   "name": "浑源",
                   "weatherCode": "101100207"
                },
                {
                   "id": "100208",
                   "name": "左云",
                   "weatherCode": "101100208"
                }
             ]
          },
          {
             "id": "1003",
             "name": "阳泉",
             "county": [
                {
                   "id": "100301",
                   "name": "阳泉",
                   "weatherCode": "101100301"
                },
                {
                   "id": "100302",
                   "name": "盂县",
                   "weatherCode": "101100302"
                },
                {
                   "id": "100303",
                   "name": "平定",
                   "weatherCode": "101100303"
                }
             ]
          },
          {
             "id": "1004",
             "name": "晋中",
             "county": [
                {
                   "id": "100401",
                   "name": "晋中",
                   "weatherCode": "101100401"
                },
                {
                   "id": "100402",
                   "name": "榆次",
                   "weatherCode": "101100402"
                },
                {
                   "id": "100403",
                   "name": "榆社",
                   "weatherCode": "101100403"
                },
                {
                   "id": "100404",
                   "name": "左权",
                   "weatherCode": "101100404"
                },
                {
                   "id": "100405",
                   "name": "和顺",
                   "weatherCode": "101100405"
                },
                {
                   "id": "100406",
                   "name": "昔阳",
                   "weatherCode": "101100406"
                },
                {
                   "id": "100407",
                   "name": "寿阳",
                   "weatherCode": "101100407"
                },
                {
                   "id": "100408",
                   "name": "太谷",
                   "weatherCode": "101100408"
                },
                {
                   "id": "100409",
                   "name": "祁县",
                   "weatherCode": "101100409"
                },
                {
                   "id": "100410",
                   "name": "平遥",
                   "weatherCode": "101100410"
                },
                {
                   "id": "100411",
                   "name": "灵石",
                   "weatherCode": "101100411"
                },
                {
                   "id": "100412",
                   "name": "介休",
                   "weatherCode": "101100412"
                }
             ]
          },
          {
             "id": "1005",
             "name": "长治",
             "county": [
                {
                   "id": "100501",
                   "name": "长治",
                   "weatherCode": "101100501"
                },
                {
                   "id": "100502",
                   "name": "黎城",
                   "weatherCode": "101100502"
                },
                {
                   "id": "100503",
                   "name": "屯留",
                   "weatherCode": "101100503"
                },
                {
                   "id": "100504",
                   "name": "潞城",
                   "weatherCode": "101100504"
                },
                {
                   "id": "100505",
                   "name": "襄垣",
                   "weatherCode": "101100505"
                },
                {
                   "id": "100506",
                   "name": "平顺",
                   "weatherCode": "101100506"
                },
                {
                   "id": "100507",
                   "name": "武乡",
                   "weatherCode": "101100507"
                },
                {
                   "id": "100508",
                   "name": "沁县",
                   "weatherCode": "101100508"
                },
                {
                   "id": "100509",
                   "name": "长子",
                   "weatherCode": "101100509"
                },
                {
                   "id": "100510",
                   "name": "沁源",
                   "weatherCode": "101100510"
                },
                {
                   "id": "100511",
                   "name": "壶关",
                   "weatherCode": "101100511"
                }
             ]
          },
          {
             "id": "1006",
             "name": "晋城",
             "county": [
                {
                   "id": "100601",
                   "name": "晋城",
                   "weatherCode": "101100601"
                },
                {
                   "id": "100602",
                   "name": "沁水",
                   "weatherCode": "101100602"
                },
                {
                   "id": "100603",
                   "name": "阳城",
                   "weatherCode": "101100603"
                },
                {
                   "id": "100604",
                   "name": "陵川",
                   "weatherCode": "101100604"
                },
                {
                   "id": "100605",
                   "name": "高平",
                   "weatherCode": "101100605"
                },
                {
                   "id": "100606",
                   "name": "泽州",
                   "weatherCode": "101100606"
                }
             ]
          },
          {
             "id": "1007",
             "name": "临汾",
             "county": [
                {
                   "id": "100701",
                   "name": "临汾",
                   "weatherCode": "101100701"
                },
                {
                   "id": "100702",
                   "name": "曲沃",
                   "weatherCode": "101100702"
                },
                {
                   "id": "100703",
                   "name": "永和",
                   "weatherCode": "101100703"
                },
                {
                   "id": "100704",
                   "name": "隰县",
                   "weatherCode": "101100704"
                },
                {
                   "id": "100705",
                   "name": "大宁",
                   "weatherCode": "101100705"
                },
                {
                   "id": "100706",
                   "name": "吉县",
                   "weatherCode": "101100706"
                },
                {
                   "id": "100707",
                   "name": "襄汾",
                   "weatherCode": "101100707"
                },
                {
                   "id": "100708",
                   "name": "蒲县",
                   "weatherCode": "101100708"
                },
                {
                   "id": "100709",
                   "name": "汾西",
                   "weatherCode": "101100709"
                },
                {
                   "id": "100710",
                   "name": "洪洞",
                   "weatherCode": "101100710"
                },
                {
                   "id": "100711",
                   "name": "霍州",
                   "weatherCode": "101100711"
                },
                {
                   "id": "100712",
                   "name": "乡宁",
                   "weatherCode": "101100712"
                },
                {
                   "id": "100713",
                   "name": "翼城",
                   "weatherCode": "101100713"
                },
                {
                   "id": "100714",
                   "name": "侯马",
                   "weatherCode": "101100714"
                },
                {
                   "id": "100715",
                   "name": "浮山",
                   "weatherCode": "101100715"
                },
                {
                   "id": "100716",
                   "name": "安泽",
                   "weatherCode": "101100716"
                },
                {
                   "id": "100717",
                   "name": "古县",
                   "weatherCode": "101100717"
                }
             ]
          },
          {
             "id": "1008",
             "name": "运城",
             "county": [
                {
                   "id": "100801",
                   "name": "运城",
                   "weatherCode": "101100801"
                },
                {
                   "id": "100802",
                   "name": "临猗",
                   "weatherCode": "101100802"
                },
                {
                   "id": "100803",
                   "name": "稷山",
                   "weatherCode": "101100803"
                },
                {
                   "id": "100804",
                   "name": "万荣",
                   "weatherCode": "101100804"
                },
                {
                   "id": "100805",
                   "name": "河津",
                   "weatherCode": "101100805"
                },
                {
                   "id": "100806",
                   "name": "新绛",
                   "weatherCode": "101100806"
                },
                {
                   "id": "100807",
                   "name": "绛县",
                   "weatherCode": "101100807"
                },
                {
                   "id": "100808",
                   "name": "闻喜",
                   "weatherCode": "101100808"
                },
                {
                   "id": "100809",
                   "name": "垣曲",
                   "weatherCode": "101100809"
                },
                {
                   "id": "100810",
                   "name": "永济",
                   "weatherCode": "101100810"
                },
                {
                   "id": "100811",
                   "name": "芮城",
                   "weatherCode": "101100811"
                },
                {
                   "id": "100812",
                   "name": "夏县",
                   "weatherCode": "101100812"
                },
                {
                   "id": "100813",
                   "name": "平陆",
                   "weatherCode": "101100813"
                }
             ]
          },
          {
             "id": "1009",
             "name": "朔州",
             "county": [
                {
                   "id": "100901",
                   "name": "朔州",
                   "weatherCode": "101100901"
                },
                {
                   "id": "100902",
                   "name": "平鲁",
                   "weatherCode": "101100902"
                },
                {
                   "id": "100903",
                   "name": "山阴",
                   "weatherCode": "101100903"
                },
                {
                   "id": "100904",
                   "name": "右玉",
                   "weatherCode": "101100904"
                },
                {
                   "id": "100905",
                   "name": "应县",
                   "weatherCode": "101100905"
                },
                {
                   "id": "100906",
                   "name": "怀仁",
                   "weatherCode": "101100906"
                }
             ]
          },
          {
             "id": "1010",
             "name": "忻州",
             "county": [
                {
                   "id": "101001",
                   "name": "忻州",
                   "weatherCode": "101101001"
                },
                {
                   "id": "101002",
                   "name": "定襄",
                   "weatherCode": "101101002"
                },
                {
                   "id": "101003",
                   "name": "五台县",
                   "weatherCode": "101101003"
                },
                {
                   "id": "101004",
                   "name": "河曲",
                   "weatherCode": "101101004"
                },
                {
                   "id": "101005",
                   "name": "偏关",
                   "weatherCode": "101101005"
                },
                {
                   "id": "101006",
                   "name": "神池",
                   "weatherCode": "101101006"
                },
                {
                   "id": "101007",
                   "name": "宁武",
                   "weatherCode": "101101007"
                },
                {
                   "id": "101008",
                   "name": "代县",
                   "weatherCode": "101101008"
                },
                {
                   "id": "101009",
                   "name": "繁峙",
                   "weatherCode": "101101009"
                },
                {
                   "id": "101010",
                   "name": "五台山",
                   "weatherCode": "101101010"
                },
                {
                   "id": "101011",
                   "name": "保德",
                   "weatherCode": "101101011"
                },
                {
                   "id": "101012",
                   "name": "静乐",
                   "weatherCode": "101101012"
                },
                {
                   "id": "101013",
                   "name": "岢岚",
                   "weatherCode": "101101013"
                },
                {
                   "id": "101014",
                   "name": "五寨",
                   "weatherCode": "101101014"
                },
                {
                   "id": "101015",
                   "name": "原平",
                   "weatherCode": "101101015"
                }
             ]
          },
          {
             "id": "1011",
             "name": "吕梁",
             "county": [
                {
                   "id": "101101",
                   "name": "吕梁",
                   "weatherCode": "101101100"
                },
                {
                   "id": "101102",
                   "name": "离石",
                   "weatherCode": "101101101"
                },
                {
                   "id": "101103",
                   "name": "临县",
                   "weatherCode": "101101102"
                },
                {
                   "id": "101104",
                   "name": "兴县",
                   "weatherCode": "101101103"
                },
                {
                   "id": "101105",
                   "name": "岚县",
                   "weatherCode": "101101104"
                },
                {
                   "id": "101106",
                   "name": "柳林",
                   "weatherCode": "101101105"
                },
                {
                   "id": "101107",
                   "name": "石楼",
                   "weatherCode": "101101106"
                },
                {
                   "id": "101108",
                   "name": "方山",
                   "weatherCode": "101101107"
                },
                {
                   "id": "101109",
                   "name": "交口",
                   "weatherCode": "101101108"
                },
                {
                   "id": "101110",
                   "name": "中阳",
                   "weatherCode": "101101109"
                },
                {
                   "id": "101111",
                   "name": "孝义",
                   "weatherCode": "101101110"
                },
                {
                   "id": "101112",
                   "name": "汾阳",
                   "weatherCode": "101101111"
                },
                {
                   "id": "101113",
                   "name": "文水",
                   "weatherCode": "101101112"
                },
                {
                   "id": "101114",
                   "name": "交城",
                   "weatherCode": "101101113"
                }
             ]
          }
       ]
    },
    {
       "id": "11",
       "name": "陕西",
       "city": [
          {
             "id": "1101",
             "name": "西安",
             "county": [
                {
                   "id": "110101",
                   "name": "西安",
                   "weatherCode": "101110101"
                },
                {
                   "id": "110102",
                   "name": "长安",
                   "weatherCode": "101110102"
                },
                {
                   "id": "110103",
                   "name": "临潼",
                   "weatherCode": "101110103"
                },
                {
                   "id": "110104",
                   "name": "蓝田",
                   "weatherCode": "101110104"
                },
                {
                   "id": "110105",
                   "name": "周至",
                   "weatherCode": "101110105"
                },
                {
                   "id": "110106",
                   "name": "户县",
                   "weatherCode": "101110106"
                },
                {
                   "id": "110107",
                   "name": "高陵",
                   "weatherCode": "101110107"
                }
             ]
          },
          {
             "id": "1102",
             "name": "咸阳",
             "county": [
                {
                   "id": "110201",
                   "name": "咸阳",
                   "weatherCode": "101110200"
                },
                {
                   "id": "110202",
                   "name": "三原",
                   "weatherCode": "101110201"
                },
                {
                   "id": "110203",
                   "name": "礼泉",
                   "weatherCode": "101110202"
                },
                {
                   "id": "110204",
                   "name": "永寿",
                   "weatherCode": "101110203"
                },
                {
                   "id": "110205",
                   "name": "淳化",
                   "weatherCode": "101110204"
                },
                {
                   "id": "110206",
                   "name": "泾阳",
                   "weatherCode": "101110205"
                },
                {
                   "id": "110207",
                   "name": "武功",
                   "weatherCode": "101110206"
                },
                {
                   "id": "110208",
                   "name": "乾县",
                   "weatherCode": "101110207"
                },
                {
                   "id": "110209",
                   "name": "彬县",
                   "weatherCode": "101110208"
                },
                {
                   "id": "110210",
                   "name": "长武",
                   "weatherCode": "101110209"
                },
                {
                   "id": "110211",
                   "name": "旬邑",
                   "weatherCode": "101110210"
                },
                {
                   "id": "110212",
                   "name": "兴平",
                   "weatherCode": "101110211"
                }
             ]
          },
          {
             "id": "1103",
             "name": "延安",
             "county": [
                {
                   "id": "110301",
                   "name": "延安",
                   "weatherCode": "101110300"
                },
                {
                   "id": "110302",
                   "name": "延长",
                   "weatherCode": "101110301"
                },
                {
                   "id": "110303",
                   "name": "延川",
                   "weatherCode": "101110302"
                },
                {
                   "id": "110304",
                   "name": "子长",
                   "weatherCode": "101110303"
                },
                {
                   "id": "110305",
                   "name": "宜川",
                   "weatherCode": "101110304"
                },
                {
                   "id": "110306",
                   "name": "富县",
                   "weatherCode": "101110305"
                },
                {
                   "id": "110307",
                   "name": "志丹",
                   "weatherCode": "101110306"
                },
                {
                   "id": "110308",
                   "name": "安塞",
                   "weatherCode": "101110307"
                },
                {
                   "id": "110309",
                   "name": "甘泉",
                   "weatherCode": "101110308"
                },
                {
                   "id": "110310",
                   "name": "洛川",
                   "weatherCode": "101110309"
                },
                {
                   "id": "110311",
                   "name": "黄陵",
                   "weatherCode": "101110310"
                },
                {
                   "id": "110312",
                   "name": "黄龙",
                   "weatherCode": "101110311"
                },
                {
                   "id": "110313",
                   "name": "吴起",
                   "weatherCode": "101110312"
                }
             ]
          },
          {
             "id": "1104",
             "name": "榆林",
             "county": [
                {
                   "id": "110401",
                   "name": "榆林",
                   "weatherCode": "101110401"
                },
                {
                   "id": "110402",
                   "name": "府谷",
                   "weatherCode": "101110402"
                },
                {
                   "id": "110403",
                   "name": "神木",
                   "weatherCode": "101110403"
                },
                {
                   "id": "110404",
                   "name": "佳县",
                   "weatherCode": "101110404"
                },
                {
                   "id": "110405",
                   "name": "定边",
                   "weatherCode": "101110405"
                },
                {
                   "id": "110406",
                   "name": "靖边",
                   "weatherCode": "101110406"
                },
                {
                   "id": "110407",
                   "name": "横山",
                   "weatherCode": "101110407"
                },
                {
                   "id": "110408",
                   "name": "米脂",
                   "weatherCode": "101110408"
                },
                {
                   "id": "110409",
                   "name": "子洲",
                   "weatherCode": "101110409"
                },
                {
                   "id": "110410",
                   "name": "绥德",
                   "weatherCode": "101110410"
                },
                {
                   "id": "110411",
                   "name": "吴堡",
                   "weatherCode": "101110411"
                },
                {
                   "id": "110412",
                   "name": "清涧",
                   "weatherCode": "101110412"
                },
                {
                   "id": "110413",
                   "name": "榆阳",
                   "weatherCode": "101110413"
                }
             ]
          },
          {
             "id": "1105",
             "name": "渭南",
             "county": [
                {
                   "id": "110501",
                   "name": "渭南",
                   "weatherCode": "101110501"
                },
                {
                   "id": "110502",
                   "name": "华县",
                   "weatherCode": "101110502"
                },
                {
                   "id": "110503",
                   "name": "潼关",
                   "weatherCode": "101110503"
                },
                {
                   "id": "110504",
                   "name": "大荔",
                   "weatherCode": "101110504"
                },
                {
                   "id": "110505",
                   "name": "白水",
                   "weatherCode": "101110505"
                },
                {
                   "id": "110506",
                   "name": "富平",
                   "weatherCode": "101110506"
                },
                {
                   "id": "110507",
                   "name": "蒲城",
                   "weatherCode": "101110507"
                },
                {
                   "id": "110508",
                   "name": "澄城",
                   "weatherCode": "101110508"
                },
                {
                   "id": "110509",
                   "name": "合阳",
                   "weatherCode": "101110509"
                },
                {
                   "id": "110510",
                   "name": "韩城",
                   "weatherCode": "101110510"
                },
                {
                   "id": "110511",
                   "name": "华阴",
                   "weatherCode": "101110511"
                }
             ]
          },
          {
             "id": "1106",
             "name": "商洛",
             "county": [
                {
                   "id": "110601",
                   "name": "商洛",
                   "weatherCode": "101110601"
                },
                {
                   "id": "110602",
                   "name": "洛南",
                   "weatherCode": "101110602"
                },
                {
                   "id": "110603",
                   "name": "柞水",
                   "weatherCode": "101110603"
                },
                {
                   "id": "110604",
                   "name": "商州",
                   "weatherCode": "101110604"
                },
                {
                   "id": "110605",
                   "name": "镇安",
                   "weatherCode": "101110605"
                },
                {
                   "id": "110606",
                   "name": "丹凤",
                   "weatherCode": "101110606"
                },
                {
                   "id": "110607",
                   "name": "商南",
                   "weatherCode": "101110607"
                },
                {
                   "id": "110608",
                   "name": "山阳",
                   "weatherCode": "101110608"
                }
             ]
          },
          {
             "id": "1107",
             "name": "安康",
             "county": [
                {
                   "id": "110701",
                   "name": "安康",
                   "weatherCode": "101110701"
                },
                {
                   "id": "110702",
                   "name": "紫阳",
                   "weatherCode": "101110702"
                },
                {
                   "id": "110703",
                   "name": "石泉",
                   "weatherCode": "101110703"
                },
                {
                   "id": "110704",
                   "name": "汉阴",
                   "weatherCode": "101110704"
                },
                {
                   "id": "110705",
                   "name": "旬阳",
                   "weatherCode": "101110705"
                },
                {
                   "id": "110706",
                   "name": "岚皋",
                   "weatherCode": "101110706"
                },
                {
                   "id": "110707",
                   "name": "平利",
                   "weatherCode": "101110707"
                },
                {
                   "id": "110708",
                   "name": "白河",
                   "weatherCode": "101110708"
                },
                {
                   "id": "110709",
                   "name": "镇坪",
                   "weatherCode": "101110709"
                },
                {
                   "id": "110710",
                   "name": "宁陕",
                   "weatherCode": "101110710"
                }
             ]
          },
          {
             "id": "1108",
             "name": "汉中",
             "county": [
                {
                   "id": "110801",
                   "name": "汉中",
                   "weatherCode": "101110801"
                },
                {
                   "id": "110802",
                   "name": "略阳",
                   "weatherCode": "101110802"
                },
                {
                   "id": "110803",
                   "name": "勉县",
                   "weatherCode": "101110803"
                },
                {
                   "id": "110804",
                   "name": "留坝",
                   "weatherCode": "101110804"
                },
                {
                   "id": "110805",
                   "name": "洋县",
                   "weatherCode": "101110805"
                },
                {
                   "id": "110806",
                   "name": "城固",
                   "weatherCode": "101110806"
                },
                {
                   "id": "110807",
                   "name": "西乡",
                   "weatherCode": "101110807"
                },
                {
                   "id": "110808",
                   "name": "佛坪",
                   "weatherCode": "101110808"
                },
                {
                   "id": "110809",
                   "name": "宁强",
                   "weatherCode": "101110809"
                },
                {
                   "id": "110810",
                   "name": "南郑",
                   "weatherCode": "101110810"
                },
                {
                   "id": "110811",
                   "name": "镇巴",
                   "weatherCode": "101110811"
                }
             ]
          },
          {
             "id": "1109",
             "name": "宝鸡",
             "county": [
                {
                   "id": "110901",
                   "name": "宝鸡",
                   "weatherCode": "101110901"
                },
                {
                   "id": "110902",
                   "name": "千阳",
                   "weatherCode": "101110903"
                },
                {
                   "id": "110903",
                   "name": "麟游",
                   "weatherCode": "101110904"
                },
                {
                   "id": "110904",
                   "name": "岐山",
                   "weatherCode": "101110905"
                },
                {
                   "id": "110905",
                   "name": "凤翔",
                   "weatherCode": "101110906"
                },
                {
                   "id": "110906",
                   "name": "扶风",
                   "weatherCode": "101110907"
                },
                {
                   "id": "110907",
                   "name": "眉县",
                   "weatherCode": "101110908"
                },
                {
                   "id": "110908",
                   "name": "太白",
                   "weatherCode": "101110909"
                },
                {
                   "id": "110909",
                   "name": "凤县",
                   "weatherCode": "101110910"
                },
                {
                   "id": "110910",
                   "name": "陇县",
                   "weatherCode": "101110911"
                },
                {
                   "id": "110911",
                   "name": "陈仓",
                   "weatherCode": "101110912"
                }
             ]
          },
          {
             "id": "1110",
             "name": "铜川",
             "county": [
                {
                   "id": "111001",
                   "name": "铜川",
                   "weatherCode": "101111001"
                },
                {
                   "id": "111002",
                   "name": "耀县",
                   "weatherCode": "101111002"
                },
                {
                   "id": "111003",
                   "name": "宜君",
                   "weatherCode": "101111003"
                },
                {
                   "id": "111004",
                   "name": "耀州",
                   "weatherCode": "101111004"
                }
             ]
          },
          {
             "id": "1111",
             "name": "杨凌",
             "county": [{
                "id": "111101",
                "name": "杨凌",
                "weatherCode": "101111101"
             }]
          }
       ]
    },
    {
       "id": "12",
       "name": "山东",
       "city": [
          {
             "id": "1201",
             "name": "济南",
             "county": [
                {
                   "id": "120101",
                   "name": "济南",
                   "weatherCode": "101120101"
                },
                {
                   "id": "120102",
                   "name": "长清",
                   "weatherCode": "101120102"
                },
                {
                   "id": "120103",
                   "name": "商河",
                   "weatherCode": "101120103"
                },
                {
                   "id": "120104",
                   "name": "章丘",
                   "weatherCode": "101120104"
                },
                {
                   "id": "120105",
                   "name": "平阴",
                   "weatherCode": "101120105"
                },
                {
                   "id": "120106",
                   "name": "济阳",
                   "weatherCode": "101120106"
                }
             ]
          },
          {
             "id": "1202",
             "name": "青岛",
             "county": [
                {
                   "id": "120201",
                   "name": "青岛",
                   "weatherCode": "101120201"
                },
                {
                   "id": "120202",
                   "name": "崂山",
                   "weatherCode": "101120202"
                },
                {
                   "id": "120203",
                   "name": "即墨",
                   "weatherCode": "101120204"
                },
                {
                   "id": "120204",
                   "name": "胶州",
                   "weatherCode": "101120205"
                },
                {
                   "id": "120205",
                   "name": "胶南",
                   "weatherCode": "101120206"
                },
                {
                   "id": "120206",
                   "name": "莱西",
                   "weatherCode": "101120207"
                },
                {
                   "id": "120207",
                   "name": "平度",
                   "weatherCode": "101120208"
                }
             ]
          },
          {
             "id": "1203",
             "name": "淄博",
             "county": [
                {
                   "id": "120301",
                   "name": "淄博",
                   "weatherCode": "101120301"
                },
                {
                   "id": "120302",
                   "name": "淄川",
                   "weatherCode": "101120302"
                },
                {
                   "id": "120303",
                   "name": "博山",
                   "weatherCode": "101120303"
                },
                {
                   "id": "120304",
                   "name": "高青",
                   "weatherCode": "101120304"
                },
                {
                   "id": "120305",
                   "name": "周村",
                   "weatherCode": "101120305"
                },
                {
                   "id": "120306",
                   "name": "沂源",
                   "weatherCode": "101120306"
                },
                {
                   "id": "120307",
                   "name": "桓台",
                   "weatherCode": "101120307"
                },
                {
                   "id": "120308",
                   "name": "临淄",
                   "weatherCode": "101120308"
                }
             ]
          },
          {
             "id": "1204",
             "name": "德州",
             "county": [
                {
                   "id": "120401",
                   "name": "德州",
                   "weatherCode": "101120401"
                },
                {
                   "id": "120402",
                   "name": "武城",
                   "weatherCode": "101120402"
                },
                {
                   "id": "120403",
                   "name": "临邑",
                   "weatherCode": "101120403"
                },
                {
                   "id": "120404",
                   "name": "陵县",
                   "weatherCode": "101120404"
                },
                {
                   "id": "120405",
                   "name": "齐河",
                   "weatherCode": "101120405"
                },
                {
                   "id": "120406",
                   "name": "乐陵",
                   "weatherCode": "101120406"
                },
                {
                   "id": "120407",
                   "name": "庆云",
                   "weatherCode": "101120407"
                },
                {
                   "id": "120408",
                   "name": "平原",
                   "weatherCode": "101120408"
                },
                {
                   "id": "120409",
                   "name": "宁津",
                   "weatherCode": "101120409"
                },
                {
                   "id": "120410",
                   "name": "夏津",
                   "weatherCode": "101120410"
                },
                {
                   "id": "120411",
                   "name": "禹城",
                   "weatherCode": "101120411"
                }
             ]
          },
          {
             "id": "1205",
             "name": "烟台",
             "county": [
                {
                   "id": "120501",
                   "name": "烟台",
                   "weatherCode": "101120501"
                },
                {
                   "id": "120502",
                   "name": "莱州",
                   "weatherCode": "101120502"
                },
                {
                   "id": "120503",
                   "name": "长岛",
                   "weatherCode": "101120503"
                },
                {
                   "id": "120504",
                   "name": "蓬莱",
                   "weatherCode": "101120504"
                },
                {
                   "id": "120505",
                   "name": "龙口",
                   "weatherCode": "101120505"
                },
                {
                   "id": "120506",
                   "name": "招远",
                   "weatherCode": "101120506"
                },
                {
                   "id": "120507",
                   "name": "栖霞",
                   "weatherCode": "101120507"
                },
                {
                   "id": "120508",
                   "name": "福山",
                   "weatherCode": "101120508"
                },
                {
                   "id": "120509",
                   "name": "牟平",
                   "weatherCode": "101120509"
                },
                {
                   "id": "120510",
                   "name": "莱阳",
                   "weatherCode": "101120510"
                },
                {
                   "id": "120511",
                   "name": "海阳",
                   "weatherCode": "101120511"
                }
             ]
          },
          {
             "id": "1206",
             "name": "潍坊",
             "county": [
                {
                   "id": "120601",
                   "name": "潍坊",
                   "weatherCode": "101120601"
                },
                {
                   "id": "120602",
                   "name": "青州",
                   "weatherCode": "101120602"
                },
                {
                   "id": "120603",
                   "name": "寿光",
                   "weatherCode": "101120603"
                },
                {
                   "id": "120604",
                   "name": "临朐",
                   "weatherCode": "101120604"
                },
                {
                   "id": "120605",
                   "name": "昌乐",
                   "weatherCode": "101120605"
                },
                {
                   "id": "120606",
                   "name": "昌邑",
                   "weatherCode": "101120606"
                },
                {
                   "id": "120607",
                   "name": "安丘",
                   "weatherCode": "101120607"
                },
                {
                   "id": "120608",
                   "name": "高密",
                   "weatherCode": "101120608"
                },
                {
                   "id": "120609",
                   "name": "诸城",
                   "weatherCode": "101120609"
                }
             ]
          },
          {
             "id": "1207",
             "name": "济宁",
             "county": [
                {
                   "id": "120701",
                   "name": "济宁",
                   "weatherCode": "101120701"
                },
                {
                   "id": "120702",
                   "name": "嘉祥",
                   "weatherCode": "101120702"
                },
                {
                   "id": "120703",
                   "name": "微山",
                   "weatherCode": "101120703"
                },
                {
                   "id": "120704",
                   "name": "鱼台",
                   "weatherCode": "101120704"
                },
                {
                   "id": "120705",
                   "name": "兖州",
                   "weatherCode": "101120705"
                },
                {
                   "id": "120706",
                   "name": "金乡",
                   "weatherCode": "101120706"
                },
                {
                   "id": "120707",
                   "name": "汶上",
                   "weatherCode": "101120707"
                },
                {
                   "id": "120708",
                   "name": "泗水",
                   "weatherCode": "101120708"
                },
                {
                   "id": "120709",
                   "name": "梁山",
                   "weatherCode": "101120709"
                },
                {
                   "id": "120710",
                   "name": "曲阜",
                   "weatherCode": "101120710"
                },
                {
                   "id": "120711",
                   "name": "邹城",
                   "weatherCode": "101120711"
                }
             ]
          },
          {
             "id": "1208",
             "name": "泰安",
             "county": [
                {
                   "id": "120801",
                   "name": "泰安",
                   "weatherCode": "101120801"
                },
                {
                   "id": "120802",
                   "name": "新泰",
                   "weatherCode": "101120802"
                },
                {
                   "id": "120803",
                   "name": "肥城",
                   "weatherCode": "101120804"
                },
                {
                   "id": "120804",
                   "name": "东平",
                   "weatherCode": "101120805"
                },
                {
                   "id": "120805",
                   "name": "宁阳",
                   "weatherCode": "101120806"
                }
             ]
          },
          {
             "id": "1209",
             "name": "临沂",
             "county": [
                {
                   "id": "120901",
                   "name": "临沂",
                   "weatherCode": "101120901"
                },
                {
                   "id": "120902",
                   "name": "莒南",
                   "weatherCode": "101120902"
                },
                {
                   "id": "120903",
                   "name": "沂南",
                   "weatherCode": "101120903"
                },
                {
                   "id": "120904",
                   "name": "苍山",
                   "weatherCode": "101120904"
                },
                {
                   "id": "120905",
                   "name": "临沭",
                   "weatherCode": "101120905"
                },
                {
                   "id": "120906",
                   "name": "郯城",
                   "weatherCode": "101120906"
                },
                {
                   "id": "120907",
                   "name": "蒙阴",
                   "weatherCode": "101120907"
                },
                {
                   "id": "120908",
                   "name": "平邑",
                   "weatherCode": "101120908"
                },
                {
                   "id": "120909",
                   "name": "费县",
                   "weatherCode": "101120909"
                },
                {
                   "id": "120910",
                   "name": "沂水",
                   "weatherCode": "101120910"
                }
             ]
          },
          {
             "id": "1210",
             "name": "菏泽",
             "county": [
                {
                   "id": "121001",
                   "name": "菏泽",
                   "weatherCode": "101121001"
                },
                {
                   "id": "121002",
                   "name": "鄄城",
                   "weatherCode": "101121002"
                },
                {
                   "id": "121003",
                   "name": "郓城",
                   "weatherCode": "101121003"
                },
                {
                   "id": "121004",
                   "name": "东明",
                   "weatherCode": "101121004"
                },
                {
                   "id": "121005",
                   "name": "定陶",
                   "weatherCode": "101121005"
                },
                {
                   "id": "121006",
                   "name": "巨野",
                   "weatherCode": "101121006"
                },
                {
                   "id": "121007",
                   "name": "曹县",
                   "weatherCode": "101121007"
                },
                {
                   "id": "121008",
                   "name": "成武",
                   "weatherCode": "101121008"
                },
                {
                   "id": "121009",
                   "name": "单县",
                   "weatherCode": "101121009"
                }
             ]
          },
          {
             "id": "1211",
             "name": "滨州",
             "county": [
                {
                   "id": "121101",
                   "name": "滨州",
                   "weatherCode": "101121101"
                },
                {
                   "id": "121102",
                   "name": "博兴",
                   "weatherCode": "101121102"
                },
                {
                   "id": "121103",
                   "name": "无棣",
                   "weatherCode": "101121103"
                },
                {
                   "id": "121104",
                   "name": "阳信",
                   "weatherCode": "101121104"
                },
                {
                   "id": "121105",
                   "name": "惠民",
                   "weatherCode": "101121105"
                },
                {
                   "id": "121106",
                   "name": "沾化",
                   "weatherCode": "101121106"
                },
                {
                   "id": "121107",
                   "name": "邹平",
                   "weatherCode": "101121107"
                }
             ]
          },
          {
             "id": "1212",
             "name": "东营",
             "county": [
                {
                   "id": "121201",
                   "name": "东营",
                   "weatherCode": "101121201"
                },
                {
                   "id": "121202",
                   "name": "河口",
                   "weatherCode": "101121202"
                },
                {
                   "id": "121203",
                   "name": "垦利",
                   "weatherCode": "101121203"
                },
                {
                   "id": "121204",
                   "name": "利津",
                   "weatherCode": "101121204"
                },
                {
                   "id": "121205",
                   "name": "广饶",
                   "weatherCode": "101121205"
                }
             ]
          },
          {
             "id": "1213",
             "name": "威海",
             "county": [
                {
                   "id": "121301",
                   "name": "威海",
                   "weatherCode": "101121301"
                },
                {
                   "id": "121302",
                   "name": "文登",
                   "weatherCode": "101121302"
                },
                {
                   "id": "121303",
                   "name": "荣成",
                   "weatherCode": "101121303"
                },
                {
                   "id": "121304",
                   "name": "乳山",
                   "weatherCode": "101121304"
                },
                {
                   "id": "121305",
                   "name": "成山头",
                   "weatherCode": "101121305"
                },
                {
                   "id": "121306",
                   "name": "石岛",
                   "weatherCode": "101121306"
                }
             ]
          },
          {
             "id": "1214",
             "name": "枣庄",
             "county": [
                {
                   "id": "121401",
                   "name": "枣庄",
                   "weatherCode": "101121401"
                },
                {
                   "id": "121402",
                   "name": "薛城",
                   "weatherCode": "101121402"
                },
                {
                   "id": "121403",
                   "name": "峄城",
                   "weatherCode": "101121403"
                },
                {
                   "id": "121404",
                   "name": "台儿庄",
                   "weatherCode": "101121404"
                },
                {
                   "id": "121405",
                   "name": "滕州",
                   "weatherCode": "101121405"
                }
             ]
          },
          {
             "id": "1215",
             "name": "日照",
             "county": [
                {
                   "id": "121501",
                   "name": "日照",
                   "weatherCode": "101121501"
                },
                {
                   "id": "121502",
                   "name": "五莲",
                   "weatherCode": "101121502"
                },
                {
                   "id": "121503",
                   "name": "莒县",
                   "weatherCode": "101121503"
                }
             ]
          },
          {
             "id": "1216",
             "name": "莱芜",
             "county": [{
                "id": "121601",
                "name": "莱芜",
                "weatherCode": "101121601"
             }]
          },
          {
             "id": "1217",
             "name": "聊城",
             "county": [
                {
                   "id": "121701",
                   "name": "聊城",
                   "weatherCode": "101121701"
                },
                {
                   "id": "121702",
                   "name": "冠县",
                   "weatherCode": "101121702"
                },
                {
                   "id": "121703",
                   "name": "阳谷",
                   "weatherCode": "101121703"
                },
                {
                   "id": "121704",
                   "name": "高唐",
                   "weatherCode": "101121704"
                },
                {
                   "id": "121705",
                   "name": "茌平",
                   "weatherCode": "101121705"
                },
                {
                   "id": "121706",
                   "name": "东阿",
                   "weatherCode": "101121706"
                },
                {
                   "id": "121707",
                   "name": "临清",
                   "weatherCode": "101121707"
                },
                {
                   "id": "121708",
                   "name": "莘县",
                   "weatherCode": "101121709"
                }
             ]
          }
       ]
    },
    {
       "id": "13",
       "name": "新疆",
       "city": [
          {
             "id": "1301",
             "name": "乌鲁木齐",
             "county": [
                {
                   "id": "130101",
                   "name": "乌鲁木齐",
                   "weatherCode": "101130101"
                },
                {
                   "id": "130102",
                   "name": "小渠子",
                   "weatherCode": "101130103"
                },
                {
                   "id": "130103",
                   "name": "达坂城",
                   "weatherCode": "101130105"
                },
                {
                   "id": "130104",
                   "name": "乌鲁木齐牧试站",
                   "weatherCode": "101130108"
                },
                {
                   "id": "130105",
                   "name": "天池",
                   "weatherCode": "101130109"
                },
                {
                   "id": "130106",
                   "name": "白杨沟",
                   "weatherCode": "101130110"
                }
             ]
          },
          {
             "id": "1302",
             "name": "克拉玛依",
             "county": [
                {
                   "id": "130201",
                   "name": "克拉玛依",
                   "weatherCode": "101130201"
                },
                {
                   "id": "130202",
                   "name": "乌尔禾",
                   "weatherCode": "101130202"
                },
                {
                   "id": "130203",
                   "name": "白碱滩",
                   "weatherCode": "101130203"
                }
             ]
          },
          {
             "id": "1303",
             "name": "石河子",
             "county": [
                {
                   "id": "130301",
                   "name": "石河子",
                   "weatherCode": "101130301"
                },
                {
                   "id": "130302",
                   "name": "炮台",
                   "weatherCode": "101130302"
                },
                {
                   "id": "130303",
                   "name": "莫索湾",
                   "weatherCode": "101130303"
                }
             ]
          },
          {
             "id": "1304",
             "name": "昌吉",
             "county": [
                {
                   "id": "130401",
                   "name": "昌吉",
                   "weatherCode": "101130401"
                },
                {
                   "id": "130402",
                   "name": "呼图壁",
                   "weatherCode": "101130402"
                },
                {
                   "id": "130403",
                   "name": "米泉",
                   "weatherCode": "101130403"
                },
                {
                   "id": "130404",
                   "name": "阜康",
                   "weatherCode": "101130404"
                },
                {
                   "id": "130405",
                   "name": "吉木萨尔",
                   "weatherCode": "101130405"
                },
                {
                   "id": "130406",
                   "name": "奇台",
                   "weatherCode": "101130406"
                },
                {
                   "id": "130407",
                   "name": "玛纳斯",
                   "weatherCode": "101130407"
                },
                {
                   "id": "130408",
                   "name": "木垒",
                   "weatherCode": "101130408"
                },
                {
                   "id": "130409",
                   "name": "蔡家湖",
                   "weatherCode": "101130409"
                }
             ]
          },
          {
             "id": "1305",
             "name": "吐鲁番",
             "county": [
                {
                   "id": "130501",
                   "name": "吐鲁番",
                   "weatherCode": "101130501"
                },
                {
                   "id": "130502",
                   "name": "托克逊",
                   "weatherCode": "101130502"
                },
                {
                   "id": "130503",
                   "name": "鄯善",
                   "weatherCode": "101130504"
                }
             ]
          },
          {
             "id": "1306",
             "name": "巴音郭楞",
             "county": [
                {
                   "id": "130601",
                   "name": "库尔勒",
                   "weatherCode": "101130601"
                },
                {
                   "id": "130602",
                   "name": "轮台",
                   "weatherCode": "101130602"
                },
                {
                   "id": "130603",
                   "name": "尉犁",
                   "weatherCode": "101130603"
                },
                {
                   "id": "130604",
                   "name": "若羌",
                   "weatherCode": "101130604"
                },
                {
                   "id": "130605",
                   "name": "且末",
                   "weatherCode": "101130605"
                },
                {
                   "id": "130606",
                   "name": "和静",
                   "weatherCode": "101130606"
                },
                {
                   "id": "130607",
                   "name": "焉耆",
                   "weatherCode": "101130607"
                },
                {
                   "id": "130608",
                   "name": "和硕",
                   "weatherCode": "101130608"
                },
                {
                   "id": "130609",
                   "name": "巴音布鲁克",
                   "weatherCode": "101130610"
                },
                {
                   "id": "130610",
                   "name": "铁干里克",
                   "weatherCode": "101130611"
                },
                {
                   "id": "130611",
                   "name": "博湖",
                   "weatherCode": "101130612"
                },
                {
                   "id": "130612",
                   "name": "塔中",
                   "weatherCode": "101130613"
                },
                {
                   "id": "130613",
                   "name": "巴仑台",
                   "weatherCode": "101130614"
                }
             ]
          },
          {
             "id": "1307",
             "name": "阿拉尔",
             "county": [{
                "id": "130701",
                "name": "阿拉尔",
                "weatherCode": "101130701"
             }]
          },
          {
             "id": "1308",
             "name": "阿克苏",
             "county": [
                {
                   "id": "130801",
                   "name": "阿克苏",
                   "weatherCode": "101130801"
                },
                {
                   "id": "130802",
                   "name": "乌什",
                   "weatherCode": "101130802"
                },
                {
                   "id": "130803",
                   "name": "温宿",
                   "weatherCode": "101130803"
                },
                {
                   "id": "130804",
                   "name": "拜城",
                   "weatherCode": "101130804"
                },
                {
                   "id": "130805",
                   "name": "新和",
                   "weatherCode": "101130805"
                },
                {
                   "id": "130806",
                   "name": "沙雅",
                   "weatherCode": "101130806"
                },
                {
                   "id": "130807",
                   "name": "库车",
                   "weatherCode": "101130807"
                },
                {
                   "id": "130808",
                   "name": "柯坪",
                   "weatherCode": "101130808"
                },
                {
                   "id": "130809",
                   "name": "阿瓦提",
                   "weatherCode": "101130809"
                }
             ]
          },
          {
             "id": "1309",
             "name": "喀什",
             "county": [
                {
                   "id": "130901",
                   "name": "喀什",
                   "weatherCode": "101130901"
                },
                {
                   "id": "130902",
                   "name": "英吉沙",
                   "weatherCode": "101130902"
                },
                {
                   "id": "130903",
                   "name": "塔什库尔干",
                   "weatherCode": "101130903"
                },
                {
                   "id": "130904",
                   "name": "麦盖提",
                   "weatherCode": "101130904"
                },
                {
                   "id": "130905",
                   "name": "莎车",
                   "weatherCode": "101130905"
                },
                {
                   "id": "130906",
                   "name": "叶城",
                   "weatherCode": "101130906"
                },
                {
                   "id": "130907",
                   "name": "泽普",
                   "weatherCode": "101130907"
                },
                {
                   "id": "130908",
                   "name": "巴楚",
                   "weatherCode": "101130908"
                },
                {
                   "id": "130909",
                   "name": "岳普湖",
                   "weatherCode": "101130909"
                },
                {
                   "id": "130910",
                   "name": "伽师",
                   "weatherCode": "101130910"
                },
                {
                   "id": "130911",
                   "name": "疏附",
                   "weatherCode": "101130911"
                },
                {
                   "id": "130912",
                   "name": "疏勒",
                   "weatherCode": "101130912"
                }
             ]
          },
          {
             "id": "1310",
             "name": "伊犁",
             "county": [
                {
                   "id": "131001",
                   "name": "伊宁",
                   "weatherCode": "101131001"
                },
                {
                   "id": "131002",
                   "name": "察布查尔",
                   "weatherCode": "101131002"
                },
                {
                   "id": "131003",
                   "name": "尼勒克",
                   "weatherCode": "101131003"
                },
                {
                   "id": "131004",
                   "name": "伊宁县",
                   "weatherCode": "101131004"
                },
                {
                   "id": "131005",
                   "name": "巩留",
                   "weatherCode": "101131005"
                },
                {
                   "id": "131006",
                   "name": "新源",
                   "weatherCode": "101131006"
                },
                {
                   "id": "131007",
                   "name": "昭苏",
                   "weatherCode": "101131007"
                },
                {
                   "id": "131008",
                   "name": "特克斯",
                   "weatherCode": "101131008"
                },
                {
                   "id": "131009",
                   "name": "霍城",
                   "weatherCode": "101131009"
                },
                {
                   "id": "131010",
                   "name": "霍尔果斯",
                   "weatherCode": "101131010"
                },
                {
                   "id": "131011",
                   "name": "奎屯",
                   "weatherCode": "101131011"
                }
             ]
          },
          {
             "id": "1311",
             "name": "塔城",
             "county": [
                {
                   "id": "131101",
                   "name": "塔城",
                   "weatherCode": "101131101"
                },
                {
                   "id": "131102",
                   "name": "裕民",
                   "weatherCode": "101131102"
                },
                {
                   "id": "131103",
                   "name": "额敏",
                   "weatherCode": "101131103"
                },
                {
                   "id": "131104",
                   "name": "和布克赛尔",
                   "weatherCode": "101131104"
                },
                {
                   "id": "131105",
                   "name": "托里",
                   "weatherCode": "101131105"
                },
                {
                   "id": "131106",
                   "name": "乌苏",
                   "weatherCode": "101131106"
                },
                {
                   "id": "131107",
                   "name": "沙湾",
                   "weatherCode": "101131107"
                }
             ]
          },
          {
             "id": "1312",
             "name": "哈密",
             "county": [
                {
                   "id": "131201",
                   "name": "哈密",
                   "weatherCode": "101131201"
                },
                {
                   "id": "131202",
                   "name": "巴里坤",
                   "weatherCode": "101131203"
                },
                {
                   "id": "131203",
                   "name": "伊吾",
                   "weatherCode": "101131204"
                }
             ]
          },
          {
             "id": "1313",
             "name": "和田",
             "county": [
                {
                   "id": "131301",
                   "name": "和田",
                   "weatherCode": "101131301"
                },
                {
                   "id": "131302",
                   "name": "皮山",
                   "weatherCode": "101131302"
                },
                {
                   "id": "131303",
                   "name": "策勒",
                   "weatherCode": "101131303"
                },
                {
                   "id": "131304",
                   "name": "墨玉",
                   "weatherCode": "101131304"
                },
                {
                   "id": "131305",
                   "name": "洛浦",
                   "weatherCode": "101131305"
                },
                {
                   "id": "131306",
                   "name": "民丰",
                   "weatherCode": "101131306"
                },
                {
                   "id": "131307",
                   "name": "于田",
                   "weatherCode": "101131307"
                }
             ]
          },
          {
             "id": "1314",
             "name": "阿勒泰",
             "county": [
                {
                   "id": "131401",
                   "name": "阿勒泰",
                   "weatherCode": "101131401"
                },
                {
                   "id": "131402",
                   "name": "哈巴河",
                   "weatherCode": "101131402"
                },
                {
                   "id": "131403",
                   "name": "吉木乃",
                   "weatherCode": "101131405"
                },
                {
                   "id": "131404",
                   "name": "布尔津",
                   "weatherCode": "101131406"
                },
                {
                   "id": "131405",
                   "name": "福海",
                   "weatherCode": "101131407"
                },
                {
                   "id": "131406",
                   "name": "富蕴",
                   "weatherCode": "101131408"
                },
                {
                   "id": "131407",
                   "name": "青河",
                   "weatherCode": "101131409"
                }
             ]
          },
          {
             "id": "1315",
             "name": "克州",
             "county": [
                {
                   "id": "131501",
                   "name": "阿图什",
                   "weatherCode": "101131501"
                },
                {
                   "id": "131502",
                   "name": "乌恰",
                   "weatherCode": "101131502"
                },
                {
                   "id": "131503",
                   "name": "阿克陶",
                   "weatherCode": "101131503"
                },
                {
                   "id": "131504",
                   "name": "阿合奇",
                   "weatherCode": "101131504"
                }
             ]
          },
          {
             "id": "1316",
             "name": "博尔塔拉",
             "county": [
                {
                   "id": "131601",
                   "name": "博乐",
                   "weatherCode": "101131601"
                },
                {
                   "id": "131602",
                   "name": "温泉",
                   "weatherCode": "101131602"
                },
                {
                   "id": "131603",
                   "name": "精河",
                   "weatherCode": "101131603"
                },
                {
                   "id": "131604",
                   "name": "阿拉山口",
                   "weatherCode": "101131606"
                }
             ]
          }
       ]
    },
    {
       "id": "14",
       "name": "西藏",
       "city": [
          {
             "id": "1401",
             "name": "拉萨",
             "county": [
                {
                   "id": "140101",
                   "name": "拉萨",
                   "weatherCode": "101140101"
                },
                {
                   "id": "140102",
                   "name": "当雄",
                   "weatherCode": "101140102"
                },
                {
                   "id": "140103",
                   "name": "尼木",
                   "weatherCode": "101140103"
                },
                {
                   "id": "140104",
                   "name": "林周",
                   "weatherCode": "101140104"
                },
                {
                   "id": "140105",
                   "name": "堆龙德庆",
                   "weatherCode": "101140105"
                },
                {
                   "id": "140106",
                   "name": "曲水",
                   "weatherCode": "101140106"
                },
                {
                   "id": "140107",
                   "name": "达孜",
                   "weatherCode": "101140107"
                },
                {
                   "id": "140108",
                   "name": "墨竹工卡",
                   "weatherCode": "101140108"
                }
             ]
          },
          {
             "id": "1402",
             "name": "日喀则",
             "county": [
                {
                   "id": "140201",
                   "name": "日喀则",
                   "weatherCode": "101140201"
                },
                {
                   "id": "140202",
                   "name": "拉孜",
                   "weatherCode": "101140202"
                },
                {
                   "id": "140203",
                   "name": "南木林",
                   "weatherCode": "101140203"
                },
                {
                   "id": "140204",
                   "name": "聂拉木",
                   "weatherCode": "101140204"
                },
                {
                   "id": "140205",
                   "name": "定日",
                   "weatherCode": "101140205"
                },
                {
                   "id": "140206",
                   "name": "江孜",
                   "weatherCode": "101140206"
                },
                {
                   "id": "140207",
                   "name": "帕里",
                   "weatherCode": "101140207"
                },
                {
                   "id": "140208",
                   "name": "仲巴",
                   "weatherCode": "101140208"
                },
                {
                   "id": "140209",
                   "name": "萨嘎",
                   "weatherCode": "101140209"
                },
                {
                   "id": "140210",
                   "name": "吉隆",
                   "weatherCode": "101140210"
                },
                {
                   "id": "140211",
                   "name": "昂仁",
                   "weatherCode": "101140211"
                },
                {
                   "id": "140212",
                   "name": "定结",
                   "weatherCode": "101140212"
                },
                {
                   "id": "140213",
                   "name": "萨迦",
                   "weatherCode": "101140213"
                },
                {
                   "id": "140214",
                   "name": "谢通门",
                   "weatherCode": "101140214"
                },
                {
                   "id": "140215",
                   "name": "岗巴",
                   "weatherCode": "101140216"
                },
                {
                   "id": "140216",
                   "name": "白朗",
                   "weatherCode": "101140217"
                },
                {
                   "id": "140217",
                   "name": "亚东",
                   "weatherCode": "101140218"
                },
                {
                   "id": "140218",
                   "name": "康马",
                   "weatherCode": "101140219"
                },
                {
                   "id": "140219",
                   "name": "仁布",
                   "weatherCode": "101140220"
                }
             ]
          },
          {
             "id": "1403",
             "name": "山南",
             "county": [
                {
                   "id": "140301",
                   "name": "山南",
                   "weatherCode": "101140301"
                },
                {
                   "id": "140302",
                   "name": "贡嘎",
                   "weatherCode": "101140302"
                },
                {
                   "id": "140303",
                   "name": "扎囊",
                   "weatherCode": "101140303"
                },
                {
                   "id": "140304",
                   "name": "加查",
                   "weatherCode": "101140304"
                },
                {
                   "id": "140305",
                   "name": "浪卡子",
                   "weatherCode": "101140305"
                },
                {
                   "id": "140306",
                   "name": "错那",
                   "weatherCode": "101140306"
                },
                {
                   "id": "140307",
                   "name": "隆子",
                   "weatherCode": "101140307"
                },
                {
                   "id": "140308",
                   "name": "泽当",
                   "weatherCode": "101140308"
                },
                {
                   "id": "140309",
                   "name": "乃东",
                   "weatherCode": "101140309"
                },
                {
                   "id": "140310",
                   "name": "桑日",
                   "weatherCode": "101140310"
                },
                {
                   "id": "140311",
                   "name": "洛扎",
                   "weatherCode": "101140311"
                },
                {
                   "id": "140312",
                   "name": "措美",
                   "weatherCode": "101140312"
                },
                {
                   "id": "140313",
                   "name": "琼结",
                   "weatherCode": "101140313"
                },
                {
                   "id": "140314",
                   "name": "曲松",
                   "weatherCode": "101140314"
                }
             ]
          },
          {
             "id": "1404",
             "name": "林芝",
             "county": [
                {
                   "id": "140401",
                   "name": "林芝",
                   "weatherCode": "101140401"
                },
                {
                   "id": "140402",
                   "name": "波密",
                   "weatherCode": "101140402"
                },
                {
                   "id": "140403",
                   "name": "米林",
                   "weatherCode": "101140403"
                },
                {
                   "id": "140404",
                   "name": "察隅",
                   "weatherCode": "101140404"
                },
                {
                   "id": "140405",
                   "name": "工布江达",
                   "weatherCode": "101140405"
                },
                {
                   "id": "140406",
                   "name": "朗县",
                   "weatherCode": "101140406"
                },
                {
                   "id": "140407",
                   "name": "墨脱",
                   "weatherCode": "101140407"
                }
             ]
          },
          {
             "id": "1405",
             "name": "昌都",
             "county": [
                {
                   "id": "140501",
                   "name": "昌都",
                   "weatherCode": "101140501"
                },
                {
                   "id": "140502",
                   "name": "丁青",
                   "weatherCode": "101140502"
                },
                {
                   "id": "140503",
                   "name": "边坝",
                   "weatherCode": "101140503"
                },
                {
                   "id": "140504",
                   "name": "洛隆",
                   "weatherCode": "101140504"
                },
                {
                   "id": "140505",
                   "name": "左贡",
                   "weatherCode": "101140505"
                },
                {
                   "id": "140506",
                   "name": "芒康",
                   "weatherCode": "101140506"
                },
                {
                   "id": "140507",
                   "name": "类乌齐",
                   "weatherCode": "101140507"
                },
                {
                   "id": "140508",
                   "name": "八宿",
                   "weatherCode": "101140508"
                },
                {
                   "id": "140509",
                   "name": "江达",
                   "weatherCode": "101140509"
                },
                {
                   "id": "140510",
                   "name": "察雅",
                   "weatherCode": "101140510"
                },
                {
                   "id": "140511",
                   "name": "贡觉",
                   "weatherCode": "101140511"
                }
             ]
          },
          {
             "id": "1406",
             "name": "那曲",
             "county": [
                {
                   "id": "140601",
                   "name": "那曲",
                   "weatherCode": "101140601"
                },
                {
                   "id": "140602",
                   "name": "尼玛",
                   "weatherCode": "101140602"
                },
                {
                   "id": "140603",
                   "name": "嘉黎",
                   "weatherCode": "101140603"
                },
                {
                   "id": "140604",
                   "name": "班戈",
                   "weatherCode": "101140604"
                },
                {
                   "id": "140605",
                   "name": "安多",
                   "weatherCode": "101140605"
                },
                {
                   "id": "140606",
                   "name": "索县",
                   "weatherCode": "101140606"
                },
                {
                   "id": "140607",
                   "name": "聂荣",
                   "weatherCode": "101140607"
                },
                {
                   "id": "140608",
                   "name": "巴青",
                   "weatherCode": "101140608"
                },
                {
                   "id": "140609",
                   "name": "比如",
                   "weatherCode": "101140609"
                },
                {
                   "id": "140610",
                   "name": "双湖",
                   "weatherCode": "101140610"
                }
             ]
          },
          {
             "id": "1407",
             "name": "阿里",
             "county": [
                {
                   "id": "140701",
                   "name": "阿里",
                   "weatherCode": "101140701"
                },
                {
                   "id": "140702",
                   "name": "改则",
                   "weatherCode": "101140702"
                },
                {
                   "id": "140703",
                   "name": "申扎",
                   "weatherCode": "101140703"
                },
                {
                   "id": "140704",
                   "name": "狮泉河",
                   "weatherCode": "101140704"
                },
                {
                   "id": "140705",
                   "name": "普兰",
                   "weatherCode": "101140705"
                },
                {
                   "id": "140706",
                   "name": "札达",
                   "weatherCode": "101140706"
                },
                {
                   "id": "140707",
                   "name": "噶尔",
                   "weatherCode": "101140707"
                },
                {
                   "id": "140708",
                   "name": "日土",
                   "weatherCode": "101140708"
                },
                {
                   "id": "140709",
                   "name": "革吉",
                   "weatherCode": "101140709"
                },
                {
                   "id": "140710",
                   "name": "措勤",
                   "weatherCode": "101140710"
                }
             ]
          }
       ]
    },
    {
       "id": "15",
       "name": "青海",
       "city": [
          {
             "id": "1501",
             "name": "西宁",
             "county": [
                {
                   "id": "150101",
                   "name": "西宁",
                   "weatherCode": "101150101"
                },
                {
                   "id": "150102",
                   "name": "大通",
                   "weatherCode": "101150102"
                },
                {
                   "id": "150103",
                   "name": "湟源",
                   "weatherCode": "101150103"
                },
                {
                   "id": "150104",
                   "name": "湟中",
                   "weatherCode": "101150104"
                }
             ]
          },
          {
             "id": "1502",
             "name": "海东",
             "county": [
                {
                   "id": "150201",
                   "name": "海东",
                   "weatherCode": "101150201"
                },
                {
                   "id": "150202",
                   "name": "乐都",
                   "weatherCode": "101150202"
                },
                {
                   "id": "150203",
                   "name": "民和",
                   "weatherCode": "101150203"
                },
                {
                   "id": "150204",
                   "name": "互助",
                   "weatherCode": "101150204"
                },
                {
                   "id": "150205",
                   "name": "化隆",
                   "weatherCode": "101150205"
                },
                {
                   "id": "150206",
                   "name": "循化",
                   "weatherCode": "101150206"
                },
                {
                   "id": "150207",
                   "name": "冷湖",
                   "weatherCode": "101150207"
                },
                {
                   "id": "150208",
                   "name": "平安",
                   "weatherCode": "101150208"
                }
             ]
          },
          {
             "id": "1503",
             "name": "黄南",
             "county": [
                {
                   "id": "150301",
                   "name": "黄南",
                   "weatherCode": "101150301"
                },
                {
                   "id": "150302",
                   "name": "尖扎",
                   "weatherCode": "101150302"
                },
                {
                   "id": "150303",
                   "name": "泽库",
                   "weatherCode": "101150303"
                },
                {
                   "id": "150305",
                   "name": "同仁",
                   "weatherCode": "101150305"
                }
             ]
          },
          {
             "id": "1504",
             "name": "海南",
             "county": [
                {
                   "id": "150401",
                   "name": "海南",
                   "weatherCode": "101150401"
                },
                {
                   "id": "150402",
                   "name": "贵德",
                   "weatherCode": "101150404"
                },
                {
                   "id": "150403",
                   "name": "兴海",
                   "weatherCode": "101150406"
                },
                {
                   "id": "150404",
                   "name": "贵南",
                   "weatherCode": "101150407"
                },
                {
                   "id": "150405",
                   "name": "同德",
                   "weatherCode": "101150408"
                },
                {
                   "id": "150406",
                   "name": "共和",
                   "weatherCode": "101150409"
                }
             ]
          },
          {
             "id": "1505",
             "name": "果洛",
             "county": [
                {
                   "id": "150501",
                   "name": "果洛",
                   "weatherCode": "101150501"
                },
                {
                   "id": "150502",
                   "name": "班玛",
                   "weatherCode": "101150502"
                },
                {
                   "id": "150503",
                   "name": "甘德",
                   "weatherCode": "101150503"
                },
                {
                   "id": "150504",
                   "name": "达日",
                   "weatherCode": "101150504"
                },
                {
                   "id": "150505",
                   "name": "久治",
                   "weatherCode": "101150505"
                },
                {
                   "id": "150506",
                   "name": "玛多",
                   "weatherCode": "101150506"
                },
                {
                   "id": "150507",
                   "name": "多县",
                   "weatherCode": "101150507"
                },
                {
                   "id": "150508",
                   "name": "玛沁",
                   "weatherCode": "101150508"
                }
             ]
          },
          {
             "id": "1506",
             "name": "玉树",
             "county": [
                {
                   "id": "150601",
                   "name": "玉树",
                   "weatherCode": "101150601"
                },
                {
                   "id": "150602",
                   "name": "称多",
                   "weatherCode": "101150602"
                },
                {
                   "id": "150603",
                   "name": "治多",
                   "weatherCode": "101150603"
                },
                {
                   "id": "150604",
                   "name": "杂多",
                   "weatherCode": "101150604"
                },
                {
                   "id": "150605",
                   "name": "囊谦",
                   "weatherCode": "101150605"
                },
                {
                   "id": "150606",
                   "name": "曲麻莱",
                   "weatherCode": "101150606"
                }
             ]
          },
          {
             "id": "1507",
             "name": "海西",
             "county": [
                {
                   "id": "150701",
                   "name": "海西",
                   "weatherCode": "101150701"
                },
                {
                   "id": "150702",
                   "name": "天峻",
                   "weatherCode": "101150708"
                },
                {
                   "id": "150703",
                   "name": "乌兰",
                   "weatherCode": "101150709"
                },
                {
                   "id": "150704",
                   "name": "茫崖",
                   "weatherCode": "101150712"
                },
                {
                   "id": "150705",
                   "name": "大柴旦",
                   "weatherCode": "101150713"
                },
                {
                   "id": "150706",
                   "name": "德令哈",
                   "weatherCode": "101150716"
                }
             ]
          },
          {
             "id": "1508",
             "name": "海北",
             "county": [
                {
                   "id": "150801",
                   "name": "海北",
                   "weatherCode": "101150801"
                },
                {
                   "id": "150802",
                   "name": "门源",
                   "weatherCode": "101150802"
                },
                {
                   "id": "150803",
                   "name": "祁连",
                   "weatherCode": "101150803"
                },
                {
                   "id": "150804",
                   "name": "海晏",
                   "weatherCode": "101150804"
                },
                {
                   "id": "150805",
                   "name": "刚察",
                   "weatherCode": "101150806"
                }
             ]
          },
          {
             "id": "1509",
             "name": "格尔木",
             "county": [
                {
                   "id": "150901",
                   "name": "格尔木",
                   "weatherCode": "101150901"
                },
                {
                   "id": "150902",
                   "name": "都兰",
                   "weatherCode": "101150902"
                }
             ]
          }
       ]
    },
    {
       "id": "16",
       "name": "甘肃",
       "city": [
          {
             "id": "1601",
             "name": "兰州",
             "county": [
                {
                   "id": "160101",
                   "name": "兰州",
                   "weatherCode": "101160101"
                },
                {
                   "id": "160102",
                   "name": "皋兰",
                   "weatherCode": "101160102"
                },
                {
                   "id": "160103",
                   "name": "永登",
                   "weatherCode": "101160103"
                },
                {
                   "id": "160104",
                   "name": "榆中",
                   "weatherCode": "101160104"
                }
             ]
          },
          {
             "id": "1602",
             "name": "定西",
             "county": [
                {
                   "id": "160201",
                   "name": "定西",
                   "weatherCode": "101160201"
                },
                {
                   "id": "160202",
                   "name": "通渭",
                   "weatherCode": "101160202"
                },
                {
                   "id": "160203",
                   "name": "陇西",
                   "weatherCode": "101160203"
                },
                {
                   "id": "160204",
                   "name": "渭源",
                   "weatherCode": "101160204"
                },
                {
                   "id": "160205",
                   "name": "临洮",
                   "weatherCode": "101160205"
                },
                {
                   "id": "160206",
                   "name": "漳县",
                   "weatherCode": "101160206"
                },
                {
                   "id": "160207",
                   "name": "岷县",
                   "weatherCode": "101160207"
                },
                {
                   "id": "160208",
                   "name": "安定",
                   "weatherCode": "101160208"
                }
             ]
          },
          {
             "id": "1603",
             "name": "平凉",
             "county": [
                {
                   "id": "160301",
                   "name": "平凉",
                   "weatherCode": "101160301"
                },
                {
                   "id": "160302",
                   "name": "泾川",
                   "weatherCode": "101160302"
                },
                {
                   "id": "160303",
                   "name": "灵台",
                   "weatherCode": "101160303"
                },
                {
                   "id": "160304",
                   "name": "崇信",
                   "weatherCode": "101160304"
                },
                {
                   "id": "160305",
                   "name": "华亭",
                   "weatherCode": "101160305"
                },
                {
                   "id": "160306",
                   "name": "庄浪",
                   "weatherCode": "101160306"
                },
                {
                   "id": "160307",
                   "name": "静宁",
                   "weatherCode": "101160307"
                },
                {
                   "id": "160308",
                   "name": "崆峒",
                   "weatherCode": "101160308"
                }
             ]
          },
          {
             "id": "1604",
             "name": "庆阳",
             "county": [
                {
                   "id": "160401",
                   "name": "西峰",
                   "weatherCode": "101160401"
                },
                {
                   "id": "160402",
                   "name": "环县",
                   "weatherCode": "101160403"
                },
                {
                   "id": "160403",
                   "name": "华池",
                   "weatherCode": "101160404"
                },
                {
                   "id": "160404",
                   "name": "合水",
                   "weatherCode": "101160405"
                },
                {
                   "id": "160405",
                   "name": "正宁",
                   "weatherCode": "101160406"
                },
                {
                   "id": "160406",
                   "name": "宁县",
                   "weatherCode": "101160407"
                },
                {
                   "id": "160407",
                   "name": "镇原",
                   "weatherCode": "101160408"
                },
                {
                   "id": "160408",
                   "name": "庆城",
                   "weatherCode": "101160409"
                }
             ]
          },
          {
             "id": "1605",
             "name": "武威",
             "county": [
                {
                   "id": "160501",
                   "name": "武威",
                   "weatherCode": "101160501"
                },
                {
                   "id": "160502",
                   "name": "民勤",
                   "weatherCode": "101160502"
                },
                {
                   "id": "160503",
                   "name": "古浪",
                   "weatherCode": "101160503"
                },
                {
                   "id": "160504",
                   "name": "天祝",
                   "weatherCode": "101160505"
                }
             ]
          },
          {
             "id": "1606",
             "name": "金昌",
             "county": [
                {
                   "id": "160601",
                   "name": "金昌",
                   "weatherCode": "101160601"
                },
                {
                   "id": "160602",
                   "name": "永昌",
                   "weatherCode": "101160602"
                }
             ]
          },
          {
             "id": "1607",
             "name": "张掖",
             "county": [
                {
                   "id": "160701",
                   "name": "张掖",
                   "weatherCode": "101160701"
                },
                {
                   "id": "160702",
                   "name": "肃南",
                   "weatherCode": "101160702"
                },
                {
                   "id": "160703",
                   "name": "民乐",
                   "weatherCode": "101160703"
                },
                {
                   "id": "160704",
                   "name": "临泽",
                   "weatherCode": "101160704"
                },
                {
                   "id": "160705",
                   "name": "高台",
                   "weatherCode": "101160705"
                },
                {
                   "id": "160706",
                   "name": "山丹",
                   "weatherCode": "101160706"
                }
             ]
          },
          {
             "id": "1608",
             "name": "酒泉",
             "county": [
                {
                   "id": "160801",
                   "name": "酒泉",
                   "weatherCode": "101160801"
                },
                {
                   "id": "160802",
                   "name": "金塔",
                   "weatherCode": "101160803"
                },
                {
                   "id": "160803",
                   "name": "阿克塞",
                   "weatherCode": "101160804"
                },
                {
                   "id": "160804",
                   "name": "瓜州",
                   "weatherCode": "101160805"
                },
                {
                   "id": "160805",
                   "name": "肃北",
                   "weatherCode": "101160806"
                },
                {
                   "id": "160806",
                   "name": "玉门",
                   "weatherCode": "101160807"
                },
                {
                   "id": "160807",
                   "name": "敦煌",
                   "weatherCode": "101160808"
                }
             ]
          },
          {
             "id": "1609",
             "name": "天水",
             "county": [
                {
                   "id": "160901",
                   "name": "天水",
                   "weatherCode": "101160901"
                },
                {
                   "id": "160902",
                   "name": "清水",
                   "weatherCode": "101160903"
                },
                {
                   "id": "160903",
                   "name": "秦安",
                   "weatherCode": "101160904"
                },
                {
                   "id": "160904",
                   "name": "甘谷",
                   "weatherCode": "101160905"
                },
                {
                   "id": "160905",
                   "name": "武山",
                   "weatherCode": "101160906"
                },
                {
                   "id": "160906",
                   "name": "张家川",
                   "weatherCode": "101160907"
                },
                {
                   "id": "160907",
                   "name": "麦积",
                   "weatherCode": "101160908"
                }
             ]
          },
          {
             "id": "1610",
             "name": "陇南",
             "county": [
                {
                   "id": "161001",
                   "name": "武都",
                   "weatherCode": "101161001"
                },
                {
                   "id": "161002",
                   "name": "成县",
                   "weatherCode": "101161002"
                },
                {
                   "id": "161003",
                   "name": "文县",
                   "weatherCode": "101161003"
                },
                {
                   "id": "161004",
                   "name": "宕昌",
                   "weatherCode": "101161004"
                },
                {
                   "id": "161005",
                   "name": "康县",
                   "weatherCode": "101161005"
                },
                {
                   "id": "161006",
                   "name": "西和",
                   "weatherCode": "101161006"
                },
                {
                   "id": "161007",
                   "name": "礼县",
                   "weatherCode": "101161007"
                },
                {
                   "id": "161008",
                   "name": "徽县",
                   "weatherCode": "101161008"
                },
                {
                   "id": "161009",
                   "name": "两当",
                   "weatherCode": "101161009"
                }
             ]
          },
          {
             "id": "1611",
             "name": "临夏",
             "county": [
                {
                   "id": "161101",
                   "name": "临夏",
                   "weatherCode": "101161101"
                },
                {
                   "id": "161102",
                   "name": "康乐",
                   "weatherCode": "101161102"
                },
                {
                   "id": "161103",
                   "name": "永靖",
                   "weatherCode": "101161103"
                },
                {
                   "id": "161104",
                   "name": "广河",
                   "weatherCode": "101161104"
                },
                {
                   "id": "161105",
                   "name": "和政",
                   "weatherCode": "101161105"
                },
                {
                   "id": "161107",
                   "name": "积石山",
                   "weatherCode": "101161107"
                }
             ]
          },
          {
             "id": "1612",
             "name": "甘南",
             "county": [
                {
                   "id": "161201",
                   "name": "合作",
                   "weatherCode": "101161201"
                },
                {
                   "id": "161202",
                   "name": "临潭",
                   "weatherCode": "101161202"
                },
                {
                   "id": "161203",
                   "name": "卓尼",
                   "weatherCode": "101161203"
                },
                {
                   "id": "161204",
                   "name": "舟曲",
                   "weatherCode": "101161204"
                },
                {
                   "id": "161205",
                   "name": "迭部",
                   "weatherCode": "101161205"
                },
                {
                   "id": "161206",
                   "name": "玛曲",
                   "weatherCode": "101161206"
                },
                {
                   "id": "161207",
                   "name": "碌曲",
                   "weatherCode": "101161207"
                },
                {
                   "id": "161208",
                   "name": "夏河",
                   "weatherCode": "101161208"
                }
             ]
          },
          {
             "id": "1613",
             "name": "白银",
             "county": [
                {
                   "id": "161301",
                   "name": "白银",
                   "weatherCode": "101161301"
                },
                {
                   "id": "161302",
                   "name": "靖远",
                   "weatherCode": "101161302"
                },
                {
                   "id": "161303",
                   "name": "会宁",
                   "weatherCode": "101161303"
                },
                {
                   "id": "161304",
                   "name": "平川",
                   "weatherCode": "101161304"
                },
                {
                   "id": "161305",
                   "name": "景泰",
                   "weatherCode": "101161305"
                }
             ]
          },
          {
             "id": "1614",
             "name": "嘉峪关",
             "county": [{
                "id": "161401",
                "name": "嘉峪关",
                "weatherCode": "101161401"
             }]
          }
       ]
    },
    {
       "id": "17",
       "name": "宁夏",
       "city": [
          {
             "id": "1701",
             "name": "银川",
             "county": [
                {
                   "id": "170101",
                   "name": "银川",
                   "weatherCode": "101170101"
                },
                {
                   "id": "170102",
                   "name": "永宁",
                   "weatherCode": "101170102"
                },
                {
                   "id": "170103",
                   "name": "灵武",
                   "weatherCode": "101170103"
                },
                {
                   "id": "170104",
                   "name": "贺兰",
                   "weatherCode": "101170104"
                }
             ]
          },
          {
             "id": "1702",
             "name": "石嘴山",
             "county": [
                {
                   "id": "170201",
                   "name": "石嘴山",
                   "weatherCode": "101170201"
                },
                {
                   "id": "170202",
                   "name": "惠农",
                   "weatherCode": "101170202"
                },
                {
                   "id": "170203",
                   "name": "平罗",
                   "weatherCode": "101170203"
                },
                {
                   "id": "170204",
                   "name": "陶乐",
                   "weatherCode": "101170204"
                }
             ]
          },
          {
             "id": "1703",
             "name": "吴忠",
             "county": [
                {
                   "id": "170301",
                   "name": "吴忠",
                   "weatherCode": "101170301"
                },
                {
                   "id": "170302",
                   "name": "同心",
                   "weatherCode": "101170302"
                },
                {
                   "id": "170303",
                   "name": "盐池",
                   "weatherCode": "101170303"
                },
                {
                   "id": "170304",
                   "name": "青铜峡",
                   "weatherCode": "101170306"
                }
             ]
          },
          {
             "id": "1704",
             "name": "固原",
             "county": [
                {
                   "id": "170401",
                   "name": "固原",
                   "weatherCode": "101170401"
                },
                {
                   "id": "170402",
                   "name": "西吉",
                   "weatherCode": "101170402"
                },
                {
                   "id": "170403",
                   "name": "隆德",
                   "weatherCode": "101170403"
                },
                {
                   "id": "170404",
                   "name": "泾源",
                   "weatherCode": "101170404"
                },
                {
                   "id": "170405",
                   "name": "彭阳",
                   "weatherCode": "101170406"
                }
             ]
          },
          {
             "id": "1705",
             "name": "中卫",
             "county": [
                {
                   "id": "170501",
                   "name": "中卫",
                   "weatherCode": "101170501"
                },
                {
                   "id": "170502",
                   "name": "中宁",
                   "weatherCode": "101170502"
                },
                {
                   "id": "170503",
                   "name": "海原",
                   "weatherCode": "101170504"
                }
             ]
          }
       ]
    },
    {
       "id": "18",
       "name": "河南",
       "city": [
          {
             "id": "1801",
             "name": "郑州",
             "county": [
                {
                   "id": "180101",
                   "name": "郑州",
                   "weatherCode": "101180101"
                },
                {
                   "id": "180102",
                   "name": "巩义",
                   "weatherCode": "101180102"
                },
                {
                   "id": "180103",
                   "name": "荥阳",
                   "weatherCode": "101180103"
                },
                {
                   "id": "180104",
                   "name": "登封",
                   "weatherCode": "101180104"
                },
                {
                   "id": "180105",
                   "name": "新密",
                   "weatherCode": "101180105"
                },
                {
                   "id": "180106",
                   "name": "新郑",
                   "weatherCode": "101180106"
                },
                {
                   "id": "180107",
                   "name": "中牟",
                   "weatherCode": "101180107"
                },
                {
                   "id": "180108",
                   "name": "上街",
                   "weatherCode": "101180108"
                }
             ]
          },
          {
             "id": "1802",
             "name": "安阳",
             "county": [
                {
                   "id": "180201",
                   "name": "安阳",
                   "weatherCode": "101180201"
                },
                {
                   "id": "180202",
                   "name": "汤阴",
                   "weatherCode": "101180202"
                },
                {
                   "id": "180203",
                   "name": "滑县",
                   "weatherCode": "101180203"
                },
                {
                   "id": "180204",
                   "name": "内黄",
                   "weatherCode": "101180204"
                },
                {
                   "id": "180205",
                   "name": "林州",
                   "weatherCode": "101180205"
                }
             ]
          },
          {
             "id": "1803",
             "name": "新乡",
             "county": [
                {
                   "id": "180301",
                   "name": "新乡",
                   "weatherCode": "101180301"
                },
                {
                   "id": "180302",
                   "name": "获嘉",
                   "weatherCode": "101180302"
                },
                {
                   "id": "180303",
                   "name": "原阳",
                   "weatherCode": "101180303"
                },
                {
                   "id": "180304",
                   "name": "辉县",
                   "weatherCode": "101180304"
                },
                {
                   "id": "180305",
                   "name": "卫辉",
                   "weatherCode": "101180305"
                },
                {
                   "id": "180306",
                   "name": "延津",
                   "weatherCode": "101180306"
                },
                {
                   "id": "180307",
                   "name": "封丘",
                   "weatherCode": "101180307"
                },
                {
                   "id": "180308",
                   "name": "长垣",
                   "weatherCode": "101180308"
                }
             ]
          },
          {
             "id": "1804",
             "name": "许昌",
             "county": [
                {
                   "id": "180401",
                   "name": "许昌",
                   "weatherCode": "101180401"
                },
                {
                   "id": "180402",
                   "name": "鄢陵",
                   "weatherCode": "101180402"
                },
                {
                   "id": "180403",
                   "name": "襄城",
                   "weatherCode": "101180403"
                },
                {
                   "id": "180404",
                   "name": "长葛",
                   "weatherCode": "101180404"
                },
                {
                   "id": "180405",
                   "name": "禹州",
                   "weatherCode": "101180405"
                }
             ]
          },
          {
             "id": "1805",
             "name": "平顶山",
             "county": [
                {
                   "id": "180501",
                   "name": "平顶山",
                   "weatherCode": "101180501"
                },
                {
                   "id": "180502",
                   "name": "郏县",
                   "weatherCode": "101180502"
                },
                {
                   "id": "180503",
                   "name": "宝丰",
                   "weatherCode": "101180503"
                },
                {
                   "id": "180504",
                   "name": "汝州",
                   "weatherCode": "101180504"
                },
                {
                   "id": "180505",
                   "name": "叶县",
                   "weatherCode": "101180505"
                },
                {
                   "id": "180506",
                   "name": "舞钢",
                   "weatherCode": "101180506"
                },
                {
                   "id": "180507",
                   "name": "鲁山",
                   "weatherCode": "101180507"
                },
                {
                   "id": "180508",
                   "name": "石龙",
                   "weatherCode": "101180508"
                }
             ]
          },
          {
             "id": "1806",
             "name": "信阳",
             "county": [
                {
                   "id": "180601",
                   "name": "信阳",
                   "weatherCode": "101180601"
                },
                {
                   "id": "180602",
                   "name": "息县",
                   "weatherCode": "101180602"
                },
                {
                   "id": "180603",
                   "name": "罗山",
                   "weatherCode": "101180603"
                },
                {
                   "id": "180604",
                   "name": "光山",
                   "weatherCode": "101180604"
                },
                {
                   "id": "180605",
                   "name": "新县",
                   "weatherCode": "101180605"
                },
                {
                   "id": "180606",
                   "name": "淮滨",
                   "weatherCode": "101180606"
                },
                {
                   "id": "180607",
                   "name": "潢川",
                   "weatherCode": "101180607"
                },
                {
                   "id": "180608",
                   "name": "固始",
                   "weatherCode": "101180608"
                },
                {
                   "id": "180609",
                   "name": "商城",
                   "weatherCode": "101180609"
                }
             ]
          },
          {
             "id": "1807",
             "name": "南阳",
             "county": [
                {
                   "id": "180701",
                   "name": "南阳",
                   "weatherCode": "101180701"
                },
                {
                   "id": "180702",
                   "name": "南召",
                   "weatherCode": "101180702"
                },
                {
                   "id": "180703",
                   "name": "方城",
                   "weatherCode": "101180703"
                },
                {
                   "id": "180704",
                   "name": "社旗",
                   "weatherCode": "101180704"
                },
                {
                   "id": "180705",
                   "name": "西峡",
                   "weatherCode": "101180705"
                },
                {
                   "id": "180706",
                   "name": "内乡",
                   "weatherCode": "101180706"
                },
                {
                   "id": "180707",
                   "name": "镇平",
                   "weatherCode": "101180707"
                },
                {
                   "id": "180708",
                   "name": "淅川",
                   "weatherCode": "101180708"
                },
                {
                   "id": "180709",
                   "name": "新野",
                   "weatherCode": "101180709"
                },
                {
                   "id": "180710",
                   "name": "唐河",
                   "weatherCode": "101180710"
                },
                {
                   "id": "180711",
                   "name": "邓州",
                   "weatherCode": "101180711"
                },
                {
                   "id": "180712",
                   "name": "桐柏",
                   "weatherCode": "101180712"
                }
             ]
          },
          {
             "id": "1808",
             "name": "开封",
             "county": [
                {
                   "id": "180801",
                   "name": "开封",
                   "weatherCode": "101180801"
                },
                {
                   "id": "180802",
                   "name": "杞县",
                   "weatherCode": "101180802"
                },
                {
                   "id": "180803",
                   "name": "尉氏",
                   "weatherCode": "101180803"
                },
                {
                   "id": "180804",
                   "name": "通许",
                   "weatherCode": "101180804"
                },
                {
                   "id": "180805",
                   "name": "兰考",
                   "weatherCode": "101180805"
                }
             ]
          },
          {
             "id": "1809",
             "name": "洛阳",
             "county": [
                {
                   "id": "180901",
                   "name": "洛阳",
                   "weatherCode": "101180901"
                },
                {
                   "id": "180902",
                   "name": "新安",
                   "weatherCode": "101180902"
                },
                {
                   "id": "180903",
                   "name": "孟津",
                   "weatherCode": "101180903"
                },
                {
                   "id": "180904",
                   "name": "宜阳",
                   "weatherCode": "101180904"
                },
                {
                   "id": "180905",
                   "name": "洛宁",
                   "weatherCode": "101180905"
                },
                {
                   "id": "180906",
                   "name": "伊川",
                   "weatherCode": "101180906"
                },
                {
                   "id": "180907",
                   "name": "嵩县",
                   "weatherCode": "101180907"
                },
                {
                   "id": "180908",
                   "name": "偃师",
                   "weatherCode": "101180908"
                },
                {
                   "id": "180909",
                   "name": "栾川",
                   "weatherCode": "101180909"
                },
                {
                   "id": "180910",
                   "name": "汝阳",
                   "weatherCode": "101180910"
                },
                {
                   "id": "180911",
                   "name": "吉利",
                   "weatherCode": "101180911"
                }
             ]
          },
          {
             "id": "1810",
             "name": "商丘",
             "county": [
                {
                   "id": "181001",
                   "name": "商丘",
                   "weatherCode": "101181001"
                },
                {
                   "id": "181002",
                   "name": "睢县",
                   "weatherCode": "101181003"
                },
                {
                   "id": "181003",
                   "name": "民权",
                   "weatherCode": "101181004"
                },
                {
                   "id": "181004",
                   "name": "虞城",
                   "weatherCode": "101181005"
                },
                {
                   "id": "181005",
                   "name": "柘城",
                   "weatherCode": "101181006"
                },
                {
                   "id": "181006",
                   "name": "宁陵",
                   "weatherCode": "101181007"
                },
                {
                   "id": "181007",
                   "name": "夏邑",
                   "weatherCode": "101181008"
                },
                {
                   "id": "181008",
                   "name": "永城",
                   "weatherCode": "101181009"
                }
             ]
          },
          {
             "id": "1811",
             "name": "焦作",
             "county": [
                {
                   "id": "181101",
                   "name": "焦作",
                   "weatherCode": "101181101"
                },
                {
                   "id": "181102",
                   "name": "修武",
                   "weatherCode": "101181102"
                },
                {
                   "id": "181103",
                   "name": "武陟",
                   "weatherCode": "101181103"
                },
                {
                   "id": "181104",
                   "name": "沁阳",
                   "weatherCode": "101181104"
                },
                {
                   "id": "181105",
                   "name": "博爱",
                   "weatherCode": "101181106"
                },
                {
                   "id": "181106",
                   "name": "温县",
                   "weatherCode": "101181107"
                },
                {
                   "id": "181107",
                   "name": "孟州",
                   "weatherCode": "101181108"
                }
             ]
          },
          {
             "id": "1812",
             "name": "鹤壁",
             "county": [
                {
                   "id": "181201",
                   "name": "鹤壁",
                   "weatherCode": "101181201"
                },
                {
                   "id": "181202",
                   "name": "浚县",
                   "weatherCode": "101181202"
                },
                {
                   "id": "181203",
                   "name": "淇县",
                   "weatherCode": "101181203"
                }
             ]
          },
          {
             "id": "1813",
             "name": "濮阳",
             "county": [
                {
                   "id": "181301",
                   "name": "濮阳",
                   "weatherCode": "101181301"
                },
                {
                   "id": "181302",
                   "name": "台前",
                   "weatherCode": "101181302"
                },
                {
                   "id": "181303",
                   "name": "南乐",
                   "weatherCode": "101181303"
                },
                {
                   "id": "181304",
                   "name": "清丰",
                   "weatherCode": "101181304"
                },
                {
                   "id": "181305",
                   "name": "范县",
                   "weatherCode": "101181305"
                }
             ]
          },
          {
             "id": "1814",
             "name": "周口",
             "county": [
                {
                   "id": "181401",
                   "name": "周口",
                   "weatherCode": "101181401"
                },
                {
                   "id": "181402",
                   "name": "扶沟",
                   "weatherCode": "101181402"
                },
                {
                   "id": "181403",
                   "name": "太康",
                   "weatherCode": "101181403"
                },
                {
                   "id": "181404",
                   "name": "淮阳",
                   "weatherCode": "101181404"
                },
                {
                   "id": "181405",
                   "name": "西华",
                   "weatherCode": "101181405"
                },
                {
                   "id": "181406",
                   "name": "商水",
                   "weatherCode": "101181406"
                },
                {
                   "id": "181407",
                   "name": "项城",
                   "weatherCode": "101181407"
                },
                {
                   "id": "181408",
                   "name": "郸城",
                   "weatherCode": "101181408"
                },
                {
                   "id": "181409",
                   "name": "鹿邑",
                   "weatherCode": "101181409"
                },
                {
                   "id": "181410",
                   "name": "沈丘",
                   "weatherCode": "101181410"
                }
             ]
          },
          {
             "id": "1815",
             "name": "漯河",
             "county": [
                {
                   "id": "181501",
                   "name": "漯河",
                   "weatherCode": "101181501"
                },
                {
                   "id": "181502",
                   "name": "临颍",
                   "weatherCode": "101181502"
                },
                {
                   "id": "181503",
                   "name": "舞阳",
                   "weatherCode": "101181503"
                }
             ]
          },
          {
             "id": "1816",
             "name": "驻马店",
             "county": [
                {
                   "id": "181601",
                   "name": "驻马店",
                   "weatherCode": "101181601"
                },
                {
                   "id": "181602",
                   "name": "西平",
                   "weatherCode": "101181602"
                },
                {
                   "id": "181603",
                   "name": "遂平",
                   "weatherCode": "101181603"
                },
                {
                   "id": "181604",
                   "name": "上蔡",
                   "weatherCode": "101181604"
                },
                {
                   "id": "181605",
                   "name": "汝南",
                   "weatherCode": "101181605"
                },
                {
                   "id": "181606",
                   "name": "泌阳",
                   "weatherCode": "101181606"
                },
                {
                   "id": "181607",
                   "name": "平舆",
                   "weatherCode": "101181607"
                },
                {
                   "id": "181608",
                   "name": "新蔡",
                   "weatherCode": "101181608"
                },
                {
                   "id": "181609",
                   "name": "确山",
                   "weatherCode": "101181609"
                },
                {
                   "id": "181610",
                   "name": "正阳",
                   "weatherCode": "101181610"
                }
             ]
          },
          {
             "id": "1817",
             "name": "三门峡",
             "county": [
                {
                   "id": "181701",
                   "name": "三门峡",
                   "weatherCode": "101181701"
                },
                {
                   "id": "181702",
                   "name": "灵宝",
                   "weatherCode": "101181702"
                },
                {
                   "id": "181703",
                   "name": "渑池",
                   "weatherCode": "101181703"
                },
                {
                   "id": "181704",
                   "name": "卢氏",
                   "weatherCode": "101181704"
                },
                {
                   "id": "181705",
                   "name": "义马",
                   "weatherCode": "101181705"
                },
                {
                   "id": "181706",
                   "name": "陕县",
                   "weatherCode": "101181706"
                }
             ]
          },
          {
             "id": "1818",
             "name": "济源",
             "county": [{
                "id": "181801",
                "name": "济源",
                "weatherCode": "101181801"
             }]
          }
       ]
    },
    {
       "id": "19",
       "name": "江苏",
       "city": [
          {
             "id": "1901",
             "name": "南京",
             "county": [
                {
                   "id": "190101",
                   "name": "南京",
                   "weatherCode": "101190101"
                },
                {
                   "id": "190102",
                   "name": "溧水",
                   "weatherCode": "101190102"
                },
                {
                   "id": "190103",
                   "name": "高淳",
                   "weatherCode": "101190103"
                },
                {
                   "id": "190104",
                   "name": "江宁",
                   "weatherCode": "101190104"
                },
                {
                   "id": "190105",
                   "name": "六合",
                   "weatherCode": "101190105"
                },
                {
                   "id": "190106",
                   "name": "江浦",
                   "weatherCode": "101190106"
                },
                {
                   "id": "190107",
                   "name": "浦口",
                   "weatherCode": "101190107"
                }
             ]
          },
          {
             "id": "1902",
             "name": "无锡",
             "county": [
                {
                   "id": "190201",
                   "name": "无锡",
                   "weatherCode": "101190201"
                },
                {
                   "id": "190202",
                   "name": "江阴",
                   "weatherCode": "101190202"
                },
                {
                   "id": "190203",
                   "name": "宜兴",
                   "weatherCode": "101190203"
                },
                {
                   "id": "190204",
                   "name": "锡山",
                   "weatherCode": "101190204"
                }
             ]
          },
          {
             "id": "1903",
             "name": "镇江",
             "county": [
                {
                   "id": "190301",
                   "name": "镇江",
                   "weatherCode": "101190301"
                },
                {
                   "id": "190302",
                   "name": "丹阳",
                   "weatherCode": "101190302"
                },
                {
                   "id": "190303",
                   "name": "扬中",
                   "weatherCode": "101190303"
                },
                {
                   "id": "190304",
                   "name": "句容",
                   "weatherCode": "101190304"
                },
                {
                   "id": "190305",
                   "name": "丹徒",
                   "weatherCode": "101190305"
                }
             ]
          },
          {
             "id": "1904",
             "name": "苏州",
             "county": [
                {
                   "id": "190401",
                   "name": "苏州",
                   "weatherCode": "101190401"
                },
                {
                   "id": "190402",
                   "name": "常熟",
                   "weatherCode": "101190402"
                },
                {
                   "id": "190403",
                   "name": "张家港",
                   "weatherCode": "101190403"
                },
                {
                   "id": "190404",
                   "name": "昆山",
                   "weatherCode": "101190404"
                },
                {
                   "id": "190405",
                   "name": "吴中",
                   "weatherCode": "101190405"
                },
                {
                   "id": "190406",
                   "name": "吴江",
                   "weatherCode": "101190407"
                },
                {
                   "id": "190407",
                   "name": "太仓",
                   "weatherCode": "101190408"
                }
             ]
          },
          {
             "id": "1905",
             "name": "南通",
             "county": [
                {
                   "id": "190501",
                   "name": "南通",
                   "weatherCode": "101190501"
                },
                {
                   "id": "190502",
                   "name": "海安",
                   "weatherCode": "101190502"
                },
                {
                   "id": "190503",
                   "name": "如皋",
                   "weatherCode": "101190503"
                },
                {
                   "id": "190504",
                   "name": "如东",
                   "weatherCode": "101190504"
                },
                {
                   "id": "190505",
                   "name": "启东",
                   "weatherCode": "101190507"
                },
                {
                   "id": "190506",
                   "name": "海门",
                   "weatherCode": "101190508"
                }
             ]
          },
          {
             "id": "1906",
             "name": "扬州",
             "county": [
                {
                   "id": "190601",
                   "name": "扬州",
                   "weatherCode": "101190601"
                },
                {
                   "id": "190602",
                   "name": "宝应",
                   "weatherCode": "101190602"
                },
                {
                   "id": "190603",
                   "name": "仪征",
                   "weatherCode": "101190603"
                },
                {
                   "id": "190604",
                   "name": "高邮",
                   "weatherCode": "101190604"
                },
                {
                   "id": "190605",
                   "name": "江都",
                   "weatherCode": "101190605"
                },
                {
                   "id": "190606",
                   "name": "邗江",
                   "weatherCode": "101190606"
                }
             ]
          },
          {
             "id": "1907",
             "name": "盐城",
             "county": [
                {
                   "id": "190701",
                   "name": "盐城",
                   "weatherCode": "101190701"
                },
                {
                   "id": "190702",
                   "name": "响水",
                   "weatherCode": "101190702"
                },
                {
                   "id": "190703",
                   "name": "滨海",
                   "weatherCode": "101190703"
                },
                {
                   "id": "190704",
                   "name": "阜宁",
                   "weatherCode": "101190704"
                },
                {
                   "id": "190705",
                   "name": "射阳",
                   "weatherCode": "101190705"
                },
                {
                   "id": "190706",
                   "name": "建湖",
                   "weatherCode": "101190706"
                },
                {
                   "id": "190707",
                   "name": "东台",
                   "weatherCode": "101190707"
                },
                {
                   "id": "190708",
                   "name": "大丰",
                   "weatherCode": "101190708"
                },
                {
                   "id": "190709",
                   "name": "盐都",
                   "weatherCode": "101190709"
                }
             ]
          },
          {
             "id": "1908",
             "name": "徐州",
             "county": [
                {
                   "id": "190801",
                   "name": "徐州",
                   "weatherCode": "101190801"
                },
                {
                   "id": "190802",
                   "name": "铜山",
                   "weatherCode": "101190802"
                },
                {
                   "id": "190803",
                   "name": "丰县",
                   "weatherCode": "101190803"
                },
                {
                   "id": "190804",
                   "name": "沛县",
                   "weatherCode": "101190804"
                },
                {
                   "id": "190805",
                   "name": "邳州",
                   "weatherCode": "101190805"
                },
                {
                   "id": "190806",
                   "name": "睢宁",
                   "weatherCode": "101190806"
                },
                {
                   "id": "190807",
                   "name": "新沂",
                   "weatherCode": "101190807"
                }
             ]
          },
          {
             "id": "1909",
             "name": "淮安",
             "county": [
                {
                   "id": "190901",
                   "name": "淮安",
                   "weatherCode": "101190901"
                },
                {
                   "id": "190902",
                   "name": "金湖",
                   "weatherCode": "101190902"
                },
                {
                   "id": "190903",
                   "name": "盱眙",
                   "weatherCode": "101190903"
                },
                {
                   "id": "190904",
                   "name": "洪泽",
                   "weatherCode": "101190904"
                },
                {
                   "id": "190905",
                   "name": "涟水",
                   "weatherCode": "101190905"
                },
                {
                   "id": "190906",
                   "name": "淮阴区",
                   "weatherCode": "101190906"
                },
                {
                   "id": "190907",
                   "name": "淮安区",
                   "weatherCode": "101190908"
                }
             ]
          },
          {
             "id": "1910",
             "name": "连云港",
             "county": [
                {
                   "id": "191001",
                   "name": "连云港",
                   "weatherCode": "101191001"
                },
                {
                   "id": "191002",
                   "name": "东海",
                   "weatherCode": "101191002"
                },
                {
                   "id": "191003",
                   "name": "赣榆",
                   "weatherCode": "101191003"
                },
                {
                   "id": "191004",
                   "name": "灌云",
                   "weatherCode": "101191004"
                },
                {
                   "id": "191005",
                   "name": "灌南",
                   "weatherCode": "101191005"
                }
             ]
          },
          {
             "id": "1911",
             "name": "常州",
             "county": [
                {
                   "id": "191101",
                   "name": "常州",
                   "weatherCode": "101191101"
                },
                {
                   "id": "191102",
                   "name": "溧阳",
                   "weatherCode": "101191102"
                },
                {
                   "id": "191103",
                   "name": "金坛",
                   "weatherCode": "101191103"
                },
                {
                   "id": "191104",
                   "name": "武进",
                   "weatherCode": "101191104"
                }
             ]
          },
          {
             "id": "1912",
             "name": "泰州",
             "county": [
                {
                   "id": "191201",
                   "name": "泰州",
                   "weatherCode": "101191201"
                },
                {
                   "id": "191202",
                   "name": "兴化",
                   "weatherCode": "101191202"
                },
                {
                   "id": "191203",
                   "name": "泰兴",
                   "weatherCode": "101191203"
                },
                {
                   "id": "191204",
                   "name": "姜堰",
                   "weatherCode": "101191204"
                },
                {
                   "id": "191205",
                   "name": "靖江",
                   "weatherCode": "101191205"
                }
             ]
          },
          {
             "id": "1913",
             "name": "宿迁",
             "county": [
                {
                   "id": "191301",
                   "name": "宿迁",
                   "weatherCode": "101191301"
                },
                {
                   "id": "191302",
                   "name": "沭阳",
                   "weatherCode": "101191302"
                },
                {
                   "id": "191303",
                   "name": "泗阳",
                   "weatherCode": "101191303"
                },
                {
                   "id": "191304",
                   "name": "泗洪",
                   "weatherCode": "101191304"
                },
                {
                   "id": "191305",
                   "name": "宿豫",
                   "weatherCode": "101191305"
                }
             ]
          }
       ]
    },
    {
       "id": "20",
       "name": "湖北",
       "city": [
          {
             "id": "2001",
             "name": "武汉",
             "county": [
                {
                   "id": "200101",
                   "name": "武汉",
                   "weatherCode": "101200101"
                },
                {
                   "id": "200102",
                   "name": "蔡甸",
                   "weatherCode": "101200102"
                },
                {
                   "id": "200103",
                   "name": "黄陂",
                   "weatherCode": "101200103"
                },
                {
                   "id": "200104",
                   "name": "新洲",
                   "weatherCode": "101200104"
                },
                {
                   "id": "200105",
                   "name": "江夏",
                   "weatherCode": "101200105"
                },
                {
                   "id": "200106",
                   "name": "东西湖",
                   "weatherCode": "101200106"
                }
             ]
          },
          {
             "id": "2002",
             "name": "襄阳",
             "county": [
                {
                   "id": "200201",
                   "name": "襄阳",
                   "weatherCode": "101200201"
                },
                {
                   "id": "200202",
                   "name": "襄州",
                   "weatherCode": "101200202"
                },
                {
                   "id": "200203",
                   "name": "保康",
                   "weatherCode": "101200203"
                },
                {
                   "id": "200204",
                   "name": "南漳",
                   "weatherCode": "101200204"
                },
                {
                   "id": "200205",
                   "name": "宜城",
                   "weatherCode": "101200205"
                },
                {
                   "id": "200206",
                   "name": "老河口",
                   "weatherCode": "101200206"
                },
                {
                   "id": "200207",
                   "name": "谷城",
                   "weatherCode": "101200207"
                },
                {
                   "id": "200208",
                   "name": "枣阳",
                   "weatherCode": "101200208"
                }
             ]
          },
          {
             "id": "2003",
             "name": "鄂州",
             "county": [
                {
                   "id": "200301",
                   "name": "鄂州",
                   "weatherCode": "101200301"
                },
                {
                   "id": "200302",
                   "name": "梁子湖",
                   "weatherCode": "101200302"
                }
             ]
          },
          {
             "id": "2004",
             "name": "孝感",
             "county": [
                {
                   "id": "200401",
                   "name": "孝感",
                   "weatherCode": "101200401"
                },
                {
                   "id": "200402",
                   "name": "安陆",
                   "weatherCode": "101200402"
                },
                {
                   "id": "200403",
                   "name": "云梦",
                   "weatherCode": "101200403"
                },
                {
                   "id": "200404",
                   "name": "大悟",
                   "weatherCode": "101200404"
                },
                {
                   "id": "200405",
                   "name": "应城",
                   "weatherCode": "101200405"
                },
                {
                   "id": "200406",
                   "name": "汉川",
                   "weatherCode": "101200406"
                },
                {
                   "id": "200407",
                   "name": "孝昌",
                   "weatherCode": "101200407"
                }
             ]
          },
          {
             "id": "2005",
             "name": "黄冈",
             "county": [
                {
                   "id": "200501",
                   "name": "黄冈",
                   "weatherCode": "101200501"
                },
                {
                   "id": "200502",
                   "name": "红安",
                   "weatherCode": "101200502"
                },
                {
                   "id": "200503",
                   "name": "麻城",
                   "weatherCode": "101200503"
                },
                {
                   "id": "200504",
                   "name": "罗田",
                   "weatherCode": "101200504"
                },
                {
                   "id": "200505",
                   "name": "英山",
                   "weatherCode": "101200505"
                },
                {
                   "id": "200506",
                   "name": "浠水",
                   "weatherCode": "101200506"
                },
                {
                   "id": "200507",
                   "name": "蕲春",
                   "weatherCode": "101200507"
                },
                {
                   "id": "200508",
                   "name": "黄梅",
                   "weatherCode": "101200508"
                },
                {
                   "id": "200509",
                   "name": "武穴",
                   "weatherCode": "101200509"
                },
                {
                   "id": "200510",
                   "name": "团风",
                   "weatherCode": "101200510"
                }
             ]
          },
          {
             "id": "2006",
             "name": "黄石",
             "county": [
                {
                   "id": "200601",
                   "name": "黄石",
                   "weatherCode": "101200601"
                },
                {
                   "id": "200602",
                   "name": "大冶",
                   "weatherCode": "101200602"
                },
                {
                   "id": "200603",
                   "name": "阳新",
                   "weatherCode": "101200603"
                },
                {
                   "id": "200604",
                   "name": "铁山",
                   "weatherCode": "101200604"
                },
                {
                   "id": "200605",
                   "name": "下陆",
                   "weatherCode": "101200605"
                },
                {
                   "id": "200606",
                   "name": "西塞山",
                   "weatherCode": "101200606"
                }
             ]
          },
          {
             "id": "2007",
             "name": "咸宁",
             "county": [
                {
                   "id": "200701",
                   "name": "咸宁",
                   "weatherCode": "101200701"
                },
                {
                   "id": "200702",
                   "name": "赤壁",
                   "weatherCode": "101200702"
                },
                {
                   "id": "200703",
                   "name": "嘉鱼",
                   "weatherCode": "101200703"
                },
                {
                   "id": "200704",
                   "name": "崇阳",
                   "weatherCode": "101200704"
                },
                {
                   "id": "200705",
                   "name": "通城",
                   "weatherCode": "101200705"
                },
                {
                   "id": "200706",
                   "name": "通山",
                   "weatherCode": "101200706"
                }
             ]
          },
          {
             "id": "2008",
             "name": "荆州",
             "county": [
                {
                   "id": "200801",
                   "name": "荆州",
                   "weatherCode": "101200801"
                },
                {
                   "id": "200802",
                   "name": "江陵",
                   "weatherCode": "101200802"
                },
                {
                   "id": "200803",
                   "name": "公安",
                   "weatherCode": "101200803"
                },
                {
                   "id": "200804",
                   "name": "石首",
                   "weatherCode": "101200804"
                },
                {
                   "id": "200805",
                   "name": "监利",
                   "weatherCode": "101200805"
                },
                {
                   "id": "200806",
                   "name": "洪湖",
                   "weatherCode": "101200806"
                },
                {
                   "id": "200807",
                   "name": "松滋",
                   "weatherCode": "101200807"
                },
                {
                   "id": "200808",
                   "name": "沙市",
                   "weatherCode": "101201406"
                }
             ]
          },
          {
             "id": "2009",
             "name": "宜昌",
             "county": [
                {
                   "id": "200901",
                   "name": "宜昌",
                   "weatherCode": "101200901"
                },
                {
                   "id": "200902",
                   "name": "远安",
                   "weatherCode": "101200902"
                },
                {
                   "id": "200903",
                   "name": "秭归",
                   "weatherCode": "101200903"
                },
                {
                   "id": "200904",
                   "name": "兴山",
                   "weatherCode": "101200904"
                },
                {
                   "id": "200905",
                   "name": "五峰",
                   "weatherCode": "101200906"
                },
                {
                   "id": "200906",
                   "name": "当阳",
                   "weatherCode": "101200907"
                },
                {
                   "id": "200907",
                   "name": "长阳",
                   "weatherCode": "101200908"
                },
                {
                   "id": "200908",
                   "name": "宜都",
                   "weatherCode": "101200909"
                },
                {
                   "id": "200909",
                   "name": "枝江",
                   "weatherCode": "101200910"
                },
                {
                   "id": "200910",
                   "name": "三峡",
                   "weatherCode": "101200911"
                },
                {
                   "id": "200911",
                   "name": "夷陵",
                   "weatherCode": "101200912"
                }
             ]
          },
          {
             "id": "2010",
             "name": "恩施",
             "county": [
                {
                   "id": "201001",
                   "name": "恩施",
                   "weatherCode": "101201001"
                },
                {
                   "id": "201002",
                   "name": "利川",
                   "weatherCode": "101201002"
                },
                {
                   "id": "201003",
                   "name": "建始",
                   "weatherCode": "101201003"
                },
                {
                   "id": "201004",
                   "name": "咸丰",
                   "weatherCode": "101201004"
                },
                {
                   "id": "201005",
                   "name": "宣恩",
                   "weatherCode": "101201005"
                },
                {
                   "id": "201006",
                   "name": "鹤峰",
                   "weatherCode": "101201006"
                },
                {
                   "id": "201007",
                   "name": "来凤",
                   "weatherCode": "101201007"
                },
                {
                   "id": "201008",
                   "name": "巴东",
                   "weatherCode": "101201008"
                }
             ]
          },
          {
             "id": "2011",
             "name": "十堰",
             "county": [
                {
                   "id": "201101",
                   "name": "十堰",
                   "weatherCode": "101201101"
                },
                {
                   "id": "201102",
                   "name": "竹溪",
                   "weatherCode": "101201102"
                },
                {
                   "id": "201103",
                   "name": "郧西",
                   "weatherCode": "101201103"
                },
                {
                   "id": "201104",
                   "name": "郧县",
                   "weatherCode": "101201104"
                },
                {
                   "id": "201105",
                   "name": "竹山",
                   "weatherCode": "101201105"
                },
                {
                   "id": "201106",
                   "name": "房县",
                   "weatherCode": "101201106"
                },
                {
                   "id": "201107",
                   "name": "丹江口",
                   "weatherCode": "101201107"
                },
                {
                   "id": "201108",
                   "name": "茅箭",
                   "weatherCode": "101201108"
                },
                {
                   "id": "201109",
                   "name": "张湾",
                   "weatherCode": "101201109"
                }
             ]
          },
          {
             "id": "2012",
             "name": "神农架",
             "county": [{
                "id": "201201",
                "name": "神农架",
                "weatherCode": "101201201"
             }]
          },
          {
             "id": "2013",
             "name": "随州",
             "county": [
                {
                   "id": "201301",
                   "name": "随州",
                   "weatherCode": "101201301"
                },
                {
                   "id": "201302",
                   "name": "广水",
                   "weatherCode": "101201302"
                }
             ]
          },
          {
             "id": "2014",
             "name": "荆门",
             "county": [
                {
                   "id": "201401",
                   "name": "荆门",
                   "weatherCode": "101201401"
                },
                {
                   "id": "201402",
                   "name": "钟祥",
                   "weatherCode": "101201402"
                },
                {
                   "id": "201403",
                   "name": "京山",
                   "weatherCode": "101201403"
                },
                {
                   "id": "201404",
                   "name": "掇刀",
                   "weatherCode": "101201404"
                },
                {
                   "id": "201405",
                   "name": "沙洋",
                   "weatherCode": "101201405"
                }
             ]
          },
          {
             "id": "2015",
             "name": "天门",
             "county": [{
                "id": "201501",
                "name": "天门",
                "weatherCode": "101201501"
             }]
          },
          {
             "id": "2016",
             "name": "仙桃",
             "county": [{
                "id": "201601",
                "name": "仙桃",
                "weatherCode": "101201601"
             }]
          },
          {
             "id": "2017",
             "name": "潜江",
             "county": [{
                "id": "201701",
                "name": "潜江",
                "weatherCode": "101201701"
             }]
          }
       ]
    },
    {
       "id": "21",
       "name": "浙江",
       "city": [
          {
             "id": "2101",
             "name": "杭州",
             "county": [
                {
                   "id": "210101",
                   "name": "杭州",
                   "weatherCode": "101210101"
                },
                {
                   "id": "210102",
                   "name": "萧山",
                   "weatherCode": "101210102"
                },
                {
                   "id": "210103",
                   "name": "桐庐",
                   "weatherCode": "101210103"
                },
                {
                   "id": "210104",
                   "name": "淳安",
                   "weatherCode": "101210104"
                },
                {
                   "id": "210105",
                   "name": "建德",
                   "weatherCode": "101210105"
                },
                {
                   "id": "210106",
                   "name": "余杭",
                   "weatherCode": "101210106"
                },
                {
                   "id": "210107",
                   "name": "临安",
                   "weatherCode": "101210107"
                },
                {
                   "id": "210108",
                   "name": "富阳",
                   "weatherCode": "101210108"
                }
             ]
          },
          {
             "id": "2102",
             "name": "湖州",
             "county": [
                {
                   "id": "210201",
                   "name": "湖州",
                   "weatherCode": "101210201"
                },
                {
                   "id": "210202",
                   "name": "长兴",
                   "weatherCode": "101210202"
                },
                {
                   "id": "210203",
                   "name": "安吉",
                   "weatherCode": "101210203"
                },
                {
                   "id": "210204",
                   "name": "德清",
                   "weatherCode": "101210204"
                }
             ]
          },
          {
             "id": "2103",
             "name": "嘉兴",
             "county": [
                {
                   "id": "210301",
                   "name": "嘉兴",
                   "weatherCode": "101210301"
                },
                {
                   "id": "210302",
                   "name": "嘉善",
                   "weatherCode": "101210302"
                },
                {
                   "id": "210303",
                   "name": "海宁",
                   "weatherCode": "101210303"
                },
                {
                   "id": "210304",
                   "name": "桐乡",
                   "weatherCode": "101210304"
                },
                {
                   "id": "210305",
                   "name": "平湖",
                   "weatherCode": "101210305"
                },
                {
                   "id": "210306",
                   "name": "海盐",
                   "weatherCode": "101210306"
                }
             ]
          },
          {
             "id": "2104",
             "name": "宁波",
             "county": [
                {
                   "id": "210401",
                   "name": "宁波",
                   "weatherCode": "101210401"
                },
                {
                   "id": "210402",
                   "name": "慈溪",
                   "weatherCode": "101210403"
                },
                {
                   "id": "210403",
                   "name": "余姚",
                   "weatherCode": "101210404"
                },
                {
                   "id": "210404",
                   "name": "奉化",
                   "weatherCode": "101210405"
                },
                {
                   "id": "210405",
                   "name": "象山",
                   "weatherCode": "101210406"
                },
                {
                   "id": "210406",
                   "name": "宁海",
                   "weatherCode": "101210408"
                },
                {
                   "id": "210407",
                   "name": "北仑",
                   "weatherCode": "101210410"
                },
                {
                   "id": "210408",
                   "name": "鄞州",
                   "weatherCode": "101210411"
                },
                {
                   "id": "210409",
                   "name": "镇海",
                   "weatherCode": "101210412"
                }
             ]
          },
          {
             "id": "2105",
             "name": "绍兴",
             "county": [
                {
                   "id": "210501",
                   "name": "绍兴",
                   "weatherCode": "101210501"
                },
                {
                   "id": "210502",
                   "name": "诸暨",
                   "weatherCode": "101210502"
                },
                {
                   "id": "210503",
                   "name": "上虞",
                   "weatherCode": "101210503"
                },
                {
                   "id": "210504",
                   "name": "新昌",
                   "weatherCode": "101210504"
                },
                {
                   "id": "210505",
                   "name": "嵊州",
                   "weatherCode": "101210505"
                }
             ]
          },
          {
             "id": "2106",
             "name": "台州",
             "county": [
                {
                   "id": "210601",
                   "name": "台州",
                   "weatherCode": "101210601"
                },
                {
                   "id": "210602",
                   "name": "玉环",
                   "weatherCode": "101210603"
                },
                {
                   "id": "210603",
                   "name": "三门",
                   "weatherCode": "101210604"
                },
                {
                   "id": "210604",
                   "name": "天台",
                   "weatherCode": "101210605"
                },
                {
                   "id": "210605",
                   "name": "仙居",
                   "weatherCode": "101210606"
                },
                {
                   "id": "210606",
                   "name": "温岭",
                   "weatherCode": "101210607"
                },
                {
                   "id": "210607",
                   "name": "洪家",
                   "weatherCode": "101210609"
                },
                {
                   "id": "210608",
                   "name": "临海",
                   "weatherCode": "101210610"
                },
                {
                   "id": "210609",
                   "name": "椒江",
                   "weatherCode": "101210611"
                },
                {
                   "id": "210610",
                   "name": "黄岩",
                   "weatherCode": "101210612"
                },
                {
                   "id": "210611",
                   "name": "路桥",
                   "weatherCode": "101210613"
                }
             ]
          },
          {
             "id": "2107",
             "name": "温州",
             "county": [
                {
                   "id": "210701",
                   "name": "温州",
                   "weatherCode": "101210701"
                },
                {
                   "id": "210702",
                   "name": "泰顺",
                   "weatherCode": "101210702"
                },
                {
                   "id": "210703",
                   "name": "文成",
                   "weatherCode": "101210703"
                },
                {
                   "id": "210704",
                   "name": "平阳",
                   "weatherCode": "101210704"
                },
                {
                   "id": "210705",
                   "name": "瑞安",
                   "weatherCode": "101210705"
                },
                {
                   "id": "210706",
                   "name": "洞头",
                   "weatherCode": "101210706"
                },
                {
                   "id": "210707",
                   "name": "乐清",
                   "weatherCode": "101210707"
                },
                {
                   "id": "210708",
                   "name": "永嘉",
                   "weatherCode": "101210708"
                },
                {
                   "id": "210709",
                   "name": "苍南",
                   "weatherCode": "101210709"
                }
             ]
          },
          {
             "id": "2108",
             "name": "丽水",
             "county": [
                {
                   "id": "210801",
                   "name": "丽水",
                   "weatherCode": "101210801"
                },
                {
                   "id": "210802",
                   "name": "遂昌",
                   "weatherCode": "101210802"
                },
                {
                   "id": "210803",
                   "name": "龙泉",
                   "weatherCode": "101210803"
                },
                {
                   "id": "210804",
                   "name": "缙云",
                   "weatherCode": "101210804"
                },
                {
                   "id": "210805",
                   "name": "青田",
                   "weatherCode": "101210805"
                },
                {
                   "id": "210806",
                   "name": "云和",
                   "weatherCode": "101210806"
                },
                {
                   "id": "210807",
                   "name": "庆元",
                   "weatherCode": "101210807"
                },
                {
                   "id": "210808",
                   "name": "松阳",
                   "weatherCode": "101210808"
                },
                {
                   "id": "210809",
                   "name": "景宁",
                   "weatherCode": "101210809"
                }
             ]
          },
          {
             "id": "2109",
             "name": "金华",
             "county": [
                {
                   "id": "210901",
                   "name": "金华",
                   "weatherCode": "101210901"
                },
                {
                   "id": "210902",
                   "name": "浦江",
                   "weatherCode": "101210902"
                },
                {
                   "id": "210903",
                   "name": "兰溪",
                   "weatherCode": "101210903"
                },
                {
                   "id": "210904",
                   "name": "义乌",
                   "weatherCode": "101210904"
                },
                {
                   "id": "210905",
                   "name": "东阳",
                   "weatherCode": "101210905"
                },
                {
                   "id": "210906",
                   "name": "武义",
                   "weatherCode": "101210906"
                },
                {
                   "id": "210907",
                   "name": "永康",
                   "weatherCode": "101210907"
                },
                {
                   "id": "210908",
                   "name": "磐安",
                   "weatherCode": "101210908"
                }
             ]
          },
          {
             "id": "2110",
             "name": "衢州",
             "county": [
                {
                   "id": "211001",
                   "name": "衢州",
                   "weatherCode": "101211001"
                },
                {
                   "id": "211002",
                   "name": "常山",
                   "weatherCode": "101211002"
                },
                {
                   "id": "211003",
                   "name": "开化",
                   "weatherCode": "101211003"
                },
                {
                   "id": "211004",
                   "name": "龙游",
                   "weatherCode": "101211004"
                },
                {
                   "id": "211005",
                   "name": "江山",
                   "weatherCode": "101211005"
                },
                {
                   "id": "211006",
                   "name": "衢江",
                   "weatherCode": "101211006"
                }
             ]
          },
          {
             "id": "2111",
             "name": "舟山",
             "county": [
                {
                   "id": "211101",
                   "name": "舟山",
                   "weatherCode": "101211101"
                },
                {
                   "id": "211102",
                   "name": "嵊泗",
                   "weatherCode": "101211102"
                },
                {
                   "id": "211103",
                   "name": "岱山",
                   "weatherCode": "101211104"
                },
                {
                   "id": "211104",
                   "name": "普陀",
                   "weatherCode": "101211105"
                },
                {
                   "id": "211105",
                   "name": "定海",
                   "weatherCode": "101211106"
                }
             ]
          }
       ]
    },
    {
       "id": "22",
       "name": "安徽",
       "city": [
          {
             "id": "2201",
             "name": "合肥",
             "county": [
                {
                   "id": "220101",
                   "name": "合肥",
                   "weatherCode": "101220101"
                },
                {
                   "id": "220102",
                   "name": "长丰",
                   "weatherCode": "101220102"
                },
                {
                   "id": "220103",
                   "name": "肥东",
                   "weatherCode": "101220103"
                },
                {
                   "id": "220104",
                   "name": "肥西",
                   "weatherCode": "101220104"
                }
             ]
          },
          {
             "id": "2202",
             "name": "蚌埠",
             "county": [
                {
                   "id": "220201",
                   "name": "蚌埠",
                   "weatherCode": "101220201"
                },
                {
                   "id": "220202",
                   "name": "怀远",
                   "weatherCode": "101220202"
                },
                {
                   "id": "220203",
                   "name": "固镇",
                   "weatherCode": "101220203"
                },
                {
                   "id": "220204",
                   "name": "五河",
                   "weatherCode": "101220204"
                }
             ]
          },
          {
             "id": "2203",
             "name": "芜湖",
             "county": [
                {
                   "id": "220301",
                   "name": "芜湖",
                   "weatherCode": "101220301"
                },
                {
                   "id": "220302",
                   "name": "繁昌",
                   "weatherCode": "101220302"
                },
                {
                   "id": "220303",
                   "name": "芜湖县",
                   "weatherCode": "101220303"
                },
                {
                   "id": "220304",
                   "name": "南陵",
                   "weatherCode": "101220304"
                }
             ]
          },
          {
             "id": "2204",
             "name": "淮南",
             "county": [
                {
                   "id": "220401",
                   "name": "淮南",
                   "weatherCode": "101220401"
                },
                {
                   "id": "220402",
                   "name": "凤台",
                   "weatherCode": "101220402"
                },
                {
                   "id": "220403",
                   "name": "潘集",
                   "weatherCode": "101220403"
                }
             ]
          },
          {
             "id": "2205",
             "name": "马鞍山",
             "county": [
                {
                   "id": "220501",
                   "name": "马鞍山",
                   "weatherCode": "101220501"
                },
                {
                   "id": "220502",
                   "name": "当涂",
                   "weatherCode": "101220502"
                }
             ]
          },
          {
             "id": "2206",
             "name": "安庆",
             "county": [
                {
                   "id": "220601",
                   "name": "安庆",
                   "weatherCode": "101220601"
                },
                {
                   "id": "220602",
                   "name": "枞阳",
                   "weatherCode": "101220602"
                },
                {
                   "id": "220603",
                   "name": "太湖",
                   "weatherCode": "101220603"
                },
                {
                   "id": "220604",
                   "name": "潜山",
                   "weatherCode": "101220604"
                },
                {
                   "id": "220605",
                   "name": "怀宁",
                   "weatherCode": "101220605"
                },
                {
                   "id": "220606",
                   "name": "宿松",
                   "weatherCode": "101220606"
                },
                {
                   "id": "220607",
                   "name": "望江",
                   "weatherCode": "101220607"
                },
                {
                   "id": "220608",
                   "name": "岳西",
                   "weatherCode": "101220608"
                },
                {
                   "id": "220609",
                   "name": "桐城",
                   "weatherCode": "101220609"
                }
             ]
          },
          {
             "id": "2207",
             "name": "宿州",
             "county": [
                {
                   "id": "220701",
                   "name": "宿州",
                   "weatherCode": "101220701"
                },
                {
                   "id": "220702",
                   "name": "砀山",
                   "weatherCode": "101220702"
                },
                {
                   "id": "220703",
                   "name": "灵璧",
                   "weatherCode": "101220703"
                },
                {
                   "id": "220704",
                   "name": "泗县",
                   "weatherCode": "101220704"
                },
                {
                   "id": "220705",
                   "name": "萧县",
                   "weatherCode": "101220705"
                }
             ]
          },
          {
             "id": "2208",
             "name": "阜阳",
             "county": [
                {
                   "id": "220801",
                   "name": "阜阳",
                   "weatherCode": "101220801"
                },
                {
                   "id": "220802",
                   "name": "阜南",
                   "weatherCode": "101220802"
                },
                {
                   "id": "220803",
                   "name": "颍上",
                   "weatherCode": "101220803"
                },
                {
                   "id": "220804",
                   "name": "临泉",
                   "weatherCode": "101220804"
                },
                {
                   "id": "220805",
                   "name": "界首",
                   "weatherCode": "101220805"
                },
                {
                   "id": "220806",
                   "name": "太和",
                   "weatherCode": "101220806"
                }
             ]
          },
          {
             "id": "2209",
             "name": "亳州",
             "county": [
                {
                   "id": "220901",
                   "name": "亳州",
                   "weatherCode": "101220901"
                },
                {
                   "id": "220902",
                   "name": "涡阳",
                   "weatherCode": "101220902"
                },
                {
                   "id": "220903",
                   "name": "利辛",
                   "weatherCode": "101220903"
                },
                {
                   "id": "220904",
                   "name": "蒙城",
                   "weatherCode": "101220904"
                }
             ]
          },
          {
             "id": "2210",
             "name": "黄山",
             "county": [
                {
                   "id": "221001",
                   "name": "黄山市",
                   "weatherCode": "101221001"
                },
                {
                   "id": "221002",
                   "name": "黄山区",
                   "weatherCode": "101221002"
                },
                {
                   "id": "221003",
                   "name": "屯溪",
                   "weatherCode": "101221003"
                },
                {
                   "id": "221004",
                   "name": "祁门",
                   "weatherCode": "101221004"
                },
                {
                   "id": "221005",
                   "name": "黟县",
                   "weatherCode": "101221005"
                },
                {
                   "id": "221006",
                   "name": "歙县",
                   "weatherCode": "101221006"
                },
                {
                   "id": "221007",
                   "name": "休宁",
                   "weatherCode": "101221007"
                },
                {
                   "id": "221008",
                   "name": "黄山风景区",
                   "weatherCode": "101221008"
                }
             ]
          },
          {
             "id": "2211",
             "name": "滁州",
             "county": [
                {
                   "id": "221101",
                   "name": "滁州",
                   "weatherCode": "101221101"
                },
                {
                   "id": "221102",
                   "name": "凤阳",
                   "weatherCode": "101221102"
                },
                {
                   "id": "221103",
                   "name": "明光",
                   "weatherCode": "101221103"
                },
                {
                   "id": "221104",
                   "name": "定远",
                   "weatherCode": "101221104"
                },
                {
                   "id": "221105",
                   "name": "全椒",
                   "weatherCode": "101221105"
                },
                {
                   "id": "221106",
                   "name": "来安",
                   "weatherCode": "101221106"
                },
                {
                   "id": "221107",
                   "name": "天长",
                   "weatherCode": "101221107"
                }
             ]
          },
          {
             "id": "2212",
             "name": "淮北",
             "county": [
                {
                   "id": "221201",
                   "name": "淮北",
                   "weatherCode": "101221201"
                },
                {
                   "id": "221202",
                   "name": "濉溪",
                   "weatherCode": "101221202"
                }
             ]
          },
          {
             "id": "2213",
             "name": "铜陵",
             "county": [{
                "id": "221301",
                "name": "铜陵",
                "weatherCode": "101221301"
             }]
          },
          {
             "id": "2214",
             "name": "宣城",
             "county": [
                {
                   "id": "221401",
                   "name": "宣城",
                   "weatherCode": "101221401"
                },
                {
                   "id": "221402",
                   "name": "泾县",
                   "weatherCode": "101221402"
                },
                {
                   "id": "221403",
                   "name": "旌德",
                   "weatherCode": "101221403"
                },
                {
                   "id": "221404",
                   "name": "宁国",
                   "weatherCode": "101221404"
                },
                {
                   "id": "221405",
                   "name": "绩溪",
                   "weatherCode": "101221405"
                },
                {
                   "id": "221406",
                   "name": "广德",
                   "weatherCode": "101221406"
                },
                {
                   "id": "221407",
                   "name": "郎溪",
                   "weatherCode": "101221407"
                }
             ]
          },
          {
             "id": "2215",
             "name": "六安",
             "county": [
                {
                   "id": "221501",
                   "name": "六安",
                   "weatherCode": "101221501"
                },
                {
                   "id": "221502",
                   "name": "霍邱",
                   "weatherCode": "101221502"
                },
                {
                   "id": "221503",
                   "name": "寿县",
                   "weatherCode": "101221503"
                },
                {
                   "id": "221504",
                   "name": "金寨",
                   "weatherCode": "101221505"
                },
                {
                   "id": "221505",
                   "name": "霍山",
                   "weatherCode": "101221506"
                },
                {
                   "id": "221506",
                   "name": "舒城",
                   "weatherCode": "101221507"
                }
             ]
          },
          {
             "id": "2216",
             "name": "巢湖",
             "county": [
                {
                   "id": "221601",
                   "name": "巢湖",
                   "weatherCode": "101221601"
                },
                {
                   "id": "221602",
                   "name": "庐江",
                   "weatherCode": "101221602"
                },
                {
                   "id": "221603",
                   "name": "无为",
                   "weatherCode": "101221603"
                },
                {
                   "id": "221604",
                   "name": "含山",
                   "weatherCode": "101221604"
                },
                {
                   "id": "221605",
                   "name": "和县",
                   "weatherCode": "101221605"
                }
             ]
          },
          {
             "id": "2217",
             "name": "池州",
             "county": [
                {
                   "id": "221701",
                   "name": "池州",
                   "weatherCode": "101221701"
                },
                {
                   "id": "221702",
                   "name": "东至",
                   "weatherCode": "101221702"
                },
                {
                   "id": "221703",
                   "name": "青阳",
                   "weatherCode": "101221703"
                },
                {
                   "id": "221704",
                   "name": "九华山",
                   "weatherCode": "101221704"
                },
                {
                   "id": "221705",
                   "name": "石台",
                   "weatherCode": "101221705"
                }
             ]
          }
       ]
    },
    {
       "id": "23",
       "name": "福建",
       "city": [
          {
             "id": "2301",
             "name": "福州",
             "county": [
                {
                   "id": "230101",
                   "name": "福州",
                   "weatherCode": "101230101"
                },
                {
                   "id": "230102",
                   "name": "闽清",
                   "weatherCode": "101230102"
                },
                {
                   "id": "230103",
                   "name": "闽侯",
                   "weatherCode": "101230103"
                },
                {
                   "id": "230104",
                   "name": "罗源",
                   "weatherCode": "101230104"
                },
                {
                   "id": "230105",
                   "name": "连江",
                   "weatherCode": "101230105"
                },
                {
                   "id": "230106",
                   "name": "永泰",
                   "weatherCode": "101230107"
                },
                {
                   "id": "230107",
                   "name": "平潭",
                   "weatherCode": "101230108"
                },
                {
                   "id": "230108",
                   "name": "长乐",
                   "weatherCode": "101230110"
                },
                {
                   "id": "230109",
                   "name": "福清",
                   "weatherCode": "101230111"
                }
             ]
          },
          {
             "id": "2302",
             "name": "厦门",
             "county": [
                {
                   "id": "230201",
                   "name": "厦门",
                   "weatherCode": "101230201"
                },
                {
                   "id": "230202",
                   "name": "同安",
                   "weatherCode": "101230202"
                }
             ]
          },
          {
             "id": "2303",
             "name": "宁德",
             "county": [
                {
                   "id": "230301",
                   "name": "宁德",
                   "weatherCode": "101230301"
                },
                {
                   "id": "230302",
                   "name": "古田",
                   "weatherCode": "101230302"
                },
                {
                   "id": "230303",
                   "name": "霞浦",
                   "weatherCode": "101230303"
                },
                {
                   "id": "230304",
                   "name": "寿宁",
                   "weatherCode": "101230304"
                },
                {
                   "id": "230305",
                   "name": "周宁",
                   "weatherCode": "101230305"
                },
                {
                   "id": "230306",
                   "name": "福安",
                   "weatherCode": "101230306"
                },
                {
                   "id": "230307",
                   "name": "柘荣",
                   "weatherCode": "101230307"
                },
                {
                   "id": "230308",
                   "name": "福鼎",
                   "weatherCode": "101230308"
                },
                {
                   "id": "230309",
                   "name": "屏南",
                   "weatherCode": "101230309"
                }
             ]
          },
          {
             "id": "2304",
             "name": "莆田",
             "county": [
                {
                   "id": "230401",
                   "name": "莆田",
                   "weatherCode": "101230401"
                },
                {
                   "id": "230402",
                   "name": "仙游",
                   "weatherCode": "101230402"
                },
                {
                   "id": "230403",
                   "name": "秀屿港",
                   "weatherCode": "101230403"
                },
                {
                   "id": "230404",
                   "name": "涵江",
                   "weatherCode": "101230404"
                },
                {
                   "id": "230405",
                   "name": "秀屿",
                   "weatherCode": "101230405"
                },
                {
                   "id": "230406",
                   "name": "荔城",
                   "weatherCode": "101230406"
                },
                {
                   "id": "230407",
                   "name": "城厢",
                   "weatherCode": "101230407"
                }
             ]
          },
          {
             "id": "2305",
             "name": "泉州",
             "county": [
                {
                   "id": "230501",
                   "name": "泉州",
                   "weatherCode": "101230501"
                },
                {
                   "id": "230502",
                   "name": "安溪",
                   "weatherCode": "101230502"
                },
                {
                   "id": "230503",
                   "name": "永春",
                   "weatherCode": "101230504"
                },
                {
                   "id": "230504",
                   "name": "德化",
                   "weatherCode": "101230505"
                },
                {
                   "id": "230505",
                   "name": "南安",
                   "weatherCode": "101230506"
                },
                {
                   "id": "230506",
                   "name": "崇武",
                   "weatherCode": "101230507"
                },
                {
                   "id": "230507",
                   "name": "惠安",
                   "weatherCode": "101230508"
                },
                {
                   "id": "230508",
                   "name": "晋江",
                   "weatherCode": "101230509"
                },
                {
                   "id": "230509",
                   "name": "石狮",
                   "weatherCode": "101230510"
                }
             ]
          },
          {
             "id": "2306",
             "name": "漳州",
             "county": [
                {
                   "id": "230601",
                   "name": "漳州",
                   "weatherCode": "101230601"
                },
                {
                   "id": "230602",
                   "name": "长泰",
                   "weatherCode": "101230602"
                },
                {
                   "id": "230603",
                   "name": "南靖",
                   "weatherCode": "101230603"
                },
                {
                   "id": "230604",
                   "name": "平和",
                   "weatherCode": "101230604"
                },
                {
                   "id": "230605",
                   "name": "龙海",
                   "weatherCode": "101230605"
                },
                {
                   "id": "230606",
                   "name": "漳浦",
                   "weatherCode": "101230606"
                },
                {
                   "id": "230607",
                   "name": "诏安",
                   "weatherCode": "101230607"
                },
                {
                   "id": "230608",
                   "name": "东山",
                   "weatherCode": "101230608"
                },
                {
                   "id": "230609",
                   "name": "云霄",
                   "weatherCode": "101230609"
                },
                {
                   "id": "230610",
                   "name": "华安",
                   "weatherCode": "101230610"
                }
             ]
          },
          {
             "id": "2307",
             "name": "龙岩",
             "county": [
                {
                   "id": "230701",
                   "name": "龙岩",
                   "weatherCode": "101230701"
                },
                {
                   "id": "230702",
                   "name": "长汀",
                   "weatherCode": "101230702"
                },
                {
                   "id": "230703",
                   "name": "连城",
                   "weatherCode": "101230703"
                },
                {
                   "id": "230704",
                   "name": "武平",
                   "weatherCode": "101230704"
                },
                {
                   "id": "230705",
                   "name": "上杭",
                   "weatherCode": "101230705"
                },
                {
                   "id": "230706",
                   "name": "永定",
                   "weatherCode": "101230706"
                },
                {
                   "id": "230707",
                   "name": "漳平",
                   "weatherCode": "101230707"
                }
             ]
          },
          {
             "id": "2308",
             "name": "三明",
             "county": [
                {
                   "id": "230801",
                   "name": "三明",
                   "weatherCode": "101230801"
                },
                {
                   "id": "230802",
                   "name": "宁化",
                   "weatherCode": "101230802"
                },
                {
                   "id": "230803",
                   "name": "清流",
                   "weatherCode": "101230803"
                },
                {
                   "id": "230804",
                   "name": "泰宁",
                   "weatherCode": "101230804"
                },
                {
                   "id": "230805",
                   "name": "将乐",
                   "weatherCode": "101230805"
                },
                {
                   "id": "230806",
                   "name": "建宁",
                   "weatherCode": "101230806"
                },
                {
                   "id": "230807",
                   "name": "明溪",
                   "weatherCode": "101230807"
                },
                {
                   "id": "230808",
                   "name": "沙县",
                   "weatherCode": "101230808"
                },
                {
                   "id": "230809",
                   "name": "尤溪",
                   "weatherCode": "101230809"
                },
                {
                   "id": "230810",
                   "name": "永安",
                   "weatherCode": "101230810"
                },
                {
                   "id": "230811",
                   "name": "大田",
                   "weatherCode": "101230811"
                }
             ]
          },
          {
             "id": "2309",
             "name": "南平",
             "county": [
                {
                   "id": "230901",
                   "name": "南平",
                   "weatherCode": "101230901"
                },
                {
                   "id": "230902",
                   "name": "顺昌",
                   "weatherCode": "101230902"
                },
                {
                   "id": "230903",
                   "name": "光泽",
                   "weatherCode": "101230903"
                },
                {
                   "id": "230904",
                   "name": "邵武",
                   "weatherCode": "101230904"
                },
                {
                   "id": "230905",
                   "name": "武夷山",
                   "weatherCode": "101230905"
                },
                {
                   "id": "230906",
                   "name": "浦城",
                   "weatherCode": "101230906"
                },
                {
                   "id": "230907",
                   "name": "建阳",
                   "weatherCode": "101230907"
                },
                {
                   "id": "230908",
                   "name": "松溪",
                   "weatherCode": "101230908"
                },
                {
                   "id": "230909",
                   "name": "政和",
                   "weatherCode": "101230909"
                },
                {
                   "id": "230910",
                   "name": "建瓯",
                   "weatherCode": "101230910"
                }
             ]
          }
       ]
    },
    {
       "id": "24",
       "name": "江西",
       "city": [
          {
             "id": "2401",
             "name": "南昌",
             "county": [
                {
                   "id": "240101",
                   "name": "南昌",
                   "weatherCode": "101240101"
                },
                {
                   "id": "240102",
                   "name": "新建",
                   "weatherCode": "101240102"
                },
                {
                   "id": "240103",
                   "name": "南昌县",
                   "weatherCode": "101240103"
                },
                {
                   "id": "240104",
                   "name": "安义",
                   "weatherCode": "101240104"
                },
                {
                   "id": "240105",
                   "name": "进贤",
                   "weatherCode": "101240105"
                }
             ]
          },
          {
             "id": "2402",
             "name": "九江",
             "county": [
                {
                   "id": "240201",
                   "name": "九江",
                   "weatherCode": "101240201"
                },
                {
                   "id": "240202",
                   "name": "瑞昌",
                   "weatherCode": "101240202"
                },
                {
                   "id": "240203",
                   "name": "庐山",
                   "weatherCode": "101240203"
                },
                {
                   "id": "240204",
                   "name": "武宁",
                   "weatherCode": "101240204"
                },
                {
                   "id": "240205",
                   "name": "德安",
                   "weatherCode": "101240205"
                },
                {
                   "id": "240206",
                   "name": "永修",
                   "weatherCode": "101240206"
                },
                {
                   "id": "240207",
                   "name": "湖口",
                   "weatherCode": "101240207"
                },
                {
                   "id": "240208",
                   "name": "彭泽",
                   "weatherCode": "101240208"
                },
                {
                   "id": "240209",
                   "name": "星子",
                   "weatherCode": "101240209"
                },
                {
                   "id": "240210",
                   "name": "都昌",
                   "weatherCode": "101240210"
                },
                {
                   "id": "240211",
                   "name": "修水",
                   "weatherCode": "101240212"
                }
             ]
          },
          {
             "id": "2403",
             "name": "上饶",
             "county": [
                {
                   "id": "240301",
                   "name": "上饶",
                   "weatherCode": "101240301"
                },
                {
                   "id": "240302",
                   "name": "鄱阳",
                   "weatherCode": "101240302"
                },
                {
                   "id": "240303",
                   "name": "婺源",
                   "weatherCode": "101240303"
                },
                {
                   "id": "240304",
                   "name": "余干",
                   "weatherCode": "101240305"
                },
                {
                   "id": "240305",
                   "name": "万年",
                   "weatherCode": "101240306"
                },
                {
                   "id": "240306",
                   "name": "德兴",
                   "weatherCode": "101240307"
                },
                {
                   "id": "240307",
                   "name": "上饶县",
                   "weatherCode": "101240308"
                },
                {
                   "id": "240308",
                   "name": "弋阳",
                   "weatherCode": "101240309"
                },
                {
                   "id": "240309",
                   "name": "横峰",
                   "weatherCode": "101240310"
                },
                {
                   "id": "240310",
                   "name": "铅山",
                   "weatherCode": "101240311"
                },
                {
                   "id": "240311",
                   "name": "玉山",
                   "weatherCode": "101240312"
                },
                {
                   "id": "240312",
                   "name": "广丰",
                   "weatherCode": "101240313"
                }
             ]
          },
          {
             "id": "2404",
             "name": "抚州",
             "county": [
                {
                   "id": "240401",
                   "name": "抚州",
                   "weatherCode": "101240401"
                },
                {
                   "id": "240402",
                   "name": "广昌",
                   "weatherCode": "101240402"
                },
                {
                   "id": "240403",
                   "name": "乐安",
                   "weatherCode": "101240403"
                },
                {
                   "id": "240404",
                   "name": "崇仁",
                   "weatherCode": "101240404"
                },
                {
                   "id": "240405",
                   "name": "金溪",
                   "weatherCode": "101240405"
                },
                {
                   "id": "240406",
                   "name": "资溪",
                   "weatherCode": "101240406"
                },
                {
                   "id": "240407",
                   "name": "宜黄",
                   "weatherCode": "101240407"
                },
                {
                   "id": "240408",
                   "name": "南城",
                   "weatherCode": "101240408"
                },
                {
                   "id": "240409",
                   "name": "南丰",
                   "weatherCode": "101240409"
                },
                {
                   "id": "240410",
                   "name": "黎川",
                   "weatherCode": "101240410"
                },
                {
                   "id": "240411",
                   "name": "东乡",
                   "weatherCode": "101240411"
                }
             ]
          },
          {
             "id": "2405",
             "name": "宜春",
             "county": [
                {
                   "id": "240501",
                   "name": "宜春",
                   "weatherCode": "101240501"
                },
                {
                   "id": "240502",
                   "name": "铜鼓",
                   "weatherCode": "101240502"
                },
                {
                   "id": "240503",
                   "name": "宜丰",
                   "weatherCode": "101240503"
                },
                {
                   "id": "240504",
                   "name": "万载",
                   "weatherCode": "101240504"
                },
                {
                   "id": "240505",
                   "name": "上高",
                   "weatherCode": "101240505"
                },
                {
                   "id": "240506",
                   "name": "靖安",
                   "weatherCode": "101240506"
                },
                {
                   "id": "240507",
                   "name": "奉新",
                   "weatherCode": "101240507"
                },
                {
                   "id": "240508",
                   "name": "高安",
                   "weatherCode": "101240508"
                },
                {
                   "id": "240509",
                   "name": "樟树",
                   "weatherCode": "101240509"
                },
                {
                   "id": "240510",
                   "name": "丰城",
                   "weatherCode": "101240510"
                }
             ]
          },
          {
             "id": "2406",
             "name": "吉安",
             "county": [
                {
                   "id": "240601",
                   "name": "吉安",
                   "weatherCode": "101240601"
                },
                {
                   "id": "240602",
                   "name": "吉安县",
                   "weatherCode": "101240602"
                },
                {
                   "id": "240603",
                   "name": "吉水",
                   "weatherCode": "101240603"
                },
                {
                   "id": "240604",
                   "name": "新干",
                   "weatherCode": "101240604"
                },
                {
                   "id": "240605",
                   "name": "峡江",
                   "weatherCode": "101240605"
                },
                {
                   "id": "240606",
                   "name": "永丰",
                   "weatherCode": "101240606"
                },
                {
                   "id": "240607",
                   "name": "永新",
                   "weatherCode": "101240607"
                },
                {
                   "id": "240608",
                   "name": "井冈山",
                   "weatherCode": "101240608"
                },
                {
                   "id": "240609",
                   "name": "万安",
                   "weatherCode": "101240609"
                },
                {
                   "id": "240610",
                   "name": "遂川",
                   "weatherCode": "101240610"
                },
                {
                   "id": "240611",
                   "name": "泰和",
                   "weatherCode": "101240611"
                },
                {
                   "id": "240612",
                   "name": "安福",
                   "weatherCode": "101240612"
                },
                {
                   "id": "240613",
                   "name": "宁冈",
                   "weatherCode": "101240613"
                }
             ]
          },
          {
             "id": "2407",
             "name": "赣州",
             "county": [
                {
                   "id": "240701",
                   "name": "赣州",
                   "weatherCode": "101240701"
                },
                {
                   "id": "240702",
                   "name": "崇义",
                   "weatherCode": "101240702"
                },
                {
                   "id": "240703",
                   "name": "上犹",
                   "weatherCode": "101240703"
                },
                {
                   "id": "240704",
                   "name": "南康",
                   "weatherCode": "101240704"
                },
                {
                   "id": "240705",
                   "name": "大余",
                   "weatherCode": "101240705"
                },
                {
                   "id": "240706",
                   "name": "信丰",
                   "weatherCode": "101240706"
                },
                {
                   "id": "240707",
                   "name": "宁都",
                   "weatherCode": "101240707"
                },
                {
                   "id": "240708",
                   "name": "石城",
                   "weatherCode": "101240708"
                },
                {
                   "id": "240709",
                   "name": "瑞金",
                   "weatherCode": "101240709"
                },
                {
                   "id": "240710",
                   "name": "于都",
                   "weatherCode": "101240710"
                },
                {
                   "id": "240711",
                   "name": "会昌",
                   "weatherCode": "101240711"
                },
                {
                   "id": "240712",
                   "name": "安远",
                   "weatherCode": "101240712"
                },
                {
                   "id": "240713",
                   "name": "全南",
                   "weatherCode": "101240713"
                },
                {
                   "id": "240714",
                   "name": "龙南",
                   "weatherCode": "101240714"
                },
                {
                   "id": "240715",
                   "name": "定南",
                   "weatherCode": "101240715"
                },
                {
                   "id": "240716",
                   "name": "寻乌",
                   "weatherCode": "101240716"
                },
                {
                   "id": "240717",
                   "name": "兴国",
                   "weatherCode": "101240717"
                },
                {
                   "id": "240718",
                   "name": "赣县",
                   "weatherCode": "101240718"
                }
             ]
          },
          {
             "id": "2408",
             "name": "景德镇",
             "county": [
                {
                   "id": "240801",
                   "name": "景德镇",
                   "weatherCode": "101240801"
                },
                {
                   "id": "240802",
                   "name": "乐平",
                   "weatherCode": "101240802"
                },
                {
                   "id": "240803",
                   "name": "浮梁",
                   "weatherCode": "101240803"
                }
             ]
          },
          {
             "id": "2409",
             "name": "萍乡",
             "county": [
                {
                   "id": "240901",
                   "name": "萍乡",
                   "weatherCode": "101240901"
                },
                {
                   "id": "240902",
                   "name": "莲花",
                   "weatherCode": "101240902"
                },
                {
                   "id": "240903",
                   "name": "上栗",
                   "weatherCode": "101240903"
                },
                {
                   "id": "240904",
                   "name": "安源",
                   "weatherCode": "101240904"
                },
                {
                   "id": "240905",
                   "name": "芦溪",
                   "weatherCode": "101240905"
                },
                {
                   "id": "240906",
                   "name": "湘东",
                   "weatherCode": "101240906"
                }
             ]
          },
          {
             "id": "2410",
             "name": "新余",
             "county": [
                {
                   "id": "241001",
                   "name": "新余",
                   "weatherCode": "101241001"
                },
                {
                   "id": "241002",
                   "name": "分宜",
                   "weatherCode": "101241002"
                }
             ]
          },
          {
             "id": "2411",
             "name": "鹰潭",
             "county": [
                {
                   "id": "241101",
                   "name": "鹰潭",
                   "weatherCode": "101241101"
                },
                {
                   "id": "241102",
                   "name": "余江",
                   "weatherCode": "101241102"
                },
                {
                   "id": "241103",
                   "name": "贵溪",
                   "weatherCode": "101241103"
                }
             ]
          }
       ]
    },
    {
       "id": "25",
       "name": "湖南",
       "city": [
          {
             "id": "2501",
             "name": "长沙",
             "county": [
                {
                   "id": "250101",
                   "name": "长沙",
                   "weatherCode": "101250101"
                },
                {
                   "id": "250102",
                   "name": "宁乡",
                   "weatherCode": "101250102"
                },
                {
                   "id": "250103",
                   "name": "浏阳",
                   "weatherCode": "101250103"
                },
                {
                   "id": "250104",
                   "name": "马坡岭",
                   "weatherCode": "101250104"
                },
                {
                   "id": "250105",
                   "name": "望城",
                   "weatherCode": "101250105"
                }
             ]
          },
          {
             "id": "2502",
             "name": "湘潭",
             "county": [
                {
                   "id": "250201",
                   "name": "湘潭",
                   "weatherCode": "101250201"
                },
                {
                   "id": "250202",
                   "name": "韶山",
                   "weatherCode": "101250202"
                },
                {
                   "id": "250203",
                   "name": "湘乡",
                   "weatherCode": "101250203"
                }
             ]
          },
          {
             "id": "2503",
             "name": "株洲",
             "county": [
                {
                   "id": "250301",
                   "name": "株洲",
                   "weatherCode": "101250301"
                },
                {
                   "id": "250302",
                   "name": "攸县",
                   "weatherCode": "101250302"
                },
                {
                   "id": "250303",
                   "name": "醴陵",
                   "weatherCode": "101250303"
                },
                {
                   "id": "250304",
                   "name": "茶陵",
                   "weatherCode": "101250305"
                },
                {
                   "id": "250305",
                   "name": "炎陵",
                   "weatherCode": "101250306"
                }
             ]
          },
          {
             "id": "2504",
             "name": "衡阳",
             "county": [
                {
                   "id": "250401",
                   "name": "衡阳",
                   "weatherCode": "101250401"
                },
                {
                   "id": "250402",
                   "name": "衡山",
                   "weatherCode": "101250402"
                },
                {
                   "id": "250403",
                   "name": "衡东",
                   "weatherCode": "101250403"
                },
                {
                   "id": "250404",
                   "name": "祁东",
                   "weatherCode": "101250404"
                },
                {
                   "id": "250405",
                   "name": "衡阳县",
                   "weatherCode": "101250405"
                },
                {
                   "id": "250406",
                   "name": "常宁",
                   "weatherCode": "101250406"
                },
                {
                   "id": "250407",
                   "name": "衡南",
                   "weatherCode": "101250407"
                },
                {
                   "id": "250408",
                   "name": "耒阳",
                   "weatherCode": "101250408"
                },
                {
                   "id": "250409",
                   "name": "南岳",
                   "weatherCode": "101250409"
                }
             ]
          },
          {
             "id": "2505",
             "name": "郴州",
             "county": [
                {
                   "id": "250501",
                   "name": "郴州",
                   "weatherCode": "101250501"
                },
                {
                   "id": "250502",
                   "name": "桂阳",
                   "weatherCode": "101250502"
                },
                {
                   "id": "250503",
                   "name": "嘉禾",
                   "weatherCode": "101250503"
                },
                {
                   "id": "250504",
                   "name": "宜章",
                   "weatherCode": "101250504"
                },
                {
                   "id": "250505",
                   "name": "临武",
                   "weatherCode": "101250505"
                },
                {
                   "id": "250506",
                   "name": "资兴",
                   "weatherCode": "101250507"
                },
                {
                   "id": "250507",
                   "name": "汝城",
                   "weatherCode": "101250508"
                },
                {
                   "id": "250508",
                   "name": "安仁",
                   "weatherCode": "101250509"
                },
                {
                   "id": "250509",
                   "name": "永兴",
                   "weatherCode": "101250510"
                },
                {
                   "id": "250510",
                   "name": "桂东",
                   "weatherCode": "101250511"
                },
                {
                   "id": "250511",
                   "name": "苏仙",
                   "weatherCode": "101250512"
                }
             ]
          },
          {
             "id": "2506",
             "name": "常德",
             "county": [
                {
                   "id": "250601",
                   "name": "常德",
                   "weatherCode": "101250601"
                },
                {
                   "id": "250602",
                   "name": "安乡",
                   "weatherCode": "101250602"
                },
                {
                   "id": "250603",
                   "name": "桃源",
                   "weatherCode": "101250603"
                },
                {
                   "id": "250604",
                   "name": "汉寿",
                   "weatherCode": "101250604"
                },
                {
                   "id": "250605",
                   "name": "澧县",
                   "weatherCode": "101250605"
                },
                {
                   "id": "250606",
                   "name": "临澧",
                   "weatherCode": "101250606"
                },
                {
                   "id": "250607",
                   "name": "石门",
                   "weatherCode": "101250607"
                },
                {
                   "id": "250608",
                   "name": "津市",
                   "weatherCode": "101250608"
                }
             ]
          },
          {
             "id": "2507",
             "name": "益阳",
             "county": [
                {
                   "id": "250701",
                   "name": "益阳",
                   "weatherCode": "101250700"
                },
                {
                   "id": "250702",
                   "name": "赫山区",
                   "weatherCode": "101250701"
                },
                {
                   "id": "250703",
                   "name": "南县",
                   "weatherCode": "101250702"
                },
                {
                   "id": "250704",
                   "name": "桃江",
                   "weatherCode": "101250703"
                },
                {
                   "id": "250705",
                   "name": "安化",
                   "weatherCode": "101250704"
                },
                {
                   "id": "250706",
                   "name": "沅江",
                   "weatherCode": "101250705"
                }
             ]
          },
          {
             "id": "2508",
             "name": "娄底",
             "county": [
                {
                   "id": "250801",
                   "name": "娄底",
                   "weatherCode": "101250801"
                },
                {
                   "id": "250802",
                   "name": "双峰",
                   "weatherCode": "101250802"
                },
                {
                   "id": "250803",
                   "name": "冷水江",
                   "weatherCode": "101250803"
                },
                {
                   "id": "250804",
                   "name": "新化",
                   "weatherCode": "101250805"
                },
                {
                   "id": "250805",
                   "name": "涟源",
                   "weatherCode": "101250806"
                }
             ]
          },
          {
             "id": "2509",
             "name": "邵阳",
             "county": [
                {
                   "id": "250901",
                   "name": "邵阳",
                   "weatherCode": "101250901"
                },
                {
                   "id": "250902",
                   "name": "隆回",
                   "weatherCode": "101250902"
                },
                {
                   "id": "250903",
                   "name": "洞口",
                   "weatherCode": "101250903"
                },
                {
                   "id": "250904",
                   "name": "新邵",
                   "weatherCode": "101250904"
                },
                {
                   "id": "250905",
                   "name": "邵东",
                   "weatherCode": "101250905"
                },
                {
                   "id": "250906",
                   "name": "绥宁",
                   "weatherCode": "101250906"
                },
                {
                   "id": "250907",
                   "name": "新宁",
                   "weatherCode": "101250907"
                },
                {
                   "id": "250908",
                   "name": "武冈",
                   "weatherCode": "101250908"
                },
                {
                   "id": "250909",
                   "name": "城步",
                   "weatherCode": "101250909"
                },
                {
                   "id": "250910",
                   "name": "邵阳县",
                   "weatherCode": "101250910"
                }
             ]
          },
          {
             "id": "2510",
             "name": "岳阳",
             "county": [
                {
                   "id": "251001",
                   "name": "岳阳",
                   "weatherCode": "101251001"
                },
                {
                   "id": "251002",
                   "name": "华容",
                   "weatherCode": "101251002"
                },
                {
                   "id": "251003",
                   "name": "湘阴",
                   "weatherCode": "101251003"
                },
                {
                   "id": "251004",
                   "name": "汨罗",
                   "weatherCode": "101251004"
                },
                {
                   "id": "251005",
                   "name": "平江",
                   "weatherCode": "101251005"
                },
                {
                   "id": "251006",
                   "name": "临湘",
                   "weatherCode": "101251006"
                }
             ]
          },
          {
             "id": "2511",
             "name": "张家界",
             "county": [
                {
                   "id": "251101",
                   "name": "张家界",
                   "weatherCode": "101251101"
                },
                {
                   "id": "251102",
                   "name": "桑植",
                   "weatherCode": "101251102"
                },
                {
                   "id": "251103",
                   "name": "慈利",
                   "weatherCode": "101251103"
                },
                {
                   "id": "251104",
                   "name": "武陵源",
                   "weatherCode": "101251104"
                }
             ]
          },
          {
             "id": "2512",
             "name": "怀化",
             "county": [
                {
                   "id": "251201",
                   "name": "怀化",
                   "weatherCode": "101251201"
                },
                {
                   "id": "251202",
                   "name": "沅陵",
                   "weatherCode": "101251203"
                },
                {
                   "id": "251203",
                   "name": "辰溪",
                   "weatherCode": "101251204"
                },
                {
                   "id": "251204",
                   "name": "靖州",
                   "weatherCode": "101251205"
                },
                {
                   "id": "251205",
                   "name": "会同",
                   "weatherCode": "101251206"
                },
                {
                   "id": "251206",
                   "name": "通道",
                   "weatherCode": "101251207"
                },
                {
                   "id": "251207",
                   "name": "麻阳",
                   "weatherCode": "101251208"
                },
                {
                   "id": "251208",
                   "name": "新晃",
                   "weatherCode": "101251209"
                },
                {
                   "id": "251209",
                   "name": "芷江",
                   "weatherCode": "101251210"
                },
                {
                   "id": "251210",
                   "name": "溆浦",
                   "weatherCode": "101251211"
                },
                {
                   "id": "251211",
                   "name": "中方",
                   "weatherCode": "101251212"
                },
                {
                   "id": "251212",
                   "name": "洪江",
                   "weatherCode": "101251213"
                }
             ]
          },
          {
             "id": "2513",
             "name": "永州",
             "county": [
                {
                   "id": "251301",
                   "name": "永州",
                   "weatherCode": "101251401"
                },
                {
                   "id": "251302",
                   "name": "祁阳",
                   "weatherCode": "101251402"
                },
                {
                   "id": "251303",
                   "name": "东安",
                   "weatherCode": "101251403"
                },
                {
                   "id": "251304",
                   "name": "双牌",
                   "weatherCode": "101251404"
                },
                {
                   "id": "251305",
                   "name": "道县",
                   "weatherCode": "101251405"
                },
                {
                   "id": "251306",
                   "name": "宁远",
                   "weatherCode": "101251406"
                },
                {
                   "id": "251307",
                   "name": "江永",
                   "weatherCode": "101251407"
                },
                {
                   "id": "251308",
                   "name": "蓝山",
                   "weatherCode": "101251408"
                },
                {
                   "id": "251309",
                   "name": "新田",
                   "weatherCode": "101251409"
                },
                {
                   "id": "251310",
                   "name": "江华",
                   "weatherCode": "101251410"
                },
                {
                   "id": "251311",
                   "name": "冷水滩",
                   "weatherCode": "101251411"
                }
             ]
          },
          {
             "id": "2514",
             "name": "湘西",
             "county": [
                {
                   "id": "251401",
                   "name": "吉首",
                   "weatherCode": "101251501"
                },
                {
                   "id": "251402",
                   "name": "保靖",
                   "weatherCode": "101251502"
                },
                {
                   "id": "251403",
                   "name": "永顺",
                   "weatherCode": "101251503"
                },
                {
                   "id": "251404",
                   "name": "古丈",
                   "weatherCode": "101251504"
                },
                {
                   "id": "251405",
                   "name": "凤凰",
                   "weatherCode": "101251505"
                },
                {
                   "id": "251406",
                   "name": "泸溪",
                   "weatherCode": "101251506"
                },
                {
                   "id": "251407",
                   "name": "龙山",
                   "weatherCode": "101251507"
                },
                {
                   "id": "251408",
                   "name": "花垣",
                   "weatherCode": "101251508"
                }
             ]
          }
       ]
    },
    {
       "id": "26",
       "name": "贵州",
       "city": [
          {
             "id": "2601",
             "name": "贵阳",
             "county": [
                {
                   "id": "260101",
                   "name": "贵阳",
                   "weatherCode": "101260101"
                },
                {
                   "id": "260102",
                   "name": "白云",
                   "weatherCode": "101260102"
                },
                {
                   "id": "260103",
                   "name": "花溪",
                   "weatherCode": "101260103"
                },
                {
                   "id": "260104",
                   "name": "乌当",
                   "weatherCode": "101260104"
                },
                {
                   "id": "260105",
                   "name": "息烽",
                   "weatherCode": "101260105"
                },
                {
                   "id": "260106",
                   "name": "开阳",
                   "weatherCode": "101260106"
                },
                {
                   "id": "260107",
                   "name": "修文",
                   "weatherCode": "101260107"
                },
                {
                   "id": "260108",
                   "name": "清镇",
                   "weatherCode": "101260108"
                },
                {
                   "id": "260109",
                   "name": "小河",
                   "weatherCode": "101260109"
                },
                {
                   "id": "260110",
                   "name": "云岩",
                   "weatherCode": "101260110"
                },
                {
                   "id": "260111",
                   "name": "南明",
                   "weatherCode": "101260111"
                }
             ]
          },
          {
             "id": "2602",
             "name": "遵义",
             "county": [
                {
                   "id": "260201",
                   "name": "遵义",
                   "weatherCode": "101260201"
                },
                {
                   "id": "260202",
                   "name": "遵义县",
                   "weatherCode": "101260202"
                },
                {
                   "id": "260203",
                   "name": "仁怀",
                   "weatherCode": "101260203"
                },
                {
                   "id": "260204",
                   "name": "绥阳",
                   "weatherCode": "101260204"
                },
                {
                   "id": "260205",
                   "name": "湄潭",
                   "weatherCode": "101260205"
                },
                {
                   "id": "260206",
                   "name": "凤冈",
                   "weatherCode": "101260206"
                },
                {
                   "id": "260207",
                   "name": "桐梓",
                   "weatherCode": "101260207"
                },
                {
                   "id": "260208",
                   "name": "赤水",
                   "weatherCode": "101260208"
                },
                {
                   "id": "260209",
                   "name": "习水",
                   "weatherCode": "101260209"
                },
                {
                   "id": "260210",
                   "name": "道真",
                   "weatherCode": "101260210"
                },
                {
                   "id": "260211",
                   "name": "正安",
                   "weatherCode": "101260211"
                },
                {
                   "id": "260212",
                   "name": "务川",
                   "weatherCode": "101260212"
                },
                {
                   "id": "260213",
                   "name": "余庆",
                   "weatherCode": "101260213"
                },
                {
                   "id": "260214",
                   "name": "汇川",
                   "weatherCode": "101260214"
                },
                {
                   "id": "260215",
                   "name": "红花岗",
                   "weatherCode": "101260215"
                }
             ]
          },
          {
             "id": "2603",
             "name": "安顺",
             "county": [
                {
                   "id": "260301",
                   "name": "安顺",
                   "weatherCode": "101260301"
                },
                {
                   "id": "260302",
                   "name": "普定",
                   "weatherCode": "101260302"
                },
                {
                   "id": "260303",
                   "name": "镇宁",
                   "weatherCode": "101260303"
                },
                {
                   "id": "260304",
                   "name": "平坝",
                   "weatherCode": "101260304"
                },
                {
                   "id": "260305",
                   "name": "紫云",
                   "weatherCode": "101260305"
                },
                {
                   "id": "260306",
                   "name": "关岭",
                   "weatherCode": "101260306"
                }
             ]
          },
          {
             "id": "2604",
             "name": "黔南",
             "county": [
                {
                   "id": "260401",
                   "name": "都匀",
                   "weatherCode": "101260401"
                },
                {
                   "id": "260402",
                   "name": "贵定",
                   "weatherCode": "101260402"
                },
                {
                   "id": "260403",
                   "name": "瓮安",
                   "weatherCode": "101260403"
                },
                {
                   "id": "260404",
                   "name": "长顺",
                   "weatherCode": "101260404"
                },
                {
                   "id": "260405",
                   "name": "福泉",
                   "weatherCode": "101260405"
                },
                {
                   "id": "260406",
                   "name": "惠水",
                   "weatherCode": "101260406"
                },
                {
                   "id": "260407",
                   "name": "龙里",
                   "weatherCode": "101260407"
                },
                {
                   "id": "260408",
                   "name": "罗甸",
                   "weatherCode": "101260408"
                },
                {
                   "id": "260409",
                   "name": "平塘",
                   "weatherCode": "101260409"
                },
                {
                   "id": "260410",
                   "name": "独山",
                   "weatherCode": "101260410"
                },
                {
                   "id": "260411",
                   "name": "三都",
                   "weatherCode": "101260411"
                },
                {
                   "id": "260412",
                   "name": "荔波",
                   "weatherCode": "101260412"
                }
             ]
          },
          {
             "id": "2605",
             "name": "黔东南",
             "county": [
                {
                   "id": "260501",
                   "name": "凯里",
                   "weatherCode": "101260501"
                },
                {
                   "id": "260502",
                   "name": "岑巩",
                   "weatherCode": "101260502"
                },
                {
                   "id": "260503",
                   "name": "施秉",
                   "weatherCode": "101260503"
                },
                {
                   "id": "260504",
                   "name": "镇远",
                   "weatherCode": "101260504"
                },
                {
                   "id": "260505",
                   "name": "黄平",
                   "weatherCode": "101260505"
                },
                {
                   "id": "260506",
                   "name": "麻江",
                   "weatherCode": "101260507"
                },
                {
                   "id": "260507",
                   "name": "丹寨",
                   "weatherCode": "101260508"
                },
                {
                   "id": "260508",
                   "name": "三穗",
                   "weatherCode": "101260509"
                },
                {
                   "id": "260509",
                   "name": "台江",
                   "weatherCode": "101260510"
                },
                {
                   "id": "260510",
                   "name": "剑河",
                   "weatherCode": "101260511"
                },
                {
                   "id": "260511",
                   "name": "雷山",
                   "weatherCode": "101260512"
                },
                {
                   "id": "260512",
                   "name": "黎平",
                   "weatherCode": "101260513"
                },
                {
                   "id": "260513",
                   "name": "天柱",
                   "weatherCode": "101260514"
                },
                {
                   "id": "260514",
                   "name": "锦屏",
                   "weatherCode": "101260515"
                },
                {
                   "id": "260515",
                   "name": "榕江",
                   "weatherCode": "101260516"
                },
                {
                   "id": "260516",
                   "name": "从江",
                   "weatherCode": "101260517"
                }
             ]
          },
          {
             "id": "2606",
             "name": "铜仁",
             "county": [
                {
                   "id": "260601",
                   "name": "铜仁",
                   "weatherCode": "101260601"
                },
                {
                   "id": "260602",
                   "name": "江口",
                   "weatherCode": "101260602"
                },
                {
                   "id": "260603",
                   "name": "玉屏",
                   "weatherCode": "101260603"
                },
                {
                   "id": "260604",
                   "name": "万山",
                   "weatherCode": "101260604"
                },
                {
                   "id": "260605",
                   "name": "思南",
                   "weatherCode": "101260605"
                },
                {
                   "id": "260606",
                   "name": "印江",
                   "weatherCode": "101260607"
                },
                {
                   "id": "260607",
                   "name": "石阡",
                   "weatherCode": "101260608"
                },
                {
                   "id": "260608",
                   "name": "沿河",
                   "weatherCode": "101260609"
                },
                {
                   "id": "260609",
                   "name": "德江",
                   "weatherCode": "101260610"
                },
                {
                   "id": "260610",
                   "name": "松桃",
                   "weatherCode": "101260611"
                }
             ]
          },
          {
             "id": "2607",
             "name": "毕节",
             "county": [
                {
                   "id": "260701",
                   "name": "毕节",
                   "weatherCode": "101260701"
                },
                {
                   "id": "260702",
                   "name": "赫章",
                   "weatherCode": "101260702"
                },
                {
                   "id": "260703",
                   "name": "金沙",
                   "weatherCode": "101260703"
                },
                {
                   "id": "260704",
                   "name": "威宁",
                   "weatherCode": "101260704"
                },
                {
                   "id": "260705",
                   "name": "大方",
                   "weatherCode": "101260705"
                },
                {
                   "id": "260706",
                   "name": "纳雍",
                   "weatherCode": "101260706"
                },
                {
                   "id": "260707",
                   "name": "织金",
                   "weatherCode": "101260707"
                },
                {
                   "id": "260708",
                   "name": "黔西",
                   "weatherCode": "101260708"
                }
             ]
          },
          {
             "id": "2608",
             "name": "六盘水",
             "county": [
                {
                   "id": "260801",
                   "name": "水城",
                   "weatherCode": "101260801"
                },
                {
                   "id": "260802",
                   "name": "六枝",
                   "weatherCode": "101260802"
                },
                {
                   "id": "260803",
                   "name": "盘县",
                   "weatherCode": "101260804"
                }
             ]
          },
          {
             "id": "2609",
             "name": "黔西南",
             "county": [
                {
                   "id": "260901",
                   "name": "兴义",
                   "weatherCode": "101260901"
                },
                {
                   "id": "260902",
                   "name": "晴隆",
                   "weatherCode": "101260902"
                },
                {
                   "id": "260903",
                   "name": "兴仁",
                   "weatherCode": "101260903"
                },
                {
                   "id": "260904",
                   "name": "贞丰",
                   "weatherCode": "101260904"
                },
                {
                   "id": "260905",
                   "name": "望谟",
                   "weatherCode": "101260905"
                },
                {
                   "id": "260906",
                   "name": "安龙",
                   "weatherCode": "101260907"
                },
                {
                   "id": "260907",
                   "name": "册亨",
                   "weatherCode": "101260908"
                },
                {
                   "id": "260908",
                   "name": "普安",
                   "weatherCode": "101260909"
                }
             ]
          }
       ]
    },
    {
       "id": "27",
       "name": "四川",
       "city": [
          {
             "id": "2701",
             "name": "成都",
             "county": [
                {
                   "id": "270101",
                   "name": "成都",
                   "weatherCode": "101270101"
                },
                {
                   "id": "270102",
                   "name": "龙泉驿",
                   "weatherCode": "101270102"
                },
                {
                   "id": "270103",
                   "name": "新都",
                   "weatherCode": "101270103"
                },
                {
                   "id": "270104",
                   "name": "温江",
                   "weatherCode": "101270104"
                },
                {
                   "id": "270105",
                   "name": "金堂",
                   "weatherCode": "101270105"
                },
                {
                   "id": "270106",
                   "name": "双流",
                   "weatherCode": "101270106"
                },
                {
                   "id": "270107",
                   "name": "郫县",
                   "weatherCode": "101270107"
                },
                {
                   "id": "270108",
                   "name": "大邑",
                   "weatherCode": "101270108"
                },
                {
                   "id": "270109",
                   "name": "蒲江",
                   "weatherCode": "101270109"
                },
                {
                   "id": "270110",
                   "name": "新津",
                   "weatherCode": "101270110"
                },
                {
                   "id": "270111",
                   "name": "都江堰",
                   "weatherCode": "101270111"
                },
                {
                   "id": "270112",
                   "name": "彭州",
                   "weatherCode": "101270112"
                },
                {
                   "id": "270113",
                   "name": "邛崃",
                   "weatherCode": "101270113"
                },
                {
                   "id": "270114",
                   "name": "崇州",
                   "weatherCode": "101270114"
                }
             ]
          },
          {
             "id": "2702",
             "name": "攀枝花",
             "county": [
                {
                   "id": "270201",
                   "name": "攀枝花",
                   "weatherCode": "101270201"
                },
                {
                   "id": "270202",
                   "name": "仁和",
                   "weatherCode": "101270202"
                },
                {
                   "id": "270203",
                   "name": "米易",
                   "weatherCode": "101270203"
                },
                {
                   "id": "270204",
                   "name": "盐边",
                   "weatherCode": "101270204"
                }
             ]
          },
          {
             "id": "2703",
             "name": "自贡",
             "county": [
                {
                   "id": "270301",
                   "name": "自贡",
                   "weatherCode": "101270301"
                },
                {
                   "id": "270302",
                   "name": "富顺",
                   "weatherCode": "101270302"
                },
                {
                   "id": "270303",
                   "name": "荣县",
                   "weatherCode": "101270303"
                }
             ]
          },
          {
             "id": "2704",
             "name": "绵阳",
             "county": [
                {
                   "id": "270401",
                   "name": "绵阳",
                   "weatherCode": "101270401"
                },
                {
                   "id": "270402",
                   "name": "三台",
                   "weatherCode": "101270402"
                },
                {
                   "id": "270403",
                   "name": "盐亭",
                   "weatherCode": "101270403"
                },
                {
                   "id": "270404",
                   "name": "安县",
                   "weatherCode": "101270404"
                },
                {
                   "id": "270405",
                   "name": "梓潼",
                   "weatherCode": "101270405"
                },
                {
                   "id": "270406",
                   "name": "北川",
                   "weatherCode": "101270406"
                },
                {
                   "id": "270407",
                   "name": "平武",
                   "weatherCode": "101270407"
                },
                {
                   "id": "270408",
                   "name": "江油",
                   "weatherCode": "101270408"
                }
             ]
          },
          {
             "id": "2705",
             "name": "南充",
             "county": [
                {
                   "id": "270501",
                   "name": "南充",
                   "weatherCode": "101270501"
                },
                {
                   "id": "270502",
                   "name": "南部",
                   "weatherCode": "101270502"
                },
                {
                   "id": "270503",
                   "name": "营山",
                   "weatherCode": "101270503"
                },
                {
                   "id": "270504",
                   "name": "蓬安",
                   "weatherCode": "101270504"
                },
                {
                   "id": "270505",
                   "name": "仪陇",
                   "weatherCode": "101270505"
                },
                {
                   "id": "270506",
                   "name": "西充",
                   "weatherCode": "101270506"
                },
                {
                   "id": "270507",
                   "name": "阆中",
                   "weatherCode": "101270507"
                }
             ]
          },
          {
             "id": "2706",
             "name": "达州",
             "county": [
                {
                   "id": "270601",
                   "name": "达州",
                   "weatherCode": "101270601"
                },
                {
                   "id": "270602",
                   "name": "宣汉",
                   "weatherCode": "101270602"
                },
                {
                   "id": "270603",
                   "name": "开江",
                   "weatherCode": "101270603"
                },
                {
                   "id": "270604",
                   "name": "大竹",
                   "weatherCode": "101270604"
                },
                {
                   "id": "270605",
                   "name": "渠县",
                   "weatherCode": "101270605"
                },
                {
                   "id": "270606",
                   "name": "万源",
                   "weatherCode": "101270606"
                },
                {
                   "id": "270607",
                   "name": "通川",
                   "weatherCode": "101270607"
                },
                {
                   "id": "270608",
                   "name": "达县",
                   "weatherCode": "101270608"
                }
             ]
          },
          {
             "id": "2707",
             "name": "遂宁",
             "county": [
                {
                   "id": "270701",
                   "name": "遂宁",
                   "weatherCode": "101270701"
                },
                {
                   "id": "270702",
                   "name": "蓬溪",
                   "weatherCode": "101270702"
                },
                {
                   "id": "270703",
                   "name": "射洪",
                   "weatherCode": "101270703"
                }
             ]
          },
          {
             "id": "2708",
             "name": "广安",
             "county": [
                {
                   "id": "270801",
                   "name": "广安",
                   "weatherCode": "101270801"
                },
                {
                   "id": "270802",
                   "name": "岳池",
                   "weatherCode": "101270802"
                },
                {
                   "id": "270803",
                   "name": "武胜",
                   "weatherCode": "101270803"
                },
                {
                   "id": "270804",
                   "name": "邻水",
                   "weatherCode": "101270804"
                },
                {
                   "id": "270805",
                   "name": "华蓥",
                   "weatherCode": "101270805"
                }
             ]
          },
          {
             "id": "2709",
             "name": "巴中",
             "county": [
                {
                   "id": "270901",
                   "name": "巴中",
                   "weatherCode": "101270901"
                },
                {
                   "id": "270902",
                   "name": "通江",
                   "weatherCode": "101270902"
                },
                {
                   "id": "270903",
                   "name": "南江",
                   "weatherCode": "101270903"
                },
                {
                   "id": "270904",
                   "name": "平昌",
                   "weatherCode": "101270904"
                }
             ]
          },
          {
             "id": "2710",
             "name": "泸州",
             "county": [
                {
                   "id": "271001",
                   "name": "泸州",
                   "weatherCode": "101271001"
                },
                {
                   "id": "271002",
                   "name": "泸县",
                   "weatherCode": "101271003"
                },
                {
                   "id": "271003",
                   "name": "合江",
                   "weatherCode": "101271004"
                },
                {
                   "id": "271004",
                   "name": "叙永",
                   "weatherCode": "101271005"
                },
                {
                   "id": "271005",
                   "name": "古蔺",
                   "weatherCode": "101271006"
                },
                {
                   "id": "271006",
                   "name": "纳溪",
                   "weatherCode": "101271007"
                }
             ]
          },
          {
             "id": "2711",
             "name": "宜宾",
             "county": [
                {
                   "id": "271101",
                   "name": "宜宾",
                   "weatherCode": "101271101"
                },
                {
                   "id": "271102",
                   "name": "宜宾县",
                   "weatherCode": "101271103"
                },
                {
                   "id": "271103",
                   "name": "南溪",
                   "weatherCode": "101271104"
                },
                {
                   "id": "271104",
                   "name": "江安",
                   "weatherCode": "101271105"
                },
                {
                   "id": "271105",
                   "name": "长宁",
                   "weatherCode": "101271106"
                },
                {
                   "id": "271106",
                   "name": "高县",
                   "weatherCode": "101271107"
                },
                {
                   "id": "271107",
                   "name": "珙县",
                   "weatherCode": "101271108"
                },
                {
                   "id": "271108",
                   "name": "筠连",
                   "weatherCode": "101271109"
                },
                {
                   "id": "271109",
                   "name": "兴文",
                   "weatherCode": "101271110"
                },
                {
                   "id": "271110",
                   "name": "屏山",
                   "weatherCode": "101271111"
                }
             ]
          },
          {
             "id": "2712",
             "name": "内江",
             "county": [
                {
                   "id": "271201",
                   "name": "内江",
                   "weatherCode": "101271201"
                },
                {
                   "id": "271202",
                   "name": "东兴",
                   "weatherCode": "101271202"
                },
                {
                   "id": "271203",
                   "name": "威远",
                   "weatherCode": "101271203"
                },
                {
                   "id": "271204",
                   "name": "资中",
                   "weatherCode": "101271204"
                },
                {
                   "id": "271205",
                   "name": "隆昌",
                   "weatherCode": "101271205"
                }
             ]
          },
          {
             "id": "2713",
             "name": "资阳",
             "county": [
                {
                   "id": "271301",
                   "name": "资阳",
                   "weatherCode": "101271301"
                },
                {
                   "id": "271302",
                   "name": "安岳",
                   "weatherCode": "101271302"
                },
                {
                   "id": "271303",
                   "name": "乐至",
                   "weatherCode": "101271303"
                },
                {
                   "id": "271304",
                   "name": "简阳",
                   "weatherCode": "101271304"
                }
             ]
          },
          {
             "id": "2714",
             "name": "乐山",
             "county": [
                {
                   "id": "271401",
                   "name": "乐山",
                   "weatherCode": "101271401"
                },
                {
                   "id": "271402",
                   "name": "犍为",
                   "weatherCode": "101271402"
                },
                {
                   "id": "271403",
                   "name": "井研",
                   "weatherCode": "101271403"
                },
                {
                   "id": "271404",
                   "name": "夹江",
                   "weatherCode": "101271404"
                },
                {
                   "id": "271405",
                   "name": "沐川",
                   "weatherCode": "101271405"
                },
                {
                   "id": "271406",
                   "name": "峨边",
                   "weatherCode": "101271406"
                },
                {
                   "id": "271407",
                   "name": "马边",
                   "weatherCode": "101271407"
                },
                {
                   "id": "271408",
                   "name": "峨眉",
                   "weatherCode": "101271408"
                },
                {
                   "id": "271409",
                   "name": "峨眉山",
                   "weatherCode": "101271409"
                }
             ]
          },
          {
             "id": "2715",
             "name": "眉山",
             "county": [
                {
                   "id": "271501",
                   "name": "眉山",
                   "weatherCode": "101271501"
                },
                {
                   "id": "271502",
                   "name": "仁寿",
                   "weatherCode": "101271502"
                },
                {
                   "id": "271503",
                   "name": "彭山",
                   "weatherCode": "101271503"
                },
                {
                   "id": "271504",
                   "name": "洪雅",
                   "weatherCode": "101271504"
                },
                {
                   "id": "271505",
                   "name": "丹棱",
                   "weatherCode": "101271505"
                },
                {
                   "id": "271506",
                   "name": "青神",
                   "weatherCode": "101271506"
                }
             ]
          },
          {
             "id": "2716",
             "name": "凉山",
             "county": [
                {
                   "id": "271601",
                   "name": "凉山",
                   "weatherCode": "101271601"
                },
                {
                   "id": "271602",
                   "name": "木里",
                   "weatherCode": "101271603"
                },
                {
                   "id": "271603",
                   "name": "盐源",
                   "weatherCode": "101271604"
                },
                {
                   "id": "271604",
                   "name": "德昌",
                   "weatherCode": "101271605"
                },
                {
                   "id": "271605",
                   "name": "会理",
                   "weatherCode": "101271606"
                },
                {
                   "id": "271606",
                   "name": "会东",
                   "weatherCode": "101271607"
                },
                {
                   "id": "271607",
                   "name": "宁南",
                   "weatherCode": "101271608"
                },
                {
                   "id": "271608",
                   "name": "普格",
                   "weatherCode": "101271609"
                },
                {
                   "id": "271609",
                   "name": "西昌",
                   "weatherCode": "101271610"
                },
                {
                   "id": "271610",
                   "name": "金阳",
                   "weatherCode": "101271611"
                },
                {
                   "id": "271611",
                   "name": "昭觉",
                   "weatherCode": "101271612"
                },
                {
                   "id": "271612",
                   "name": "喜德",
                   "weatherCode": "101271613"
                },
                {
                   "id": "271613",
                   "name": "冕宁",
                   "weatherCode": "101271614"
                },
                {
                   "id": "271614",
                   "name": "越西",
                   "weatherCode": "101271615"
                },
                {
                   "id": "271615",
                   "name": "甘洛",
                   "weatherCode": "101271616"
                },
                {
                   "id": "271616",
                   "name": "雷波",
                   "weatherCode": "101271617"
                },
                {
                   "id": "271617",
                   "name": "美姑",
                   "weatherCode": "101271618"
                },
                {
                   "id": "271618",
                   "name": "布拖",
                   "weatherCode": "101271619"
                }
             ]
          },
          {
             "id": "2717",
             "name": "雅安",
             "county": [
                {
                   "id": "271701",
                   "name": "雅安",
                   "weatherCode": "101271701"
                },
                {
                   "id": "271702",
                   "name": "名山",
                   "weatherCode": "101271702"
                },
                {
                   "id": "271703",
                   "name": "荥经",
                   "weatherCode": "101271703"
                },
                {
                   "id": "271704",
                   "name": "汉源",
                   "weatherCode": "101271704"
                },
                {
                   "id": "271705",
                   "name": "石棉",
                   "weatherCode": "101271705"
                },
                {
                   "id": "271706",
                   "name": "天全",
                   "weatherCode": "101271706"
                },
                {
                   "id": "271707",
                   "name": "芦山",
                   "weatherCode": "101271707"
                },
                {
                   "id": "271708",
                   "name": "宝兴",
                   "weatherCode": "101271708"
                }
             ]
          },
          {
             "id": "2718",
             "name": "甘孜",
             "county": [
                {
                   "id": "271801",
                   "name": "甘孜",
                   "weatherCode": "101271801"
                },
                {
                   "id": "271802",
                   "name": "康定",
                   "weatherCode": "101271802"
                },
                {
                   "id": "271803",
                   "name": "泸定",
                   "weatherCode": "101271803"
                },
                {
                   "id": "271804",
                   "name": "丹巴",
                   "weatherCode": "101271804"
                },
                {
                   "id": "271806",
                   "name": "雅江",
                   "weatherCode": "101271806"
                },
                {
                   "id": "271807",
                   "name": "道孚",
                   "weatherCode": "101271807"
                },
                {
                   "id": "271808",
                   "name": "炉霍",
                   "weatherCode": "101271808"
                },
                {
                   "id": "271809",
                   "name": "新龙",
                   "weatherCode": "101271809"
                },
                {
                   "id": "271810",
                   "name": "德格",
                   "weatherCode": "101271810"
                },
                {
                   "id": "271811",
                   "name": "白玉",
                   "weatherCode": "101271811"
                },
                {
                   "id": "271812",
                   "name": "石渠",
                   "weatherCode": "101271812"
                },
                {
                   "id": "271813",
                   "name": "色达",
                   "weatherCode": "101271813"
                },
                {
                   "id": "271814",
                   "name": "理塘",
                   "weatherCode": "101271814"
                },
                {
                   "id": "271815",
                   "name": "巴塘",
                   "weatherCode": "101271815"
                },
                {
                   "id": "271816",
                   "name": "乡城",
                   "weatherCode": "101271816"
                },
                {
                   "id": "271817",
                   "name": "稻城",
                   "weatherCode": "101271817"
                },
                {
                   "id": "271818",
                   "name": "得荣",
                   "weatherCode": "101271818"
                }
             ]
          },
          {
             "id": "2719",
             "name": "阿坝",
             "county": [
                {
                   "id": "271901",
                   "name": "阿坝",
                   "weatherCode": "101271901"
                },
                {
                   "id": "271902",
                   "name": "汶川",
                   "weatherCode": "101271902"
                },
                {
                   "id": "271903",
                   "name": "理县",
                   "weatherCode": "101271903"
                },
                {
                   "id": "271904",
                   "name": "茂县",
                   "weatherCode": "101271904"
                },
                {
                   "id": "271905",
                   "name": "松潘",
                   "weatherCode": "101271905"
                },
                {
                   "id": "271906",
                   "name": "九寨沟",
                   "weatherCode": "101271906"
                },
                {
                   "id": "271907",
                   "name": "金川",
                   "weatherCode": "101271907"
                },
                {
                   "id": "271908",
                   "name": "小金",
                   "weatherCode": "101271908"
                },
                {
                   "id": "271909",
                   "name": "黑水",
                   "weatherCode": "101271909"
                },
                {
                   "id": "271910",
                   "name": "马尔康",
                   "weatherCode": "101271910"
                },
                {
                   "id": "271911",
                   "name": "壤塘",
                   "weatherCode": "101271911"
                },
                {
                   "id": "271912",
                   "name": "若尔盖",
                   "weatherCode": "101271912"
                },
                {
                   "id": "271913",
                   "name": "红原",
                   "weatherCode": "101271913"
                }
             ]
          },
          {
             "id": "2720",
             "name": "德阳",
             "county": [
                {
                   "id": "272001",
                   "name": "德阳",
                   "weatherCode": "101272001"
                },
                {
                   "id": "272002",
                   "name": "中江",
                   "weatherCode": "101272002"
                },
                {
                   "id": "272003",
                   "name": "广汉",
                   "weatherCode": "101272003"
                },
                {
                   "id": "272004",
                   "name": "什邡",
                   "weatherCode": "101272004"
                },
                {
                   "id": "272005",
                   "name": "绵竹",
                   "weatherCode": "101272005"
                },
                {
                   "id": "272006",
                   "name": "罗江",
                   "weatherCode": "101272006"
                }
             ]
          },
          {
             "id": "2721",
             "name": "广元",
             "county": [
                {
                   "id": "272101",
                   "name": "广元",
                   "weatherCode": "101272101"
                },
                {
                   "id": "272102",
                   "name": "旺苍",
                   "weatherCode": "101272102"
                },
                {
                   "id": "272103",
                   "name": "青川",
                   "weatherCode": "101272103"
                },
                {
                   "id": "272104",
                   "name": "剑阁",
                   "weatherCode": "101272104"
                },
                {
                   "id": "272105",
                   "name": "苍溪",
                   "weatherCode": "101272105"
                }
             ]
          }
       ]
    },
    {
       "id": "28",
       "name": "广东",
       "city": [
          {
             "id": "2801",
             "name": "广州",
             "county": [
                {
                   "id": "280101",
                   "name": "广州",
                   "weatherCode": "101280101"
                },
                {
                   "id": "280102",
                   "name": "番禺",
                   "weatherCode": "101280102"
                },
                {
                   "id": "280103",
                   "name": "从化",
                   "weatherCode": "101280103"
                },
                {
                   "id": "280104",
                   "name": "增城",
                   "weatherCode": "101280104"
                },
                {
                   "id": "280105",
                   "name": "花都",
                   "weatherCode": "101280105"
                }
             ]
          },
          {
             "id": "2802",
             "name": "韶关",
             "county": [
                {
                   "id": "280201",
                   "name": "韶关",
                   "weatherCode": "101280201"
                },
                {
                   "id": "280202",
                   "name": "乳源",
                   "weatherCode": "101280202"
                },
                {
                   "id": "280203",
                   "name": "始兴",
                   "weatherCode": "101280203"
                },
                {
                   "id": "280204",
                   "name": "翁源",
                   "weatherCode": "101280204"
                },
                {
                   "id": "280205",
                   "name": "乐昌",
                   "weatherCode": "101280205"
                },
                {
                   "id": "280206",
                   "name": "仁化",
                   "weatherCode": "101280206"
                },
                {
                   "id": "280207",
                   "name": "南雄",
                   "weatherCode": "101280207"
                },
                {
                   "id": "280208",
                   "name": "新丰",
                   "weatherCode": "101280208"
                },
                {
                   "id": "280209",
                   "name": "曲江",
                   "weatherCode": "101280209"
                },
                {
                   "id": "280210",
                   "name": "浈江",
                   "weatherCode": "101280210"
                },
                {
                   "id": "280211",
                   "name": "武江",
                   "weatherCode": "101280211"
                }
             ]
          },
          {
             "id": "2803",
             "name": "惠州",
             "county": [
                {
                   "id": "280301",
                   "name": "惠州",
                   "weatherCode": "101280301"
                },
                {
                   "id": "280302",
                   "name": "博罗",
                   "weatherCode": "101280302"
                },
                {
                   "id": "280303",
                   "name": "惠阳",
                   "weatherCode": "101280303"
                },
                {
                   "id": "280304",
                   "name": "惠东",
                   "weatherCode": "101280304"
                },
                {
                   "id": "280305",
                   "name": "龙门",
                   "weatherCode": "101280305"
                }
             ]
          },
          {
             "id": "2804",
             "name": "梅州",
             "county": [
                {
                   "id": "280401",
                   "name": "梅州",
                   "weatherCode": "101280401"
                },
                {
                   "id": "280402",
                   "name": "兴宁",
                   "weatherCode": "101280402"
                },
                {
                   "id": "280403",
                   "name": "蕉岭",
                   "weatherCode": "101280403"
                },
                {
                   "id": "280404",
                   "name": "大埔",
                   "weatherCode": "101280404"
                },
                {
                   "id": "280405",
                   "name": "丰顺",
                   "weatherCode": "101280406"
                },
                {
                   "id": "280406",
                   "name": "平远",
                   "weatherCode": "101280407"
                },
                {
                   "id": "280407",
                   "name": "五华",
                   "weatherCode": "101280408"
                },
                {
                   "id": "280408",
                   "name": "梅县",
                   "weatherCode": "101280409"
                }
             ]
          },
          {
             "id": "2805",
             "name": "汕头",
             "county": [
                {
                   "id": "280501",
                   "name": "汕头",
                   "weatherCode": "101280501"
                },
                {
                   "id": "280502",
                   "name": "潮阳",
                   "weatherCode": "101280502"
                },
                {
                   "id": "280503",
                   "name": "澄海",
                   "weatherCode": "101280503"
                },
                {
                   "id": "280504",
                   "name": "南澳",
                   "weatherCode": "101280504"
                }
             ]
          },
          {
             "id": "2806",
             "name": "深圳",
             "county": [{
                "id": "280601",
                "name": "深圳",
                "weatherCode": "101280601"
             }]
          },
          {
             "id": "2807",
             "name": "珠海",
             "county": [
                {
                   "id": "280701",
                   "name": "珠海",
                   "weatherCode": "101280701"
                },
                {
                   "id": "280702",
                   "name": "斗门",
                   "weatherCode": "101280702"
                },
                {
                   "id": "280703",
                   "name": "金湾",
                   "weatherCode": "101280703"
                }
             ]
          },
          {
             "id": "2808",
             "name": "佛山",
             "county": [
                {
                   "id": "280801",
                   "name": "佛山",
                   "weatherCode": "101280800"
                },
                {
                   "id": "280802",
                   "name": "顺德",
                   "weatherCode": "101280801"
                },
                {
                   "id": "280803",
                   "name": "三水",
                   "weatherCode": "101280802"
                },
                {
                   "id": "280804",
                   "name": "南海",
                   "weatherCode": "101280803"
                },
                {
                   "id": "280805",
                   "name": "高明",
                   "weatherCode": "101280804"
                }
             ]
          },
          {
             "id": "2809",
             "name": "肇庆",
             "county": [
                {
                   "id": "280901",
                   "name": "肇庆",
                   "weatherCode": "101280901"
                },
                {
                   "id": "280902",
                   "name": "广宁",
                   "weatherCode": "101280902"
                },
                {
                   "id": "280903",
                   "name": "四会",
                   "weatherCode": "101280903"
                },
                {
                   "id": "280904",
                   "name": "德庆",
                   "weatherCode": "101280905"
                },
                {
                   "id": "280905",
                   "name": "怀集",
                   "weatherCode": "101280906"
                },
                {
                   "id": "280906",
                   "name": "封开",
                   "weatherCode": "101280907"
                },
                {
                   "id": "280907",
                   "name": "高要",
                   "weatherCode": "101280908"
                }
             ]
          },
          {
             "id": "2810",
             "name": "湛江",
             "county": [
                {
                   "id": "281001",
                   "name": "湛江",
                   "weatherCode": "101281001"
                },
                {
                   "id": "281002",
                   "name": "吴川",
                   "weatherCode": "101281002"
                },
                {
                   "id": "281003",
                   "name": "雷州",
                   "weatherCode": "101281003"
                },
                {
                   "id": "281004",
                   "name": "徐闻",
                   "weatherCode": "101281004"
                },
                {
                   "id": "281005",
                   "name": "廉江",
                   "weatherCode": "101281005"
                },
                {
                   "id": "281006",
                   "name": "赤坎",
                   "weatherCode": "101281006"
                },
                {
                   "id": "281007",
                   "name": "遂溪",
                   "weatherCode": "101281007"
                },
                {
                   "id": "281008",
                   "name": "坡头",
                   "weatherCode": "101281008"
                },
                {
                   "id": "281009",
                   "name": "霞山",
                   "weatherCode": "101281009"
                },
                {
                   "id": "281010",
                   "name": "麻章",
                   "weatherCode": "101281010"
                }
             ]
          },
          {
             "id": "2811",
             "name": "江门",
             "county": [
                {
                   "id": "281101",
                   "name": "江门",
                   "weatherCode": "101281101"
                },
                {
                   "id": "281102",
                   "name": "开平",
                   "weatherCode": "101281103"
                },
                {
                   "id": "281103",
                   "name": "新会",
                   "weatherCode": "101281104"
                },
                {
                   "id": "281104",
                   "name": "恩平",
                   "weatherCode": "101281105"
                },
                {
                   "id": "281105",
                   "name": "台山",
                   "weatherCode": "101281106"
                },
                {
                   "id": "281106",
                   "name": "蓬江",
                   "weatherCode": "101281107"
                },
                {
                   "id": "281107",
                   "name": "鹤山",
                   "weatherCode": "101281108"
                },
                {
                   "id": "281108",
                   "name": "江海",
                   "weatherCode": "101281109"
                }
             ]
          },
          {
             "id": "2812",
             "name": "河源",
             "county": [
                {
                   "id": "281201",
                   "name": "河源",
                   "weatherCode": "101281201"
                },
                {
                   "id": "281202",
                   "name": "紫金",
                   "weatherCode": "101281202"
                },
                {
                   "id": "281203",
                   "name": "连平",
                   "weatherCode": "101281203"
                },
                {
                   "id": "281204",
                   "name": "和平",
                   "weatherCode": "101281204"
                },
                {
                   "id": "281205",
                   "name": "龙川",
                   "weatherCode": "101281205"
                },
                {
                   "id": "281206",
                   "name": "东源",
                   "weatherCode": "101281206"
                }
             ]
          },
          {
             "id": "2813",
             "name": "清远",
             "county": [
                {
                   "id": "281301",
                   "name": "清远",
                   "weatherCode": "101281301"
                },
                {
                   "id": "281302",
                   "name": "连南",
                   "weatherCode": "101281302"
                },
                {
                   "id": "281303",
                   "name": "连州",
                   "weatherCode": "101281303"
                },
                {
                   "id": "281304",
                   "name": "连山",
                   "weatherCode": "101281304"
                },
                {
                   "id": "281305",
                   "name": "阳山",
                   "weatherCode": "101281305"
                },
                {
                   "id": "281306",
                   "name": "佛冈",
                   "weatherCode": "101281306"
                },
                {
                   "id": "281307",
                   "name": "英德",
                   "weatherCode": "101281307"
                },
                {
                   "id": "281308",
                   "name": "清新",
                   "weatherCode": "101281308"
                }
             ]
          },
          {
             "id": "2814",
             "name": "云浮",
             "county": [
                {
                   "id": "281401",
                   "name": "云浮",
                   "weatherCode": "101281401"
                },
                {
                   "id": "281402",
                   "name": "罗定",
                   "weatherCode": "101281402"
                },
                {
                   "id": "281403",
                   "name": "新兴",
                   "weatherCode": "101281403"
                },
                {
                   "id": "281404",
                   "name": "郁南",
                   "weatherCode": "101281404"
                },
                {
                   "id": "281405",
                   "name": "云安",
                   "weatherCode": "101281406"
                }
             ]
          },
          {
             "id": "2815",
             "name": "潮州",
             "county": [
                {
                   "id": "281501",
                   "name": "潮州",
                   "weatherCode": "101281501"
                },
                {
                   "id": "281502",
                   "name": "饶平",
                   "weatherCode": "101281502"
                },
                {
                   "id": "281503",
                   "name": "潮安",
                   "weatherCode": "101281503"
                }
             ]
          },
          {
             "id": "2816",
             "name": "东莞",
             "county": [{
                "id": "281601",
                "name": "东莞",
                "weatherCode": "101281601"
             }]
          },
          {
             "id": "2817",
             "name": "中山",
             "county": [{
                "id": "281701",
                "name": "中山",
                "weatherCode": "101281701"
             }]
          },
          {
             "id": "2818",
             "name": "阳江",
             "county": [
                {
                   "id": "281801",
                   "name": "阳江",
                   "weatherCode": "101281801"
                },
                {
                   "id": "281802",
                   "name": "阳春",
                   "weatherCode": "101281802"
                },
                {
                   "id": "281803",
                   "name": "阳东",
                   "weatherCode": "101281803"
                },
                {
                   "id": "281804",
                   "name": "阳西",
                   "weatherCode": "101281804"
                }
             ]
          },
          {
             "id": "2819",
             "name": "揭阳",
             "county": [
                {
                   "id": "281901",
                   "name": "揭阳",
                   "weatherCode": "101281901"
                },
                {
                   "id": "281902",
                   "name": "揭西",
                   "weatherCode": "101281902"
                },
                {
                   "id": "281903",
                   "name": "普宁",
                   "weatherCode": "101281903"
                },
                {
                   "id": "281904",
                   "name": "惠来",
                   "weatherCode": "101281904"
                },
                {
                   "id": "281905",
                   "name": "揭东",
                   "weatherCode": "101281905"
                }
             ]
          },
          {
             "id": "2820",
             "name": "茂名",
             "county": [
                {
                   "id": "282001",
                   "name": "茂名",
                   "weatherCode": "101282001"
                },
                {
                   "id": "282002",
                   "name": "高州",
                   "weatherCode": "101282002"
                },
                {
                   "id": "282003",
                   "name": "化州",
                   "weatherCode": "101282003"
                },
                {
                   "id": "282004",
                   "name": "电白",
                   "weatherCode": "101282004"
                },
                {
                   "id": "282005",
                   "name": "信宜",
                   "weatherCode": "101282005"
                },
                {
                   "id": "282006",
                   "name": "茂港",
                   "weatherCode": "101282006"
                }
             ]
          },
          {
             "id": "2821",
             "name": "汕尾",
             "county": [
                {
                   "id": "282101",
                   "name": "汕尾",
                   "weatherCode": "101282101"
                },
                {
                   "id": "282102",
                   "name": "海丰",
                   "weatherCode": "101282102"
                },
                {
                   "id": "282103",
                   "name": "陆丰",
                   "weatherCode": "101282103"
                },
                {
                   "id": "282104",
                   "name": "陆河",
                   "weatherCode": "101282104"
                }
             ]
          }
       ]
    },
    {
       "id": "29",
       "name": "云南",
       "city": [
          {
             "id": "2901",
             "name": "昆明",
             "county": [
                {
                   "id": "290101",
                   "name": "昆明",
                   "weatherCode": "101290101"
                },
                {
                   "id": "290102",
                   "name": "东川",
                   "weatherCode": "101290103"
                },
                {
                   "id": "290103",
                   "name": "寻甸",
                   "weatherCode": "101290104"
                },
                {
                   "id": "290104",
                   "name": "晋宁",
                   "weatherCode": "101290105"
                },
                {
                   "id": "290105",
                   "name": "宜良",
                   "weatherCode": "101290106"
                },
                {
                   "id": "290106",
                   "name": "石林",
                   "weatherCode": "101290107"
                },
                {
                   "id": "290107",
                   "name": "呈贡",
                   "weatherCode": "101290108"
                },
                {
                   "id": "290108",
                   "name": "富民",
                   "weatherCode": "101290109"
                },
                {
                   "id": "290109",
                   "name": "嵩明",
                   "weatherCode": "101290110"
                },
                {
                   "id": "290110",
                   "name": "禄劝",
                   "weatherCode": "101290111"
                },
                {
                   "id": "290111",
                   "name": "安宁",
                   "weatherCode": "101290112"
                },
                {
                   "id": "290112",
                   "name": "太华山",
                   "weatherCode": "101290113"
                }
             ]
          },
          {
             "id": "2902",
             "name": "大理",
             "county": [
                {
                   "id": "290201",
                   "name": "大理",
                   "weatherCode": "101290201"
                },
                {
                   "id": "290202",
                   "name": "云龙",
                   "weatherCode": "101290202"
                },
                {
                   "id": "290203",
                   "name": "漾濞",
                   "weatherCode": "101290203"
                },
                {
                   "id": "290204",
                   "name": "永平",
                   "weatherCode": "101290204"
                },
                {
                   "id": "290205",
                   "name": "宾川",
                   "weatherCode": "101290205"
                },
                {
                   "id": "290206",
                   "name": "弥渡",
                   "weatherCode": "101290206"
                },
                {
                   "id": "290207",
                   "name": "祥云",
                   "weatherCode": "101290207"
                },
                {
                   "id": "290208",
                   "name": "巍山",
                   "weatherCode": "101290208"
                },
                {
                   "id": "290209",
                   "name": "剑川",
                   "weatherCode": "101290209"
                },
                {
                   "id": "290210",
                   "name": "洱源",
                   "weatherCode": "101290210"
                },
                {
                   "id": "290211",
                   "name": "鹤庆",
                   "weatherCode": "101290211"
                },
                {
                   "id": "290212",
                   "name": "南涧",
                   "weatherCode": "101290212"
                }
             ]
          },
          {
             "id": "2903",
             "name": "红河",
             "county": [
                {
                   "id": "290301",
                   "name": "红河",
                   "weatherCode": "101290301"
                },
                {
                   "id": "290302",
                   "name": "石屏",
                   "weatherCode": "101290302"
                },
                {
                   "id": "290303",
                   "name": "建水",
                   "weatherCode": "101290303"
                },
                {
                   "id": "290304",
                   "name": "弥勒",
                   "weatherCode": "101290304"
                },
                {
                   "id": "290305",
                   "name": "元阳",
                   "weatherCode": "101290305"
                },
                {
                   "id": "290306",
                   "name": "绿春",
                   "weatherCode": "101290306"
                },
                {
                   "id": "290307",
                   "name": "开远",
                   "weatherCode": "101290307"
                },
                {
                   "id": "290308",
                   "name": "个旧",
                   "weatherCode": "101290308"
                },
                {
                   "id": "290309",
                   "name": "蒙自",
                   "weatherCode": "101290309"
                },
                {
                   "id": "290310",
                   "name": "屏边",
                   "weatherCode": "101290310"
                },
                {
                   "id": "290311",
                   "name": "泸西",
                   "weatherCode": "101290311"
                },
                {
                   "id": "290312",
                   "name": "金平",
                   "weatherCode": "101290312"
                }
             ]
          },
          {
             "id": "2904",
             "name": "曲靖",
             "county": [
                {
                   "id": "290401",
                   "name": "曲靖",
                   "weatherCode": "101290401"
                },
                {
                   "id": "290402",
                   "name": "沾益",
                   "weatherCode": "101290402"
                },
                {
                   "id": "290403",
                   "name": "陆良",
                   "weatherCode": "101290403"
                },
                {
                   "id": "290404",
                   "name": "富源",
                   "weatherCode": "101290404"
                },
                {
                   "id": "290405",
                   "name": "马龙",
                   "weatherCode": "101290405"
                },
                {
                   "id": "290406",
                   "name": "师宗",
                   "weatherCode": "101290406"
                },
                {
                   "id": "290407",
                   "name": "罗平",
                   "weatherCode": "101290407"
                },
                {
                   "id": "290408",
                   "name": "会泽",
                   "weatherCode": "101290408"
                },
                {
                   "id": "290409",
                   "name": "宣威",
                   "weatherCode": "101290409"
                }
             ]
          },
          {
             "id": "2905",
             "name": "保山",
             "county": [
                {
                   "id": "290501",
                   "name": "保山",
                   "weatherCode": "101290501"
                },
                {
                   "id": "290502",
                   "name": "龙陵",
                   "weatherCode": "101290503"
                },
                {
                   "id": "290503",
                   "name": "施甸",
                   "weatherCode": "101290504"
                },
                {
                   "id": "290504",
                   "name": "昌宁",
                   "weatherCode": "101290505"
                },
                {
                   "id": "290505",
                   "name": "腾冲",
                   "weatherCode": "101290506"
                }
             ]
          },
          {
             "id": "2906",
             "name": "文山",
             "county": [
                {
                   "id": "290601",
                   "name": "文山",
                   "weatherCode": "101290601"
                },
                {
                   "id": "290602",
                   "name": "西畴",
                   "weatherCode": "101290602"
                },
                {
                   "id": "290603",
                   "name": "马关",
                   "weatherCode": "101290603"
                },
                {
                   "id": "290604",
                   "name": "麻栗坡",
                   "weatherCode": "101290604"
                },
                {
                   "id": "290605",
                   "name": "砚山",
                   "weatherCode": "101290605"
                },
                {
                   "id": "290606",
                   "name": "丘北",
                   "weatherCode": "101290606"
                },
                {
                   "id": "290607",
                   "name": "广南",
                   "weatherCode": "101290607"
                },
                {
                   "id": "290608",
                   "name": "富宁",
                   "weatherCode": "101290608"
                }
             ]
          },
          {
             "id": "2907",
             "name": "玉溪",
             "county": [
                {
                   "id": "290701",
                   "name": "玉溪",
                   "weatherCode": "101290701"
                },
                {
                   "id": "290702",
                   "name": "澄江",
                   "weatherCode": "101290702"
                },
                {
                   "id": "290703",
                   "name": "江川",
                   "weatherCode": "101290703"
                },
                {
                   "id": "290704",
                   "name": "通海",
                   "weatherCode": "101290704"
                },
                {
                   "id": "290705",
                   "name": "华宁",
                   "weatherCode": "101290705"
                },
                {
                   "id": "290706",
                   "name": "新平",
                   "weatherCode": "101290706"
                },
                {
                   "id": "290707",
                   "name": "易门",
                   "weatherCode": "101290707"
                },
                {
                   "id": "290708",
                   "name": "峨山",
                   "weatherCode": "101290708"
                },
                {
                   "id": "290709",
                   "name": "元江",
                   "weatherCode": "101290709"
                }
             ]
          },
          {
             "id": "2908",
             "name": "楚雄",
             "county": [
                {
                   "id": "290801",
                   "name": "楚雄",
                   "weatherCode": "101290801"
                },
                {
                   "id": "290802",
                   "name": "大姚",
                   "weatherCode": "101290802"
                },
                {
                   "id": "290803",
                   "name": "元谋",
                   "weatherCode": "101290803"
                },
                {
                   "id": "290804",
                   "name": "姚安",
                   "weatherCode": "101290804"
                },
                {
                   "id": "290805",
                   "name": "牟定",
                   "weatherCode": "101290805"
                },
                {
                   "id": "290806",
                   "name": "南华",
                   "weatherCode": "101290806"
                },
                {
                   "id": "290807",
                   "name": "武定",
                   "weatherCode": "101290807"
                },
                {
                   "id": "290808",
                   "name": "禄丰",
                   "weatherCode": "101290808"
                },
                {
                   "id": "290809",
                   "name": "双柏",
                   "weatherCode": "101290809"
                },
                {
                   "id": "290810",
                   "name": "永仁",
                   "weatherCode": "101290810"
                }
             ]
          },
          {
             "id": "2909",
             "name": "普洱",
             "county": [
                {
                   "id": "290901",
                   "name": "普洱",
                   "weatherCode": "101290901"
                },
                {
                   "id": "290902",
                   "name": "景谷",
                   "weatherCode": "101290902"
                },
                {
                   "id": "290903",
                   "name": "景东",
                   "weatherCode": "101290903"
                },
                {
                   "id": "290904",
                   "name": "澜沧",
                   "weatherCode": "101290904"
                },
                {
                   "id": "290905",
                   "name": "墨江",
                   "weatherCode": "101290906"
                },
                {
                   "id": "290906",
                   "name": "江城",
                   "weatherCode": "101290907"
                },
                {
                   "id": "290907",
                   "name": "孟连",
                   "weatherCode": "101290908"
                },
                {
                   "id": "290908",
                   "name": "西盟",
                   "weatherCode": "101290909"
                },
                {
                   "id": "290909",
                   "name": "镇沅",
                   "weatherCode": "101290911"
                },
                {
                   "id": "290910",
                   "name": "宁洱",
                   "weatherCode": "101290912"
                }
             ]
          },
          {
             "id": "2910",
             "name": "昭通",
             "county": [
                {
                   "id": "291001",
                   "name": "昭通",
                   "weatherCode": "101291001"
                },
                {
                   "id": "291002",
                   "name": "鲁甸",
                   "weatherCode": "101291002"
                },
                {
                   "id": "291003",
                   "name": "彝良",
                   "weatherCode": "101291003"
                },
                {
                   "id": "291004",
                   "name": "镇雄",
                   "weatherCode": "101291004"
                },
                {
                   "id": "291005",
                   "name": "威信",
                   "weatherCode": "101291005"
                },
                {
                   "id": "291006",
                   "name": "巧家",
                   "weatherCode": "101291006"
                },
                {
                   "id": "291007",
                   "name": "绥江",
                   "weatherCode": "101291007"
                },
                {
                   "id": "291008",
                   "name": "永善",
                   "weatherCode": "101291008"
                },
                {
                   "id": "291009",
                   "name": "盐津",
                   "weatherCode": "101291009"
                },
                {
                   "id": "291010",
                   "name": "大关",
                   "weatherCode": "101291010"
                },
                {
                   "id": "291011",
                   "name": "水富",
                   "weatherCode": "101291011"
                }
             ]
          },
          {
             "id": "2911",
             "name": "临沧",
             "county": [
                {
                   "id": "291101",
                   "name": "临沧",
                   "weatherCode": "101291101"
                },
                {
                   "id": "291102",
                   "name": "沧源",
                   "weatherCode": "101291102"
                },
                {
                   "id": "291103",
                   "name": "耿马",
                   "weatherCode": "101291103"
                },
                {
                   "id": "291104",
                   "name": "双江",
                   "weatherCode": "101291104"
                },
                {
                   "id": "291105",
                   "name": "凤庆",
                   "weatherCode": "101291105"
                },
                {
                   "id": "291106",
                   "name": "永德",
                   "weatherCode": "101291106"
                },
                {
                   "id": "291107",
                   "name": "云县",
                   "weatherCode": "101291107"
                },
                {
                   "id": "291108",
                   "name": "镇康",
                   "weatherCode": "101291108"
                }
             ]
          },
          {
             "id": "2912",
             "name": "怒江",
             "county": [
                {
                   "id": "291201",
                   "name": "怒江",
                   "weatherCode": "101291201"
                },
                {
                   "id": "291202",
                   "name": "福贡",
                   "weatherCode": "101291203"
                },
                {
                   "id": "291203",
                   "name": "兰坪",
                   "weatherCode": "101291204"
                },
                {
                   "id": "291204",
                   "name": "泸水",
                   "weatherCode": "101291205"
                },
                {
                   "id": "291205",
                   "name": "六库",
                   "weatherCode": "101291206"
                },
                {
                   "id": "291206",
                   "name": "贡山",
                   "weatherCode": "101291207"
                }
             ]
          },
          {
             "id": "2913",
             "name": "迪庆",
             "county": [
                {
                   "id": "291301",
                   "name": "香格里拉",
                   "weatherCode": "101291301"
                },
                {
                   "id": "291302",
                   "name": "德钦",
                   "weatherCode": "101291302"
                },
                {
                   "id": "291303",
                   "name": "维西",
                   "weatherCode": "101291303"
                },
                {
                   "id": "291304",
                   "name": "中甸",
                   "weatherCode": "101291304"
                }
             ]
          },
          {
             "id": "2914",
             "name": "丽江",
             "county": [
                {
                   "id": "291401",
                   "name": "丽江",
                   "weatherCode": "101291401"
                },
                {
                   "id": "291402",
                   "name": "永胜",
                   "weatherCode": "101291402"
                },
                {
                   "id": "291403",
                   "name": "华坪",
                   "weatherCode": "101291403"
                },
                {
                   "id": "291404",
                   "name": "宁蒗",
                   "weatherCode": "101291404"
                }
             ]
          },
          {
             "id": "2915",
             "name": "德宏",
             "county": [
                {
                   "id": "291501",
                   "name": "德宏",
                   "weatherCode": "101291501"
                },
                {
                   "id": "291502",
                   "name": "陇川",
                   "weatherCode": "101291503"
                },
                {
                   "id": "291503",
                   "name": "盈江",
                   "weatherCode": "101291504"
                },
                {
                   "id": "291504",
                   "name": "瑞丽",
                   "weatherCode": "101291506"
                },
                {
                   "id": "291505",
                   "name": "梁河",
                   "weatherCode": "101291507"
                },
                {
                   "id": "291506",
                   "name": "潞西",
                   "weatherCode": "101291508"
                }
             ]
          },
          {
             "id": "2916",
             "name": "西双版纳",
             "county": [
                {
                   "id": "291601",
                   "name": "景洪",
                   "weatherCode": "101291601"
                },
                {
                   "id": "291602",
                   "name": "勐海",
                   "weatherCode": "101291603"
                },
                {
                   "id": "291603",
                   "name": "勐腊",
                   "weatherCode": "101291605"
                }
             ]
          }
       ]
    },
    {
       "id": "30",
       "name": "广西",
       "city": [
          {
             "id": "3001",
             "name": "南宁",
             "county": [
                {
                   "id": "300101",
                   "name": "南宁",
                   "weatherCode": "101300101"
                },
                {
                   "id": "300102",
                   "name": "邕宁",
                   "weatherCode": "101300103"
                },
                {
                   "id": "300103",
                   "name": "横县",
                   "weatherCode": "101300104"
                },
                {
                   "id": "300104",
                   "name": "隆安",
                   "weatherCode": "101300105"
                },
                {
                   "id": "300105",
                   "name": "马山",
                   "weatherCode": "101300106"
                },
                {
                   "id": "300106",
                   "name": "上林",
                   "weatherCode": "101300107"
                },
                {
                   "id": "300107",
                   "name": "武鸣",
                   "weatherCode": "101300108"
                },
                {
                   "id": "300108",
                   "name": "宾阳",
                   "weatherCode": "101300109"
                }
             ]
          },
          {
             "id": "3002",
             "name": "崇左",
             "county": [
                {
                   "id": "300201",
                   "name": "崇左",
                   "weatherCode": "101300201"
                },
                {
                   "id": "300202",
                   "name": "天等",
                   "weatherCode": "101300202"
                },
                {
                   "id": "300203",
                   "name": "龙州",
                   "weatherCode": "101300203"
                },
                {
                   "id": "300204",
                   "name": "凭祥",
                   "weatherCode": "101300204"
                },
                {
                   "id": "300205",
                   "name": "大新",
                   "weatherCode": "101300205"
                },
                {
                   "id": "300206",
                   "name": "扶绥",
                   "weatherCode": "101300206"
                },
                {
                   "id": "300207",
                   "name": "宁明",
                   "weatherCode": "101300207"
                }
             ]
          },
          {
             "id": "3003",
             "name": "柳州",
             "county": [
                {
                   "id": "300301",
                   "name": "柳州",
                   "weatherCode": "101300301"
                },
                {
                   "id": "300302",
                   "name": "柳城",
                   "weatherCode": "101300302"
                },
                {
                   "id": "300303",
                   "name": "鹿寨",
                   "weatherCode": "101300304"
                },
                {
                   "id": "300304",
                   "name": "柳江",
                   "weatherCode": "101300305"
                },
                {
                   "id": "300305",
                   "name": "融安",
                   "weatherCode": "101300306"
                },
                {
                   "id": "300306",
                   "name": "融水",
                   "weatherCode": "101300307"
                },
                {
                   "id": "300307",
                   "name": "三江",
                   "weatherCode": "101300308"
                }
             ]
          },
          {
             "id": "3004",
             "name": "来宾",
             "county": [
                {
                   "id": "300401",
                   "name": "来宾",
                   "weatherCode": "101300401"
                },
                {
                   "id": "300402",
                   "name": "忻城",
                   "weatherCode": "101300402"
                },
                {
                   "id": "300403",
                   "name": "金秀",
                   "weatherCode": "101300403"
                },
                {
                   "id": "300404",
                   "name": "象州",
                   "weatherCode": "101300404"
                },
                {
                   "id": "300405",
                   "name": "武宣",
                   "weatherCode": "101300405"
                },
                {
                   "id": "300406",
                   "name": "合山",
                   "weatherCode": "101300406"
                }
             ]
          },
          {
             "id": "3005",
             "name": "桂林",
             "county": [
                {
                   "id": "300501",
                   "name": "桂林",
                   "weatherCode": "101300501"
                },
                {
                   "id": "300502",
                   "name": "龙胜",
                   "weatherCode": "101300503"
                },
                {
                   "id": "300503",
                   "name": "永福",
                   "weatherCode": "101300504"
                },
                {
                   "id": "300504",
                   "name": "临桂",
                   "weatherCode": "101300505"
                },
                {
                   "id": "300505",
                   "name": "兴安",
                   "weatherCode": "101300506"
                },
                {
                   "id": "300506",
                   "name": "灵川",
                   "weatherCode": "101300507"
                },
                {
                   "id": "300507",
                   "name": "全州",
                   "weatherCode": "101300508"
                },
                {
                   "id": "300508",
                   "name": "灌阳",
                   "weatherCode": "101300509"
                },
                {
                   "id": "300509",
                   "name": "阳朔",
                   "weatherCode": "101300510"
                },
                {
                   "id": "300510",
                   "name": "恭城",
                   "weatherCode": "101300511"
                },
                {
                   "id": "300511",
                   "name": "平乐",
                   "weatherCode": "101300512"
                },
                {
                   "id": "300512",
                   "name": "荔浦",
                   "weatherCode": "101300513"
                },
                {
                   "id": "300513",
                   "name": "资源",
                   "weatherCode": "101300514"
                }
             ]
          },
          {
             "id": "3006",
             "name": "梧州",
             "county": [
                {
                   "id": "300601",
                   "name": "梧州",
                   "weatherCode": "101300601"
                },
                {
                   "id": "300602",
                   "name": "藤县",
                   "weatherCode": "101300602"
                },
                {
                   "id": "300603",
                   "name": "苍梧",
                   "weatherCode": "101300604"
                },
                {
                   "id": "300604",
                   "name": "蒙山",
                   "weatherCode": "101300605"
                },
                {
                   "id": "300605",
                   "name": "岑溪",
                   "weatherCode": "101300606"
                }
             ]
          },
          {
             "id": "3007",
             "name": "贺州",
             "county": [
                {
                   "id": "300701",
                   "name": "贺州",
                   "weatherCode": "101300701"
                },
                {
                   "id": "300702",
                   "name": "昭平",
                   "weatherCode": "101300702"
                },
                {
                   "id": "300703",
                   "name": "富川",
                   "weatherCode": "101300703"
                },
                {
                   "id": "300704",
                   "name": "钟山",
                   "weatherCode": "101300704"
                }
             ]
          },
          {
             "id": "3008",
             "name": "贵港",
             "county": [
                {
                   "id": "300801",
                   "name": "贵港",
                   "weatherCode": "101300801"
                },
                {
                   "id": "300802",
                   "name": "桂平",
                   "weatherCode": "101300802"
                },
                {
                   "id": "300803",
                   "name": "平南",
                   "weatherCode": "101300803"
                }
             ]
          },
          {
             "id": "3009",
             "name": "玉林",
             "county": [
                {
                   "id": "300901",
                   "name": "玉林",
                   "weatherCode": "101300901"
                },
                {
                   "id": "300902",
                   "name": "博白",
                   "weatherCode": "101300902"
                },
                {
                   "id": "300903",
                   "name": "北流",
                   "weatherCode": "101300903"
                },
                {
                   "id": "300904",
                   "name": "容县",
                   "weatherCode": "101300904"
                },
                {
                   "id": "300905",
                   "name": "陆川",
                   "weatherCode": "101300905"
                },
                {
                   "id": "300906",
                   "name": "兴业",
                   "weatherCode": "101300906"
                }
             ]
          },
          {
             "id": "3010",
             "name": "百色",
             "county": [
                {
                   "id": "301001",
                   "name": "百色",
                   "weatherCode": "101301001"
                },
                {
                   "id": "301002",
                   "name": "那坡",
                   "weatherCode": "101301002"
                },
                {
                   "id": "301003",
                   "name": "田阳",
                   "weatherCode": "101301003"
                },
                {
                   "id": "301004",
                   "name": "德保",
                   "weatherCode": "101301004"
                },
                {
                   "id": "301005",
                   "name": "靖西",
                   "weatherCode": "101301005"
                },
                {
                   "id": "301006",
                   "name": "田东",
                   "weatherCode": "101301006"
                },
                {
                   "id": "301007",
                   "name": "平果",
                   "weatherCode": "101301007"
                },
                {
                   "id": "301008",
                   "name": "隆林",
                   "weatherCode": "101301008"
                },
                {
                   "id": "301009",
                   "name": "西林",
                   "weatherCode": "101301009"
                },
                {
                   "id": "301010",
                   "name": "乐业",
                   "weatherCode": "101301010"
                },
                {
                   "id": "301011",
                   "name": "凌云",
                   "weatherCode": "101301011"
                },
                {
                   "id": "301012",
                   "name": "田林",
                   "weatherCode": "101301012"
                }
             ]
          },
          {
             "id": "3011",
             "name": "钦州",
             "county": [
                {
                   "id": "301101",
                   "name": "钦州",
                   "weatherCode": "101301101"
                },
                {
                   "id": "301102",
                   "name": "浦北",
                   "weatherCode": "101301102"
                },
                {
                   "id": "301103",
                   "name": "灵山",
                   "weatherCode": "101301103"
                }
             ]
          },
          {
             "id": "3012",
             "name": "河池",
             "county": [
                {
                   "id": "301201",
                   "name": "河池",
                   "weatherCode": "101301201"
                },
                {
                   "id": "301202",
                   "name": "天峨",
                   "weatherCode": "101301202"
                },
                {
                   "id": "301203",
                   "name": "东兰",
                   "weatherCode": "101301203"
                },
                {
                   "id": "301204",
                   "name": "巴马",
                   "weatherCode": "101301204"
                },
                {
                   "id": "301205",
                   "name": "环江",
                   "weatherCode": "101301205"
                },
                {
                   "id": "301206",
                   "name": "罗城",
                   "weatherCode": "101301206"
                },
                {
                   "id": "301207",
                   "name": "宜州",
                   "weatherCode": "101301207"
                },
                {
                   "id": "301208",
                   "name": "凤山",
                   "weatherCode": "101301208"
                },
                {
                   "id": "301209",
                   "name": "南丹",
                   "weatherCode": "101301209"
                },
                {
                   "id": "301210",
                   "name": "都安",
                   "weatherCode": "101301210"
                },
                {
                   "id": "301211",
                   "name": "大化",
                   "weatherCode": "101301211"
                }
             ]
          },
          {
             "id": "3013",
             "name": "北海",
             "county": [
                {
                   "id": "301301",
                   "name": "北海",
                   "weatherCode": "101301301"
                },
                {
                   "id": "301302",
                   "name": "合浦",
                   "weatherCode": "101301302"
                },
                {
                   "id": "301303",
                   "name": "涠洲岛",
                   "weatherCode": "101301303"
                }
             ]
          },
          {
             "id": "3014",
             "name": "防城港",
             "county": [
                {
                   "id": "301401",
                   "name": "防城港",
                   "weatherCode": "101301401"
                },
                {
                   "id": "301402",
                   "name": "上思",
                   "weatherCode": "101301402"
                },
                {
                   "id": "301404",
                   "name": "防城",
                   "weatherCode": "101301405"
                }
             ]
          }
       ]
    },
    {
       "id": "31",
       "name": "海南",
       "city": [
          {
             "id": "3101",
             "name": "海口",
             "county": [{
                "id": "310101",
                "name": "海口",
                "weatherCode": "101310101"
             }]
          },
          {
             "id": "3102",
             "name": "三亚",
             "county": [{
                "id": "310201",
                "name": "三亚",
                "weatherCode": "101310201"
             }]
          },
          {
             "id": "3103",
             "name": "东方",
             "county": [{
                "id": "310301",
                "name": "东方",
                "weatherCode": "101310202"
             }]
          },
          {
             "id": "3104",
             "name": "临高",
             "county": [{
                "id": "310401",
                "name": "临高",
                "weatherCode": "101310203"
             }]
          },
          {
             "id": "3105",
             "name": "澄迈",
             "county": [{
                "id": "310501",
                "name": "澄迈",
                "weatherCode": "101310204"
             }]
          },
          {
             "id": "3106",
             "name": "儋州",
             "county": [{
                "id": "310601",
                "name": "儋州",
                "weatherCode": "101310205"
             }]
          },
          {
             "id": "3107",
             "name": "昌江",
             "county": [{
                "id": "310701",
                "name": "昌江",
                "weatherCode": "101310206"
             }]
          },
          {
             "id": "3108",
             "name": "白沙",
             "county": [{
                "id": "310801",
                "name": "白沙",
                "weatherCode": "101310207"
             }]
          },
          {
             "id": "3109",
             "name": "琼中",
             "county": [{
                "id": "310901",
                "name": "琼中",
                "weatherCode": "101310208"
             }]
          },
          {
             "id": "3110",
             "name": "定安",
             "county": [{
                "id": "311001",
                "name": "定安",
                "weatherCode": "101310209"
             }]
          },
          {
             "id": "3111",
             "name": "屯昌",
             "county": [{
                "id": "311101",
                "name": "屯昌",
                "weatherCode": "101310210"
             }]
          },
          {
             "id": "3112",
             "name": "琼海",
             "county": [{
                "id": "311201",
                "name": "琼海",
                "weatherCode": "101310211"
             }]
          },
          {
             "id": "3113",
             "name": "文昌",
             "county": [{
                "id": "311301",
                "name": "文昌",
                "weatherCode": "101310212"
             }]
          },
          {
             "id": "3114",
             "name": "保亭",
             "county": [{
                "id": "311401",
                "name": "保亭",
                "weatherCode": "101310214"
             }]
          },
          {
             "id": "3115",
             "name": "万宁",
             "county": [{
                "id": "311501",
                "name": "万宁",
                "weatherCode": "101310215"
             }]
          },
          {
             "id": "3116",
             "name": "陵水",
             "county": [{
                "id": "311601",
                "name": "陵水",
                "weatherCode": "101310216"
             }]
          },
          {
             "id": "3117",
             "name": "西沙",
             "county": [{
                "id": "311701",
                "name": "西沙",
                "weatherCode": "101310217"
             }]
          },
          {
             "id": "3118",
             "name": "南沙",
             "county": [{
                "id": "311801",
                "name": "南沙",
                "weatherCode": "101310220"
             }]
          },
          {
             "id": "3119",
             "name": "乐东",
             "county": [{
                "id": "311901",
                "name": "乐东",
                "weatherCode": "101310221"
             }]
          },
          {
             "id": "3120",
             "name": "五指山",
             "county": [{
                "id": "312001",
                "name": "五指山",
                "weatherCode": "101310222"
             }]
          }
       ]
    },
    {
       "id": "32",
       "name": "香港",
       "city": [{
          "id": "3201",
          "name": "香港",
          "county": [
             {
                "id": "320101",
                "name": "香港",
                "weatherCode": "101320101"
             },
             {
                "id": "320102",
                "name": "九龙",
                "weatherCode": "101320102"
             },
             {
                "id": "320103",
                "name": "新界",
                "weatherCode": "101320103"
             }
          ]
       }]
    },
    {
       "id": "33",
       "name": "澳门",
       "city": [{
          "id": "3301",
          "name": "澳门",
          "county": [
             {
                "id": "330101",
                "name": "澳门",
                "weatherCode": "101330101"
             },
             {
                "id": "330102",
                "name": "氹仔岛",
                "weatherCode": "101330102"
             },
             {
                "id": "330103",
                "name": "路环岛",
                "weatherCode": "101330103"
             }
          ]
       }]
    },
    {
       "id": "34",
       "name": "台湾",
       "city": [
          {
             "id": "3401",
             "name": "台北",
             "county": [
                {
                   "id": "340101",
                   "name": "台北",
                   "weatherCode": "101340101"
                },
                {
                   "id": "340102",
                   "name": "桃园",
                   "weatherCode": "101340102"
                },
                {
                   "id": "340103",
                   "name": "新竹",
                   "weatherCode": "101340103"
                },
                {
                   "id": "340104",
                   "name": "宜兰",
                   "weatherCode": "101340104"
                }
             ]
          },
          {
             "id": "3402",
             "name": "高雄",
             "county": [
                {
                   "id": "340201",
                   "name": "高雄",
                   "weatherCode": "101340201"
                },
                {
                   "id": "340202",
                   "name": "嘉义",
                   "weatherCode": "101340202"
                },
                {
                   "id": "340203",
                   "name": "台南",
                   "weatherCode": "101340203"
                },
                {
                   "id": "340204",
                   "name": "台东",
                   "weatherCode": "101340204"
                },
                {
                   "id": "340205",
                   "name": "屏东",
                   "weatherCode": "101340205"
                }
             ]
          },
          {
             "id": "3403",
             "name": "台中",
             "county": [
                {
                   "id": "340301",
                   "name": "台中",
                   "weatherCode": "101340401"
                },
                {
                   "id": "340302",
                   "name": "苗栗",
                   "weatherCode": "101340402"
                },
                {
                   "id": "340303",
                   "name": "彰化",
                   "weatherCode": "101340403"
                },
                {
                   "id": "340304",
                   "name": "南投",
                   "weatherCode": "101340404"
                },
                {
                   "id": "340305",
                   "name": "花莲",
                   "weatherCode": "101340405"
                },
                {
                   "id": "340306",
                   "name": "云林",
                   "weatherCode": "101340406"
                }
             ]
          }
       ]
    }
 ]`
)
