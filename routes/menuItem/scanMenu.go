package menuItem

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sashabaranov/go-openai"
	"server.simplifycontrol.com/firebase"
	"server.simplifycontrol.com/helpers"
)

func ScanMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
		return
	}

	// Limit upload size (e.g., 10MB)
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		fmt.Println("ParseMultipartForm error:", err)
		helpers.SendJSONError(w, "Error while parsing multi part form", http.StatusBadRequest)
		return
	}

	// Check if the form contains files
	if r.MultipartForm == nil || r.MultipartForm.File == nil {
		helpers.SendJSONError(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	// Retrieve multiple files
	files, ok := r.MultipartForm.File["files"]
	if !ok || len(files) == 0 {
		helpers.SendJSONError(w, "No files uploaded or not ok", http.StatusBadRequest)
		return
	}

	menuImageLinks, err := firebase.UploadToStorageAndGetPublicLinks(files)
	if err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Error while uploading files", http.StatusBadRequest)
		return
	}

	var formattedTexts []string
	for _, menuImageLink := range menuImageLinks {
		imgData, err := helpers.DetectTextURI(menuImageLink)
		if err != nil {
			fmt.Println(err)
			helpers.SendJSONError(w, "Error in scanning ocr", http.StatusBadRequest)
			return
		}
		data, err := helpers.GetGPTResponse("gpt-4o-mini", []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are an AI designed to extract menu items from a given text. Your task is to identify and organize menu categories, item names, and their corresponding prices. Follow these rules: Identify distinct categories such as Main Course, Appetizers, Beverages, Desserts, etc. Extract each menu item along with its price. Format the output in a structured list with categories, items, and prices. Maintain the original order and wording of items. Ignore any non-menu-related text."},
			{Role: "user", Content: fmt.Sprintf("Give me a list of menu Items along with prices from %s", imgData)},
		})
		if err != nil {
			fmt.Println(err)
			helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
			return
		}
		formattedTexts = append(formattedTexts, data)
	}

	var results []interface{}
	for _, formattedText := range formattedTexts {
		data, err := helpers.GetGPTResponse("gpt-4o-mini", []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are an text parser assistant that analyzes the contents of a text of menu items and responds in JSON array format of given type."},
			{Role: "user", Content: fmt.Sprintf("%s and give an JSON array of type { name: string;\n description: string;\n photo?: string;\n isVeg: boolean; // veg, non-veg \n category: string; // heading or such as Indian, Italian, Starters, Main course, etc. \n cookingTime: string; \n varieties: { \n name: string; // half, full, default, etc. \n price: string; // 190 \n }[]; \n}", formattedText)},
		})
		if err != nil {
			fmt.Println(err)
			helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
			return
		}
		results = append(results, data)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": results})
}
