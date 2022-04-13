package mono

import (
	"reflect"
	"testing"
)

func TestConnectService_GetAssets(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name  string
		c     *ConnectService
		args  args
		want  interface{}
		want1 interface{}
	}{
		// name: "Status code",
		// c: *ConnectService,
		// args: "",
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.GetAssets(tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectService.GetAssets() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ConnectService.GetAssets() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
