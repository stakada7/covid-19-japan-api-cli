package cmd

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestGetPrefectures(t *testing.T) {
	tests := []struct {
		name string

		mockResponseHeaderFile string
		mockResponseBodyFile   string

		expectedMethod      string
		expectedRequestPath string
		expectedRawQuery    string

		want               []*Prefecture
		wantErr            bool
		expectedErrMessage string
	}{
		{
			name: "success",

			mockResponseHeaderFile: "testdata/GetPrefectures/success-header",
			mockResponseBodyFile:   "testdata/GetPrefectures/success-body",

			expectedMethod:      http.MethodGet,
			expectedRequestPath: "/api/v1/prefectures",
			expectedRawQuery:    "",

			want: []*Prefecture{
				{ID: 1, NameJa: "北海道", NameEn: "Hokkaido", Lat: 43.46722222, Lng: 142.8277778},
				{ID: 2, NameJa: "青森", NameEn: "Aomori", Lat: 40.78027778, Lng: 140.83194440000003},
				{ID: 3, NameJa: "岩手", NameEn: "Iwate", Lat: 39.59138889, Lng: 141.3625},
				{ID: 4, NameJa: "宮城", NameEn: "Miyagi", Lat: 38.44555556, Lng: 140.9280556},
				{ID: 5, NameJa: "秋田", NameEn: "Akita", Lat: 39.7475, Lng: 140.4086111},
				{ID: 6, NameJa: "山形", NameEn: "Yamagata", Lat: 38.44638889, Lng: 140.1027778},
				{ID: 7, NameJa: "福島", NameEn: "Fukushima", Lat: 37.37888889, Lng: 140.22527780000001},
				{ID: 8, NameJa: "茨城", NameEn: "Ibaraki", Lat: 36.30638889, Lng: 140.3186111},
				{ID: 9, NameJa: "栃木", NameEn: "Tochigi", Lat: 36.68916667, Lng: 139.81916669999998},
				{ID: 10, NameJa: "群馬", NameEn: "Gunma", Lat: 36.50388889, Lng: 138.9852778},
				{ID: 11, NameJa: "埼玉", NameEn: "Saitama", Lat: 35.99666667, Lng: 139.34777780000002},
				{ID: 12, NameJa: "千葉", NameEn: "Chiba", Lat: 35.51277778, Lng: 140.2038889},
				{ID: 13, NameJa: "東京", NameEn: "Tokyo", Lat: 35.702313700000005, Lng: 139.5803228},
				{ID: 14, NameJa: "神奈川", NameEn: "Kanagawa", Lat: 35.41416667, Lng: 139.34027780000002},
				{ID: 15, NameJa: "新潟", NameEn: "Niigata", Lat: 37.51888889, Lng: 138.91722219999997},
				{ID: 16, NameJa: "富山", NameEn: "Toyama", Lat: 36.63611111, Lng: 137.2680556},
				{ID: 17, NameJa: "石川", NameEn: "Ishikawa", Lat: 36.76583333, Lng: 136.7713889},
				{ID: 18, NameJa: "福井", NameEn: "Fukui", Lat: 35.84666667, Lng: 136.22722219999997},
				{ID: 19, NameJa: "山梨", NameEn: "Yamanashi", Lat: 35.61222222, Lng: 138.6116667},
				{ID: 20, NameJa: "長野", NameEn: "Nagano", Lat: 36.13, Lng: 138.04388889999998},
				{ID: 21, NameJa: "岐阜", NameEn: "Gifu", Lat: 35.7775, Lng: 137.055},
				{ID: 22, NameJa: "静岡", NameEn: "Shizuoka", Lat: 35.01694444, Lng: 138.33},
				{ID: 23, NameJa: "愛知", NameEn: "Aichi", Lat: 35.03444444, Lng: 137.215},
				{ID: 24, NameJa: "三重", NameEn: "Mie", Lat: 34.51361111, Lng: 136.3813889},
				{ID: 25, NameJa: "滋賀", NameEn: "Shiga", Lat: 35.21527778, Lng: 136.13805559999997},
				{ID: 26, NameJa: "京都", NameEn: "Kyoto", Lat: 35.25194444, Lng: 135.4458333},
				{ID: 27, NameJa: "大阪", NameEn: "Osaka", Lat: 34.62277778, Lng: 135.5111111},
				{ID: 28, NameJa: "兵庫", NameEn: "Hyogo", Lat: 35.03694444, Lng: 134.8286111},
				{ID: 29, NameJa: "奈良", NameEn: "Nara", Lat: 34.31555556, Lng: 135.8713889},
				{ID: 30, NameJa: "和歌山", NameEn: "Wakayama", Lat: 33.90944444, Lng: 135.5133333},
				{ID: 31, NameJa: "鳥取", NameEn: "Tottori", Lat: 35.36055556, Lng: 133.8516667},
				{ID: 32, NameJa: "島根", NameEn: "Shimane", Lat: 35.07305556, Lng: 132.55944440000002},
				{ID: 33, NameJa: "岡山", NameEn: "Okayama", Lat: 34.90083333, Lng: 133.8152778},
				{ID: 34, NameJa: "広島", NameEn: "Hiroshima", Lat: 34.60361111, Lng: 132.7875},
				{ID: 35, NameJa: "山口", NameEn: "Yamaguchi", Lat: 34.19861111, Lng: 131.575},
				{ID: 36, NameJa: "徳島", NameEn: "Tokushima", Lat: 33.91805556, Lng: 134.2430556},
				{ID: 37, NameJa: "香川", NameEn: "Kagawa", Lat: 34.24305556, Lng: 133.99666670000002},
				{ID: 38, NameJa: "愛媛", NameEn: "Ehime", Lat: 33.62194444, Lng: 132.8558333},
				{ID: 39, NameJa: "高知", NameEn: "Kōchi", Lat: 33.42111111, Lng: 133.36666670000002},
				{ID: 40, NameJa: "福岡", NameEn: "Fukuoka", Lat: 33.5225, Lng: 130.66805559999997},
				{ID: 41, NameJa: "佐賀", NameEn: "Saga", Lat: 33.28527778, Lng: 130.11694440000002},
				{ID: 42, NameJa: "長崎", NameEn: "Nagasaki", Lat: 33.2275, Lng: 129.61416670000003},
				{ID: 43, NameJa: "熊本", NameEn: "Kumamoto", Lat: 32.615, Lng: 130.7563889},
				{ID: 44, NameJa: "大分", NameEn: "Oita", Lat: 33.19916667, Lng: 131.43416670000002},
				{ID: 45, NameJa: "宮崎", NameEn: "Miyazaki", Lat: 32.19083333, Lng: 131.3005556},
				{ID: 46, NameJa: "鹿児島", NameEn: "Kagoshima", Lat: 31.01277778, Lng: 130.4241667},
				{ID: 47, NameJa: "沖縄", NameEn: "Okinawa", Lat: 25.77111111, Lng: 126.64},
			},
			wantErr:            false,
			expectedErrMessage: "",
		},
	}

	_, mockServerURL := newMockServer()
	cli := newTestClient(mockServerURL)

	for _, tt := range tests {
		// mux.HandleFunc(tt.expectedRequestPath, func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Fprint(w, tt.want)
		// })

		prefectures, err := cli.GetPrefectures(context.Background())
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

			if len(prefectures) != len(tt.want) {
				t.Fatalf("response prefectures wrong. want=%+v, got=%+v", tt.want, prefectures)
			}

			for i, expected := range tt.want {
				actual := prefectures[i]
				if actual.ID != expected.ID || actual.NameJa != expected.NameJa || actual.NameEn != expected.NameEn || actual.Lat != expected.Lat || actual.Lng != expected.Lng {
					t.Fatalf("response prefectures wrong. want=%+v, got=%+v", tt.want, prefectures)
				}
			}
		}

	}
}
