package version1

import (
	"github.com/pip-services-users/pip-clients-activities-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-grpc-go/clients"
)

type ActivitiesGrpcClientV1 struct {
	clients.GrpcClient
}

func NewActivitiesGrpcClientV1() *ActivitiesGrpcClientV1 {
	return &ActivitiesGrpcClientV1{
		GrpcClient: *clients.NewGrpcClient("activities_v1.Activities"),
	}
}

func (c *ActivitiesGrpcClientV1) GetPartyActivities(correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result *data.DataPage, err error) {
	req := &protos.PartyActivityPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.PartyActivityPageReply)
	err = c.Call("get_party_activities", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toPartyActivityPage(reply.Page)

	return result, nil
}

func (c *ActivitiesGrpcClientV1) LogPartyActivity(correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error) {
	req := &protos.PartyActivityLogRequest{
		CorrelationId: correlationId,
		Activity:      fromPartyActivity(activity),
	}

	reply := new(protos.PartyActivityObjectReply)
	err = c.Call("log_party_activity", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toPartyActivity(reply.Activity)

	return result, nil
}

func (c *ActivitiesGrpcClientV1) BatchPartyActivities(correlationId string, activities []*PartyActivityV1) error {
	req := &protos.PartyActivityBatchRequest{
		CorrelationId: correlationId,
		Activities:    fromPartyActivities(activities),
	}

	reply := new(protos.PartyActivityOnlyErrorReply)
	err := c.Call("batch_party_activities", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *ActivitiesGrpcClientV1) DeletePartyActivities(correlationId string, filter *data.FilterParams) error {
	req := &protos.PartyActivityDeleteRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}

	reply := new(protos.PartyActivityOnlyErrorReply)
	err := c.Call("delete_party_activities", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
