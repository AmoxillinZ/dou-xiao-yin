/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 17:24
 */

package service

import "testing"

func TestTokenVerify(t *testing.T) {
	type args struct {
		userId int
		token  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TokenVerify(tt.args.userId, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("TokenVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
