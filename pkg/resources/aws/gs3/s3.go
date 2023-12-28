package gs3

type Client interface {
	WriteToObject(string, []byte) error
}
