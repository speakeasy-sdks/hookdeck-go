<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
)

func main() {
    s := hookdeck.New(
        hookdeck.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Attempts.GetAttempt(ctx, operations.GetAttemptRequest{
        ID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventAttempt != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->