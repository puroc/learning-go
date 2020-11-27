package mock

type RetrieverImpl struct {
	Contents string
}

func (r RetrieverImpl) Get(url string) string {
	return r.Contents
}
