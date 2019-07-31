package codec

const (
	None       = "none"
	Javascript = "raw"
	Protobuf   = "protobuf"
	IntoyunTLV = "intoyuntlv"
)

type CodecPlugin struct {
	name string
}

type Coder interface {
	Decode(appId string, data []byte) (props map[string]interface{}, err error)
	Encode(appId string, props map[string]interface{}) (data []byte, err error)
}
