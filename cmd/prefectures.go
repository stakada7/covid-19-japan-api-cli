package cmd

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/spf13/cobra"
)

// prefecturesCmd represents the prefectures command
var prefecturesCmd = &cobra.Command{
	Use:   "prefectures",
	Short: "Get total number of cases and deaths per prefecture ",
	Long:  `Get total number of cases and deaths per prefecture along with the name (in Japanese and English) and coordinates (lattitude and longitude) of the prefecture.`,
	RunE:  runPrefecturesCmd,
}

func init() {
	rootCmd.AddCommand(prefecturesCmd)

}

func runPrefecturesCmd(cmd *cobra.Command, args []string) error {

	cli, err := newDefaultClient()
	if err != nil {
		return fmt.Errorf("newClient failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	prefectures, err := cli.GetPrefectures(ctx)
	if err != nil {
		return fmt.Errorf("GetPrefectures failed: res = %+v", prefectures)
	}

	for _, prefecture := range prefectures {
		fmt.Printf("ID:%d NameJa:%s NameEn:%s Lat:%g Lng:%g Cases:%d Deaths:%d\n", prefecture.ID, prefecture.NameJa, prefecture.NameEn, prefecture.Lat, prefecture.Lng, prefecture.Cases, prefecture.Deaths)
	}

	return nil

}

// GetPrefectures ...
func (cli *Client) GetPrefectures(ctx context.Context) ([]*Prefecture, error) {

	subPath := path.Join("api", "v1", "prefectures")
	httpRequest, err := cli.newRequest(ctx, http.MethodGet, subPath, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := cli.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var prefectures []*Prefecture
	if err := decodeBody(httpResponse, &prefectures); err != nil {
		return nil, err
	}

	return prefectures, nil

}
