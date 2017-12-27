package api

import (
	"testing"
)

func TestReadCredentialsFromJSON(t *testing.T) {
	r := &readCredentialImpl{}

	t.Run("存在するjsonファイルを指定した場合に正しく構造体に値が設定されること", func(t *testing.T) {
		t.Parallel()

		// テスト対象処理の実行
		c, err := r.fromJSON("./creds_test.json")
		if err != nil {
			t.Fatalf("エラーが発生しました。/// %v", err)
		}

		if c.Cid != "credential_id" {
			t.Errorf("cidの内容が想定と異なります。 /// %s", c.Cid)
		}

		if c.Csecret != "credential_secret" {
			t.Errorf("csecretの内容が想定と異なります。 /// %s", c.Csecret)
		}
	})

	t.Run("存在しないjsonファイルを指定した場合にエラーが返却されること", func(t *testing.T) {
		t.Parallel()

		// テスト対象処理の実行
		_, err := r.fromJSON("./creds_nofile.json")
		if err == nil {
			t.Fatalf("エラーが返却されませんでした。")
		} else {
			t.Logf("エラー内容 /// %v", err)
		}
	})
}
