package version1

import (
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IActivitiesClientV1 interface {
	GetPartyActivities(correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result *data.DataPage, err error)

	LogPartyActivity(correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error)

	BatchPartyActivities(correlationId string, activities []*PartyActivityV1) error

	DeletePartyActivities(correlationId string, filter *data.FilterParams) error
}
