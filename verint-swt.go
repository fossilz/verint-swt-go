package verint_swt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func GenerateToken(apiKeyId string, apiKeyValue string, httpMethod string, webServiceEndpoint string) string{
	var u, _ = url.Parse(webServiceEndpoint)
	path := u.Path
	random := RandString(16)

	salt := base64url(random)

	loc, _ := time.LoadLocation("UTC")
	issuedAt := time.Now().In(loc).Format(time.RFC3339)

	headers := ""

	stringToSign := salt + "\n" + httpMethod + "\n" + path + "\n" + issuedAt + "\n" + headers + "\n"

	key, _ := base64.StdEncoding.DecodeString(debase64url(apiKeyValue))

	hash := hmac.New(sha256.New, key)

	hash.Write([]byte(stringToSign))

	b := hash.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(b)

	verintAuthId := "Vrnt-1-HMAC-SHA256"

	authHeaderValue := verintAuthId + " salt=" + salt + ",iat=" + issuedAt + ",kid=" + apiKeyId + ",sig=" + urlConvertBase64(signature)

	return authHeaderValue

}


func debase64url(input string) string {
	s := [3]string{"=","=","="}
	output0 := input + s[(len(input) + 3) % 4]
	output1 := strings.ReplaceAll(output0, "-", "+")
	output2 := strings.ReplaceAll(output1, "_", "/")
	return output2
}


func base64url(input string) string {
	base64String := base64.StdEncoding.EncodeToString([]byte(input))
	return urlConvertBase64(base64String)

}

func urlConvertBase64(input string) string {
	output1 := strings.ReplaceAll(input, "=","")
	output2 := strings.ReplaceAll(output1, "+", "-")
	return strings.ReplaceAll(output2, "/", "_")
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
