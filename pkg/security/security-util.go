package security

import (
	"errors"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/accessor"
	_const "github.com/imminoglobulin/e-commerce-backend/file-service/pkg/const"
	"net/http"
)

var customerAccessor = new(accessor.CustomerServiceAccessor)

func GetCurrentCustomerId(req http.Request) string {
	return req.Header.Get(_const.CUSTOMER_ID_HEADER)
}

func GetCurrentTenantId(req http.Request) string {
	return req.Header.Get(_const.TENANT_ID_HEADER)
}

func GetCurrentCustomer(req http.Request) (*accessor.CustomerResource, error) {
	customerId := GetCurrentCustomerId(req)

	if customerId == "" {
		return nil, errors.New("Authentication not found.")
	}

	return customerAccessor.GetCustomerById(customerId)
}
