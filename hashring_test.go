package hashring

import (
	"reflect"
	"testing"
)

func TestNewHashRing(t *testing.T) {
	tests := []struct {
		name    string
		want    *HashRing
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHashRing()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHashRing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashRing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashRing_AddNode(t *testing.T) {
	type args struct {
		nodeName   string
		nodeWeight int64
	}
	tests := []struct {
		name    string
		h       *HashRing
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddNode(tt.args.nodeName, tt.args.nodeWeight); (err != nil) != tt.wantErr {
				t.Errorf("HashRing.AddNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashRing_RemoveNode(t *testing.T) {
	type args struct {
		nodeName string
	}
	tests := []struct {
		name    string
		h       *HashRing
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.RemoveNode(tt.args.nodeName); (err != nil) != tt.wantErr {
				t.Errorf("HashRing.RemoveNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashRing_GetNode(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    *HashRing
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.GetNode(tt.args.key); got != tt.want {
				t.Errorf("HashRing.GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashRing_AddNodes(t *testing.T) {
	type args struct {
		nodeWeightMap map[string]int64
	}
	tests := []struct {
		name    string
		h       *HashRing
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddNodes(tt.args.nodeWeightMap); (err != nil) != tt.wantErr {
				t.Errorf("HashRing.AddNodes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
