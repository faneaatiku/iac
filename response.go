package iac

type Response struct {
	Messages []Message   `json:"messages,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata Metadata    `json:"metadata,omitempty"`
}

func (r *Response) SetData(data interface{}) {
	r.Data = data
}

func (r *Response) SetMetadata(metadata Metadata) {
	r.Metadata = metadata
}

func (r *Response) AddMessage(msg Message) {
	r.Messages = append(r.Messages, msg)
}

func NewResponse(data interface{}) *Response {
	return &Response{Data: data}
}

func NewEmptyResponse() *Response {
	return &Response{}
}

func BuildResponse(mutators ...ResponseMutator) *Response {
	resp := Response{}
	for _, m := range mutators {
		m.MutateResponse(&resp)
	}

	return &resp
}

type ResponseMutator interface {
	MutateResponse(mutator *Response)
}

type ResponseMutatorFunc func(resp *Response)

func (f ResponseMutatorFunc) MutateResponse(resp *Response) {
	f(resp)
}

func WithData(data interface{}) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Data = data
	})
}

func WithMetadata(metadata Metadata) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Metadata = metadata
	})
}

func WithMessages(messages []Message) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Messages = messages
	})
}

func WithMsg(message Message) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Messages = append(resp.Messages, message)
	})
}

func WithGenericMsg(messageText string) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Messages = append(resp.Messages, *NewGenericMessage(messageText))
	})
}

func WithContextualMsg(messageText, messagePath string) ResponseMutator {
	return ResponseMutatorFunc(func(resp *Response) {
		resp.Messages = append(resp.Messages, *NewContextualMessage(messageText, messagePath))
	})
}
