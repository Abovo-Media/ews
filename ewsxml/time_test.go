package ewsxml

import (
	"reflect"
	"testing"
	"time"
)

func Test_getRFC3339Offset(t *testing.T) {
	riyadh, _ := time.LoadLocation("Asia/Riyadh")
	marquesas, _ := time.LoadLocation("Pacific/Marquesas")

	type args struct {
		time time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name:    "test with timezone +03:00",
			args:    args{time: time.Now().In(riyadh)},
			wantErr: false,
			want:    "+03:00",
		},
		{
			name:    "test with timezone -09:30",
			args:    args{time: time.Now().In(marquesas)},
			wantErr: false,
			want:    "-09:30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRFC3339Offset(tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSoapFault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSoapFault() got = %v, want %v", got, tt.want)
			}
		})
	}

}
