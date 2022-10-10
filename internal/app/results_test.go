package app

import (
	"reflect"
	"testing"
)

func TestNewResultsMap(t *testing.T) {
	tests := []struct {
		name    string
		domains []string
		//want    ResultsMap
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "empty list",
			domains: []string{},
			wantErr: true,
		},
		{
			name:    "no duplicates",
			domains: []string{"client://aa.com", "client://bb.org", "client://cc.net"},
			want: map[string]string{
				"client://aa.com": "Not Processed",
				"client://bb.org": "Not Processed",
				"client://cc.net": "Not Processed",
			},
		},
		{
			name:    "one duplicate",
			domains: []string{"client://aa.com", "client://bb.org", "client://bb.org", "client://cc.net"},
			want: map[string]string{
				"client://aa.com": "Not Processed",
				"client://bb.org": "Not Processed",
				"client://cc.net": "Not Processed",
			},
		},
		{
			name:    "all duplicated",
			domains: []string{"client://aa.com", "client://aa.com", "client://bb.org", "client://bb.org", "client://cc.net", "client://cc.net"},
			want: map[string]string{
				"client://aa.com": "Not Processed",
				"client://bb.org": "Not Processed",
				"client://cc.net": "Not Processed",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResultsMap(tt.domains)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResultsMap() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultsMap() got = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
