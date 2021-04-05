package unit

import (
	"testing"

	lib "actiontest1/lib"

	"github.com/cloud-barista/cb-spider/interface/api"
)

func TestLib(t *testing.T) {
	total := lib.SumLib1(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSpider(t *testing.T) {

	cim := api.NewCloudInfoManager()
	defer cim.Close()

	err := cim.SetConfigPath("./grpc_conf.yaml")
	if err != nil {
		t.Errorf("error raised: %v", err)
	}

	err = cim.Open()
	if err != nil {
		t.Errorf("error raised: %v", err)
	}
	defer cim.Close()

	result, err := cim.ListCloudOS()
	if err != nil {
		t.Errorf("error raised: %v", err)
	}

	t.Logf("result : %v\n", result)

}
