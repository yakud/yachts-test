package gds

type RestAuthentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RestYachtReservationsRequest struct {
	Credentials RestAuthentication `json:"credentials"`
	PeriodFrom  string             `json:"periodFrom"`
	PeriodTo    string             `json:"periodTo"`
}
