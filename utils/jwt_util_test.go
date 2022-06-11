package utils

import (
	"fmt"
	"testing"
)

func TestGetIdFromToken(t *testing.T) {
	id, err := GetIdFromToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3bIjoiMjMzMzMzIiwiZXhwIjoxNjUzODc2MDExLCJpYXQiOjE2NTMyNzEyMTEsImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.ZFqz1Pc7VfrNjWBUYKmav4WVsWkcIB4qC5CxActqA5Q")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}
