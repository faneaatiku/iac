package iac

const (
	TypeGeneric    = "generic"    //the message is referring to the entire request
	TypeContextual = "contextual" //the message is referring to a certain part of the request that made it fail
)

// Message - represents an error on the response
// The Message.Type can be either TypeGeneric or TypeContextual
//
// # The Text.Text is a description of the error that occurred
//
// The Message.Path is the path of the request where the error occurred. It should be used only when the
// Message is of type TypeContextual
// Example:
// Assuming the JSON submitted is
//
//	{
//		"name": "John",
//		"surname": "Doe",
//		"address": {
//			"street": ""
//	}
//
// and that the street address is mandatory then the json path for this error would be $.address.street
// https://goessner.net/articles/JsonPath/
type Message struct {
	Type string `json:"type"`
	Text string `json:"message"`
	Path string `json:"path,omitempty"`
}

func NewGenericMessage(text string) *Message {
	return &Message{Text: text, Type: TypeGeneric}
}

func NewContextualMessage(text, path string) *Message {
	return &Message{Type: TypeContextual, Text: text, Path: path}
}

func BuildResponseMessage(mutators ...MessageMutator) *Message {
	msg := Message{}
	for _, m := range mutators {
		m.MutateMessage(&msg)
	}

	return &msg
}

type MessageMutator interface {
	MutateMessage(mutator *Message)
}

type MessageMutatorFunc func(message *Message)

func (f MessageMutatorFunc) MutateMessage(message *Message) {
	f(message)
}

func WithType(mType string) MessageMutator {
	return MessageMutatorFunc(func(message *Message) {
		//override with a valid one if it's not a known one
		if mType != TypeGeneric && mType != TypeContextual {
			mType = TypeGeneric
		}

		message.Type = mType
	})
}

func WithText(mText string) MessageMutator {
	return MessageMutatorFunc(func(message *Message) {
		message.Text = mText
	})
}

func WithPath(mPath string) MessageMutator {
	return MessageMutatorFunc(func(message *Message) {
		message.Path = mPath
	})
}
