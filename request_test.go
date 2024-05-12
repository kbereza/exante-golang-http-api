package exante

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_libTransport_RoundTrip(t *testing.T) {
	type fields struct {
		underlyingTransport http.RoundTripper
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &libTransport{
				underlyingTransport: tt.fields.underlyingTransport,
			}
			got, err := l.RoundTrip(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("libTransport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("libTransport.RoundTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serialize(t *testing.T) {
	type args struct {
		data  []byte
		model interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantQuotes *[]Quote
		wantErr    bool
	}{
		{"#1", args{[]byte(`[{"ask":[{"price":"182.52","size":"6E+2"}],"bid":[{"price":"182.5","size":"1.6E+3"}],"symbolId":"AAPL.NASDAQ","timestamp":1715385592452}]`),
			NewQuote()}, &[]Quote{
			{Timestamp: 1715385592452, SymbolID: "AAPL.NASDAQ",
				Ask: []Ask{{Price: "182.52", Size: "6E+2"}},
				Bid: []Bid{{Price: "182.5", Size: "1.6E+3"}},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := serialize(tt.args.data, &tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.model, tt.wantQuotes) {
				t.Errorf("equal() = %v, want %v", tt.args.model, tt.wantQuotes)
			}
		})
	}
}
