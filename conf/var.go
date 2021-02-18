package conf

type Configure struct {
	Conf []Conf `toml:"conf"`
}

type Conf struct {
	Method string `toml:"method"`
	URI    string `toml:"uri"`
	Status int    `toml:"status"`
	Header Header `toml:"header"`
	Body   Body   `toml:"body"`
}

type Body struct {
	Type string `toml:"type"`
	Data Data   `toml:"data"`
}

type Header map[string]string
type Data interface{}
