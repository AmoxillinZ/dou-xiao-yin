package mapper

import (
	"MyProject/src/model"
	"fmt"
	"reflect"
	"testing"
)

func TestGetVideoById(t *testing.T) {
	fmt.Println(GetVideos()[1])
}

func TestGetVideoById1(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *model.Video
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVideoById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideoById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVideos(t *testing.T) {
	tests := []struct {
		name string
		want []*model.Video
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVideos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideos() = %v, want %v", got, tt.want)
			}
		})
	}
}
