package api

import "github.com/gin-gonic/gin"

// API implements the ipfs pinning service api
type API struct {
	router *gin.Engine
}

// NewAPI instantiates a new instance of the pinning service API
func NewAPI( /*TODO(postables): take in database*/ ) *API {
	return &API{router: NewRouter()}
}
