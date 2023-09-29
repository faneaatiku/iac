package iac

type Metadata struct {
	PreviousPage int64 `json:"previous_page,omitempty"`
	CurrentPage  int64 `json:"current_page,omitempty"`
	NextPage     int64 `json:"next_page,omitempty"`
	TotalItems   int64 `json:"total_items,omitempty"`
}

func NewMetadata(page int64) *Metadata {
	return &Metadata{CurrentPage: page}
}

func BuildResponseMetadata(mutators ...MetadataMutator) *Metadata {
	meta := Metadata{}
	for _, m := range mutators {
		m.MutateMetadata(&meta)
	}

	return &meta
}

type MetadataMutator interface {
	MutateMetadata(mutator *Metadata)
}

type MetadataMutatorFunc func(metadata *Metadata)

func (f MetadataMutatorFunc) MutateMetadata(metadata *Metadata) {
	f(metadata)
}

func WithPreviousPage(page int64) MetadataMutator {
	return MetadataMutatorFunc(func(metadata *Metadata) {
		metadata.PreviousPage = page
	})
}

func WithCurrentPage(page int64) MetadataMutator {
	return MetadataMutatorFunc(func(metadata *Metadata) {
		metadata.CurrentPage = page
	})
}

func WithNextPage(page int64) MetadataMutator {
	return MetadataMutatorFunc(func(metadata *Metadata) {
		metadata.NextPage = page
	})
}

func WithTotalItems(total int64) MetadataMutator {
	return MetadataMutatorFunc(func(metadata *Metadata) {
		metadata.TotalItems = total
	})
}
