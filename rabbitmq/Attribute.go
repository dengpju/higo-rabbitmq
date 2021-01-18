package rabbitmq

const (
	USER_NAME = "user_name"
	PASSWORD  = "password"
	HOST      = "host"
	POST      = "post"
)

type Attributes []*attribute

func (this Attributes) Find(name string) interface{} {
	for _, p := range this {
		if p.name == name {
			return p.value
		}
	}
	return nil
}

func (this Attributes) String(name string) string {
	val := this.Find(name)
	if val != nil {
		return val.(string)
	}
	return ""
}

type attribute struct {
	name  string
	value interface{}
}

func Attribute(name string, value interface{}) *attribute {
	return &attribute{name, value}
}

func Username(value string) *attribute {
	return Attribute(USER_NAME, value)
}

func Password(value string) *attribute {
	return Attribute(PASSWORD, value)
}

func Host(value string) *attribute {
	return Attribute(HOST, value)
}

func Post(value string) *attribute {
	return Attribute(POST, value)
}
