package common

import "fmt"

type Complanint struct {
	SN        string // 投诉编号
	Brand     string // 品牌
	BrandID   string // 品牌id
	Series    string // 车系
	SeriesID  string //车系id
	Model     string // 型号
	ModelID   string // 型号id
	Sketch    string // 简述
	SketchURL string // 简述连接
	Date      string // 投诉日期
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
