package bus

import (
	"sort"
	"time"
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"

	"github.com/gofiber/websocket/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		CreateBusEntry(data dto.CreateBusDto) (dto.CreateBusResponse, error)
		LoginDriver(data dto.DriverLoginDto) (dto.DriverLoginResponse, error)
		DeleteBus(id string) error
		EditBus(data dto.EditBusDto, id string, token string) (dto.EditBusResponse, error)
		TrackBusLocation(query dto.BusLocationQuery, c *websocket.Conn) (dto.BusLocationMessage, error)
		StreamBusLocation() []dto.TrackLocationResponse
		BusInfo(id string) (dto.BusInfoResponse, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateBusEntry(data dto.CreateBusDto) (dto.CreateBusResponse, error) {
	var (
		bus      *dto.Bus
		response dto.CreateBusResponse
	)

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password), bcrypt.DefaultCost,
	)

	if err != nil {
		v.shared.Logger.Errorf("error when encrypting password, err: %s", err.Error())
		return response, err
	}

	bus = &dto.Bus{
		Number:   data.Number,
		Plate:    data.Plate,
		Route:    data.Route,
		Username: data.Username,
		Password: string(encryptedPassword),
	}

	err = v.application.BusService.Create(bus)
	if err != nil {
		v.shared.Logger.Errorf("error when inserting bus to database, err: %s", err.Error())
		return response, err
	}

	response = bus.ToCreateBusResponse()

	return response, nil
}

func (v *viewService) LoginDriver(data dto.DriverLoginDto) (dto.DriverLoginResponse, error) {
	var (
		bus      = &dto.Bus{}
		response dto.DriverLoginResponse
	)

	err := v.application.BusService.FindByUsername(data.Username, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when finding bus by username, err: %s", err.Error())
		return response, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(bus.Password),
		[]byte(data.Password),
	)
	if err != nil {
		v.shared.Logger.Errorf("wrong password, err: %s", err.Error())
		return response, err
	}

	token, err := common.NewJWT(bus.Username, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when creating jwt, err: %s", err.Error())
		return response, err
	}

	response = bus.ToDriverLoginResponse(token)

	return response, nil
}

func (v *viewService) DeleteBus(id string) error {
	err := v.application.BusService.Delete(id)
	if err != nil {
		v.shared.Logger.Errorf("error when deleteing bus, err: %s", err.Error())
		return err
	}
	return nil
}

func (v *viewService) EditBus(data dto.EditBusDto, id string, token string) (dto.EditBusResponse, error) {
	var (
		bus      = &dto.Bus{}
		response dto.EditBusResponse
	)

	username, err := common.ExtractTokenData(token, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when extract jwt, err: %s", err.Error())
	}

	err = v.application.BusService.FindByUsername(username, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when checking username, err: %s", err.Error())
		return response, err
	}

	err = v.application.BusService.FindById(id, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when finding bus, err: %s", err.Error())
		return response, err
	}

	bus.FillBusEdit(data)

	err = v.application.BusService.Save(bus)
	if err != nil {
		v.shared.Logger.Errorf("error when saving update, err: %s", err.Error())
		return response, err
	}

	response = bus.ToEditBusResponnse()

	return response, nil
}

func (v *viewService) TrackBusLocation(query dto.BusLocationQuery, c *websocket.Conn) (dto.BusLocationMessage, error) {
	var (
		data = dto.BusLocationMessage{}
		bus  = dto.Bus{}
	)

	username, err := common.ExtractTokenData(query.Token, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when parsing jwt, err: %s", err.Error())
		return data, err
	}

	err = v.application.BusService.FindByUsername(username, &bus)
	if err != nil {
		return data, err
	}

	if err := c.ReadJSON(&data); err != nil {
		v.shared.Logger.Errorf("error when receiving websocket message, err: %s", err.Error())
		return data, err
	}

	location := dto.BusLocation{
		BusID:     bus.ID,
		Lat:       data.Lat,
		Long:      data.Long,
		Timestamp: time.Now(),
		Speed:     data.Speed,
		Heading:   data.Heading,
	}

	go func() {
		v.application.BusService.InsertBusLocation(&location)
		v.shared.Logger.Infof("insert bus location, data: %s", location)
	}()

	if err := c.WriteJSON(data); err != nil {
		v.shared.Logger.Errorf("error when writing message, err: %s", err.Error())
		return data, err
	}

	return data, nil
}

func (v *viewService) StreamBusLocation() []dto.TrackLocationResponse {
	var (
		bus      = []dto.Bus{}
		response = make([]dto.TrackLocationResponse, 0)
	)

	err := v.application.BusService.FindAllBus(&bus)
	if err != nil {
		v.shared.Logger.Errorf("error when finding all bus, err: %s", err.Error())
		return response
	}

	for _, d := range bus {
		parsedData := dto.TrackLocationResponse{
			ID:       d.ID,
			Number:   d.Number,
			Status:   d.Status,
			Route:    d.Route,
			Plate:    d.Plate,
			IsActive: d.IsActive,
		}
		location := dto.BusLocation{}
		err = v.application.BusService.FindBusLatestLocation(d.ID, &location)
		if err != nil {
			v.shared.Logger.Errorf("error when finding bus latest location, err: %s", err.Error())
			continue
		}
		parsedData.Lat = location.Lat
		parsedData.Long = location.Long
		parsedData.Speed = location.Speed
		parsedData.Heading = location.Heading

		response = append(response, parsedData)
	}

	return response
}

func (v *viewService) BusInfo(id string) (dto.BusInfoResponse, error) {
	var (
		res               dto.BusInfoResponse
		terminal          = dto.Terminal{}
		busLatestLocation []dto.TrackLocationResponse
		busInfo           = make([]dto.BusInfo, 0)
	)

	err := v.application.TerminalService.GetById(id, &terminal)
	if err != nil {
		v.shared.Logger.Errorf("error when finding all terminal by id, err: %s", err.Error())
		return res, err
	}

	busLatestLocation = v.StreamBusLocation()

	for _, b := range busLatestLocation {
		distance := common.Distance(b.Lat, b.Long, terminal.Lat, terminal.Long)
		estimate := int(distance/b.Speed) * 3600
		busInfo = append(busInfo, dto.BusInfo{
			ID:       b.ID,
			Number:   b.Number,
			Plate:    b.Plate,
			Status:   b.Status,
			Route:    b.Route,
			Estimate: estimate,
		})
	}

	sort.Slice(busInfo, func(i, j int) bool {
		return busInfo[i].Estimate < busInfo[j].Estimate
	})

	res.Bus = busInfo

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
