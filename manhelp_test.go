package manhelp

import "testing"

func TestAddManHelper(t *testing.T) {
	hi := HelpInfo{}
	hi.Alias = []string{"t", "-t", "--t"}
	hi.FullName = []string{"buildtime", "-buildtime", "--buildtime"}
	hi.ExecuteFunc = func() {
	}
	hi2 := hi
	hi2.FullName = []string{"-xifej__feifje"}
	hi2.Alias = []string{"xifej"}
	type args struct {
		newHelpers ManHelper
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {"error", args{hi}, true},
		{"normal", args{hi2}, false},
	}
	Main()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddManHelper(tt.args.newHelpers); (err != nil) != tt.wantErr {
				t.Errorf("AddManHelper(%v) error = %v, wantErr %v", tt.args.newHelpers, err, tt.wantErr)
			}
		})
	}
}
