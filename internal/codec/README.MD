##Codec: Encode/Decode

负责特殊数据类型[]byte到map的转换

####支持的类型为:

	None = "none"
	Javascript = "raw"
	Protobuf = "protobuf"
	IntoyunTLV = "intoyuntlv"

接口:
```go
package codec

type Coder interface {
    Decode(data []byte) (map[string]interface{}, error)
    Encode(map[string]interface{}) ([]byte, error)
}
```
