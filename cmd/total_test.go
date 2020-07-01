package cmd

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestGetTotal(t *testing.T) {
	tests := []struct {
		name string

		mockResponseHeaderFile string
		mockResponseBodyFile   string

		expectedMethod      string
		expectedRequestPath string
		expectedRawQuery    string

		want               Total
		wantErr            bool
		expectedErrMessage string
	}{
		{
			name: "success",

			mockResponseHeaderFile: "../testdata/GetTotal/success-header",
			mockResponseBodyFile:   "../testdata/GetTotal/success-body",

			expectedMethod:      http.MethodGet,
			expectedRequestPath: "/api/v1/total",
			expectedRawQuery:    "",

			want: Total{
				Date: 20200402, Pcr: 32002, Positive: 2306, Symptom: 1693, Symptomless: 227, SymtomConfirming: 386, Hospitalize: 1757, Mild: 882, Severe: 62, Confirming: 418, Waiting: 9, Discharge: 489, Death: 60,
			},
			wantErr:            false,
			expectedErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, teardown := setup(t, tt.mockResponseHeaderFile, tt.mockResponseBodyFile, tt.expectedMethod, tt.expectedRequestPath, tt.expectedRawQuery)
			defer teardown()

			_, err := cli.GetTotal(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Fatalf("response error should not be non-nil. got=nil")
				}

				if !strings.Contains(err.Error(), tt.expectedErrMessage) {
					t.Fatalf("response erro message wtong. '%s' is expected to contain '%s'", err.Error(), tt.expectedErrMessage)
				}
			} else {
				if err != nil {
					t.Fatalf("response error should be nil. got=%s", err.Error())
				}

			}

		})
	}
}
