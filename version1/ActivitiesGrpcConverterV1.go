package version1

import (
	"encoding/json"

	"github.com/pip-services-users/pip-clients-activities-go/protos"
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromStringValueMap(value *data.StringValueMap) map[string]string {
	if value == nil {
		return nil
	}

	return value.Value()
}

func toStringValueMap(value map[string]string) *data.StringValueMap {
	if value == nil {
		return nil
	}

	return data.NewStringValueMapFromMaps(value)
}

func fromReference(reference *ReferenceV1) *protos.Reference {
	if reference == nil {
		return nil
	}

	obj := &protos.Reference{
		Id:   reference.Id,
		Type: reference.Type,
		Name: reference.Name,
	}

	return obj
}

func toReference(obj *protos.Reference) *ReferenceV1 {
	if obj == nil {
		return nil
	}

	reference := &ReferenceV1{
		Id:   obj.Id,
		Type: obj.Type,
		Name: obj.Name,
	}

	return reference
}

func fromReferences(references []*ReferenceV1) []*protos.Reference {
	if references == nil {
		return nil
	}

	obj := make([]*protos.Reference, len(references))

	for i, v := range references {
		obj[i] = fromReference(v)
	}

	return obj
}

func toReferences(obj []*protos.Reference) []*ReferenceV1 {
	if obj == nil {
		return nil
	}

	references := make([]*ReferenceV1, len(obj))

	for i, v := range obj {
		references[i] = toReference(v)
	}

	return references
}

func fromPartyActivity(activity *PartyActivityV1) *protos.PartyActivity {
	if activity == nil {
		return nil
	}

	obj := &protos.PartyActivity{
		Id:         activity.Id,
		OrgId:      activity.OrgId,
		Time:       convert.StringConverter.ToString(activity.Time),
		Type:       activity.Type,
		Party:      fromReference(activity.Party),
		RefItem:    fromReference(activity.RefItem),
		RefParents: fromReferences(activity.RefParents),
		RefParty:   fromReference(activity.RefParty),
		Details:    fromStringValueMap(activity.Details),
	}

	return obj
}

func toPartyActivity(obj *protos.PartyActivity) *PartyActivityV1 {
	if obj == nil {
		return nil
	}

	activity := &PartyActivityV1{
		Id:         obj.Id,
		OrgId:      obj.OrgId,
		Time:       convert.DateTimeConverter.ToDateTime(obj.Time),
		Type:       obj.Type,
		Party:      toReference(obj.Party),
		RefItem:    toReference(obj.RefItem),
		RefParents: toReferences(obj.RefParents),
		RefParty:   toReference(obj.RefParty),
		Details:    toStringValueMap(obj.Details),
	}

	return activity
}

func fromPartyActivityPage(page *data.DataPage) *protos.PartyActivityPage {
	if page == nil {
		return nil
	}

	obj := &protos.PartyActivityPage{
		Total: *page.Total,
		Data:  make([]*protos.PartyActivity, len(page.Data)),
	}

	for i, v := range page.Data {
		activity := v.(*PartyActivityV1)
		obj.Data[i] = fromPartyActivity(activity)
	}

	return obj
}

func toPartyActivityPage(obj *protos.PartyActivityPage) *data.DataPage {
	if obj == nil {
		return nil
	}

	activities := make([]interface{}, len(obj.Data))

	for i, v := range obj.Data {
		activities[i] = toPartyActivity(v)
	}

	page := data.NewDataPage(&obj.Total, activities)

	return page
}

func fromPartyActivities(activities []*PartyActivityV1) []*protos.PartyActivity {
	if activities == nil {
		return nil
	}

	obj := make([]*protos.PartyActivity, len(activities))

	for i, v := range activities {
		obj[i] = fromPartyActivity(v)
	}

	return obj
}

func toPartyActivities(obj []*protos.PartyActivity) []*PartyActivityV1 {
	if obj == nil {
		return nil
	}

	activities := make([]*PartyActivityV1, len(obj))

	for i, v := range obj {
		activities[i] = toPartyActivity(v)
	}

	return activities
}
