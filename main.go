package main

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/oauth2/jws"
)

func main() {
	jws.Verify("maliciousToken", nil)

	maliciousHTML := strings.Repeat("<template>", 100000)

	html.Parse(strings.NewReader(maliciousHTML))
}
