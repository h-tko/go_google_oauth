package api

import (
	"testing"
)

type readCredentialTest struct {
}

func (r *readCredentialTest) fromJSON(f string) (Credentials, error) {
	c := Credentials{
		Cid:     "test",
		Csecret: "test",
	}

	return c, nil
}

func TestAuth(t *testing.T) {
	g := NewGoogleOAuth(&readCredentialTest{})

	t.Run("認証urlが正しく返却されること", func(t *testing.T) {
		t.Parallel()

		url, err := g.Auth()
		if err != nil {
			t.Fatalf("エラーが発生しました。 /// %v", err)
		}

		if url == "" {
			t.Error("戻り値が空文字です。")
		}

		t.Logf("url = %s", url)
	})
}

func TestCallback(t *testing.T) {
	t.Fatal("this test is not coding yet.")
}
