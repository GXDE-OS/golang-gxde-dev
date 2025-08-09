/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package calendar

import (
	"strings"
	"time"
	"fmt"
)

type Day struct {
	Year, Month, Day int
}

var solarFestival = map[int]string{
	101:  "元旦",
	110:  "中国人民警察节",
	202:  "世界湿地日",
	210:  "国际气象节",
	214:  "情人节",
	301:  "国际海豹日",
	303:  "全国爱耳日",
	305:  "学雷锋纪念日",
	308:  "妇女节",
	312:  "植树节 孙中山逝世纪念日",
	314:  "国际警察日",
	315:  "消费者权益日",
	317:  "中国国医节 国际航海日",
	321:  "世界森林日 消除种族歧视国际日 世界儿歌日",
	322:  "世界水日",
	323:  "世界气象日",
	324:  "世界防治结核病日",
	325:  "全国中小学生安全教育日",
	330:  "巴勒斯坦国土日",
	401:  "愚人节",
	407:  "世界卫生日",
	422:  "世界地球日",
	423:  "世界图书和版权日",
	424:  "亚非新闻工作者日",
	501:  "劳动节",
	504:  "青年节",
	505:  "碘缺乏病防治日",
	508:  "世界红十字日",
	512:  "国际护士节",
	515:  "国际家庭日",
	517:  "世界电信日",
	518:  "国际博物馆日",
	520:  "全国学生营养日",
	522:  "国际生物多样性日",
	523:  "国际牛奶日",
	531:  "世界无烟日",
	601:  "国际儿童节",
	605:  "世界环境日",
	606:  "全国爱眼日",
	617:  "防治荒漠化和干旱日",
	623:  "国际奥林匹克日",
	625:  "全国土地日",
	626:  "国际禁毒日",
	701:  "香港回归纪念日 中共诞辰 世界建筑日",
	702:  "国际体育记者日",
	707:  "抗日战争纪念日",
	711:  "世界人口日",
	730:  "非洲妇女日",
	801:  "建军节",
	803:  "男人节",
	804:  "国际云豹日",
	808:  "全民健身日 国际猫咪日 中国男子节(爸爸节)",
	809:  "世界土著人民国际日",
	810:  "国际狮子日",
	811:  "世界钢鼓日 全国肢残人活动日",
	812:  "国际青年节 世界大象日",
	813:  "国际狼日",
	814:  "绿色情人节",
	815:  "抗日战争胜利纪念",
	818:  "中国人力资源日",
	819:  "中国医师节 世界人道主义日 世界摄影日",
	820:  "世界蚊子日",
	825:  "全国残疾预防日",
	826:  "律师咨询日",
	830:  "国际鲸鲨日",
	903:  "中国人民抗日战争胜利纪念日",
	908:  "国际扫盲日 国际新闻工作者日",
	909:  "毛泽东逝世纪念",
	910:  "中国教师节",
	914:  "世界清洁地球日",
	916:  "国际臭氧层保护日",
	918:  "九一八事变纪念日",
	920:  "国际爱牙日",
	927:  "世界旅游日",
	928:  "孔子诞辰",
	930:  "中国烈士纪念日",
	1001: "国庆节 世界音乐日 国际老人节",
	1002: "国际和平与民主自由斗争日",
	1004: "世界动物日",
	1006: "老人节",
	1008: "全国高血压日 世界视觉日",
	1009: "世界邮政日 万国邮联日",
	1010: "辛亥革命纪念日 世界精神卫生日",
	1013: "世界保健日 国际教师节",
	1014: "世界标准日",
	1015: "国际盲人节(白手杖节)",
	1016: "世界粮食日",
	1017: "世界消除贫困日",
	1022: "世界传统医药日",
	1024: "联合国日 世界发展信息日",
	1031: "世界勤俭日 万圣夜",
	1107: "十月社会主义革命纪念日",
	1108: "中国记者日",
	1109: "全国消防安全宣传教育日",
	1110: "世界青年节",
	1111: "国际科学与和平周(本日所属的一周)",
	1112: "孙中山诞辰纪念日",
	1114: "世界糖尿病日",
	1117: "国际大学生节 世界学生节",
	1121: "世界问候日 世界电视日",
	1129: "国际声援巴勒斯坦人民国际日",
	1201: "世界艾滋病日",
	1203: "世界残疾人日",
	1205: "国际经济和社会发展志愿人员日",
	1208: "国际儿童电视日",
	1209: "世界足球日",
	1210: "世界人权日",
	1212: "西安事变纪念日",
	1213: "南京大屠杀纪念日",
	1220: "澳门回归纪念",
	1221: "国际篮球日",
	1224: "平安夜",
	1225: "圣诞节",
	1226: "毛泽东诞辰纪念",
	1231: "deepin15 正式版发布纪念日",
}

// 动态节日定义表（月份 -> 规则 -> 节日名称）
var dynamicFestivals = map[int]map[string]string{
	5: {
		"second sunday": "母亲节", // 五月第二个星期日
	},
	6: {
		"third sunday": "父亲节", // 六月第三个星期日
	},
	11: {
		"fourth thursday": "感恩节", // 十一月第四个星期四
	},
}

// 动态节日规则处理器
func (d *Day) dynamicFestival() string {
	rules, ok := dynamicFestivals[d.Month]
	if !ok {
		return ""
	}

	date := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	
	// 尝试所有规则
	for rule, name := range rules {
		if matchDynamicRule(date, rule) {
			return name
		}
	}
	
	return ""
}

// 判断日期是否匹配动态规则
func matchDynamicRule(date time.Time, rule string) bool {
	// 解析规则：序数 + 星期几
	var ord, weekday string
	if _, err := fmt.Sscanf(rule, "%s %s", &ord, &weekday); err != nil {
		return false
	}
	
	// 获取目标星期几
	targetWeekday := parseWeekday(weekday)
	if targetWeekday == -1 || date.Weekday() != time.Weekday(targetWeekday) {
		return false
	}
	
	// 获取序数位置
	ordinal := parseOrdinal(ord)
	if ordinal == 0 {
		return false
	}
	
	// 计算该日期是当月的第几个目标星期
	return getWeekPosition(date) == ordinal
}

// 解析星期几
func parseWeekday(wd string) int {
	switch strings.ToLower(wd) {
	case "sunday": return 0
	case "monday": return 1
	case "tuesday": return 2
	case "wednesday": return 3
	case "thursday": return 4
	case "friday": return 5
	case "saturday": return 6
	}
	return -1
}

// 解析序数
func parseOrdinal(ord string) int {
	switch strings.ToLower(ord) {
	case "first": return 1
	case "second": return 2
	case "third": return 3
	case "fourth": return 4
	case "fifth": return 5
	case "last": return -1 // 特殊处理最后一周
	}
	return 0
}

// 计算日期是当月的第几个目标星期
func getWeekPosition(date time.Time) int {
	// 获取当月第一天
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	
	// 计算第一个目标星期几的日期
	firstTarget := firstDay
	for firstTarget.Weekday() != date.Weekday() {
		firstTarget = firstTarget.AddDate(0, 0, 1)
	}
	
	// 计算当前日期是第几个目标星期
	daysDiff := date.Day() - firstTarget.Day()
	weekNum := (daysDiff / 7) + 1
	
	// 如果是最后一周，特殊处理
	if weekNum == 5 || weekNum == 6 {
		// 检查是否是最后一周
		lastDay := firstDay.AddDate(0, 1, -1)
		daysLeft := lastDay.Day() - date.Day()
		if daysLeft < 7 {
			return -1 // 标记为最后一周
		}
	}
	
	return weekNum
}

func (d *Day) Festival() string {
	// 检查固定节日
	key := d.Month*100 + d.Day
	if name, ok := solarFestival[key]; ok {
		return name
	}
	// 检查动态节日
	if name := d.dynamicFestival(); name != "" {
		return name
	}
	return ""
}
