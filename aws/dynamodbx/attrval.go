package dynamodbx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sky91/lets-go/aws/dynamox"
	"github.com/sky91/lets-go/gox"
)

// deprecated
type IsAttrVal interface {
	isAttrVal()
}

// deprecated
type AttrVal struct {
	Val IsAttrVal
}

// deprecated
func (thisV AttrVal) ToAttributeValue() types.AttributeValue {
	return thisV.Val.(interface{ ToAttributeValue() types.AttributeValue }).ToAttributeValue()
}

// deprecated
func (thisV AttrVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(thisV.Val)
}

// deprecated
func (thisP *AttrVal) UnmarshalJSON(data []byte) error {
	var val map[string]json.RawMessage
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	if len(val) != 1 {
		return fmt.Errorf("invalid AttrVal: [%s]", data)
	}
	for k := range val {
		switch k {
		case "B":
			var av AttrValBinary
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "BOOL":
			var av AttrValBool
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "BS":
			var av AttrValBinarySet
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "L":
			var av AttrValList
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "M":
			var av AttrValMap
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "NULL":
			var av AttrValNull
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "N":
			var av AttrValNumber
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "NS":
			var av AttrValNumberSet
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "SS":
			var av AttrValStringSet
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		case "S":
			var av AttrValString
			if err := json.Unmarshal(data, &av); err != nil {
				return err
			}
			*thisP = AttrVal{Val: av}
		default:
			return fmt.Errorf("invalid AttrVal: [%s]", data)
		}
		return nil
	}
	return fmt.Errorf("invalid AttrVal: [%s]", data)
}

// deprecated
type AttrValBinaryStruct struct {
	Val []byte `json:"B"`
}

// deprecated
type AttrValBinary []byte

// deprecated
func (thisV AttrValBinary) isAttrVal() {}

// deprecated
func (thisV AttrValBinary) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberB())
}

// deprecated
func (thisV AttrValBinary) ToAttributeValueMemberB() types.AttributeValueMemberB {
	return types.AttributeValueMemberB{Value: thisV}
}

// deprecated
func (thisV AttrValBinary) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValBinaryStruct{Val: thisV})
}

// deprecated
func (thisP *AttrValBinary) UnmarshalJSON(data []byte) error {
	var v AttrValBinaryStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

// deprecated
type AttrValBoolStruct struct {
	Val bool `json:"BOOL"`
}

// deprecated
type AttrValBool bool

// deprecated
func (thisV AttrValBool) isAttrVal() {}

// deprecated
func (thisV AttrValBool) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberBOOL())
}

// deprecated
func (thisV AttrValBool) ToAttributeValueMemberBOOL() types.AttributeValueMemberBOOL {
	return types.AttributeValueMemberBOOL{Value: bool(thisV)}
}

// deprecated
func (thisV AttrValBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValBoolStruct{Val: bool(thisV)})
}

// deprecated
func (thisP *AttrValBool) UnmarshalJSON(data []byte) error {
	var v AttrValBoolStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValBool(v.Val)
	return nil
}

// deprecated
type AttrValBinarySetStruct struct {
	Val [][]byte `json:"BS"`
}

// deprecated
type AttrValBinarySet map[string]struct{}

// deprecated
func (thisV AttrValBinarySet) isAttrVal() {}

// deprecated
func (thisV AttrValBinarySet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberBS())
}

// deprecated
func (thisV AttrValBinarySet) ToAttributeValueMemberBS() types.AttributeValueMemberBS {
	val := make([][]byte, 0, len(thisV))
	for k := range thisV {
		val = append(val, []byte(k))
	}
	return types.AttributeValueMemberBS{Value: val}
}

// deprecated
func (thisV AttrValBinarySet) MarshalJSON() ([]byte, error) {
	val := make([][]byte, 0, len(thisV))
	for k := range thisV {
		val = append(val, []byte(k))
	}
	return json.Marshal(AttrValBinarySetStruct{Val: val})
}

// deprecated
func (thisP *AttrValBinarySet) UnmarshalJSON(data []byte) error {
	var v AttrValBinarySetStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = make(map[string]struct{}, len(v.Val))
	for _, val := range v.Val {
		(*thisP)[string(val)] = struct{}{}
	}
	return nil
}

// deprecated
type AttrValListStruct struct {
	Val []AttrVal `json:"L"`
}

// deprecated
type AttrValList []AttrVal

// deprecated
func (thisV AttrValList) isAttrVal() {}

// deprecated
func (thisV AttrValList) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberL())
}

// deprecated
func (thisV AttrValList) ToAttributeValueMemberL() types.AttributeValueMemberL {
	val := make([]types.AttributeValue, 0, len(thisV))
	for _, v := range thisV {
		val = append(val, v.ToAttributeValue())
	}
	return types.AttributeValueMemberL{Value: val}
}

// deprecated
func (thisV AttrValList) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValListStruct{Val: thisV})
}

// deprecated
func (thisP *AttrValList) UnmarshalJSON(data []byte) error {
	var v AttrValListStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

// deprecated
type AttrValMapStruct struct {
	Val map[string]AttrVal `json:"M"`
}

// deprecated
type AttrValMap map[string]AttrVal

// deprecated
func (thisV AttrValMap) isAttrVal() {}

// deprecated
func (thisV AttrValMap) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberM())
}

// deprecated
func (thisV AttrValMap) ToAttributeValueMemberM() types.AttributeValueMemberM {
	val := make(map[string]types.AttributeValue, len(thisV))
	for k, v := range thisV {
		val[k] = v.ToAttributeValue()
	}
	return types.AttributeValueMemberM{Value: val}
}

// deprecated
func (thisV AttrValMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValMapStruct{Val: thisV})
}

// deprecated
func (thisP *AttrValMap) UnmarshalJSON(data []byte) error {
	var v AttrValMapStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

// deprecated
type AttrValNullStruct struct {
	Val bool `json:"NULL"`
}

// deprecated
type AttrValNull struct{}

// deprecated
func (thisV AttrValNull) isAttrVal() {}

// deprecated
func (thisV AttrValNull) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberNULL())
}

// deprecated
func (thisV AttrValNull) ToAttributeValueMemberNULL() types.AttributeValueMemberNULL {
	return types.AttributeValueMemberNULL{Value: true}
}

// deprecated
func (thisV AttrValNull) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValNullStruct{Val: true})
}

// deprecated
func (thisP *AttrValNull) UnmarshalJSON(data []byte) error {
	var v AttrValNullStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if !v.Val {
		return errors.New("null value must be true")
	}
	return nil
}

// deprecated
type AttrValNumberStruct struct {
	Val string `json:"N"`
}

// deprecated
type AttrValNumber string

// deprecated
func (thisV AttrValNumber) isAttrVal() {}

// deprecated
func (thisV AttrValNumber) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberN())
}

// deprecated
func (thisV AttrValNumber) ToAttributeValueMemberN() types.AttributeValueMemberN {
	return types.AttributeValueMemberN{Value: string(thisV)}
}

// deprecated
func (thisV AttrValNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValNumberStruct{Val: string(thisV)})
}

// deprecated
func (thisP *AttrValNumber) UnmarshalJSON(data []byte) error {
	var v AttrValNumberStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValNumber(v.Val)
	return nil
}

// deprecated
type AttrValNumberSetStruct struct {
	Val []string `json:"NS"`
}

// deprecated
type AttrValNumberSet map[string]struct{}

// deprecated
func (thisV AttrValNumberSet) isAttrVal() {}

// deprecated
func (thisV AttrValNumberSet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberNS())
}

// deprecated
func (thisV AttrValNumberSet) ToAttributeValueMemberNS() types.AttributeValueMemberNS {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return types.AttributeValueMemberNS{Value: val}
}

// deprecated
func (thisV AttrValNumberSet) MarshalJSON() ([]byte, error) {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return json.Marshal(AttrValNumberSetStruct{Val: val})
}

// deprecated
func (thisP *AttrValNumberSet) UnmarshalJSON(data []byte) error {
	var v AttrValNumberSetStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = make(map[string]struct{}, len(v.Val))
	for _, val := range v.Val {
		(*thisP)[val] = struct{}{}
	}
	return nil
}

// deprecated
type AttrValStringSetStruct struct {
	Val []string `json:"SS"`
}

// deprecated
type AttrValStringSet map[string]struct{}

// deprecated
func (thisV AttrValStringSet) isAttrVal() {}

// deprecated
func (thisV AttrValStringSet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberSS())
}

// deprecated
func (thisV AttrValStringSet) ToAttributeValueMemberSS() types.AttributeValueMemberSS {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return types.AttributeValueMemberSS{Value: val}
}

// deprecated
func (thisV AttrValStringSet) MarshalJSON() ([]byte, error) {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return json.Marshal(AttrValStringSetStruct{Val: val})
}

// deprecated
func (thisP *AttrValStringSet) UnmarshalJSON(data []byte) error {
	var v AttrValStringSetStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = make(map[string]struct{}, len(v.Val))
	for _, val := range v.Val {
		(*thisP)[val] = struct{}{}
	}
	return nil
}

// deprecated
type AttrValStringStruct struct {
	Val string `json:"S"`
}

// deprecated
type AttrValString string

// deprecated
func (thisV AttrValString) isAttrVal() {}

// deprecated
func (thisV AttrValString) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberS())
}

// deprecated
func (thisV AttrValString) ToAttributeValueMemberS() types.AttributeValueMemberS {
	return types.AttributeValueMemberS{Value: string(thisV)}
}

// deprecated
func (thisV AttrValString) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValStringStruct{Val: string(thisV)})
}

// deprecated
func (thisP *AttrValString) UnmarshalJSON(data []byte) error {
	var v AttrValStringStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValString(v.Val)
	return nil
}

type AttrKey = dynamox.AttrKey

const (
	AttrKeyPk = dynamox.AttrKeyPk
	AttrKeySk = dynamox.AttrKeySk
)

// deprecated
type Record = dynamox.Record

// deprecated
type StrAttr = dynamox.StringAttr

// deprecated
type StrSeqRecord = dynamox.StringPositionalRecord

// deprecated
func ParseStrSeqRecord(value string) (StrSeqRecord, error) {
	return dynamox.ParseStringPositionalRecord(value)
}
