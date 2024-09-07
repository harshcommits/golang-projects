package learning

import "testing"

func TestIsEmployee(t *testing.T) {
	if IsEmployee("harsh@linkedin.com") == false {
		t.Errorf("Expected false but got true")
	}
}
