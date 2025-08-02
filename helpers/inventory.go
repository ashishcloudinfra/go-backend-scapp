package helpers

import (
	"fmt"
	"log"

	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func GetInventoryByStatus(companyId string, status string) ([]types.InventoryItemDescription, error) {
	query := fmt.Sprintf(`
		SELECT 
			eq.id, 
			eq.name, 
			eq.img, 
			eq.type, 
			eq.instructions, 
			COUNT(Item.status) AS count
		FROM %s AS eq
		JOIN %s ON Item.equipmentId = eq.id
		WHERE Item.companyId = $1 AND Item.status = $2
		GROUP BY eq.id, eq.name, eq.img, eq.type, eq.instructions`,
		database.INVENTORY_EQUIPMENT_TABLE_NAME,
		database.INVENTORY_ITEM_TABLE_NAME,
	)

	rows, err := database.Query(query, companyId, status)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var result []types.InventoryItemDescription
	for rows.Next() {
		var item types.InventoryItemDescription
		err := rows.Scan(&item.Id, &item.Name, &item.Img, &item.Type, &item.Instructions, &item.Count)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}
