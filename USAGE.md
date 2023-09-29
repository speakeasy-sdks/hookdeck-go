<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    id := "female"

    ctx := context.Background()
    res, err := s.Attempt.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventAttempt != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->