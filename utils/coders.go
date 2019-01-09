package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/clbanning/anyxml"
)

func DecodeMessageRes(r *http.Response, data interface{}) error {
	if r.Header.Get("Content-type") == "application/json" {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		fmt.Println(data)
		if err != nil {
			fmt.Println(err)
		}
	} else if r.Header.Get("Content-type") == "application/xml" {
		decoder := xml.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return errors.New("Invalid format")
	}
	return nil
}

func DecodeMessageReq(r *http.Request, data interface{}) error {
	if r.Header.Get("Content-type") == "application/json" {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		fmt.Println(data)
		if err != nil {
			fmt.Println(err)
		}
	} else if r.Header.Get("Content-type") == "application/xml" {
		decoder := xml.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return errors.New("Invalid format")
	}
	return nil
}

func EncodeResponse(data interface{}, xlmJson string) ([]byte, error) {
	switch xlmJson {
	case "json":
		return json.Marshal(data)
	case "xml":
		anyxml.XMLEscapeChars(false)
		return anyxml.XmlIndent(data, "", "  ", "doc")
	default:
		return make([]byte, 0), errors.New("Invalid encode type")
	}
}
