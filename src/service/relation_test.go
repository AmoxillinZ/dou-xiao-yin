/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:30
 */

package service

import "testing"

func TestFollowAction(t *testing.T) {
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
			if err := FollowAction(tt.args.userId, tt.args.toUserId); (err != nil) != tt.wantErr {
				t.Errorf("FollowAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRelationAction(t *testing.T) {
	type args struct {
		userId     int
		toUserId   int
		actionType int
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
			if err := RelationAction(tt.args.userId, tt.args.toUserId, tt.args.actionType); (err != nil) != tt.wantErr {
				t.Errorf("RelationAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnFollowAction(t *testing.T) {
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
			if err := UnFollowAction(tt.args.userId, tt.args.toUserId); (err != nil) != tt.wantErr {
				t.Errorf("UnFollowAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
