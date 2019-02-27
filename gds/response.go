package gds

type RestStatus struct {
	Status    string `json:"status"`
	ErrorCode int    `json:"errorCode,omitempty"`
}

type RestPeriodRange struct {
	PeriodFrom string `json:"periodFrom"`
	PeriodTo   string `json:"periodTo"`
}

type RestCharterBaseList struct {
	RestStatus
	Bases []RestCharterBase `json:"bases"`
}

type RestCharterBase struct {
	CompanyId uint64 `json:"companyId"`
}

type RestCharterCompanyList struct {
	RestStatus
	Companies []RestCharterCompany `json:"companies"`
}

type RestCharterCompany struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type RestYachtList struct {
	RestStatus
	Yachts []RestYacht `json:"yachts"`
}

type RestYacht struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	CompanyId    uint64 `json:"companyId"`
	YachtModelId uint64 `json:"yachtModelId"`
}

type RestYachtModelList struct {
	RestStatus
	Models []RestYachtModel `json:"models"`
}

type RestYachtModel struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	YachtBuilderId uint64 `json:"yachtBuilderId"`
}

type RestYachtBuilderList struct {
	RestStatus
	Builders []RestYachtBuilder `json:"builders"`
}

type RestYachtBuilder struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type RestYachtReservationList struct {
	RestStatus
	Reservations []RestYachtReservation `json:"reservations"`
}

type RestYachtReservation struct {
	Id                uint64 `json:"id"`
	YachtId           uint64 `json:"yachtId"`
	ReservationStatus string `json:"reservationStatus"`

	RestPeriodRange
}
