package test

import (
	"bou.ke/monkey"
	"mygo/internal/server/model"
	"mygo/internal/server/service"
	"testing"
)

func TestSearch(t *testing.T) {
	returnList := []*model.Transaction{
		{
			Title:       "紧急修理",
			Description: "地点：蒙德-（法拉正准备运货去醉汉峡附近，向路过的商队兜售货物，但是货车却被丘丘人袭击，货物和他自己都陷入了危险的境地…）",
		},
		{
			Title:       "圆滚滚的大团骚乱",
			Description: "地点：蒙德-晨曦酒庄（史莱姆突然在晨曦酒庄出现，在酒庄附近的特纳似乎被卷入其中…）",
		},
		{
			Title:       "猫的留影",
			Description: "地点：稻妻-浅濑神社（在浅濑神社工作的大岛纯平似乎正在为木雕的事而发愁…）",
		},
		{
			Title:       "稻妻时尚入门",
			Description: "地点：稻妻-稻妻城（町街小仓屋）（「小仓屋」的店主澪正在犯愁…）",
		},
		{
			Title:       "必须精进的武艺",
			Description: "地点：稻妻-稻妻城（天守阁北海滩）（一名叫做朝仓的家伙，似乎正在找人帮他锻炼武艺…）",
		},
		{
			Title:       "全能美食队·烹饪对决",
			Description: "地点：稻妻-镇守之森（全能美食队的旭东和龟井宗久似乎正在准备进行一场烹饪对决...）",
		},
	}

	monkey.Patch(model.GetPassedTransactions, func() ([]*model.Transaction, error) {
		return returnList, nil
	})

	service.Seg.LoadDict()

	t.Log("----------Case 1----------")
	t.Log("Search for 稻妻")
	response, err := service.SearchTransactions("稻妻")
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Log(r.Title)
		t.Log(r.Description + "\n")
	}

	t.Log("----------Case 2----------")
	t.Log("Search for 锻炼武艺")
	response, err = service.SearchTransactions("锻炼武艺")
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Log(r.Title)
		t.Log(r.Description + "\n")
	}
}
