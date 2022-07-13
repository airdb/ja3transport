package ja3transport_test

import (
	"io/ioutil"
	"testing"

	"github.com/airdb/ja3transport"
)

// const DefaultJA3Sig string = "771,4865-4866-4867-49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-18-16-11-51-45-43-10-21,29-23-24-25,0"
// const JA3erURL = "https://ja3er.com/json"
const JA3erURL = "https://dean.airdb.dev/"

func TestJa3(t *testing.T) {
	ja3str := ja3transport.BadJA3
	client, err := ja3transport.New(ja3str)
	if err != nil {
		t.Fatalf("new ja3 transport error: %v", err)
	}
	resp, err := client.Get(JA3erURL)
	if err != nil {
		t.Fatalf("http get error: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body error: %v", err)
	}

	t.Log(string(body))
	t.Log(ja3str)
}
