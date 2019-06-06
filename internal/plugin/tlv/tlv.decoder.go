package tlv

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
)


func DecodeData(tag *Tag, data []byte) {
	switch tag.TagType {
	case 0:// bool
		tmp := BytesToInt(data[0:1])
		tag.TagData = fmt.Sprintf("%d", tmp)
		tag.TagTypeLabel = "bool"
		break
	case 1:
		//number
		tmp := BytesToInt(data)
		tag.TagData = fmt.Sprintf("%d", tmp)
		tag.TagTypeLabel = "number"
		break
	case 2:
		// enum/num
		tmp := BytesToInt(data[0:1])
		tag.TagData = fmt.Sprintf("%d", tmp)
		tag.TagTypeLabel = "enum"
		break
	case 3:
		// txt
		tag.TagData = string(data)
		tag.TagTypeLabel = "txt"
		break
	default:
		tag.TagData = string(data)
		tag.TagTypeLabel = "extra"
		break
	}
}

func DecodeTlvTag(b []byte) (Tag, int32) {
	buf := bytes.NewBuffer(b)
	var tagSize int32 = 0
	// tag
	var tag = Tag{}

	tag.TagId = 1
	// id
	var idLow uint8
	var idHigh uint8
	_ = binary.Read(buf, binary.BigEndian, &idLow)
	tag.TagId = int32(idLow)
	tagSize++
	if idLow >= 0x80 {
		_ = binary.Read(buf, binary.BigEndian, &idHigh)
		tag.TagId = (tag.TagId << 8) + int32(idHigh)
		tagSize++
	}

	// type
	var tagType uint8
	_ = binary.Read(buf, binary.BigEndian, &tagType)
	tag.TagType = int32(tagType)
	tagSize++

	// length
	var lenLow uint8
	var lenHigh uint8
	_ = binary.Read(buf, binary.BigEndian, &lenLow)
	tag.TagLength = int32(lenLow)
	tagSize++
	if tag.TagLength >= 0x80 {
		_ = binary.Read(buf, binary.BigEndian, &lenHigh)
		tag.TagLength = (tag.TagLength << 8) + int32(lenHigh)
		tagSize++
	}

	// content
	data := make([]byte, tag.TagLength)
	_, _ = buf.Read(data)
	var t = &tag
	DecodeData(t, data)
	tagSize += tag.TagLength
	return tag, tagSize
}

func DecodeTlv(data []byte) []Tag {
	tlvs := make([]byte, len(data))
	copy(tlvs, data)
	res := list.New()
	for {
		tag, length := DecodeTlvTag(tlvs)

		res.PushBack(tag)
		if length < int32(len(tlvs)) {
			tlvs = tlvs[length:]
		} else {
			break
		}
	}
	tlvTags := make([]Tag, res.Len())
	i := 0
	for e := res.Front(); e != nil; e = e.Next() {
		tmp := e.Value
		t := tmp.(Tag)
		tlvTags[i] = t
		i++
	}
	return tlvTags
}
