package libpx

type ProviderConfig struct {
	Name  string
	Value interface{}
}

type Configs []ProviderConfig

type ProviderConfigOption func(*Configs)

func WithProviderConfig(name string, value interface{}) ProviderConfigOption {
	return func(c *Configs) {

		var cs []ProviderConfig
		if c != nil {
			cs = []ProviderConfig(*c)
		} else {
			cs = make([]ProviderConfig, 0)
		}

		cs = append(cs, ProviderConfig{Name: name, Value: value})
		aux := Configs(cs)

		c = &aux
	}
}
