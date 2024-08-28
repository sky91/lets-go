package dynamox

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sky91/lets-go/gox/collection"
	"strconv"
	"strings"
)

type Record map[string]types.AttributeValue

func (thisV Record) MarshalJSON() ([]byte, error) {
	m := make(map[string]AttributeValueWrapper)
	for k, v := range thisV {
		m[k] = AttributeValueWrapper{Val: v}
	}
	return json.Marshal(m)
}

func (thisP *Record) UnmarshalJSON(b []byte) error {
	var m map[string]AttributeValueWrapper
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	if *thisP == nil {
		*thisP = make(Record, len(m))
	}
	for k, v := range m {
		(*thisP)[k] = v.Val
	}
	return nil
}

func (thisV Record) GetAsB(key AttrKey) (val *types.AttributeValueMemberB, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberB)
	return
}
func (thisV Record) GetAsBOOL(key AttrKey) (val *types.AttributeValueMemberBOOL, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberBOOL)
	return
}
func (thisV Record) GetAsBS(key AttrKey) (val *types.AttributeValueMemberBS, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberBS)
	return
}
func (thisV Record) GetAsL(key AttrKey) (val *types.AttributeValueMemberL, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberL)
	return
}
func (thisV Record) GetAsM(key AttrKey) (val *types.AttributeValueMemberM, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberM)
	return
}
func (thisV Record) GetAsNULL(key AttrKey) (val *types.AttributeValueMemberNULL, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberNULL)
	return
}
func (thisV Record) GetAsN(key AttrKey) (val *types.AttributeValueMemberN, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberN)
	return
}
func (thisV Record) GetAsNS(key AttrKey) (val *types.AttributeValueMemberNS, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberNS)
	return
}
func (thisV Record) GetAsS(key AttrKey) (val *types.AttributeValueMemberS, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberS)
	return
}
func (thisV Record) GetAsSS(key AttrKey) (val *types.AttributeValueMemberSS, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	val, typeOk = rawVal.(*types.AttributeValueMemberSS)
	return
}

func (thisV Record) GetAsBool(key AttrKey) (val bool, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsBOOL(key)
	if exists && typeOk {
		val = v.Value
	}
	return
}

func (thisV Record) GetAsString(key AttrKey) (val string, exists, typeOk bool) {
	rawVal, exists := thisV[string(key)]
	if !exists {
		return
	}
	switch v := rawVal.(type) {
	case *types.AttributeValueMemberS:
		val, typeOk = v.Value, true
	case *types.AttributeValueMemberN:
		val, typeOk = v.Value, true
	}
	return
}

func (thisV Record) GetAsInt(key AttrKey) (val int, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsN(key)
	if exists && typeOk {
		var err error
		val, err = strconv.Atoi(v.Value)
		typeOk = err == nil
	}
	return
}
func (thisV Record) GetAsIntN(key AttrKey, bitSize int) (val int64, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsN(key)
	if exists && typeOk {
		var err error
		val, err = strconv.ParseInt(v.Value, 10, bitSize)
		typeOk = err == nil
	}
	return
}
func (thisV Record) GetAsInt64(key AttrKey) (val int64, exists, typeOk bool) {
	return thisV.GetAsIntN(key, 64)
}
func (thisV Record) GetAsInt32(key AttrKey) (val int32, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsIntN(key, 32)
	if exists && typeOk {
		val = int32(v)
	}
	return
}
func (thisV Record) GetAsUintN(key AttrKey, bitSize int) (val uint64, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsN(key)
	if exists && typeOk {
		var err error
		val, err = strconv.ParseUint(v.Value, 10, bitSize)
		typeOk = err == nil
	}
	return
}
func (thisV Record) GetAsUint64(key AttrKey) (val uint64, exists, typeOk bool) {
	return thisV.GetAsUintN(key, 64)
}
func (thisV Record) GetAsUint32(key AttrKey) (val uint32, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsUintN(key, 32)
	if exists && typeOk {
		val = uint32(v)
	}
	return
}

func (thisV Record) GetAsRecord(key AttrKey) (record Record, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsM(key)
	if exists && typeOk {
		record = v.Value
	}
	return
}

func (thisV Record) GetAsStringSlice(key AttrKey) (val []string, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsSS(key)
	if exists && typeOk {
		val = v.Value
	}
	return
}

func (thisV Record) GetAsStringSet(key AttrKey) (val map[string]struct{}, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsStringSlice(key)
	if exists && typeOk {
		val = collection.Slice2Set(v)
	}
	return
}

func (thisV Record) GetAsSlice(key AttrKey) (val []types.AttributeValue, exists, typeOk bool) {
	v, exists, typeOk := thisV.GetAsL(key)
	if exists && typeOk {
		val = v.Value
	}
	return
}

func (thisV Record) SubRecord(keys ...AttrKey) Record {
	sub := make(Record, len(keys))
	for _, key := range keys {
		if v, exists := thisV[string(key)]; exists {
			sub[string(key)] = v
		}
	}
	return sub
}

func (thisV Record) SetAttributeValue(key AttrKey, val types.AttributeValue) {
	thisV[string(key)] = val
}

func (thisV Record) SetBool(key AttrKey, val bool) {
	thisV.SetAttributeValue(key, &types.AttributeValueMemberBOOL{Value: val})
}

func (thisV Record) SetString(key AttrKey, val string) {
	thisV.SetAttributeValue(key, &types.AttributeValueMemberS{Value: val})
}

func (thisV Record) SetNumber(key AttrKey, val string) {
	thisV.SetAttributeValue(key, &types.AttributeValueMemberN{Value: val})
}
func (thisV Record) SetInt(key AttrKey, val int) {
	thisV.SetNumber(key, strconv.Itoa(val))
}
func (thisV Record) SetInt64(key AttrKey, val int64) {
	thisV.SetNumber(key, strconv.FormatInt(val, 10))
}
func (thisV Record) SetUint64(key AttrKey, val uint64) {
	thisV.SetNumber(key, strconv.FormatUint(val, 10))
}

func (thisV Record) SetSlice(key AttrKey, val []types.AttributeValue) {
	thisV.SetAttributeValue(key, &types.AttributeValueMemberL{Value: val})
}

type StringAttr struct {
	Key AttrKey
	Val string
}

type StringPositionalRecord []StringAttr

func ParseStringPositionalRecord(value string) (StringPositionalRecord, error) {
	if !strings.HasPrefix(value, "/") {
		return nil, fmt.Errorf("invalid Attr: [%s]", value)
	}
	split := strings.Split(value[1:], "/")
	values := make(StringPositionalRecord, len(split))
	for i, s := range split {
		kvSplit := strings.SplitN(s, "+", 2)
		if len(kvSplit) != 2 {
			return nil, fmt.Errorf("invalid Attr: [%s]", value)
		}
		values[i].Key = AttrKey(kvSplit[0])
		values[i].Val = kvSplit[1]
	}
	return values, nil
}

func (thisV StringPositionalRecord) Find(key AttrKey) (string, bool) {
	for _, attr := range thisV {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func (thisV StringPositionalRecord) String() string {
	sb := strings.Builder{}
	sb.Grow(64)
	for _, attr := range thisV {
		_ = sb.WriteByte('/')
		_, _ = sb.WriteString(string(attr.Key))
		_ = sb.WriteByte('+')
		_, _ = sb.WriteString(attr.Val)
	}
	return sb.String()
}
