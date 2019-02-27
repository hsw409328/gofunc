/**
 * Author: haoshuaiwei 
 * Date: 2019-02-27 17:20 
 */

package go_jwt

import (
	"testing"
)

func TestNewJwtObject(t *testing.T) {
	o := NewJwtObject()
	t.Log(string(o.SignKey))
	o1 := NewJwtObject()
	t.Log(string(o1.SignKey))
}

func TestJwtSetSignKey(t *testing.T) {
	JwtSetSignKey("aaaaa")
	t.Log(SignKey)
}

func TestJwtObject_CreateJwtToken(t *testing.T) {
	o := NewJwtObject()
	t.Log(o.CreateJwtToken(CustomClaimsPayload{
		Host: "192.168.1.1",
	}))
}
