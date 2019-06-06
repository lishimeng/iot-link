package tlv

import (
	"fmt"
	"encoding/json"
)

type Tag struct {

	TagId        int32  `json:"tagId"`
	TagType      int32  `json:"tagType"`
	TagTypeLabel string `json:"tagTypeLabel"`
	TagLength    int32  `json:"tagLength"`
	TagData      string `json:"tagData"`
}

type TagFrame struct {
	Tags []Tag `json:"tags"`
}

func DecodeJson(tlvContent []byte) string {
	tags := DecodeTlv(tlvContent)
	fmt.Print("convert json\n")
	frame := TagFrame{}
	frame.Tags = tags
	s, err := json.Marshal(frame)
	if err != nil {
		fmt.Print(err)
		fmt.Print("\n")
	}
	return fmt.Sprint(string(s))
}

func DecodeToFrame(tlvContent []byte) *TagFrame {
	tags := DecodeTlv(tlvContent)
	frame := TagFrame{}
	frame.Tags = tags
	return &frame
}

func DecodeToJson(tlvContent []byte) string {
	frame := DecodeToFrame(tlvContent)
	s, err := json.Marshal(frame)
	if err != nil {
		return "{}"
	}
	return fmt.Sprint(string(s))
}


func Encode(frame *TagFrame) string{
	return EncodeTlv(frame)
}