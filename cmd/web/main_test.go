package main

import "testing"

func TestSetupComponents(t *testing.T) {
	_, err := setupComponents()
	if err != nil {
		t.Errorf("Failed running setupComponents() function => %v", err)
	}
}
