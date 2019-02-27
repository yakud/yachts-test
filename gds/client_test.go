package gds

import (
	"testing"
	"time"
)

func client() *Client {
	return NewClient(&ClientConfig{
		Entrypoint: "http://ws.nausys.com/",
		Login:      "rest83@TTTTT",
		Password:   "Rest59Tb",
	})
}

func TestClient_CharterBaseList(t *testing.T) {
	list, err := client().CharterBaseList()
	if err != nil {
		t.Error(err)
	}

	if len(list.Bases) == 0 {
		t.Errorf("Empty bases list")
	}
}

func TestClient_CharterCompanies(t *testing.T) {
	list, err := client().CharterCompanies()
	if err != nil {
		t.Error(err)
	}

	if len(list.Companies) == 0 {
		t.Errorf("Empty companies list")
	}
}

func TestClient_YachtBuilders(t *testing.T) {
	list, err := client().YachtBuilders()
	if err != nil {
		t.Error(err)
	}

	if len(list.Builders) == 0 {
		t.Errorf("Empty builders list")
	}
}

func TestClient_Yachts(t *testing.T) {
	companiesList, err := client().CharterCompanies()
	if err != nil {
		t.Error(err)
	}

	if len(companiesList.Companies) == 0 {
		t.Errorf("Empty companies list")
	}

	list, err := client().Yachts(companiesList.Companies[0].Id)
	if err != nil {
		t.Error(err)
	}

	if len(list.Yachts) == 0 {
		t.Errorf("Empty yachts list")
	}
}

func TestClient_YachtsFail(t *testing.T) {
	_, err := client().Yachts(0)
	if err == nil {
		t.Error("Should be error")
	}
}

func TestClient_YachtModels(t *testing.T) {
	list, err := client().YachtModels()
	if err != nil {
		t.Error(err)
	}

	if len(list.Models) == 0 {
		t.Errorf("Empty models list")
	}
}

func TestClient_YachtReservation(t *testing.T) {
	list, err := client().YachtReservation(&RestYachtReservationsRequest{
		Credentials: client().Credentials(),
		PeriodFrom:  time.Now().Format("02.01.2006"),
		PeriodTo:    time.Now().Format("02.01.2006"),
	})
	if err != nil {
		t.Error(err)
	}

	if list.ErrorCode != 0 {
		t.Errorf("Request fail. Error code: %d", list.ErrorCode)
	}
}

func TestClient_YachtReservationFail(t *testing.T) {
	_, err := client().YachtReservation(&RestYachtReservationsRequest{
		Credentials: client().Credentials(),
		PeriodFrom:  time.Now().Format("02.01"),
		PeriodTo:    time.Now().Format("02.01.2006"),
	})
	if err == nil {
		t.Error("Should be error")
	}
}

func TestClient_YachtReservationFail2(t *testing.T) {
	_, err := client().YachtReservation(nil)
	if err == nil {
		t.Error("Should be error")
	}
}
