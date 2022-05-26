/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:22
 */

package mapper

import "testing"

func TestAddToRelation(t *testing.T) {
	type args struct {
		userId   int
		toUserId int
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
			if err := AddToRelation(tt.args.userId, tt.args.toUserId); (err != nil) != tt.wantErr {
				t.Errorf("AddToRelation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFromRelation(t *testing.T) {
	type args struct {
		userId   int
		toUserId int
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
			if err := DeleteFromRelation(tt.args.userId, tt.args.toUserId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFromRelation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
