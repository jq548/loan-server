package job

import (
	"loan-server/config"
	"loan-server/db"
	"loan-server/service"
	"testing"
)

var database, _ = db.Init(&config.Db{Dsn: "dev:XASIFSNXO!353@tcp(8.210.246.101:3306)/loan?charset=utf8mb4&parseTime=True&loc=UTC"})
var bscService, _ = service.NewBscChainService(&config.Bsc{
	"https://bsc-testnet-rpc.publicnode.com",
	97,
	"0xd851D918C4970F91453f5Cf50CD59e6f38aE6D5b",
	"0xD856fEc774FA5E7CA8561DE9ef852cb0D94AFE77",
	"",
	"0xe1354798516b08D65160CA5CB2C409b166699013",
	"",
}, database)
var job = NewJob(nil, bscService, database, "")

func TestJob_StartCalculateRate(t *testing.T) {
	job.StartCalculateRate()
}
