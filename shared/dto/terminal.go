package dto

import (
	"strings"
)

type (
	Terminal struct {
		ID          uint    `gorm:"primaryKey;autoIncrement"`
		Name        string  `gorm:"colum:name;uniqueIndex:route_name_pair"`
		Route       Route   `gorm:"column:route;uniqueIndex:route_name_pair"`
		PlaceAround string  `gorm:"column:place_around"`
		Long        float64 `gorm:"column:longitude"`
		Lat         float64 `gorm:"column:latitude"`
	}

	VisitedTerminal struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Past bool   `json:"past"`
	}

	// GetTerminalInfoResponse GetTerminalInfoResponse
	GetTerminalInfoResponse struct {
		Name            string            `json:"name"`
		Route           Route             `json:"route"`
		RelatedPlace    []string          `json:"relatedPlace"`
		RelatedTerminal []VisitedTerminal `json:"relatedTerminal"`
	}

	TerminalListWithDistance struct {
		ID       uint    `json:"id"`
		Distance float64 `json:"distance"`
		Name     string  `json:"name"`
		Next     string  `json:"next"`
		Route    Route   `json:"route"`
	}

	// GetAllTerminalDto GetAllTerminalDto
	GetAllTerminalDto struct {
		Long float64 `json:"long" validate:"required"`
		Lat  float64 `json:"lat" validate:"required"`
	}

	// GetAllTerminalResponse GetAllTerminalResponse
	GetAllTerminalResponse struct {
		Terminals []TerminalListWithDistance `json:"terminal"`
	}
)

func (t *Terminal) ToTerminalInfo(terminal []Terminal) GetTerminalInfoResponse {
	var (
		res                      = GetTerminalInfoResponse{}
		visitedTerminal          = make([]VisitedTerminal, 0)
		isCurrentTerminalVisited = true
	)

	res.Name = t.Name
	res.RelatedPlace = strings.Split(t.PlaceAround, ",")
	res.Route = t.Route

	for _, v := range terminal {
		vt := VisitedTerminal{
			ID:   v.ID,
			Name: v.Name,
			Past: isCurrentTerminalVisited,
		}

		visitedTerminal = append(visitedTerminal, vt)

		if v.ID == t.ID {
			isCurrentTerminalVisited = false
		}
	}

	res.RelatedTerminal = visitedTerminal

	return res
}

func (t *Terminal) Seeder() []Terminal {
	return []Terminal{
		{
			Name:        "Asrama UI",
			Route:       RED,
			PlaceAround: "Wisma Makara,Hall Sabha Widya",
			Lat:         -6.348373127525387,
			Long:        106.8297679527903,
		},
		{
			Name:        "Menwa",
			Route:       RED,
			PlaceAround: "Halte Transjakarta UI Depok",
			Lat:         -6.353465386293707,
			Long:        106.83182325822173,
		},
		{
			Name:        "Stasiun UI",
			Route:       RED,
			PlaceAround: "Apartemen Taman Melati",
			Lat:         -6.361046716889507,
			Long:        106.8317240044786,
		},
		{
			Name:        "FH",
			Route:       RED,
			PlaceAround: "Pintu Belakang Rel (Barel), Masjid UI,Perpustakaan UI,Balai Sebaguna Purnowo Prawiro UI,Fasilkom (Gedung Lama)",
			Lat:         -6.364864762651361,
			Long:        106.83223079221105,
		},
		{
			Name:        "Balairung",
			Route:       RED,
			PlaceAround: "Stasiun Pondok Cina,Makara Art Center",
			Lat:         -6.368205271318413,
			Long:        106.83184387661237,
		},
		{
			Name:        "RIK",
			Route:       RED,
			PlaceAround: "FKM",
			Lat:         -6.370190241285223,
			Long:        106.83109626794518,
		},
		{
			Name:        "RSUI",
			Route:       RED,
			PlaceAround: "FKM",
			Lat:         -6.371697905932189,
			Long:        106.8293758480366,
		},
		{
			Name:        "FIK",
			Route:       RED,
			PlaceAround: "Fasilkom (Gedung Baru)",
			Lat:         -6.371101272862191,
			Long:        106.82696734342873,
		},
		{
			Name:        "FMIPA",
			Route:       RED,
			PlaceAround: "",
			Lat:         -6.369838377677364,
			Long:        106.82575903066468,
		},
		{
			Name:        "SOR",
			Route:       RED,
			PlaceAround: "Politeknik Teknik Jakarta (PNJ),Gymnasium",
			Lat:         -6.367004060974791,
			Long:        106.82448615509534,
		},
		{
			Name:        "Vokasi",
			Route:       RED,
			PlaceAround: "Pusat Kegiatan Mahasiswa (Pusgiwa),Stadion, Career Development UI",
			Lat:         -6.366114158411598,
			Long:        106.82167086626085,
		},
		{
			Name:        "FT",
			Route:       RED,
			PlaceAround: "Pintu Kukusan-Teknik (Kutek)",
			Lat:         -6.361069834121701,
			Long:        106.82321257394592,
		},
		{
			Name:        "FEB",
			Route:       RED,
			PlaceAround: "",
			Lat:         -6.359443306471211,
			Long:        106.82575218376806,
		},
		{
			Name:        "FIB",
			Route:       RED,
			PlaceAround: "",
			Lat:         -6.361133065901284,
			Long:        106.82970210098532,
		},
		{
			Name:        "FISIP",
			Route:       RED,
			PlaceAround: "Fasilkom (Gedung Lama)",
			Lat:         -6.361723672481245,
			Long:        106.83030996654941,
		},
		{
			Name:        "F.Psi",
			Route:       RED,
			PlaceAround: "",
			Lat:         -6.362172631787366,
			Long:        106.83083040668357,
		}, {
			Name:        "Asrama UI",
			Route:       BLUE,
			PlaceAround: "Wisma Makara,Hall Sabha Widya",
			Lat:         -6.348373127525387,
			Long:        106.8297679527903,
		},
		{
			Name:        "Menwa",
			Route:       BLUE,
			PlaceAround: "Halte Transjakarta UI Depok",
			Lat:         -6.3534610177474,
			Long:        106.83162029695444,
		},
		{
			Name:        "Stasiun UI",
			Route:       BLUE,
			PlaceAround: "Apartemen Taman Melati",
			Lat:         -6.36086929545325,
			Long:        106.83146112622818,
		},
		{
			Name:        "F.Psi",
			Route:       BLUE,
			PlaceAround: "",
			Lat:         -6.362850786328479,
			Long:        106.83116675399012,
		},
		{
			Name:        "FISIP",
			Route:       BLUE,
			PlaceAround: "Fasilkom (Gedung Lama)",
			Lat:         -6.361835631548166,
			Long:        106.83016512726645,
		},
		{
			Name:        "FIB",
			Route:       BLUE,
			PlaceAround: "",
			Lat:         -6.361143942325545,
			Long:        106.82947600448857,
		},
		{
			Name:        "FEB",
			Route:       BLUE,
			PlaceAround: "",
			Lat:         -6.359626440076783,
			Long:        106.82572631991094,
		},
		{
			Name:        "FT",
			Route:       BLUE,
			PlaceAround: "Pintu Kukusan-Teknik (Kutek)",
			Lat:         -6.361277803365007,
			Long:        106.82333110948572,
		},
		{
			Name:        "Vokasi",
			Route:       BLUE,
			PlaceAround: "Pusat Kegiatan Mahasiswa (Pusgiwa),Stadion,Career Development UI",
			Lat:         -6.3659382798442765,
			Long:        106.82177091590128,
		},
		{
			Name:        "SOR",
			Route:       BLUE,
			PlaceAround: "Politeknik Teknik Jakarta (PNJ),Gymnasium",
			Lat:         -6.366780044175628,
			Long:        106.82385382123188,
		},
		{
			Name:        "FMIPA",
			Route:       BLUE,
			PlaceAround: "",
			Lat:         -6.369756748019911,
			Long:        106.8259792966702,
		},
		{
			Name:        "FIK",
			Route:       BLUE,
			PlaceAround: "Fasilkom (Gedung Baru)",
			Lat:         -6.371061340660264,
			Long:        106.82719280923317,
		},
		{
			Name:        "FKM",
			Route:       BLUE,
			PlaceAround: "RSUI",
			Lat:         -6.3714849070313475,
			Long:        106.82925853982198,
		},
		{
			Name:        "RIK",
			Route:       BLUE,
			PlaceAround: "",
			Lat:         -6.370075618390656,
			Long:        106.8308615746626,
		},
		{
			Name:        "Balairung",
			Route:       BLUE,
			PlaceAround: "Stasiun Pondok Cina,Makara Art Center",
			Lat:         -6.36809398005471,
			Long:        106.83161722995663,
		},
		{
			Name:        "Masjid UI",
			Route:       BLUE,
			PlaceAround: "Pintu Belakang Rel (Barel), Masjid UI,Perpustakaan UI,Balai Sebaguna Purnowo Prawiro UI,Fasilkom (Gedung Lama)",
			Lat:         -6.365574974922631,
			Long:        106.83203831176702,
		},
	}
}
