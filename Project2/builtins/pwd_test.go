package builtins

import "testing"

func TestPwd(t *testing.T) {
	type args struct {
		args []string
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
			if err := Pwd(tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Pwd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
