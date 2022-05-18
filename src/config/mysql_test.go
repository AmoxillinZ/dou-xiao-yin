/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/18 20:25
 */

package config

import (
	"testing"
)

func TestGetDefaultDb(t *testing.T) {
	InitDefaultDbEngine()
}

func TestInitDefaultDbEngine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDefaultDbEngine()
		})
	}
}

func TestGetDefaultDb1(t *testing.T) {
	GetDefaultDb()
}

func TestInitDefaultDbEngine1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDefaultDbEngine()
		})
	}
}
