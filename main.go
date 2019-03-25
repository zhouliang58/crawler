package main

import (
	"crawler/poi"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var keys = [6]string{"批发", "商铺", "公司", "市场", "机构", "场馆"}
var headers = [5]string{"名称", "地址", "电话", "经度", "纬度"}
var region = "广州"

// 百度申请：应用AK
var ak = ""

func main() {
	file := xlsx.NewFile()
	var totalNum = 0
	for _, key := range keys {
		// 每个关键词一个sheet
		sheet, _ := file.AddSheet(key)
		row := sheet.AddRow()
		for _, v := range headers {
			cell := row.AddCell()
			cell.Value = string(v)
			fmt.Println(cell.Value)
		}
		fmt.Printf("关键词:%s\n", key)
		for pageNum := 0; pageNum < 20; pageNum++ {
			url := "http://api.map.baidu.com/place/v2/search?query=" + key + "&region=" + region +
				"&city_limit=true&page_size=20&page_num=" + strconv.Itoa(pageNum) +
				"&output=json&ak=" + ak +
				// 百度申请：SHA1 包名
				"&mcode=CB:EB:2A:70:86:D7:9D:DD:CA:40:E1:9F:8F:74:A1:DE:17:F0:51:9C;com.example.hello_world"
			resp, _ := http.Get(url)
			body, _ := ioutil.ReadAll(resp.Body)
			poiResponse := &poi.PoiResponse{}
			_ = json.Unmarshal(body, poiResponse)
			fmt.Println(poiResponse.Total)
			fmt.Println(poiResponse.Message)
			if poiResponse.Total == 0 {
				break
			}
			for _, v := range poiResponse.Results {
				totalNum++
				fmt.Printf("关键词:%s, 第%d条数据\n", key, totalNum)
				row := sheet.AddRow()
				cell1 := row.AddCell()
				cell1.Value = v.Name
				cell2 := row.AddCell()
				cell2.Value = v.Address
				cell3 := row.AddCell()
				cell3.Value = v.Telephone
				cell4 := row.AddCell()
				cell4.Value = strconv.FormatFloat(v.Location.Lgt, 'f', -1, 64)
				cell5 := row.AddCell()
				cell5.Value = strconv.FormatFloat(v.Location.Lat, 'f', -1, 64)
			}
			_ = resp.Body.Close()
			if totalNum%20 == 0 {
				time.Sleep(time.Duration(2) * time.Second)
			}
			if totalNum%100 == 0 {
				time.Sleep(time.Duration(5) * time.Second)
			}
			if totalNum%200 == 0 {
				time.Sleep(time.Duration(7) * time.Second)
			}
		}
	}
	fmt.Println(totalNum)
	_ = file.Save("广州商家信息.xlsx")
}
