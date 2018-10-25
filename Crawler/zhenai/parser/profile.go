package parser

import (
	"Crawler/engine"
	"Crawler/model"
	"regexp"
	"strconv"
)

var AgeRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span>([\d]+)</td>`)
var MarriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span>([^<]+)</td>`)
var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span>([^<]+)</td>`)
func Profile(contents []byte, name string) engine.ParseResult {
	//fmt.Println(contents)


	profile := model.Profile{}

	//昵称
	profile.Name = name

	// 婚况
	profile.Marriage = extractString(contents,MarriageRe)

	// 性别
	profile.Gender = extractString(contents,genderRe)

	// 收入
	profile.Income = extractString(contents,incomeRe)

	//教育
	profile.Education = extractString(contents,educationRe)

	//职业
	profile.Occupation = extractString(contents,occupationRe)

	//户口
	profile.Hokou = extractString(contents,hokouRe)

	//星座
	profile.Xinzuo = extractString(contents,xinzuoRe)

	//是否有房
	profile.House = extractString(contents,HouseRe)

	//是否有车
	profile.Car = extractString(contents,CarRe)


	// 年龄
	age, err := strconv.Atoi(extractString(contents,AgeRe))

	if err != nil{
		panic(err)
	}
	profile.Age = age

	//身高
	height, err := strconv.Atoi(extractString(contents,heightRe))

	if err != nil{
		panic(err)
	}
	profile.Height = height

	//体重
	weight, err := strconv.Atoi(extractString(contents,weightRe))
	if err != nil{
		weight = 0
	}
	profile.Weight = weight

	//fmt.Printf("profile:%v \n",profile)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
	//for _,item := range match{
	//	result.Items = append(result.Items,string(item[2]))
	//	result.Requests = append(result.Requests,engine.Request{
	//		Url: string(item[1]),
	//		ParserFunc:engine.NilParser,
	//
	//	})
	//
	//	//fmt.Printf("User: %s Url: %s\n",item[2],item[1])
	//}

}

func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	} else {
		return ""
	}
}
