package handlers

import "github.com/heronhoga/simplearn/backend/services"

type DocumentHandler struct {
	service *services.DocumentService
}

func NewDocumentHandler(service *services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		service: service,
	}
}
