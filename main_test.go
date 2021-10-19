package main

import "testing"

func Test_levenshtein(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"trivial", args{a: "", b: ""}, 0},
		{"simple", args{a: "aa", b: "a"}, 1},
		{"simple 1", args{a: "test", b: "testar"}, 2},
		{"simple 2", args{a: "clone", b: "clock"}, 2},
		{"large ", args{a: "cloneclonerclonest", b: "clonestclonerclone"}, 4},
		{"larger ", args{a: large1, b: large2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levenshtein(tt.args.a, tt.args.b)
			t.Log(got)
			if got != tt.want {
				t.Errorf("levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

const large1 = "estxrdcyfvugibjnmokä,ö.sxtdcfvghkbmöl,ä.xtsdcfgvdjhbnkm,löxtdhfcbkj möl,äterdycfvtugbhnjmklrdy tfiuopvgbhkjnml,"
const large2 = "estxrdcyfvubjnmokä,ö.sxtdcfvghkbmöl,ä.xtsdcfgvjhbnkm,löxtdhfcbkj möl,äterdycfvtugbhnjmwklrdy tfiuopvgbhkjnml,"
