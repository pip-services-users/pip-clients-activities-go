package version1

type ReferenceV1 struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func EmptyReferenceV1() *ReferenceV1 {
	return &ReferenceV1{}
}

func NewReferenceV1(id string, typ string, name string) *ReferenceV1 {
	return &ReferenceV1{
		Id:   id,
		Type: typ,
		Name: name,
	}
}
