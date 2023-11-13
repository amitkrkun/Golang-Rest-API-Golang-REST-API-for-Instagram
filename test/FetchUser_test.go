package unittesting

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	user "nimeshjohari02.com/restapi/user"
)
func JSONBytesEqual(a, b []byte) (bool, error) {
    var j, j2 interface{}
    if err := json.Unmarshal(a, &j); err != nil {
        return false, err
    }
    if err := json.Unmarshal(b, &j2); err != nil {
        return false, err
    }
    return reflect.DeepEqual(j2, j), nil
}
func TestFetchUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/getUserById", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "7a62cbbc-7f70-48db-b740-7be5fef57328")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.GetUserById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected :=`{"Email":"\"nimeshjohari95@gmail.com\"","Password":"$2a$14$C0pdzLah2gIHdttH2kbiOOf55mwHgEdlewV1Jlt2nyK8E8Jo.PSga","Posts":null,"_id":"6161ad1923c42f4188930f8a","id":"7a62cbbc-7f70-48db-b740-7be5fef57328","name":"\"NJ\""}`
	received := rr.Body.String()
	same,err:=JSONBytesEqual([]byte(expected),[]byte(received));
	if err!=nil{
		t.Errorf("Error in comparing JSON") 
	}	
	if !same {
		t.Errorf("handler returned unexpected body: got %v want %v",
			received, expected)
	}
}

