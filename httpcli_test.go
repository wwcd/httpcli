package httpcli

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDo(t *testing.T) {
	expected := []byte(`{"foo":"bar"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Error("make req failed", err)
	}

	rsp, err := Do(context.Background(), req)
	if err != nil {
		t.Error("do failed", err)
	}
	defer rsp.Body.Close()

	actual, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Error("read body failed", err)
	}
	if !bytes.Equal(expected, actual) {
		t.Error("expected neq actual")
	}
}
