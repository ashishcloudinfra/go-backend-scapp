package scripts

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
)

func GetStockValue(tickerSymbol string) (float64, error) {
	wd, _ := os.Getwd()
	var scriptName, interpreter string

	if os.Getenv("ENV") == "local" {
		scriptName, interpreter = "runScript.sh", "bash"
	} else {
		scriptName, interpreter = "stockPrice.py", "python3"
	}

	scriptPath := filepath.Join(wd, "scripts", "stocks", scriptName)
	cmd := exec.Command(interpreter, scriptPath, tickerSymbol)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error executing script: %w", err)
	}

	outputLines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var result float64
	for _, line := range outputLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if value, err := strconv.ParseFloat(line, 64); err == nil {
			result = value
			return result, nil
		}
	}
	return 0, fmt.Errorf("error converting output to float: no valid numeric output found")
}

func GetStockPriceData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tickerSymbol := vars["tickerSymbol"]

	if tickerSymbol == "" {
		helpers.SendJSONError(w, "tickerSymbol is required", http.StatusBadRequest)
		helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": 0})
		return
	}

	i, err := GetStockValue(tickerSymbol)
	if err != nil {
		helpers.SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": i})
}
