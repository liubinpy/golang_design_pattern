package factory1

import (
	"reflect"
	"testing"
)

func TestCreateFruit(t *testing.T) {
	type args struct {
		kind string
	}

	tests := []struct {
		name string
		args args
		want Fruit
	}{
		{
			name: "apple",
			args: args{kind: "apple"},
			want: &Apple{},
		},
		{
			name: "pear",
			args: args{kind: "pear"},
			want: &Pear{},
		},
	}

	f := factory{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.CreateFruit(tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
