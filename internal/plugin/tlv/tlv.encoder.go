package tlv

import (
	"encoding/json"
	"fmt"
)

func EncodeTlvTxt(data string) string {
	var tlv string
	var tagFrame TagFrame
	tlv = "1" // TODO
	byteArray := []byte(data)
	err := json.Unmarshal(byteArray[0:], &tagFrame)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		tlv += EncodeTlv(&tagFrame)
	}
	return tlv
}

func EncodeTlv(frame *TagFrame) string {
	tags := frame.Tags
	var tlv string
	for i := 0; i < len(tags); i++ {
		tagContent := encodeTlvTag(&tags[i])
		tlv += tagContent
	}
	return tlv
}

func encodeTlvTag(tag *Tag) string {
	var content string
	tagId := tag.TagId
	content += IntToString(tagId)
	content += IntToString(tag.TagType)
	content += IntToString(tag.TagLength)

	switch tag.TagType {
	case 0:// bool

		n := StrToInt(tag.TagData)
		content += string(IntToBytes(n, tag.TagLength))
		break
	case 1:// number
		n := StrToInt(tag.TagData)
		content += string(IntToBytes(n, tag.TagLength))
		break
	case 2:// enum
		n := StrToInt(tag.TagData)
		content += string(IntToBytes(n, tag.TagLength))
		break
	case 3://txt
		content += tag.TagData
		break
	default:
		content += tag.TagData
		break
	}

	return content
}