package auth

type Authority struct {
	Name      string                                 `json:"name"`
	Authority func(args map[string]interface{}) bool `json:"authority"`
}

// args -> context

var (
	AUTHORITY_BROWSE = Authority{
		"browse",
		func(args map[string]interface{}) bool {
			return true
		},
	}
	AUTHORITY_EDIT = Authority{
		"edit",
		func(args map[string]interface{}) bool {
			return args["user"] == nil
		},
	}
)
