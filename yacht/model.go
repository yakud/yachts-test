package yacht

import "time"

type Model struct {
	Id              uint64    `json:"id"`
	BuilderName     string    `json:"builder_name"`
	ModelName       string    `json:"model_name"`
	OwnerName       string    `json:"owner_name"`
	ReservationFrom time.Time `json:"reservation_from"`
	ReservationTo   time.Time `json:"reservation_to"`
}
