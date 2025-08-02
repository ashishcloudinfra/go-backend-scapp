package budget

import "github.com/gorilla/mux"

func RegisterBudgetRoutes(mux *mux.Router) {
	mux.HandleFunc("/{userId}/budgetItem", AddBudgetItem).Methods("POST")
	mux.HandleFunc("/{userId}/budgetCategory", AddBudgetCategory).Methods("POST")
	mux.HandleFunc("/{userId}/copyBudget", CopyBudgetOfGivenMonthAndYear).Methods("POST")
	mux.HandleFunc("/budgetCategoryType", AddBudgetCategoryType).Methods("POST")
	mux.HandleFunc("/{userId}/budgetItems/m/{month}/y/{year}", GetBudgetItems).Methods("GET")
	mux.HandleFunc("/{userId}/budgetCategories/m/{month}/y/{year}", GetBudgetCategories).Methods("GET")
	mux.HandleFunc("/budgetCategoryTypes", GetBudgetCategoryTypes).Methods("GET")
	mux.HandleFunc("/{userId}/budgetStats", GetBudgetStats).Methods("GET")
	mux.HandleFunc("/{userId}/rawStats", GetRawStats).Methods("GET")
	mux.HandleFunc("/budgetItem/{itemId}", UpdateBudgetItem).Methods("PUT")
	mux.HandleFunc("/budgetCategory/{categoryId}", UpdateBudgetCategory).Methods("PUT")
	mux.HandleFunc("/budgetItem/{itemId}", DeleteBudgetItem).Methods("DELETE")
	mux.HandleFunc("/budgetCategory/{categoryId}", DeleteBudgetCategory).Methods("DELETE")
}
