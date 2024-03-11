package commonlib

import "github.com/gorilla/schema"

var decoder *schema.Decoder

func GetSchemaDecoder() *schema.Decoder {
	if decoder == nil {
		decoder = schema.NewDecoder()
	}
	return decoder
}
