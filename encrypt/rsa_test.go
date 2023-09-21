package encrypt

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"testing"
)

var PrivateKey = []byte("MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCjqCF4amX74fOhVRLSPdYXELQ7U5I9b4iOV/E+biweiOC4cTm5QWaiaN5nj8bk35NC7vvRhN0nptdufn0nNO8BmFrRY1t8yAklVXik3B7eAG9aB1azktRkOiB16bHQwXQt5LmvI+4TLis8X8w726zRozRFzMqefSX2CK5GFNnim6AA0CPB64GfGgE0VR/xHfvfjrpL3Zz8ILjeKRQCgw9MK0/PIIF8f6vk0/XvJ4eulVMG+vUlp81AsnhDthNHDv93aSgdKBlfwe5h/HkWIdSCLJVkpifbM5p5MXMJPOE5+wKkt/PbG/X+U1IzE3o4TMTAVKLuQaworNyLI4UcPNuRAgMBAAECggEAB5lP7IpsL9Te/he4JwDuBuYkqDiKKsYZ/hoBPpfY/pX2cTn4pyxtOo9J/F8FqizjZpE4KhNUPXoWuK+w2fgwPM+5L83AOrwDCLO+RVFfMYmtoe2DThU8PWchiLgQJJRG87aSFJrDfCqBoW0FS6Y4kQdV7wPCSLeg4pqmlOR09XuXHN2NPT/rO8Lr8TnpPrdfkq5/a6AmwYn2LYo80KyrkJ9TMD9VTjJvcj0W6xY3uKQ0cSuIuxJX8gGLWgjJv02rqprkRLdJyBfqdxJrSoKV0ZusS/iCJlXg6NU6fWyTWCRcQdiAevaI41D64hAJJgW5HYZM7qMEbO4dr5UfXzrG7QKBgQDx/yZnkpLHQ6KaE2d+aQKFXbPrHBbUepW4JjAa5DXo4dTQxraVTsx/DQpZ/jTEeAWMy5ZAabbCff74OAWMHXc8tCprN7Vwn02mI+HDSOxGsy55U5JRBKl8w2qzI8v4AE6HyaDFREiPthx6zzytOi1PP89vyA1o2JAVuZOmf9WkjwKBgQCtIHtfPaKNH6OCtcMbYTA309+DCnW8gHkbg1uf81sCzgnJsJ/4VRfn4sNe2Yyl3U79OEhdBbcTpg5+/MmIP5lcm9OR8KyUyA0YSnsW+Y4HRb1BXXtOT+7hVK5+Y+ESjj8qKQNM64p5OSqgSGuaw38vPBb0T67qZQO9K1eXk+TN3wKBgBv7Xtt3DVXRMkoSmFL7nXkCXr1+3/ztt8Y1VDan0Lv+8Pa7I32cQPUn8tx5EmztW+bt24/TKXjPzN2yzKFo/tKcGpOPO4GsSfQ+Derg0cFTErqguTh5C4gLqJjSONGSZ4I4BEkLHkDP0/c3Y1K8eWTCgh/wx+wprm90p+gGvBNtAoGAMbwSFRM0vkvngiZLLXNnEbKpFBEOL5/Mqs26paGYdJ7SCwHVgtaXLoNjUr02fXOtPGtNxoNcy6U5ptth3eU/Xm5ZgiRcv8UUlfTXlYsNdSNgsgVz5dRqsIrOMfrpbpY0qRztGMzVk+uLRk5nsycUQ3KEuZymiCmKwG5SFHZlFYsCgYAlYKIu7+boZW5fsO2ACuC25nbnVIzQXloKzIh7tllEZQZjaRBGjouMVGA0sZFc5O/z5xurXgU92zVzu88VTZcKgZxwzAxe+RvR7wiu8DOMKUWE/qR5/kqaXymxSYarfPv7E82Sce/o0QxoSUlQvhVpxmARQQZ2jUtAgFz/PZEWbw==")

// openssl genrsa -out private.key 2048
// openssl rsa -in private.key -pubout -out public.key
func TestName(t *testing.T) {
	paramerStr := `app_id=20135234674&biz_content={"total_amount":"10.08", "buyer_id":"2088123456781234", "discount_amount":""}&method=alipay.trade.create&sign_type=RSA2&timestamp=2019-01-01 08:09:33`
	rsa := NewRSAMethod(crypto.SHA256)
	rsa.SetPrivateKey(PrivateKey, x509.ParsePKCS8PrivateKey)
	dataByte, err := rsa.Sign([]byte(paramerStr))
	fmt.Println(err)
	data := base64.StdEncoding.EncodeToString(dataByte)
	fmt.Println(data)
}
