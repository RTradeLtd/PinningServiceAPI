package api

import "github.com/gin-gonic/gin"

// API implements the ipfs pinning service api
type API struct {
	router *gin.Engine
}

// NewAPI instantiates a new instance of the pinning service API
func NewAPI( /*TODO(postables): take in database*/ ) *API {
	api := &API{}
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
