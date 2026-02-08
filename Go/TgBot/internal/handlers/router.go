package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterT struct {
	router *mux.Router
}

func NewRouter(profileHandler *ProfileHandlerT) *RouterT {
	router := mux.NewRouter()

	router.HandleFunc("/start", profileHandler.CreateProfile).Methods("POST")
	router.HandleFunc("/addCategory", profileHandler.AddCategory).Methods("POST")
	router.HandleFunc("/categories", profileHandler.GetAllCategories).Methods("POST")
	router.HandleFunc("/category/delete", profileHandler.DeleteCategory).Methods("POST")
	router.HandleFunc("/add", profileHandler.AddExpense).Methods("POST")
	router.HandleFunc("/today", profileHandler.TodayExpense).Methods("POST")
	router.HandleFunc("/week", profileHandler.WeekExpense).Methods("POST")
	/*router.HandleFunc("/month", profileHandler.).Methods("POST")
	router.HandleFunc("/stats", profileHandler.).Methods("POST")*/

	return &RouterT{router: router}
}

func (s *RouterT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
