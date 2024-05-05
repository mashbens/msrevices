package resp

import (
	"email_verifi_service/business/emailverifi"
)

type Resp struct {
	// TotalAllData int
	// CurrentPage  int
	// PerPage      int
	// TotalPage    int
	Result RespDomain `json:"result"`
}

type RespDomain struct {
	ID int
}

func FromService(v emailverifi.Result) Resp {
	return Resp{
		Result: FromServiceDomain(v.Domain),
	}
}
func FromServiceDomain(u emailverifi.Domain) RespDomain {
	return RespDomain{
		ID: int(u.ID),
	}
}
