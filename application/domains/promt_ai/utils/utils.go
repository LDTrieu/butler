package utils

type ApiHeader map[string]string

func (a ApiHeader) AddHeader(key, value string) {
	a[key] = value
}
