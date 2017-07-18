package options

import (
	"log"
	"testing"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func ConnectDbOptionsAgent() dbproto.DbOptionsAgentClient {
	conn, err := grpc.Dial("127.0.0.1:6012", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// DbClient := dbproto.DbOptionsAgentClient{}
	DbClient := dbproto.NewDbOptionsAgentClient(conn)
	return DbClient
}

func TestQueryClientBannerLIst(t *testing.T) {
	client := ConnectDbOptionsAgent()

	arg := &dbproto.QueryClientBannerArg{
		Location: dbproto.QueryClientBannerArg_Location_Discover,
	}
	res, err := client.QueryClientBannerList(context.Background(), arg)
	if err != nil {
		t.Fatalf("QueryClientBannerList %v\n", err)
		return
	}
	log.Printf("res is %+v\n", res)
}

func TestQueryBannerById(t *testing.T) {
	client := ConnectDbOptionsAgent()

	arg := &dbproto.BannerId{
		Id: 5,
	}

	res, err := client.QueryBannerById(context.Background(), arg)
	if err != nil {
		t.Errorf("QueryBannerById %v\n", err)
		return
	}
	log.Printf("res is %+v\n", res)
}
