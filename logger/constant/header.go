package constant

import "github.com/go-http-utils/headers"

type Header string

const (
	USER_ID         Header = "Userid"
	ACCEPT_LANGUAGE Header = headers.AcceptLanguage
	CORRELATION_ID  Header = "Correlationid"
)
