package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type DirectpayService service

func (c *DirectpayService) MakeDirectPay(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/", userID)
	resp := &models.Income{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err

}

/*
"message": "Successfully fetched linked accounts.",
  "data": [
    {
      "accountNumber": "005****331",
      "institution": {
        "_id": "5f2d08bf60b92e2888287703",
        "name": "Access Bank",
        "bankCode": "044",
        "icon": "https://mono-public-bucket.s3.eu-west-2.amazonaws.com/images/access-bank-icon.png"
      }
    },
    {
      "accountNumber": "015****747",
      "institution": {
        "_id": "5f2d08be60b92e2888287702",
        "name": "GTBank",
        "bankCode": "058",
        "icon": "https://mono-public-bucket.s3.eu-west-2.amazonaws.com/images/gtbank-icon.png"
      }
    },
    {
      "accountNumber": "080****534",
      "institution": {
        "_id": "5f6de706800d071e5566ef7d",
        "name": "Access Bank",
        "bankCode": "044",
        "icon": "https://mono-public-bucket.s3.eu-west-2.amazonaws.com/images/access-bank-icon.png"
      }
    },
    {
      "accountNumber": "200****577",
      "institution": {
        "_id": "5f2d08bf60b92e2888287704",
        "name": "KudaBank",
        "bankCode": "090267",
        "icon": "https://mono-public-bucket.s3.eu-west-2.amazonaws.com/images/kuda-bank-icon.png"
      }
    }
  ]
}
*/
