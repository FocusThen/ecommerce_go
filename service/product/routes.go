package product

import (
	"fmt"
	"net/http"

	"github.com/focusthen/ecommerce_go/types"
	"github.com/focusthen/ecommerce_go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProducts).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProducts(w http.ResponseWriter, r *http.Request){
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %+v", errors))
		return
	}

  err := h.store.CreateProduct(types.Product{
    Name: payload.Name,
    Description:  payload.Description,
    Image: payload.Image,
    Price: payload.Price,
    Quantity: payload.Quantity,
  })
  if err !=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
  }

  utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "product created"})
}
