package appcontext_test

import (
	"context"
	"testing"

	"github.com/ituoga/appcontext"
)

func TestCore(t *testing.T) {
	t.Run("With and Get", func(t *testing.T) {
		var myKey = appcontext.Key[string]{}

		ctx := context.Background()
		ctx = appcontext.With(ctx, myKey, "testValue")

		value, ok := appcontext.Get(ctx, myKey)
		if !ok || value != "testValue" {
			t.Errorf("Expected 'testValue', got '%v'", value)
		}
	})

	t.Run("Get non-existent key", func(t *testing.T) {
		var anotherKey = appcontext.Key[string]{}

		ctx := context.Background()
		_, ok := appcontext.Get(ctx, anotherKey)
		if ok {
			t.Error("Expected false for non-existent key")
		}
	})
}
