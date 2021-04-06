package unit

import (
	"fmt"
	"testing"

	lib "actiontest1/lib"

	"github.com/cloud-barista/cb-spider/api-runtime/grpc-runtime/logger"
	"github.com/cloud-barista/cb-spider/interface/api"
)

func TestLib(t *testing.T) {
	total := lib.SumLib1(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

}

func TestSpider(t *testing.T) {
	logger := logger.NewLogger()

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

	t.Errorf("t.Errorf result : %v\n", result)
	t.Logf("t.Logf result : %v\n", result)

	fmt.Printf("fmt result : %v\n", result)

	logger.Debug("logger debug print...\n")
	logger.Info("logger info print...\n")
	logger.Error("logger error print...\n")

}
