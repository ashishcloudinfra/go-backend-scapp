package helpers

import (
	"regexp"

	"server.simplifycontrol.com/types"
)

var DEFAULT_UUID = "00000000-0000-0000-0000-000000000000"

// jwtBypassRules is a slice of rules that define which routes and methods bypass JWT authentication.
var JwtBypassRules = []types.BypassRule{
	{
		Pattern: regexp.MustCompile(`^/login`),
		Methods: []string{"POST"}, // only bypass JWT check for POST /login
	},
	{
		Pattern: regexp.MustCompile(`^/companyTypes`),
		Methods: []string{"GET"},
	},
	{
		Pattern: regexp.MustCompile(`^/signup`),
		Methods: []string{"POST"},
	},
	{
		Pattern: regexp.MustCompile(`^/price`),
		Methods: []string{}, // empty means bypass for any method on /price
	},
	{
		Pattern: regexp.MustCompile(`^/health`),
		Methods: []string{"GET"},
	},
}
