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

type RestFreeYachtsRequest struct {
	Credentials RestAuthentication `json:"credentials"`
	PeriodFrom  string             `json:"periodFrom"`
	PeriodTo    string             `json:"periodTo"`
	Yachts      []uint64           `json:"yachts"`
}
