package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("1. HMAC for hashing: ", hashString("password1234!"))
	fmt.Println("2. encoding base 64: ", base64enc())

	http.HandleFunc("/some-req", func(w http.ResponseWriter, r *http.Request) {})
}

// 1. HMAC - when we want to save data on user machine and make sure user has not temper with that data
// as we use secret o create the hash only we can modify it
func hashString(input string) string {
	hash := hmac.New(sha256.New, []byte("somekey"))
	io.WriteString(hash, input)                // we write input into the hash
	output := fmt.Sprintf("%x", hash.Sum(nil)) // write the output as hexadecimal
	return output
}

// 2. Base64 encoding
// cookie values that are  breaking the RFC 2616 standard should be encoding. As some special characters can cause issues with the cookies
func base64enc() string {
	input := "I'm sick of following my dreams, man. I'm just going to ask where they're going and hook up with ’em later. —Mitch Hedberg"
	// encodingStandard := "ABCDEFGHIJKLMNOPQRSTUVWXYabcd" // we can provide our encoding standard or use the default one
	// encoded := base64.NewEncoding(encodingStandard).EncodeToString([]byte(input))
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}

// 3. web storage
// those are separated into:
// a) cookies - js/bakend
// b) session storage - only js
// c) local storage - only js

// 4. context - when request comes in, context can be crated and we add something like session variable, and pass this thru the system. Params and limits can be set on the specific context. We can cancel processes related to the context. context is used in app engine

func singleRequst(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()                          // create a context from request
	ctx = context.WithValue(ctx, "userID", 784) // we should be careful what we put in context. Only things related to the request. Putting a lot of things and storing them in context is a bad idea
}

func anotherRequest(ctx context.Context) int {
	uid := ctx.Value("userID").(int) // we can restrive passed values. PS. .(int) is for asserting that the value is a int
	//ctx, cancel := context.WithTimeout(ctx, 1*time.Second) we could shut down the process after some time
	// defer cancel()
	return uid
}

// 5. https and tls
// whe using ssl cert we want to use http.ListenAndServeTLS() which will accept cert and a private key using for communication encrytption
