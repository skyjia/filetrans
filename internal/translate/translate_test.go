package translate

import (
	"testing"
	"time"
)

func TestTranslator_Translate(t *testing.T) {
	type fields struct {
		appID string
		key   string
		delay time.Duration
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		// {
		// 	name: "t1",
		// 	fields: fields{
		// 		appID: "20210707000882113",
		// 		key:   "_xfaR1zsBEOk7H1Yi60b",
		// 		delay: 1000 * time.Millisecond,
		// 	},
		// 	args: args{
		// 		query: "hello",
		// 	},
		// 	want: "你好",
		// },
		// {
		// 	name: "t2",
		// 	fields: fields{
		// 		appID: "20210707000882113",
		// 		key:   "_xfaR1zsBEOk7H1Yi60b",
		// 		delay: 1000 * time.Millisecond,
		// 	},
		// 	args: args{
		// 		query: "hello,sky",
		// 	},
		// 	want: "你好，天空",
		// },
		{
			name: "t3",
			fields: fields{
				appID: "20210707000882113",
				key:   "_xfaR1zsBEOk7H1Yi60b",
				delay: 1000 * time.Millisecond,
			},
			args: args{
				query: "hello_world water",
			},
			want: "你好_世界水",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Translator{
				appID: tt.fields.appID,
				key:   tt.fields.key,
				delay: tt.fields.delay,
			}
			if got := tr.Translate(tt.args.query); got != tt.want {
				t.Errorf("Translator.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
