package api

import (
	"fmt"
	"net/http"

	"github.com/tc1io/clean-arch-demo/ent"
	"github.com/tc1io/clean-arch-demo/uc"

	"github.com/julienschmidt/httprouter"
)

func MakePlaceOrder(placeOrder uc.SubmitOrderUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// turn payload into order...
		order := ent.Order{}
		trackingId, err := placeOrder(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Tracking ID: %s", trackingId)
	}
}
