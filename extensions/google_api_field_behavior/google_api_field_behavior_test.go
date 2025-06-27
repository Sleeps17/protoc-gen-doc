package extensions_test

import (
	"testing"

	"github.com/Sleeps17/protoc-gen-doc/extensions"
	. "github.com/Sleeps17/protoc-gen-doc/extensions/google_api_field_behavior"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func TestTransform(t *testing.T) {
	// Test with single behavior
	behavior := annotations.FieldBehavior_REQUIRED
	transformed := extensions.Transform(map[string]interface{}{"google.api.field_behavior": behavior})
	require.NotEmpty(t, transformed)

	fieldBehavior := transformed["google.api.field_behavior"].(FieldBehaviorExtension)
	assert.Equal(t, []string{"REQUIRED"}, fieldBehavior.Behaviors)

	// Test with multiple behaviors
	behaviors := []annotations.FieldBehavior{
		annotations.FieldBehavior_REQUIRED,
		annotations.FieldBehavior_OUTPUT_ONLY,
	}
	transformed = extensions.Transform(map[string]interface{}{"google.api.field_behavior": behaviors})
	require.NotEmpty(t, transformed)

	fieldBehavior = transformed["google.api.field_behavior"].(FieldBehaviorExtension)
	assert.Equal(t, []string{"REQUIRED", "OUTPUT_ONLY"}, fieldBehavior.Behaviors)

	// Test with nil payload
	transformed = extensions.Transform(map[string]interface{}{"google.api.field_behavior": nil})
	assert.Empty(t, transformed)

	// Test with wrong type
	transformed = extensions.Transform(map[string]interface{}{"google.api.field_behavior": "not a behavior"})
	assert.Empty(t, transformed)
}
