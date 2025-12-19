package core_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/openlogi/poc_stock/pkg/basic/core"
	"testing"
)

func TestNewIdentifier_Valid(t *testing.T) {
	validUUIDs := []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b811-9dad-11d1-80b4-00c04fd430c8",
		"00000000-0000-0000-0000-000000000000",
	}

	for _, uuidStr := range validUUIDs {
		id, err := core.NewIdentifier(uuidStr)
		if err != nil {
			t.Errorf("Valid UUID should not return error: %s, got error: %v", uuidStr, err)
		}
		if diff := cmp.Diff(uuidStr, id.String()); diff != "" {
			t.Errorf("Identifier mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestNewIdentifier_Invalid(t *testing.T) {
	invalidUUIDs := []string{
		"not-a-uuid",
		"550e8400-e29b-41d4-a716",              // too short
		"550e8400-e29b-41d4-a716-44665544000G", // invalid character
		"",                                     // empty
		"123",
	}

	for _, uuidStr := range invalidUUIDs {
		id, err := core.NewIdentifier(uuidStr)
		if err == nil {
			t.Errorf("Invalid UUID should return error: %s", uuidStr)
		}
		if err != nil {
			if diff := cmp.Diff("", id.String()); diff != "" {
				t.Errorf("Empty identifier expected for invalid input (-want +got):\n%s", diff)
			}
			t.Logf("Invalid UUID '%s' error: %s", uuidStr, err.Error())
		}
	}
}

func TestIdentifier_String(t *testing.T) {
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"
	id, _ := core.NewIdentifier(uuidStr)
	
	if diff := cmp.Diff(uuidStr, id.String()); diff != "" {
		t.Errorf("Identifier.String() mismatch (-want +got):\n%s", diff)
	}
}
