package lfu

import "testing"

func TestOptions(t *testing.T) {
	t.Parallel()

	resultSet := processOptions(Page(1), Page(2), Limit(1), From(1), To(1), Extended(true))

	actual := resultSet.urlParams.Encode()
	expected := "extended=1&from=1&limit=1&page=2&to=1"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
