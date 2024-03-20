package rsa

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type RSAMethod struct {
	h          crypto.Hash
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSAMethod(h crypto.Hash) *RSAMethod {
	var nRSA = &RSAMethod{}
	nRSA.h = h
	return nRSA
}

func (this *RSAMethod) Sign(data []byte) ([]byte, error) {
	var h = this.h.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	digest := h.Sum(nil)
	return rsa.SignPKCS1v15(nil, this.privateKey, this.h, digest)
}

func (this *RSAMethod) Verify(data []byte, signature []byte) error {
	var err error
	var h = this.h.New()
	if _, err = h.Write(data); err != nil {
		return err
	}
	var hashed = h.Sum(nil)
	if signature, err = base64.StdEncoding.DecodeString(string(signature)); err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(this.publicKey, this.h, hashed, signature)
}

func (this *RSAMethod) SetPrivateKey(data []byte, parseFunc func(der []byte) (key any, err error)) error {
	n, err := rsaKeyDecode(data)
	if err != nil {
		return err
	}
	privateKey, err := parseFunc(n)
	if err != nil {
		return err
	}
	this.privateKey = privateKey.(*rsa.PrivateKey)
	return nil
}

func (this *RSAMethod) SetPublicKey(data []byte, parseFunc func(der []byte) (value any, err error)) error {
	n, err := rsaKeyDecode(data)
	if err != nil {
		return err
	}
	publicKey, err := parseFunc(n)
	if err != nil {
		return err
	}
	this.publicKey = publicKey.(*rsa.PublicKey)
	return nil
}

func rsaKeyDecode(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("invalid private key")
	}
	if data[0] == '-' {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, errors.New("invalid private key")
		}
		return block.Bytes, nil
	}
	return base64decode(data)
}

func base64decode(data []byte) ([]byte, error) {
	var dBuf = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dBuf, data)
	return dBuf[:n], err
}
