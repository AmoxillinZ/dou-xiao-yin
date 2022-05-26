/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 17:24
 */

package service

import (
	"fmt"
	"testing"
)

func TestTokenVerify(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3b3JkIjoiMjMzMzMzIiwiZXhwIjoxNjU0MTYxNTk5LCJpYXQiOjE2NTM1NTY3OTksImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.aUbGxzQ2uzWmhIoA13AEcp-7DAULK48YEoLKX0ksAY4"
	err := TokenVerify(0, token)
	fmt.Println(err)
}
