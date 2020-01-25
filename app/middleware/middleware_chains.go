package middleware

import (
	"github.com/gorilla/csrf"
	"github.com/windstep/go-web-skeleton/app/middleware/chain"
)

var ApiMiddlewareChain = chain.New(
	[]chain.Constructor{},
	[]chain.ConstructorFunc{
		LoggerMiddleware,
		JsonMiddleware,
	},
)

var WebMiddlewareChain = chain.New(
	[]chain.Constructor{
		csrf.Protect([]byte("12345678910111213141151617181920")),
	},
	[]chain.ConstructorFunc{
		LoggerMiddleware,
	},
)
