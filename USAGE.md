<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "",
				Username: "",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{
		Dir: operations.CreateDirArrayOfgetAttemptsQueryParam2(
			[]operations.GetAttemptsQueryParam2{
				operations.GetAttemptsQueryParam2Asc,
			},
		),
		EventID: operations.CreateEventIDArrayOfstr(
			[]string{
				"string",
			},
		),
		OrderBy: operations.CreateOrderByArrayOfqueryParam2(
			[]operations.QueryParam2{
				operations.QueryParam2CreatedAt,
			},
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->