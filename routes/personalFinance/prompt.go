package personalfinance

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sashabaranov/go-openai"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/types"
)

func Promt(w http.ResponseWriter, r *http.Request) {
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

	var reqBody types.PersonalFinancePromptReqBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	stats, err := helpers.GetRawUserStats(userId)
	if err != nil {
		log.Printf("GetBudgetCategoryTypes error: %v", err)
		helpers.SendJSONError(w, "Error fetching stats", http.StatusInternalServerError)
		return
	}

	assets, err := helpers.GetAssetItems(userId)
	if err != nil {
		log.Printf("GetAssetItems error: %v", err)
		helpers.SendJSONError(w, "Error fetching assets", http.StatusInternalServerError)
		return
	}

	const systemPrompt = "" +
		"Role and Scope:\n" +
		"You are an AI assistant specializing in personal finance, budgeting, and money management.\n" +
		"Your primary function is to help users plan their budgets, track expenses, and make informed\n" +
		"decisions about saving, spending, and investing.\n\n" +
		"Remember all data is in Indian rupees.\n\n" +
		"Tone and Communication Style:\n" +
		"Use clear, concise, and jargon-free language that is accessible to a wide range of users.\n" +
		"Keep your explanations and steps straightforward and actionable. Provide relevant examples\n" +
		"or hypothetical scenarios to illustrate your points whenever beneficial.\n\n" +
		"Use below data to answer user question:\n" +
		"%s"

	data, err := helpers.GetGPTResponse("gpt-4o-mini", []openai.ChatCompletionMessage{
		{Role: "system", Content: fmt.Sprintf(systemPrompt, map[string]interface{}{
			"budgetData": stats,
			"assetData":  assets,
		})},
		{Role: "user", Content: reqBody.Question},
	})
	if err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": data})
}
