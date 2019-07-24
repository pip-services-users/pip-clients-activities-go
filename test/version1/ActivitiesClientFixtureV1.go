package test_version1

import (
	"testing"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/data"

	"github.com/pip-services-users/pip-clients-activities-go/version1"
	"github.com/stretchr/testify/assert"
)

type ActivitiesClientFixtureV1 struct {
	Client version1.IActivitiesClientV1
}

var ACTIVITY = &version1.PartyActivityV1{
	Id:   "",
	Type: "test",
	Time: time.Now(),
	Party: &version1.ReferenceV1{
		Id:   "1",
		Type: "party",
		Name: "Test User",
	},
	RefItem: &version1.ReferenceV1{
		Id:   "2",
		Type: "party",
		Name: "Admin User",
	},
	RefParents: []*version1.ReferenceV1{},
	RefParty:   nil,
	Details:    nil,
}

func NewActivitiesClientFixtureV1(client version1.IActivitiesClientV1) *ActivitiesClientFixtureV1 {
	return &ActivitiesClientFixtureV1{
		Client: client,
	}
}

func (c *ActivitiesClientFixtureV1) clear() {
	c.Client.DeletePartyActivities("", nil)
}

func (c *ActivitiesClientFixtureV1) TestBatchPartyActivities(t *testing.T) {
	c.clear()
	defer c.clear()

	// Log an activity batch
	err := c.Client.BatchPartyActivities("", []*version1.PartyActivityV1{ACTIVITY, ACTIVITY, ACTIVITY})
	assert.Nil(t, err)

	// Get activities
	page, err1 := c.Client.GetPartyActivities("", data.NewFilterParamsFromTuples("party_id", "1"), nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) > 2)

	activity := page.Data[0].(*version1.PartyActivityV1)

	assert.NotNil(t, activity.Time)
	assert.Equal(t, activity.Type, ACTIVITY.Type)
	assert.Equal(t, activity.Party.Name, ACTIVITY.Party.Name)
}
