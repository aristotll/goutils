package stringutils

import "testing"

func TestFromInt64UsingHex(t *testing.T) {
	type args struct {
		in int64
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			name: "test1",
			args: args{
				in: 123456789,
			},
			wantRet: "75bcd15",
		},
		{
			name: "test2",
			args: args{
				in: 0,
			},
			wantRet: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := FromInt64UsingHex(tt.args.in); gotRet != tt.wantRet {
				t.Errorf("FromInt64UsingHex() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestToInt64UsingHex(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int64
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				str: "75bcd15",
			},
			wantRet: 123456789,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				str: "0",
			},
			wantRet: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := ToInt64UsingHex(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64UsingHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("ToInt64UsingHex() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
