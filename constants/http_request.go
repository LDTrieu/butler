package constants

const METHOD_GET = "GET"
const METHOD_POST = "POST"
const METHOD_PUT = "PUT"
const METHOD_DELETE = "DELETE"

const (
	_      = iota //blank identifier
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)
