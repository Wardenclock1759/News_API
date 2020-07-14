package models

type Tag struct {
	ID   string `json:"tag_id"`
	Name string `json:"tag_name"`
}

var tagStorage = NewTagController()

func NewTagController() *map[string]Tag {
	tag1 := Tag{
		ID:   "id0",
		Name: "None",
	}

	tag2 := Tag{
		ID:   "id1",
		Name: "Sport",
	}

	tag3 := Tag{
		ID:   "id2",
		Name: "Celebrity",
	}

	res := map[string]Tag{}

	res[tag1.ID] = tag1
	res[tag2.ID] = tag2
	res[tag3.ID] = tag3

	return &res
}

func AddTag(tag Tag) Tag {
	s := *tagStorage
	tag.ID = GenerateID()
	s[tag.ID] = tag
	return tag
}

func GetOrCreateTagByName(name string) *Tag {
	if name == "" {
		return GetOrCreateTagByName("None")
	}
	storage := *tagStorage
	for _, _tag := range storage {
		if _tag.Name == name {
			return &_tag
		}
	}
	tag := AddTag(Tag{"", name})
	return &tag
}

func GetTags() []Tag {
	tags := make([]Tag, len(*tagStorage))

	i := 0
	for _, tag := range *tagStorage {
		tags[i] = tag
		i++
	}

	return tags
}

func GetTagByID(id string) (*Tag, bool) {
	storage := *tagStorage
	tag, ok := storage[id]
	return &tag, ok
}
