package gds

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/yakud/gramework"

	"github.com/valyala/fasthttp"
)

const ContentTypeRequest = "application/json;charset=UTF-8"
const StatusOK = "OK"

type methodPath string

const (
	PathCharterBases     methodPath = "/CBMS-external/rest/catalogue/v6/charterBases"
	PathCharterCompanies methodPath = "/CBMS-external/rest/catalogue/v6/charterCompanies"
	PathYachts           methodPath = "/CBMS-external/rest/catalogue/v6/yachts/%d"
	PathYachtModels      methodPath = "/CBMS-external/rest/catalogue/v6/yachtModels"
	PathYachtBuilders    methodPath = "/CBMS-external/rest/catalogue/v6/yachtBuilders"
	PathYachtReservation methodPath = "/CBMS-external/rest/yachtReservation/v6/reservations"
	PathFreeYachts       methodPath = "/CBMS-external/rest/yachtReservation/v6/freeYachts"
)

var defaultClient fasthttp.Client

type Api struct {
	config *ApiConfig
	client *fasthttp.Client
}

func (t *Api) CharterBaseList() (*RestCharterBaseList, error) {
	code, body, err := t.Do(PathCharterBases, nil)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestCharterBaseList{
		Bases: make([]RestCharterBase, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s", resp.Status)
	}

	return resp, nil
}

func (t *Api) CharterCompanies() (*RestCharterCompanyList, error) {
	code, body, err := t.Do(PathCharterCompanies, nil)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestCharterCompanyList{
		Companies: make([]RestCharterCompany, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s", resp.Status)
	}

	return resp, nil
}

func (t *Api) YachtBuilders() (*RestYachtBuilderList, error) {
	code, body, err := t.Do(PathYachtBuilders, nil)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestYachtBuilderList{
		Builders: make([]RestYachtBuilder, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s", resp.Status)
	}

	return resp, nil
}

func (t *Api) Yachts(charterCompanyId uint64) (*RestYachtList, error) {
	path := methodPath(fmt.Sprintf(string(PathYachts), charterCompanyId))
	code, body, err := t.Do(path, nil)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestYachtList{
		Yachts: make([]RestYacht, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s", resp.Status)
	}

	return resp, nil
}

func (t *Api) YachtModels() (*RestYachtModelList, error) {
	code, body, err := t.Do(PathYachtModels, nil)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestYachtModelList{
		Models: make([]RestYachtModel, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s", resp.Status)
	}

	return resp, nil
}

func (t *Api) YachtReservation(req *RestYachtReservationsRequest) (*RestYachtReservationList, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	req.Credentials = t.Credentials()

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	code, body, err := t.Do(PathYachtReservation, reqJson)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestYachtReservationList{
		Reservations: make([]RestYachtReservation, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s %-v", resp.Status, resp)
	}

	return resp, nil
}

func (t *Api) FreeYachts(req *RestFreeYachtsRequest) (*RestFreeYachtList, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	req.Credentials = t.Credentials()

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	code, body, err := t.Do(PathFreeYachts, reqJson)
	if err != nil {
		return nil, err
	}

	if code != fasthttp.StatusOK {
		return nil, errors.Errorf("Bad status code: %d", code)
	}

	resp := &RestFreeYachtList{
		FreeYachts: make([]RestFreeYacht, 0),
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp.Status != StatusOK {
		return nil, errors.Errorf("Bad status code: %s %-v", resp.Status, resp)
	}

	return resp, nil
}

// Do make request to black box with params
func (t *Api) Do(path methodPath, body []byte) (int, []byte, error) {
	req := fasthttp.AcquireRequest()

	uri := req.URI()
	uri.Update(t.config.Entrypoint)
	uri.SetPath(string(path))

	req.Header.SetMethod(gramework.POST)
	req.Header.SetContentType(ContentTypeRequest)

	if body == nil || len(body) == 0 {
		var err error
		body, err = t.makeAuthCredentials()
		if err != nil {
			return 0, nil, err
		}
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()

	var err error
	if err = t.client.Do(req, resp); err != nil {
		err = fmt.Errorf("client request %s error: %s", uri, err.Error())
	}

	code, body := resp.StatusCode(), resp.Body()

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return code, body, err
}

func (t *Api) makeAuthCredentials() ([]byte, error) {
	// @TODO: cache it
	return json.Marshal(t.Credentials())
}

func (t *Api) Credentials() RestAuthentication {
	return RestAuthentication{
		Username: t.config.Login,
		Password: t.config.Password,
	}
}

func NewApi(config *ApiConfig) *Api {
	return &Api{
		config: config,
		client: &defaultClient,
	}
}
