package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

type APIResponse struct {
	Data []NavData `json:"data"` // use appropriate type for your data
}

type NavData struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

func FetchStockLatestData(code string) {
	websiteUrl := fmt.Sprintf("https://www.google.com/finance/quote/%s:NSE", code)
	resp, err := http.Get(websiteUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// 2. Create a goquery document from the response body.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// 3. Use goquery (CSS-style selectors) to find specific elements.
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Printf("Paragraph %d: %s\n", i, text)
	})
}

func FetchMFLatestNavData(code string) (float64, error) {
	url := fmt.Sprintf("https://api.mfapi.in/mf/%s/latest", code)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	// Set headers if needed
	req.Header.Set("Accept", "application/json")
	// For example, if you require an authorization token:
	// req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0, err
	}

	var apiResp APIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return 0, err
	}

	i, err := strconv.ParseFloat(apiResp.Data[0].Nav, 64)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, err
	}

	return i, nil
}

func InsertAssetType(assetTypeReqBody types.AssetTypeReqBody) error {
	query := fmt.Sprintf(`INSERT INTO %s (name) VALUES ($1)`, database.ASSET_TYPE_TABLE_NAME)
	_, err := database.Execute(query, assetTypeReqBody.Name)
	if err != nil {
		return err
	}

	return nil
}

func GetAssetTypes() ([]string, error) {
	// Build the SELECT query
	query := fmt.Sprintf(`SELECT name FROM %s`, database.ASSET_TYPE_TABLE_NAME)

	// Execute the query
	rows, err := database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	// Ensure the rows are closed once we're done
	defer rows.Close()

	var assetTypes []string

	// Iterate through the rows
	for rows.Next() {
		var assetType string
		// Adjust the columns you scan based on your actual table schema
		if err := rows.Scan(&assetType); err != nil {
			return nil, fmt.Errorf("error scanning asset type: %w", err)
		}

		assetTypes = append(assetTypes, assetType)
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	// Return the slice of asset types and no error
	return assetTypes, nil
}

func InsertAssetItem(userId string, assetItemReq types.AssetItemReqBody) error {
	// Adjust this constant or use the actual table name directly
	const assetItemTable = "AssetItem"

	// Build your INSERT statement. Note how we match columns to the fields in AssetItemReqBody.
	query := fmt.Sprintf(`
		INSERT INTO %s (
			name,
			assetType,
			code,
			avgBuyValue,
			currentValue,
			totalUnits,
			pctIncPerYear,
			iamId
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, assetItemTable)

	encyptedAvgBuyValue, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.AvgBuyValue))
	if err != nil {
		return err
	}
	encyptedCurrentValue, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.CurrentValue))
	if err != nil {
		return err
	}
	encyptedTotalUnits, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.TotalUnits))
	if err != nil {
		return err
	}
	encyptedPctIncPerYear, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.PctIncPerYear))
	if err != nil {
		return err
	}

	_, err = database.Execute(
		query,
		assetItemReq.Name,
		assetItemReq.AssetType,
		assetItemReq.Code,
		encyptedAvgBuyValue,
		encyptedCurrentValue,
		encyptedTotalUnits,
		encyptedPctIncPerYear,
		userId,
	)
	if err != nil {
		log.Printf("Error inserting asset item: %v", err)
		return err
	}

	return nil
}

func GetAssetItems(userId string) ([]types.AssetItem, error) {
	// Replace this with the table name where you're storing AssetItem rows
	const assetItemTable = "AssetItem"

	// Build the SELECT query. Adjust the columns to match your table schema as needed.
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			assetType,
			code,
			avgBuyValue,
			currentValue,
			totalUnits,
			pctIncPerYear,
			iamId
		FROM %s where iamId = $1
	`, assetItemTable)

	// Execute the query (adjust to your DB helper or global DB instance)
	rows, err := database.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var assetItems []types.AssetItem

	// Iterate through the result set
	for rows.Next() {
		var item types.AssetItem
		var encryptedAvgBuyValue, encryptedCurrentValue, encryptedTotalUnits, encryptedPctIncPerYear string

		// Scan the row into our item
		if err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.AssetType,
			&item.Code,
			&encryptedAvgBuyValue,
			&encryptedCurrentValue,
			&encryptedTotalUnits,
			&encryptedPctIncPerYear,
			&item.IamId,
		); err != nil {
			return nil, fmt.Errorf("error scanning asset item: %w", err)
		}
		decryptedAvgBuyValue, err := secrets.DecryptAESGCM(encryptedAvgBuyValue)
		if err != nil {
			return nil, err
		}
		item.AvgBuyValue, _ = strconv.ParseFloat(decryptedAvgBuyValue, 64)
		decryptedCurrentValue, err := secrets.DecryptAESGCM(encryptedCurrentValue)
		if err != nil {
			return nil, err
		}
		item.CurrentValue, _ = strconv.ParseFloat(decryptedCurrentValue, 64)
		decryptedTotalUnits, err := secrets.DecryptAESGCM(encryptedTotalUnits)
		if err != nil {
			return nil, err
		}
		item.TotalUnits, _ = strconv.ParseFloat(decryptedTotalUnits, 64)
		decryptedPctIncPerYear, err := secrets.DecryptAESGCM(encryptedPctIncPerYear)
		if err != nil {
			return nil, err
		}
		item.PctIncPerYear, _ = strconv.ParseFloat(decryptedPctIncPerYear, 64)

		assetItems = append(assetItems, item)
	}

	// Check for any row iteration error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	// Return the slice of AssetItem and no error
	return assetItems, nil
}

func UpdateAssetItem(userId string, assetId string, assetItemReq types.AssetItemReqBody) error {
	const assetItemTable = "AssetItem"

	// Build the UPDATE query, mapping each column to a placeholder
	query := fmt.Sprintf(`
			UPDATE %s
			SET
				name          = $1,
				assetType     = $2,
				code          = $3,
				avgBuyValue   = $4,
				currentValue  = $5,
				totalUnits    = $6,
				pctIncPerYear = $7
			WHERE
					iamId = $8 AND id = $9
	`, assetItemTable)

	encyptedAvgBuyValue, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.AvgBuyValue))
	if err != nil {
		return err
	}
	encyptedCurrentValue, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.CurrentValue))
	if err != nil {
		return err
	}
	encyptedTotalUnits, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.TotalUnits))
	if err != nil {
		return err
	}
	encyptedPctIncPerYear, err := secrets.EncryptAESGCM(FloatToString(assetItemReq.PctIncPerYear))
	if err != nil {
		return err
	}

	// Execute the UPDATE statement with the necessary parameters
	res, err := database.Execute(
		query,
		assetItemReq.Name,
		assetItemReq.AssetType,
		assetItemReq.Code,
		encyptedAvgBuyValue,
		encyptedCurrentValue,
		encyptedTotalUnits,
		encyptedPctIncPerYear,
		userId,
		assetId,
	)
	if err != nil {
		log.Printf("Error updating asset item: %v\n", err)
		return err
	}

	// Optional: Check how many rows were affected, in case you want to confirm successful update
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated; item with id=%s and iamId=%s not found", assetId, userId)
	}

	return nil
}

func DeleteAssetItem(userId, assetId string) error {
	const assetItemTable = "AssetItem"

	// Build the DELETE query. We constrain the delete by iamId and id.
	query := fmt.Sprintf(`
			DELETE FROM %s
			WHERE iamId = $1 AND id = $2
	`, assetItemTable)

	// Execute the DELETE.
	// If "id" is an integer, pass assetIDInt instead of assetId.
	res, err := database.Execute(query, userId, assetId)
	if err != nil {
		log.Printf("Error deleting asset item: %v\n", err)
		return err
	}

	// Optionally check how many rows were affected.
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected error: %v\n", err)
		return err
	}
	if rowsAffected == 0 {
		// This means no records matched the iamId + id.
		return fmt.Errorf("no asset found for userId=%s and assetId=%s", userId, assetId)
	}

	return nil
}
