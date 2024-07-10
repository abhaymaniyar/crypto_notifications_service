package enums

type Type string

var (
	HIGH_CHANGE Type = "HIGH_CHANGE"
	ORDER       Type = "ORDER"
)

func (t Type) IsValid() bool {
	switch t {
	case HIGH_CHANGE, ORDER:
		return true
	}

	return false
}
