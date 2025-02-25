package controllers

import (
	"fmt"
	"net/http"
)

//NOTE: function names needs to be in PascalCase to be able to exported by default

// this function for root endpoint
func RootHandler(res http.ResponseWriter, req *http.Request)  {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(res, "Go world!\n")
}