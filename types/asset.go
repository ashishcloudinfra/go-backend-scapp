package types

type AssetTypeList string

const (
	Cash       AssetTypeList = "Cash"
	Gold       AssetTypeList = "Gold"
	Silver     AssetTypeList = "Silver"
	MutualFund AssetTypeList = "Mutual fund"
)

type AssetTypeReqBody struct {
	Name string `json:"name"`
}

type AssetType struct {
	Name string `json:"name"`
}

type AssetItemReqBody struct {
	Name          string  `db:"name" binding:"required"`
	AssetType     string  `db:"assetType" binding:"required"`
	Code          string  `db:"code"`
	AvgBuyValue   float64 `db:"avgBuyValue"`
	CurrentValue  float64 `db:"currentValue"`
	TotalUnits    float64 `db:"totalUnits"`
	PctIncPerYear float64 `db:"pctIncPerYear"`
}

type AssetItem struct {
	Id            string  `db:"id" binding:"required"`
	Name          string  `db:"name" binding:"required"`
	AssetType     string  `db:"assetType" binding:"required"`
	Code          string  `db:"code"`
	AvgBuyValue   float64 `db:"avgBuyValue"`
	CurrentValue  float64 `db:"currentValue"`
	TotalUnits    float64 `db:"totalUnits"`
	PctIncPerYear float64 `db:"pctIncPerYear"`
	IamId         string  `db:"iamId"`
}
