package version1

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

type PartyActivityV1 struct {
	/* Identification */
	Id string `json:"id"`

	/* Identification fields */
	Time  time.Time    `json:"time"`
	Type  string       `json:"type"`
	Party *ReferenceV1 `json:"party"`

	/* References objects (notes, goals, etc.) */
	RefItem    *ReferenceV1   `json:"ref_item"`
	RefParents []*ReferenceV1 `json:"ref_parents"`
	RefParty   *ReferenceV1   `json:"ref_party"`

	/* Other details like % of progress or new status */
	Details *data.StringValueMap `json:"details"`
}

func EmptyPartyActivityV1() *PartyActivityV1 {
	return &PartyActivityV1{}
}

func NewPartyActivityV1(id string, typ string, party *ReferenceV1,
	refItem *ReferenceV1, refParents []*ReferenceV1,
	refParty *ReferenceV1, details *data.StringValueMap) *PartyActivityV1 {
	return &PartyActivityV1{
		Id:         id,
		Time:       time.Now(),
		Type:       typ,
		Party:      party,
		RefItem:    refItem,
		RefParents: refParents,
		RefParty:   refParty,
		Details:    details,
	}
}
