package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/samber/lo"
	"golang.org/x/crypto/sha3"
	"io"
)

type TokenAlg struct {
	tokenCipherBlock cipher.Block
	macShake         sha3.ShakeHash
}

func (thisP *TokenAlg) Init(rootKey []byte) {
	tokenAesKeyShake := sha3.NewCShake256(nil, []byte("token_aes_key"))
	lo.Must0(tokenAesKeyShake.Write(rootKey))
	aesKey := make([]byte, 16)
	lo.Must0(io.ReadFull(tokenAesKeyShake, aesKey))
	thisP.tokenCipherBlock = lo.Must(aes.NewCipher(aesKey))

	thisP.macShake = sha3.NewCShake256(nil, []byte("mac_shake"))
	lo.Must0(thisP.macShake.Write(rootKey))
}

func (thisP *TokenAlg) EncodeToken(tokenData []byte) []byte {
	return thisP.encodeTokenV0(tokenData)
}

func (thisP *TokenAlg) DecodeToken(encodedToken []byte) ([]byte, bool) {
	if len(encodedToken) < 1 {
		return nil, false
	}
	switch encodedToken[0] {
	case 0:
		return thisP.decodeTokenV0(encodedToken)
	}
	return nil, false
}

func (thisP *TokenAlg) encodeTokenV0(tokenData []byte) []byte {
	padLen := thisP.tokenCipherBlock.BlockSize() - (len(tokenData) % thisP.tokenCipherBlock.BlockSize())
	encodedToken := make([]byte, 9+thisP.tokenCipherBlock.BlockSize()+len(tokenData)+padLen)
	encodedToken[0] = 0
	macIvEnc := encodedToken[1:]
	mac := macIvEnc[:8]
	ivEnc := macIvEnc[len(mac):]

	iv := ivEnc[:thisP.tokenCipherBlock.BlockSize()]
	lo.Must0(rand.Read(iv))

	enc := ivEnc[len(iv):]
	copy(enc, tokenData)
	for i := len(tokenData); i < len(enc); i++ {
		enc[i] = byte(padLen)
	}
	cipher.NewCBCEncrypter(thisP.tokenCipherBlock, iv).CryptBlocks(enc, enc)

	shake := thisP.macShake.Clone()
	lo.Must0(shake.Write(ivEnc))
	lo.Must0(io.ReadFull(shake, mac))
	return encodedToken
}

func (thisP *TokenAlg) decodeTokenV0(encodedToken []byte) ([]byte, bool) {
	if len(encodedToken) < 9 {
		return nil, false
	}
	if encodedToken[0] != 0 {
		return nil, false
	}

	macIvEnc := encodedToken[1:]
	mac := macIvEnc[:8]
	ivEnc := macIvEnc[len(mac):]

	shake := thisP.macShake.Clone()
	lo.Must0(shake.Write(ivEnc))
	calcMac := [8]byte{}
	lo.Must0(io.ReadFull(shake, calcMac[:]))
	if !bytes.Equal(mac, calcMac[:]) {
		return nil, false
	}

	if len(ivEnc) < thisP.tokenCipherBlock.BlockSize() {
		return nil, false
	}

	iv := ivEnc[:thisP.tokenCipherBlock.BlockSize()]
	enc := ivEnc[len(iv):]
	if len(enc)%thisP.tokenCipherBlock.BlockSize() != 0 {
		return nil, false
	}

	cipher.NewCBCDecrypter(thisP.tokenCipherBlock, iv).CryptBlocks(enc, enc)
	padLen := int(enc[len(enc)-1])
	if padLen > len(enc) {
		return nil, false
	}
	return enc[:len(enc)-padLen], true
}
