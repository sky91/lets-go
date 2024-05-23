package dynamodbx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sky91/lets-go/gox"
)

type IsAttrVal interface {
	isAttrVal()
}

type AttrVal struct {
	Val IsAttrVal
}

func (thisV AttrVal) ToAttributeValue() types.AttributeValue {
	return thisV.Val.(interface{ ToAttributeValue() types.AttributeValue }).ToAttributeValue()
}

func (thisV AttrVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(thisV.Val)
}

func (thisP *AttrVal) UnmarshalJSON(data []byte) error {
	var val map[string]struct{}
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

type AttrValBinaryStruct struct {
	Val []byte `json:"B"`
}
type AttrValBinary []byte

func (thisV AttrValBinary) isAttrVal() {}
func (thisV AttrValBinary) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberB())
}
func (thisV AttrValBinary) ToAttributeValueMemberB() types.AttributeValueMemberB {
	return types.AttributeValueMemberB{Value: thisV}
}
func (thisV AttrValBinary) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValBinaryStruct{Val: thisV})
}
func (thisP *AttrValBinary) UnmarshalJSON(data []byte) error {
	var v AttrValBinaryStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

type AttrValBoolStruct struct {
	Val bool `json:"BOOL"`
}
type AttrValBool bool

func (thisV AttrValBool) isAttrVal() {}
func (thisV AttrValBool) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberBOOL())
}
func (thisV AttrValBool) ToAttributeValueMemberBOOL() types.AttributeValueMemberBOOL {
	return types.AttributeValueMemberBOOL{Value: bool(thisV)}
}
func (thisV AttrValBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValBoolStruct{Val: bool(thisV)})
}
func (thisP *AttrValBool) UnmarshalJSON(data []byte) error {
	var v AttrValBoolStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValBool(v.Val)
	return nil
}

type AttrValBinarySetStruct struct {
	Val [][]byte `json:"BS"`
}
type AttrValBinarySet map[string]struct{}

func (thisV AttrValBinarySet) isAttrVal() {}
func (thisV AttrValBinarySet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberBS())
}
func (thisV AttrValBinarySet) ToAttributeValueMemberBS() types.AttributeValueMemberBS {
	val := make([][]byte, 0, len(thisV))
	for k := range thisV {
		val = append(val, []byte(k))
	}
	return types.AttributeValueMemberBS{Value: val}
}
func (thisV AttrValBinarySet) MarshalJSON() ([]byte, error) {
	val := make([][]byte, 0, len(thisV))
	for k := range thisV {
		val = append(val, []byte(k))
	}
	return json.Marshal(AttrValBinarySetStruct{Val: val})
}
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

type AttrValListStruct struct {
	Val []AttrVal `json:"L"`
}
type AttrValList []AttrVal

func (thisV AttrValList) isAttrVal() {}
func (thisV AttrValList) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberL())
}
func (thisV AttrValList) ToAttributeValueMemberL() types.AttributeValueMemberL {
	val := make([]types.AttributeValue, 0, len(thisV))
	for _, v := range thisV {
		val = append(val, v.ToAttributeValue())
	}
	return types.AttributeValueMemberL{Value: val}
}
func (thisV AttrValList) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValListStruct{Val: thisV})
}
func (thisP *AttrValList) UnmarshalJSON(data []byte) error {
	var v AttrValListStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

type AttrValMapStruct struct {
	Val map[string]AttrVal `json:"M"`
}
type AttrValMap map[string]AttrVal

func (thisV AttrValMap) isAttrVal() {}
func (thisV AttrValMap) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberM())
}
func (thisV AttrValMap) ToAttributeValueMemberM() types.AttributeValueMemberM {
	val := make(map[string]types.AttributeValue, len(thisV))
	for k, v := range thisV {
		val[k] = v.ToAttributeValue()
	}
	return types.AttributeValueMemberM{Value: val}
}
func (thisV AttrValMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValMapStruct{Val: thisV})
}
func (thisP *AttrValMap) UnmarshalJSON(data []byte) error {
	var v AttrValMapStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = v.Val
	return nil
}

type AttrValNullStruct struct {
	Val bool `json:"NULL"`
}
type AttrValNull struct{}

func (thisV AttrValNull) isAttrVal() {}
func (thisV AttrValNull) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberNULL())
}
func (thisV AttrValNull) ToAttributeValueMemberNULL() types.AttributeValueMemberNULL {
	return types.AttributeValueMemberNULL{Value: true}
}
func (thisV AttrValNull) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValNullStruct{Val: true})
}
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

type AttrValNumberStruct struct {
	Val string `json:"N"`
}
type AttrValNumber string

func (thisV AttrValNumber) isAttrVal() {}
func (thisV AttrValNumber) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberN())
}
func (thisV AttrValNumber) ToAttributeValueMemberN() types.AttributeValueMemberN {
	return types.AttributeValueMemberN{Value: string(thisV)}
}
func (thisV AttrValNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValNumberStruct{Val: string(thisV)})
}
func (thisP *AttrValNumber) UnmarshalJSON(data []byte) error {
	var v AttrValNumberStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValNumber(v.Val)
	return nil
}

type AttrValNumberSetStruct struct {
	Val []string `json:"NS"`
}
type AttrValNumberSet map[string]struct{}

func (thisV AttrValNumberSet) isAttrVal() {}
func (thisV AttrValNumberSet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberNS())
}
func (thisV AttrValNumberSet) ToAttributeValueMemberNS() types.AttributeValueMemberNS {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return types.AttributeValueMemberNS{Value: val}
}
func (thisV AttrValNumberSet) MarshalJSON() ([]byte, error) {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return json.Marshal(AttrValNumberSetStruct{Val: val})
}
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

type AttrValStringSetStruct struct {
	Val []string `json:"SS"`
}
type AttrValStringSet map[string]struct{}

func (thisV AttrValStringSet) isAttrVal() {}
func (thisV AttrValStringSet) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberSS())
}
func (thisV AttrValStringSet) ToAttributeValueMemberSS() types.AttributeValueMemberSS {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return types.AttributeValueMemberSS{Value: val}
}
func (thisV AttrValStringSet) MarshalJSON() ([]byte, error) {
	val := make([]string, 0, len(thisV))
	for k := range thisV {
		val = append(val, k)
	}
	return json.Marshal(AttrValStringSetStruct{Val: val})
}
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

type AttrValStringStruct struct {
	Val string `json:"S"`
}
type AttrValString string

func (thisV AttrValString) isAttrVal() {}
func (thisV AttrValString) ToAttributeValue() types.AttributeValue {
	return gox.New(thisV.ToAttributeValueMemberS())
}
func (thisV AttrValString) ToAttributeValueMemberS() types.AttributeValueMemberS {
	return types.AttributeValueMemberS{Value: string(thisV)}
}
func (thisV AttrValString) MarshalJSON() ([]byte, error) {
	return json.Marshal(AttrValStringStruct{Val: string(thisV)})
}
func (thisP *AttrValString) UnmarshalJSON(data []byte) error {
	var v AttrValStringStruct
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*thisP = AttrValString(v.Val)
	return nil
}

type AttrKey string

type Record map[AttrKey]AttrVal
