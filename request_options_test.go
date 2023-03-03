package lfu

import "testing"

func TestOptions(t *testing.T) {
	t.Parallel()

	resultSet := processOptions(Page(1), Page(2))

	actual := resultSet.urlParams.Encode()
	expected := "page=2"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
