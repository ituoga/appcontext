# appcontext

usage: 

```go
package main

import "github.com/ituoga/appcontext"

var localeKey = appcontext.Key[string]("localKey")

func WithLocale(ctx context.Context, locale string) context.Context {
	return appcontext.With(ctx, localeKey, locale)
}

func Locale(ctx context.Context) (string, bool) {
	return appcontext.Get(ctx, localeKey)
}

```

or check `appcontext_test.go`
