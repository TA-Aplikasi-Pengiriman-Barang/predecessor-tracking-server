package terminal

import (
	"sort"
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"
)

type (
	ViewService interface {
		GetTerminalInfo(id string) (dto.GetTerminalInfoResponse, error)
		GetAllTerminalSorted(data dto.GetAllTerminalDto) (dto.GetAllTerminalResponse, error)
		GetTwoClosesTerminal(data dto.GetAllTerminalDto) (dto.GetAllTerminalResponse, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetTerminalInfo(id string) (dto.GetTerminalInfoResponse, error) {
	var (
		response           dto.GetTerminalInfoResponse
		terminal           = &dto.Terminal{}
		allTerminalInRoute = &[]dto.Terminal{}
	)

	err := v.application.TerminalService.GetById(id, terminal)
	if err != nil {
		v.shared.Logger.Errorf("error when finding terminal by id, err: %s", err.Error())
		return response, err
	}

	err = v.application.TerminalService.GetAllByRoute(terminal.Route, allTerminalInRoute)
	if err != nil {
		v.shared.Logger.Errorf("error when finding all terminal by route, err: %s", err.Error())
		return response, err
	}

	response = terminal.ToTerminalInfo(*allTerminalInRoute)

	return response, nil
}

func (v *viewService) GetAllTerminalSorted(data dto.GetAllTerminalDto) (dto.GetAllTerminalResponse, error) {
	var (
		res                dto.GetAllTerminalResponse
		terminals          = []dto.Terminal{}
		terminalListSorted = make([]dto.TerminalListWithDistance, 0)
	)

	err := v.application.TerminalService.GetAllTerminal(&terminals)
	if err != nil {
		v.shared.Logger.Errorf("error when getting all terminal, err: %s", err.Error())
		return res, err
	}

	for i := 0; i < len(terminals); i++ {
		next := terminals[0]
		if i != len(terminals)-1 {
			next = terminals[i+1]
		}

		terminalSorted := v.getTerminalDistance(data, terminals[i], next.Name)

		terminalListSorted = append(terminalListSorted, terminalSorted)
	}

	sort.Slice(terminalListSorted, func(i, j int) bool {
		return terminalListSorted[i].Distance < terminalListSorted[j].Distance
	})

	res = dto.GetAllTerminalResponse{
		Terminals: terminalListSorted,
	}

	return res, nil
}

func (v *viewService) GetTwoClosesTerminal(data dto.GetAllTerminalDto) (dto.GetAllTerminalResponse, error) {
	resp, err := v.GetAllTerminalSorted(data)
	if err != nil {
		return resp, err
	}

	resp.Terminals = resp.Terminals[0:2]

	return resp, nil
}

func (v *viewService) getTerminalDistance(data dto.GetAllTerminalDto, terminal dto.Terminal, next string) dto.TerminalListWithDistance {
	res := dto.TerminalListWithDistance{
		ID:    terminal.ID,
		Name:  terminal.Name,
		Route: terminal.Route,
		Next:  next,
	}

	res.Distance = common.Distance(data.Lat, data.Long, terminal.Lat, terminal.Long)

	return res
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
