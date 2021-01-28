package accessor

import (
	"encoding/json"
	"errors"
	"fmt"
	_extstore "github.com/bilalkocoglu/eureka-client/store"
	_const "github.com/imminoglobulin/e-commerce-backend/file-service/pkg/const"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/helper"
	"github.com/rs/zerolog/log"
)

type CustomerResource struct {
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Mail        string   `json:"mail"`
	Password    string   `json:"password"`
	Authorities []string `json:"authorities"`
}

type CustomerServiceAccessor struct {
}

func (c CustomerServiceAccessor) GetCustomerById(customerId string) (*CustomerResource, error) {
	path := "/api/customer/by-id"
	url := ""
	if _extstore.RegisteredServices != nil {
		url = _extstore.RegisteredServices.GetServiceUrl(_const.CUSTOMER_SERVICE)
	}

	if url == "" {
		return nil, errors.New(fmt.Sprintf("%s url not found registry", _const.CUSTOMER_SERVICE))
	}

	fullUrl := url + path
	log.Info().Msg("GetCustomerById -> service url: " + fullUrl)

	params := make(map[string]string)
	params["id"] = customerId

	err, response := helper.MakeGetCall(fullUrl, nil, params, true)

	if err != nil {
		return nil, err
	}

	var customer CustomerResource
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&customer)

	if err != nil {
		log.Err(err)
		return nil, err
	} else {
		return &customer, nil
	}
}
