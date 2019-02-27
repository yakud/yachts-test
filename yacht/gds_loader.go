package yacht

import (
	"fmt"
	"time"

	"github.com/yakud/yachts-test/gds"
)

type GDSLoader struct {
	api *gds.Api
}

func (t *GDSLoader) LoadTo(storage *StorageES) error {
	// Load companies
	companiesList, err := t.api.CharterCompanies()
	if err != nil {
		return err
	}

	// Load Builders
	fmt.Println("All builers list:")
	buildersResp, err := t.api.YachtBuilders()
	if err != nil {
		return err
	}

	builders := make(map[uint64]gds.RestYachtBuilder)
	for _, builder := range buildersResp.Builders {
		builders[builder.Id] = builder
		//fmt.Printf("%+v\n", builder)
	}

	// Load models
	fmt.Println("All models list:")
	modelsResp, err := t.api.YachtModels()
	if err != nil {
		return err
	}

	models := make(map[uint64]gds.RestYachtModel)
	for _, model := range modelsResp.Models {
		models[model.Id] = model
		//fmt.Printf("%+v\n", model)
	}

	// Load reservations
	fmt.Println("Yacht reservation:")
	reservationsResp, err := t.api.YachtReservation(&gds.RestYachtReservationsRequest{
		PeriodFrom: time.Now().Format("02.01.2006"),
		PeriodTo:   time.Now().Add(time.Hour * 24 * 7).Format("02.01.2006"),
		//Credentials: api.Credentials(),
	})
	if err != nil {
		return err
	}

	reservations := make(map[uint64]gds.RestYachtReservation)
	for _, reservation := range reservationsResp.Reservations {
		reservations[reservation.YachtId] = reservation
		//fmt.Printf("%+v\n", reservation)
	}

	fmt.Println("All companies and yachts list:")
	yachts := make([]*Model, 0)
	for _, co := range companiesList.Companies {
		// Load yachts
		yachtsList, err := t.api.Yachts(co.Id)
		if err != nil {
			return err
		}

		for _, y := range yachtsList.Yachts {
			model, ok := models[y.YachtModelId]
			if !ok {
				model = gds.RestYachtModel{Name: "UNDEFINED"}
			}

			builder, ok := builders[model.YachtBuilderId]
			if !ok {
				builder = gds.RestYachtBuilder{Name: "UNDEFINED"}
			}

			yachtModel := &Model{
				Id:          y.Id,
				BuilderName: builder.Name,
				ModelName:   model.Name,
				OwnerName:   co.Name,
			}

			reservation, ok := reservations[y.Id]
			if ok {
				yachtModel.ReservationFrom, err = time.Parse("02.01.2006 15:04", reservation.PeriodFrom)
				if err != nil {
					return err
				}
				yachtModel.ReservationTo, err = time.Parse("02.01.2006 15:04", reservation.PeriodTo)
				if err != nil {
					return err
				}
			}

			yachts = append(yachts, yachtModel)

			if err := storage.Add(yachtModel); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewGDSLoader(gdsApi *gds.Api) *GDSLoader {
	return &GDSLoader{
		api: gdsApi,
	}
}
