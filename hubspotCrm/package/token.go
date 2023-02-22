package crm_integration

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func TokenRefresh() string {

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("POST")
	req.SetRequestURI("https://Companies.zoho.com/oauth/v2/token")
	req.Header.SetContentType("application/x-www-form-urlencoded")

	req.SetBody([]byte("refresh_token=1000.1101451cb48320dfcae15ac2cc2a0a4d.675f0a34d48914ed090db90d612d524e&client_id=1000.V8PD2J4VLP14NZ2R2S3PPYKIXRVSHV&client_secret=4e8ce48ab1201ce2a10569bbae8c15920b27314410&grant_type=refresh_token"))
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
	var token Token
	json.Unmarshal(body, &token)
	return token.Access_token
}
