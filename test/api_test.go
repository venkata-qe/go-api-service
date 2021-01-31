package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func Test_GetUs90210_StatusCodeShouldEqual200(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	assert.Equal(t, 201, resp.StatusCode())
}

func Test_GetUs90210_StatusCodeShouldEqual200_WithoutAssert(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
}

func Test_GetUs90210_ContentTypeShouldEqualApplicationJson(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
}

func Test_GetUs90210_CountryShouldEqualUnitedStates(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	response := LocationResponse{}

	err := json.Unmarshal(resp.Body(), &response)

	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, "United State", response.Country)
}

type LocationResponse struct {
	Country string `json:"country"`
}
