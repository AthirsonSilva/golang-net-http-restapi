package main

import "testing"

func TestSetupComponents(t *testing.T) {
	err := setupComponents()
	if err != nil {
		t.Errorf("Failed running setupComponents() function => %v", err)
	}
}
