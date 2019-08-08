package api

import (
	"github.com/RTradeLtd/rtfs/v2"
	"github.com/gin-gonic/gin"
)

// API implements the ipfs pinning service api
type API struct {
	router *gin.Engine
	ipfs rtfs.Manager
}

// NewAPI instantiates a new instance of the pinning service API
func NewAPI(manager rtfs.Manager) *API {
	api := &API{ipfs: manager}
	api.InitializeRouter()
	return api
}

// Run starts the api and access incoming connections
func (a *API) Run(address string) error {
	return a.router.Run(address)
}

// InitializeRouter initializes the gin router engine
func (a *API) InitializeRouter() {
	router := gin.Default()

	root := router.Group("/")
	pins := root.Group("pins")
	{
		pins.GET("", a.PinsGet)
		pins.POST("", a.PinsPost)
		pins.DELETE("/:cid", a.PinsCidDelete)
		pins.GET("/:cid", a.PinsCidGet)
		pins.POST("/:cid", a.PinsCidPost)
	}

	a.router = router
}
