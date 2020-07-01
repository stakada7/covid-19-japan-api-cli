package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/spf13/cobra"
)

// totalCmd represents the total command
var totalCmd = &cobra.Command{
	Use:   "total",
	Short: "Get total number of cases and deaths in Japan",
	Long:  `Get total number of cases and deaths in Japan.`,
	RunE:  runTotalCmd,
}

func init() {
	rootCmd.AddCommand(totalCmd)

}

func runTotalCmd(cmd *cobra.Command, args []string) error {

	cli, err := newDefaultClient()
	if err != nil {
		return fmt.Errorf("newClient failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	total, err := cli.GetTotal(ctx)
	if err != nil {
		return fmt.Errorf("GetTotal failed: res = %+v", total)
	}

	fmt.Printf("Date:%d ", total.Date)
	fmt.Printf("Pcr:%d ", total.Pcr)
	fmt.Printf("Positive:%d ", total.Positive)
	fmt.Printf("Symptom:%d ", total.Symptom)
	fmt.Printf("Symptomless:%d ", total.Symptomless)
	fmt.Printf("SymtomConfirming:%d ", total.SymtomConfirming)
	fmt.Printf("Hospitalize:%d ", total.Hospitalize)
	fmt.Printf("Mild:%d ", total.Mild)
	fmt.Printf("Severe:%d ", total.Severe)
	fmt.Printf("Confirming:%d ", total.Confirming)
	fmt.Printf("Waiting:%d ", total.Waiting)
	fmt.Printf("Discharge:%d ", total.Discharge)
	fmt.Printf("Death%d\n", total.Death)

	return nil
}

// GetTotal ...
func (c *Client) GetTotal(ctx context.Context) (*Total, error) {

	relativePath := path.Join("api", "v1", "total")

	req, err := c.newRequest(ctx, http.MethodGet, relativePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var total *Total
	code, err := c.doRequest(req, &total)

	switch code {
	case http.StatusOK:
		return total, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}
