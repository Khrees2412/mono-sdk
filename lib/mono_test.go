package mono

import (
	"testing"
)

var cl *Client

func init() {
	cl.secret = "test_sk_l5J0JlLVbRZMsCZzjD6b"
}

func TestView360(t *testing.T) {

	bvn := &BVN{
		Bvn: "22426081430",
	}
	res, err := cl.View360(bvn)
	if err != nil {
		t.Error(err)
	}
	if len(res.Data) < 1 {
		t.Error("no accounts were found")
	}

}
