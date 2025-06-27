package extensions

import (
	"github.com/Sleeps17/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/annotations"
)

// FieldBehaviorExtension contains the field behavior values set by the (google.api.field_behavior) field option extension.
type FieldBehaviorExtension struct {
	Behaviors []string `json:"behaviors"`
}

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		// Try different possible types
		switch v := payload.(type) {
		case []annotations.FieldBehavior:
			var behaviors []string
			for _, behavior := range v {
				behaviors = append(behaviors, behavior.String())
			}
			return FieldBehaviorExtension{Behaviors: behaviors}
		case annotations.FieldBehavior:
			return FieldBehaviorExtension{Behaviors: []string{v.String()}}
		case []interface{}:
			var behaviors []string
			for _, item := range v {
				if behavior, ok := item.(annotations.FieldBehavior); ok {
					behaviors = append(behaviors, behavior.String())
				}
			}
			if len(behaviors) > 0 {
				return FieldBehaviorExtension{Behaviors: behaviors}
			}
		}

		return nil
	})
}
