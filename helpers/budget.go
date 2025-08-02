package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func TransformData(stats []types.BudgetStatsByMonthAndYear) []types.BudgetCategoryTypeStats {
	type key struct {
		Type  string
		Month int
		Year  int
	}
	results := make(map[key]*types.BudgetCategoryTypeStats)
	categorySets := make(map[key]map[string]struct{})

	for _, stat := range stats {
		k := key{Type: stat.Type, Month: stat.Month, Year: stat.Year}
		if _, exists := results[k]; !exists {
			results[k] = &types.BudgetCategoryTypeStats{
				CategoryType: stat.Type,
				BgColor:      stat.BgColor,
				TextColor:    stat.TextColor,
				Month:        stat.Month,
				Year:         stat.Year,
				ItemCount:    0,
				Categories:   "",
			}
			categorySets[k] = make(map[string]struct{})
		}

		results[k].ItemCount++
		categorySets[k][stat.CategoryName] = struct{}{}

		results[k].TotalActualAmount += stat.ActualAmount
	}

	var finalResults []types.BudgetCategoryTypeStats
	for k, result := range results {
		categories := make([]string, 0, len(categorySets[k]))
		for category := range categorySets[k] {
			categories = append(categories, category)
		}
		result.Categories = strings.Join(categories, ", ")
		finalResults = append(finalResults, *result)
	}

	return finalResults
}

// GetBudgetCategories retrieves budget categories for a user in a specific month and year.
func GetBudgetCategories(userId string, month string, year string) ([]types.BudgetCategory, error) {
	query := fmt.Sprintf(
		`SELECT id, categoryName, categoryDescription, month, year, parentId, categoryTypeId
		 FROM %s 
		 WHERE iamId = $1 AND month = $2 AND year = $3`,
		database.BUDGET_CATEGORY_TABLE_NAME,
	)

	rows, err := database.Query(query, userId, month, year)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var categories []types.BudgetCategory

	for rows.Next() {
		var category types.BudgetCategory
		var parentID sql.NullString

		err := rows.Scan(
			&category.Id,
			&category.CategoryName,
			&category.CategoryDescription,
			&category.Month,
			&category.Year,
			&parentID,
			&category.CategoryTypeId,
		)
		if err != nil {
			log.Printf("Error scanning data: %v", err)
			return nil, err
		}

		// Convert sql.NullString to the struct field
		if parentID.Valid {
			category.ParentId = parentID.String
		} else {
			category.ParentId = ""
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading rows: %v", err)
		return nil, err
	}

	return categories, nil
}

// InsertBudgetCategory inserts a new budget category record into the database.
func InsertBudgetCategory(budgetCategoryReq types.BudgetCategoryReqBody, userId string) (string, error) {
	// Prepare the parentId for insertion
	var parentId interface{}
	if budgetCategoryReq.ParentId == "" {
		parentId = nil
	} else {
		parentId = budgetCategoryReq.ParentId
	}

	// Build the query
	query := fmt.Sprintf(`
			INSERT INTO %s 
			(categoryName, categoryDescription, month, year, parentId, categoryTypeId, iamId) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) Returning id
	`, database.BUDGET_CATEGORY_TABLE_NAME)

	// Execute the query
	rows, err := database.Query(query,
		budgetCategoryReq.CategoryName,
		budgetCategoryReq.CategoryDescription,
		budgetCategoryReq.Month,
		budgetCategoryReq.Year,
		parentId,
		budgetCategoryReq.CategoryTypeId,
		userId,
	)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		return "", err
	}

	var categories []string
	for rows.Next() {
		var category string

		err := rows.Scan(
			&category,
		)
		if err != nil {
			log.Printf("Error scanning data: %v", err)
			return "", err
		}

		categories = append(categories, category)
	}

	return categories[0], nil
}

// GetBudgetItems retrieves budget items for a given user, month, and year.
func GetBudgetItems(userId string, month string, year string) ([]types.BudgetItem, error) {
	query := fmt.Sprintf(
		`SELECT
			id,
			categoryId,
			itemName,
			description,
			allocatedAmount,
			actualAmount,
			currencyCode,
			status,
			month,
			year
		FROM %s
		WHERE iamId = $1 AND month = $2 AND year = $3`,
		database.BUDGET_ITEM_TABLE_NAME,
	)

	rows, err := database.Query(query, userId, month, year)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var items []types.BudgetItem
	for rows.Next() {
		var item types.BudgetItem
		var encryptedAllocatedAmount, encryptedActualAmount string

		if scanErr := rows.Scan(
			&item.Id,
			&item.CategoryID,
			&item.ItemName,
			&item.Description,
			&encryptedAllocatedAmount,
			&encryptedActualAmount,
			&item.CurrencyCode,
			&item.Status,
			&item.Month,
			&item.Year,
		); scanErr != nil {
			log.Printf("Error scanning data: %v", scanErr)
			return nil, scanErr
		}

		decryptedAllocated, err := secrets.DecryptAESGCM(encryptedAllocatedAmount)
		if err != nil {
			log.Printf("Error decrypting allocated amount: %v", err)
			return nil, err
		}
		item.AllocatedAmount, _ = strconv.ParseFloat(decryptedAllocated, 64)

		decryptedActual, err := secrets.DecryptAESGCM(encryptedActualAmount)
		if err != nil {
			log.Printf("Error decrypting actual amount: %v", err)
			return nil, err
		}
		item.ActualAmount, _ = strconv.ParseFloat(decryptedActual, 64)

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading rows: %v", err)
		return nil, err
	}

	return items, nil
}

// GetBudgetItems retrieves budget items for a given user, month, and year.
func GetBudgetItemsForGivenCategoryId(userId string, month string, year string, categoryId string) ([]types.BudgetItem, error) {
	query := fmt.Sprintf(
		`SELECT
			id,
			categoryId,
			itemName,
			description,
			allocatedAmount,
			actualAmount,
			currencyCode,
			status,
			month,
			year
		FROM %s
		WHERE iamId = $1 AND month = $2 AND year = $3 and categoryId = $4`,
		database.BUDGET_ITEM_TABLE_NAME,
	)

	rows, err := database.Query(query, userId, month, year, categoryId)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var items []types.BudgetItem
	for rows.Next() {
		var item types.BudgetItem
		var encryptedAllocatedAmount, encryptedActualAmount string
		if scanErr := rows.Scan(
			&item.Id,
			&item.CategoryID,
			&item.ItemName,
			&item.Description,
			&encryptedAllocatedAmount,
			&encryptedActualAmount,
			&item.CurrencyCode,
			&item.Status,
			&item.Month,
			&item.Year,
		); scanErr != nil {
			log.Printf("Error scanning data: %v", scanErr)
			return nil, scanErr
		}

		decryptedAllocated, err := secrets.DecryptAESGCM(encryptedAllocatedAmount)
		if err != nil {
			log.Printf("Error decrypting allocated amount: %v", err)
			return nil, err
		}
		item.AllocatedAmount, _ = strconv.ParseFloat(decryptedAllocated, 64)

		decryptedActual, err := secrets.DecryptAESGCM(encryptedActualAmount)
		if err != nil {
			log.Printf("Error decrypting actual amount: %v", err)
			return nil, err
		}
		item.ActualAmount, _ = strconv.ParseFloat(decryptedActual, 64)

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading rows: %v", err)
		return nil, err
	}

	return items, nil
}

// InsertBudgetItem inserts a new budget item into the database.
func InsertBudgetItem(budgetItemReq types.BudgetItemReqBody, userId string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			categoryId,
			itemName,
			description,
			allocatedAmount,
			actualAmount,
			currencyCode,
			status,
			iamId,
			month,
			year
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, database.BUDGET_ITEM_TABLE_NAME)

	encyptedAllocatedAmount, err := secrets.EncryptAESGCM(FloatToString(budgetItemReq.AllocatedAmount))
	if err != nil {
		return err
	}
	encyptedActualAmount, err := secrets.EncryptAESGCM(FloatToString(budgetItemReq.ActualAmount))
	if err != nil {
		return err
	}

	_, err = database.Execute(
		query,
		budgetItemReq.CategoryID,
		budgetItemReq.ItemName,
		budgetItemReq.Description,
		encyptedAllocatedAmount,
		encyptedActualAmount,
		budgetItemReq.CurrencyCode,
		budgetItemReq.Status,
		userId,
		budgetItemReq.Month,
		budgetItemReq.Year,
	)
	if err != nil {
		log.Printf("Error inserting budget item: %v", err)
		return err
	}

	return nil
}

func CopyBudgetCategoriesAndItems(userId string, copyBudgetReq types.CopyBudgetReqBody) error {
	// Get all categories from old month/year
	categories, err := GetBudgetCategories(
		userId,
		strconv.Itoa(copyBudgetReq.OldMonth),
		strconv.Itoa(copyBudgetReq.OldYear),
	)
	if err != nil {
		return fmt.Errorf("error querying budget categories: %w", err)
	}

	// For each category, insert into the new month/year, then copy its items
	for _, category := range categories {
		categoryId, err := InsertBudgetCategory(types.BudgetCategoryReqBody{
			CategoryName:        category.CategoryName,
			CategoryDescription: category.CategoryDescription,
			Month:               copyBudgetReq.CurrentMonth,
			Year:                copyBudgetReq.CurrentYear,
			CategoryTypeId:      category.CategoryTypeId,
			ParentId:            "",
		}, userId)
		if err != nil {
			return fmt.Errorf("error inserting budget category: %w", err)
		}

		items, err := GetBudgetItemsForGivenCategoryId(
			userId,
			strconv.Itoa(copyBudgetReq.OldMonth),
			strconv.Itoa(copyBudgetReq.OldYear),
			category.Id,
		)
		if err != nil {
			return fmt.Errorf("error querying budget items for category %s: %w", category.Id, err)
		}

		// Insert each item into the newly created category for the new month/year
		for _, item := range items {
			err := InsertBudgetItem(types.BudgetItemReqBody{
				CategoryID:      categoryId,
				ItemName:        item.ItemName,
				Description:     item.Description,
				AllocatedAmount: item.AllocatedAmount,
				ActualAmount:    item.ActualAmount,
				Month:           copyBudgetReq.CurrentMonth,
				Year:            copyBudgetReq.CurrentYear,
				Status:          item.Status,
				CurrencyCode:    item.CurrencyCode,
			}, userId)
			if err != nil {
				return fmt.Errorf("error inserting budget item: %w", err)
			}
		}
	}

	return nil
}

// GetBudgetCategoryTypes retrieves all budget category types from the database.
func GetBudgetCategoryTypes() ([]types.BudgetCategoryType, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, database.BUDGET_CATEGORY_TYPE_TABLE_NAME)
	rows, err := database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var categoryTypes []types.BudgetCategoryType
	for rows.Next() {
		var categoryType types.BudgetCategoryType
		if err := rows.Scan(
			&categoryType.Id,
			&categoryType.Type,
			&categoryType.BgColor,
			&categoryType.TextColor,
		); err != nil {
			return nil, fmt.Errorf("error scanning category type: %w", err)
		}

		categoryTypes = append(categoryTypes, categoryType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	return categoryTypes, nil
}

func DeleteBudgetCategoriesOfGivenMonthAndYear(month string, year string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE month = $1 and year = $2`, database.BUDGET_CATEGORY_TABLE_NAME)
	_, err := database.Execute(query, month, year)
	if err != nil {
		return fmt.Errorf("error deleting budget category: %w", err)
	}

	return nil
}

func GetBudgetStatsByMonthAndYear(userId string) ([]types.BudgetCategoryTypeStats, error) {
	query := fmt.Sprintf(`SELECT
    bItem.id AS itemId,
    bItem.itemName,
    bItem.actualAmount,
    bItem.status,
    bItem.month,
    bItem.year,
    bCat.categoryName,
    bCatType.type,
		bCatType.bgColor,
		bCatType.textColor
	FROM
    %s bItem
	JOIN
    %s bCat
    ON bItem.categoryId = bCat.id
	JOIN
    %s bCatType
    ON bCat.categoryTypeId = bCatType.id
	WHERE bItem.iamId = $1
	`, database.BUDGET_ITEM_TABLE_NAME, database.BUDGET_CATEGORY_TABLE_NAME, database.BUDGET_CATEGORY_TYPE_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting budget stats: %w", err)
	}
	defer rows.Close()

	var stats []types.BudgetStatsByMonthAndYear
	for rows.Next() {
		var stat types.BudgetStatsByMonthAndYear
		var encryptedActualAmount string
		if err := rows.Scan(
			&stat.ItemID,
			&stat.ItemName,
			&encryptedActualAmount,
			&stat.Status,
			&stat.Month,
			&stat.Year,
			&stat.CategoryName,
			&stat.Type,
			&stat.BgColor,
			&stat.TextColor,
		); err != nil {
			return nil, fmt.Errorf("error scanning category type: %w", err)
		}

		decryptedActual, err := secrets.DecryptAESGCM(encryptedActualAmount)
		if err != nil {
			log.Printf("Error decrypting actual amount: %v", err)
			return nil, err
		}
		stat.ActualAmount, _ = strconv.ParseFloat(decryptedActual, 64)

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	return TransformData(stats), nil
}

func GetRawUserStats(userId string) ([]types.RawUserStats, error) {
	query := fmt.Sprintf(`SELECT 
    bct.type,
		bc.categoryName,
		bi.itemName,
		bi.actualAmount,
		bi.month,
		bi.year
	FROM 
		%s bct
		INNER JOIN %s bc ON bct.id = bc.categoryTypeId
		INNER JOIN %s bi ON bc.id = bi.categoryId
	WHERE 
		bi.status = 'active' and bi.iamId = $1
	ORDER BY 
    bc.year, bc.month, bct.type, bc.categoryName;
	`, database.BUDGET_CATEGORY_TYPE_TABLE_NAME, database.BUDGET_CATEGORY_TABLE_NAME, database.BUDGET_ITEM_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error deleting budget category: %w", err)
	}
	defer rows.Close()

	var stats []types.RawUserStats
	for rows.Next() {
		var stat types.RawUserStats
		var encryptedActualAmount string
		if err := rows.Scan(
			&stat.Type,
			&stat.CategoryName,
			&stat.ItemName,
			&encryptedActualAmount,
			&stat.Month,
			&stat.Year,
		); err != nil {
			return nil, fmt.Errorf("error scanning category type: %w", err)
		}

		decryptedActual, err := secrets.DecryptAESGCM(encryptedActualAmount)
		if err != nil {
			log.Printf("Error decrypting actual amount: %v", err)
			return nil, err
		}
		stat.ActualAmount, _ = strconv.ParseFloat(decryptedActual, 64)

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	return stats, nil
}
