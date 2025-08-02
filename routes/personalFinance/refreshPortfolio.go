package personalfinance

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/routes/scripts"
	"server.simplifycontrol.com/types"
)

func RefreshPortfolioUtil(userId string) error {
	assetItems, err := helpers.GetAssetItems(userId)
	if err != nil {
		return err
	}

	for _, assetItem := range assetItems {
		if assetItem.AssetType == "Mutual fund" {
			nav, err := helpers.FetchMFLatestNavData(assetItem.Code)
			if err != nil {
				return err
			}

			if assetItem.CurrentValue != nav {
				err := helpers.UpdateAssetItem(assetItem.IamId, assetItem.Id, types.AssetItemReqBody{
					Name:          assetItem.Name,
					AssetType:     assetItem.AssetType,
					Code:          assetItem.Code,
					PctIncPerYear: assetItem.PctIncPerYear,
					AvgBuyValue:   assetItem.AvgBuyValue,
					CurrentValue:  nav,
					TotalUnits:    assetItem.TotalUnits,
				})
				if err != nil {
					return err
				}
			}
		}

		if assetItem.AssetType == "Stock" || assetItem.AssetType == "ETF" {
			price, err := scripts.GetStockValue(assetItem.Code)
			if err != nil {
				return err
			}

			if assetItem.CurrentValue != price {
				err := helpers.UpdateAssetItem(assetItem.IamId, assetItem.Id, types.AssetItemReqBody{
					Name:          assetItem.Name,
					AssetType:     assetItem.AssetType,
					Code:          assetItem.Code,
					PctIncPerYear: assetItem.PctIncPerYear,
					AvgBuyValue:   assetItem.AvgBuyValue,
					CurrentValue:  price,
					TotalUnits:    assetItem.TotalUnits,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func RefreshPortfolio(w http.ResponseWriter, r *http.Request) {
	// Ensure it’s a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	err := RefreshPortfolioUtil(userId)
	if err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Error refreshing portfolio", http.StatusBadRequest)
		return
	}
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": "success"})
}
