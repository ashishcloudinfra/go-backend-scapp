package types

type InventoryItemDescriptionRequestBody struct {
	Name         string `json:"name"`
	Img          string `json:"img"`
	Type         string `json:"type"`
	Instructions string `json:"instructions"`
}

type AddItemRequestBody struct {
	Status      string `json:"status"`
	EquipmentId string `json:"equipmentId"`
}

type UpdateItemRequestBody struct {
	PrevStatus  string `json:"prevStatus"`
	NewStatus   string `json:"newStatus"`
	EquipmentId string `json:"equipmentId"`
	Count       int    `json:"count"`
}

type InventoryItemDescription struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Img          string `json:"img"`
	Type         string `json:"type"`
	Instructions string `json:"instructions"`
	Count        int    `json:"count"`
}

type Equipment struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Img          string `json:"img"`
	Type         string `json:"type"`
	Instructions string `json:"instructions"`
}
