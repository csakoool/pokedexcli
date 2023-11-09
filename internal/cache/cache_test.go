package cache

import (
	"fmt"
	"testing"
	"time"
)

const interval = 5 * time.Second

func TestAddErrors(t *testing.T) {
	cache := NewCache(interval)

	err := cache.Add("", []byte{1, 2})
	expectedError := "Key was not provided to Cache.Add method"
	if err.Error() != expectedError {
		t.Errorf("expected to get error '%s' but got '%s'", expectedError, err)
		return
	}

	err = cache.Add("test", []byte{})
	expectedError = "There's nothing to cache - value is empty"
	if err.Error() != expectedError {
		t.Errorf("expected to get error '%s' but got '%s'", expectedError, err)
		return
	}
}

func TestAddGet(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
	}

	for index, testCase := range cases {
		t.Run(fmt.Sprintf("Test case %v", index), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(testCase.key, testCase.value)
			val, ok := cache.Get(testCase.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(testCase.value) {
				t.Errorf("expected to find correct value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const waitTime = interval + 5*time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
