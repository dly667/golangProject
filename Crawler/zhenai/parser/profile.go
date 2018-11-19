package parser

import (
	"Crawler/engine"
	"Crawler/model"



	"log"

	//"strconv"
	"gopkg.in/xmlpath.v2"
	"strings"
)

var AgeRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[2]`)
var heightRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[4]`)
var weightRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[5]`)
var MarriageRe  = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[1]`)
//var genderRe = xmlpath.MustCompile(`<td><span class="label">性别：</span>([^<]+)</td>`)
var incomeRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[7]`)
var educationRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[9]`)
var occupationRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[8]`)
//var hokouRe = xmlpath.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = xmlpath.MustCompile(`//*[@id="app"]/div[1]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[3]`)
//var HouseRe = xmlpath.MustCompile(`<td><span class="label">住房条件：</span>([^<]+)</td>`)
//var CarRe = xmlpath.MustCompile(`<td><span class="label">是否购车：</span>([^<]+)</td>`)
func Profile(contents []byte, name string) engine.ParseResult {


	profile := model.Profile{}

	//昵称
	profile.Name = name

	// 婚况
	profile.Marriage = extractString(contents,MarriageRe)
	//
	//// 性别
	//profile.Gender = extractString(contents,genderRe)
	//
	//
	//
	// 收入
	profile.Income = extractString(contents,incomeRe)

	//教育
	profile.Education = extractString(contents,educationRe)

	//职业
	profile.Occupation = extractString(contents,occupationRe)

	//户口
	//profile.Hokou = extractString(contents,hokouRe)

	//星座
	profile.Xinzuo = extractString(contents,xinzuoRe)
	//
	////是否有房
	//profile.House = extractString(contents,HouseRe)
	//
	////是否有车
	//profile.Car = extractString(contents,CarRe)
	//
	//
	// 年龄
	age := extractString(contents,AgeRe)

	//if err != nil{
	//	log.Print("age: ",err)
	//}
	profile.Age = age

	//身高
	height := extractString(contents,heightRe)

	//if err != nil{
	//	log.Print("height: ",err)
	//}
	profile.Height = height

	//体重
	weight:= extractString(contents,weightRe)

	profile.Weight = weight


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

func extractString(contents []byte, xlp *xmlpath.Path) string{

	root, err := xmlpath.ParseHTML(strings.NewReader(string(contents)))

	if err != nil {
		log.Fatal(err)
	}

	if value, ok := xlp.String(root); ok {

		return value
	}
	return ""
}
