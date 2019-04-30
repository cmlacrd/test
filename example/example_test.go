package test

import (
	"net/http"
	"bytes"
	"io"
	"log"
	"encoding/json"
	
)

//create a new http client 
func NewClient(host string) *HTTPClient{
	return &HTTPClient{
		Host: host,
	}
}

//send
func (ht *HTTPClient) Send() {
	client :=  &http.Client{
			Timeout: time.Second * 30,
	}

	resp, err := client.Do(ht.Request)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	ht.Result = DecodeTest(resp.Body)

}

//decoding json
func DecodeTest(body io.Reader) string {

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	s := buf.String()

	return s
}

//check function
func (ht *HTTPClient) IsValid() bool{
	
	if ht.Result != ht.Expected {
		return false
	}
	 
	return true
}

//example how to work all functions

func TestSignInSuccess(t *testing.T){	
	
	json := `{"email": "test@test.com",
			"pass":"123456"}`

	url := "http://localhost"

	req, err := http.NewRequest("POST", url, strings.NewReader(json))

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Logf("Error: %v", err)
	}

	client := NewClient(url)

	client.Request = req

	client.Send()

	if !client.IsValid(){
		t.Logf("Result: %v ", client.Result)
		t.Fail()
		return
	}
  
}
