package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tc1io/clean-arch-demo/adapter"
	"github.com/tc1io/clean-arch-demo/uc"
)

func Register(oq adapter.OrderQueue, app *httprouter.Router) {
	submitOrderUseCase := uc.MakeSubmitOrderUseCase(oq)
	app.POST("/orders", MakePlaceOrder(submitOrderUseCase))
}
