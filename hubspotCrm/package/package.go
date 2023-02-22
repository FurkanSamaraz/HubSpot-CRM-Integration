package crm_integration

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type HTTP_METHODS string

// * Enum string for HTTTP Request Methods
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type CRM_BASE struct { // https://crm.zoho.com/crm/private/json/Emails/getMyRecords?authtoken=Auth Token&scope=crmapi
	AuthToken     string       `json:"authToken"`
	Scope         string       `json:"scope"`
	URL           string       `json:"url"` // https://crm.zoho.com/crm/private/json/<moduleName>/<actionName>
	Fields        string       `json:"id"`
	RequestMethod HTTP_METHODS `json:"requestMethod"`
	Ctx           *fiber.Client
}

type CRM_REQUEST struct {
	Body       []byte
	ModuleName string
	ActionName string
}

type Token struct {
	Access_token string `json:"access_token"`
	Api_domain   string `json:"api_domain"`
	Token_type   string `json:"token_type"`
	Expires_in   string `json:"expires_in"`
}

func (cb *CRM_BASE) Request(requestBody CRM_REQUEST) ([]byte, error) {

	// Fetch URL
	requestURLs := fmt.Sprintf("%s/%s/%s%s", cb.GetURL(), cb.Scope, requestBody.ModuleName, cb.Fields)
	var err error
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(string(cb.RequestMethod))
	req.SetRequestURI(requestURLs)
	//cb.AuthToken = "Zoho-oauthtoken 1000.7fac4199aae188b2aded5b6e8f24f33e.2312147c72c98b9e2df4cdeeae669897"
	req.Header.Add("Authorization", "Zoho-oauthtoken "+TokenRefresh())

	//fmt.Println(requestURLs)
	switch cb.GetRequestMethod() {
	case GET:

	case DELETE:

	case POST, PUT:
		req.SetBody(requestBody.Body)

	default:
		panic("Invalid Request Method")

	}
	if err := a.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := a.Bytes() // ...
	if errs != nil {
		for _, err := range errs {
			log.Println(err)
		}
		panic("Getting Multiple Errors")
	}

	return body, err
}

func (cb *CRM_BASE) GetURL() string {
	return cb.URL
}

func (cb *CRM_BASE) GetAuthToken() string {
	return cb.AuthToken
}

func (cb *CRM_BASE) GetScope() string {
	return cb.Scope
}

func (cb *CRM_BASE) GetRequestMethod() HTTP_METHODS {
	return cb.RequestMethod
}
