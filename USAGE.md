<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->