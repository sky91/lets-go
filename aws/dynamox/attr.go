package dynamox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type AttributeValueWrapper struct {
	Val types.AttributeValue
}

func (thisV AttributeValueWrapper) MarshalJSON() ([]byte, error) {
	if thisV.Val == nil {
		return []byte("null"), nil
	}
	switch val := thisV.Val.(type) {
	case *types.AttributeValueMemberB:
		return json.Marshal(map[string][]byte{"B": val.Value})
	case *types.AttributeValueMemberBOOL:
		return json.Marshal(map[string]bool{"BOOL": val.Value})
	case *types.AttributeValueMemberBS:
		return json.Marshal(map[string][][]byte{"BS": val.Value})
	case *types.AttributeValueMemberL:
		list := make([]AttributeValueWrapper, len(val.Value))
		for i, v := range val.Value {
			list[i].Val = v
		}
		return json.Marshal(map[string][]AttributeValueWrapper{"L": list})
	case *types.AttributeValueMemberM:
		m := make(map[string]AttributeValueWrapper, len(val.Value))
		for k, v := range val.Value {
			m[k] = AttributeValueWrapper{Val: v}
		}
		return json.Marshal(map[string]map[string]AttributeValueWrapper{"M": m})
	case *types.AttributeValueMemberN:
		return json.Marshal(map[string]string{"N": val.Value})
	case *types.AttributeValueMemberNS:
		return json.Marshal(map[string][]string{"NS": val.Value})
	case *types.AttributeValueMemberNULL:
		return []byte(`{"NULL":true}`), nil
	case *types.AttributeValueMemberS:
		return json.Marshal(map[string]string{"S": val.Value})
	case *types.AttributeValueMemberSS:
		return json.Marshal(map[string][]string{"SS": val.Value})
	}
	return nil, fmt.Errorf("unknown type: %T", thisV.Val)
}

func (thisP *AttributeValueWrapper) UnmarshalJSON(b []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(b))
	token, err := decoder.Token()
	if err != nil {
		return errors.Wrap(err, "Token() error")
	}
	if token != json.Delim('{') {
		return fmt.Errorf("wrong Token: [%+v]", token)
	}

	token, err = decoder.Token()
	if err != nil {
		return errors.Wrap(err, "Token() error")
	}
	if token == json.Delim('}') {
		thisP.Val = nil
		return nil
	}

	k, ok := token.(string)
	if !ok {
		return fmt.Errorf("wrong Token: [%+v]", token)
	}

	switch k {
	case "B":
		var val []byte
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberB{Value: val}
	case "BOOL":
		var val bool
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberBOOL{Value: val}
	case "BS":
		var val [][]byte
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberBS{Value: val}
	case "L":
		var val []AttributeValueWrapper
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		list := make([]types.AttributeValue, len(val))
		for i, value := range val {
			list[i] = value.Val
		}
		thisP.Val = &types.AttributeValueMemberL{Value: list}
	case "M":
		var val map[string]AttributeValueWrapper
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		m := make(map[string]types.AttributeValue, len(val))
		for key, v := range val {
			m[key] = v.Val
		}
		thisP.Val = &types.AttributeValueMemberM{Value: m}
	case "NULL":
		thisP.Val = &types.AttributeValueMemberNULL{Value: true}
	case "N":
		var val string
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberN{Value: val}
	case "NS":
		var val []string
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberNS{Value: val}
	case "S":
		var val string
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberS{Value: val}
	case "SS":
		var val []string
		if err = decoder.Decode(&val); err != nil {
			return err
		}
		thisP.Val = &types.AttributeValueMemberSS{Value: val}
	default:
		return fmt.Errorf("invalid key: [%s]", k)
	}
	return nil
}

func (thisV AttributeValueWrapper) AsB() (*types.AttributeValueMemberB, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberB)
	return val, ok
}
func (thisV AttributeValueWrapper) AsBOOL() (*types.AttributeValueMemberBOOL, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberBOOL)
	return val, ok
}
func (thisV AttributeValueWrapper) AsBS() (*types.AttributeValueMemberBS, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberBS)
	return val, ok
}
func (thisV AttributeValueWrapper) AsL() (*types.AttributeValueMemberL, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberL)
	return val, ok
}
func (thisV AttributeValueWrapper) AsM() (*types.AttributeValueMemberM, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberM)
	return val, ok
}
func (thisV AttributeValueWrapper) AsNULL() (*types.AttributeValueMemberNULL, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberNULL)
	return val, ok
}
func (thisV AttributeValueWrapper) AsN() (*types.AttributeValueMemberN, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberN)
	return val, ok
}
func (thisV AttributeValueWrapper) AsNS() (*types.AttributeValueMemberNS, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberNS)
	return val, ok
}
func (thisV AttributeValueWrapper) AsS() (*types.AttributeValueMemberS, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberS)
	return val, ok
}
func (thisV AttributeValueWrapper) AsSS() (*types.AttributeValueMemberSS, bool) {
	val, ok := thisV.Val.(*types.AttributeValueMemberSS)
	return val, ok
}

type AttrKey string

const (
	AttrKeyPk    AttrKey = "PK"
	AttrKeySk    AttrKey = "SK"
	AttrKeyClass AttrKey = "C"
)
