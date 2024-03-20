package rsa

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var privateKey = []byte("MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCIqbQ3khuapPW+3T76EwL/C3ZuThVfs+OPktFhse6AFMzmdcsk9tzhRsKPp+Gh+dygRuGsV00XIl/lfsxBzFxqX0ZyQEngarQqUal1VLGAvGekzOMyw9WVufN4z1F7FyNMSihc9XR3+JHJX1THwTEek6dJGm82LtvzDEA8HvaEBpZXcswEyIg/XGtj32qha/9dBhsCIdQE44+fx5LAks9bS7uGk6GbFBFYZhwAL+Q2946SAHs7dCQlEEvIy4cZ2FrBEJN5KDMW/RKZOQMNCGxYl3eauLPF9L9oocaBJQ1nGCx0PNdVhmO0RLkIIe4VagYJA16hnIAD7oWdkw7ic/upAgMBAAECggEBAId8Ftpw00UA/VB3bxPk/Td0WbkJWE9Eu/l+iG3eRKBd/rULWpYO1vPPftuEiEBYwc1z+A8vjZG91mxixIUG4SfjxKdH7PW7U1oRE3rqt/70yZusNCID3B9P9nxyrEjnq2raqiFehlTZ7U5CYH6YnOW/ZD0pI9AHrK7MGnKAD8zad5i73RoY19enRvKqcmgy42C4U7WebuiK4JB/PK8FZFckPOZCRJjKIF/oVDsNsGKCAEAIGsx6+rnfUykWtylakdciSjmquFwOkj85yNLn6g4a27FUPtjD9FduRyYbcPhyEyeICF0GrBPF1tnKGPLbHxT4tn1o6ZUOMnWMZHDvtsECgYEAxRrws9Rttujhz+O3K57icSPffKXJ5BvnUwUMsOVtao8N9Lm5zdedD2v6EXTDubFFiXBea28WniBs6N2TRdKPh+zSF5w1eJw7YfGx2wOfRyezRpzYsm1Z3bzVILrffdSBclqThs1VdwvqynD3afdHIFKct0aPeSNk92joCCN31/0CgYEAsX9h95KJLYkbBlGMKhHByopt+ybT5HjPUIdAKlCF6hVaRYmSb5/Jrud8dpYRs8fle8Dwriyz/U6YRWXEXbXasQACxvZx9MBcGosET7glXuoLMR32w07qvH6CP0gU9WsUo+cm2TQ87gDfCTYv6UQGTMVb0AJ+he58gfHpGYL/1B0CgYEAuzB6qE52Y9+HlQeOTb73Drwi4x4QPrLBXfbNtrIs16ZEesZnzzWV06+YLjvWJeRVRdGR2jYGJOZmiEDmPMlbxpsCa6nOnlzBiKMhZf5qpgjuGYGbythPIGbVgd/3oGjRBuk+cidPo1+N+/SNA/Xzdp7+dpBscX9LxCu9MP1+M4ECgYBdTCeiilOwq7CY0aRrOIo0fC7zJKWcIiRWn8EKfOjm8fJROs7N/Z91YBf7/UWwyhHFytS3uKejLfy7/ZIJ79zTzL2o90FO9Q10pj3N6W07Rqo7VxSt9H1ONvEfcjQSDAb9YDL7WInuUGbr4J+ubSMil1p7K9R1cbXgV9e+oSz3wQKBgQC2i3AUSE170HYAcWMG+pORReUZXzGlJdIA3/XddiXSL/cj+KWbRmfYM7GZAtedNywIhILqzm0A2h7k1wjoblC/asUn/AFQX3L0cXBJXtYgdLBL2p91WyH06WXrR0Wi9DDPHrv2AvituZGP6eVsS/X2uVa9GU00Q4nnghx0EKV39A==")
var publicKey = []byte("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiKm0N5IbmqT1vt0++hMC/wt2bk4VX7Pjj5LRYbHugBTM5nXLJPbc4UbCj6fhofncoEbhrFdNFyJf5X7MQcxcal9GckBJ4Gq0KlGpdVSxgLxnpMzjMsPVlbnzeM9RexcjTEooXPV0d/iRyV9Ux8ExHpOnSRpvNi7b8wxAPB72hAaWV3LMBMiIP1xrY99qoWv/XQYbAiHUBOOPn8eSwJLPW0u7hpOhmxQRWGYcAC/kNveOkgB7O3QkJRBLyMuHGdhawRCTeSgzFv0SmTkDDQhsWJd3mrizxfS/aKHGgSUNZxgsdDzXVYZjtES5CCHuFWoGCQNeoZyAA+6FnZMO4nP7qQIDAQAB")

// openssl genrsa -out private.key 2048
// openssl rsa -in private.key -pubout -out public.key
func Test_signRSA(t *testing.T) {
	paramStr := `app_id=20135234674&biz_content={"total_amount":"10.08", "buyer_id":"2088123456781234", "discount_amount":""}&method=alipay.trade.create&sign_type=RSA2&timestamp=2019-01-01 08:09:33`
	rsa := NewRSAMethod(crypto.SHA256)
	err := rsa.SetPrivateKey(privateKey, x509.ParsePKCS8PrivateKey)
	assert.NoError(t, err)
	dataByte, err := rsa.Sign([]byte(paramStr))
	assert.NoError(t, err)
	data := base64.StdEncoding.EncodeToString(dataByte)
	fmt.Println(data)
}

func Test_verifyRSA(t *testing.T) {
	signValue := []byte("ehTrzV9y5v29/gRs0hcFgWQi3DnMLPVUJTrPOb7ILEA8yZPly3YDUUnUAkoi4cPJHAE1gjY4g3nqljpbuBvwNGZJOPCnDOvrZ7DD3YLOrwHrqt1CjXtyfM7i6dNwyiqL0Jpf8D5N3ou2cE4thIWHJMSg10v7ef54SVulkiS2jEREpZxHFTmnX7Jjm5FKI/3m+/ne81gjPiZO5527dFDQpBZE7EWHP1qz76J8GOQBiqcbCyZHMWx9m3njHpTtfINpqCT/dF8+W3dextnKs5f69nC+5bXh80qTiPB9Nhe+1cq8QCq21Ii7YL0zpSkbKpbQUVhk51Hfnaz51p2fVuhjcQ==")
	paramStr := `app_id=20135234674&biz_content={"total_amount":"10.08", "buyer_id":"2088123456781234", "discount_amount":""}&method=alipay.trade.create&sign_type=RSA2&timestamp=2019-01-01 08:09:33`
	rsa := NewRSAMethod(crypto.SHA256)
	err := rsa.SetPublicKey(publicKey, x509.ParsePKIXPublicKey)
	assert.NoError(t, err)
	err = rsa.Verify([]byte(paramStr), signValue)
	assert.NoError(t, err)
}
