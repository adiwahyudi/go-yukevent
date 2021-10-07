package request

import "yukevent/business/events"

type Events struct {
	Name             string `json:"name"`
	Description      string `json:"desc"`
	Syarat_Ketentuan string `json:"s&k"`
	Location         string `json:"loc"`
	Event_Date       string `json:"event_date"`
	Event_Time       string `json:"event_time"`
	Close_Register   string `json:"close_regist"`
	Capacity         int    `json:"capacity"`
	Poster           string `json:"poster"`
	Price            int    `json:"price"`
}

func (req *Events) ToDomain() *events.Domain {
	return &events.Domain{
		Name:             req.Name,
		Description:      req.Description,
		Syarat_Ketentuan: req.Syarat_Ketentuan,
		Location:         req.Location,
		Event_Date:       req.Event_Date,
		Event_Time:       req.Event_Time,
		Close_Register:   req.Close_Register,
		Capacity:         req.Capacity,
		Poster:           req.Poster,
		Price:            req.Price,
	}
}
