package infoverify

import (
	"errors"
	"fmt"
	"log"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/infoverify/apix"
	"github.com/caojunxyz/mimi-server/proto"
	"golang.org/x/net/context"
)

var ErrVerifyFailed = errors.New("信息核验失败!")

type VerifyVendor interface {
	VerifyIdcard(idno, realname string) *dbproto.Idcard
	VerifyBankcard4(idno, realname, bankcardno, phone string) *dbproto.Bankcard
	// TODO: 银行卡通过转账1分钱来验证
}

type VerifyServer struct {
	vendor VerifyVendor
	dbc    dbproto.DbThirdApiAgentClient
}

func NewServer(c dbproto.DbThirdApiAgentClient) *VerifyServer {
	return &VerifyServer{
		vendor: &apix.Apix{},
		dbc:    c,
	}
}

func (srv *VerifyServer) VerifyIdcard(ctx context.Context, arg *proto.InfoVerifyRequest) (*proto.Idcard, error) {
	idno := arg.GetIdcardNo()
	realname := arg.GetRealname()
	idcard, err := srv.dbc.QueryIdcard(ctx, &dbproto.StringValue{Value: idno})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := &proto.Idcard{}
	if idcard.Id == 0 {
		idcard = srv.vendor.VerifyIdcard(idno, realname)
		if idcard != nil {
			_, err = srv.dbc.InsertIdcard(ctx, idcard)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}

	if idcard.Cardno == idno && idcard.Realname == realname {
		result.Cardno = idno
		result.Realname = realname
	}
	return result, nil
}

func (srv *VerifyServer) VerifyBankcard(ctx context.Context, arg *proto.InfoVerifyRequest) (*proto.Bankcard, error) {
	idno := arg.GetIdcardNo()
	realname := arg.GetRealname()
	bankcardNo := arg.GetBankcardNo()
	phone := arg.GetPhone()
	bankcard, err := srv.dbc.QueryBankcard(ctx, &dbproto.StringValue{Value: bankcardNo})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := &proto.Bankcard{}
	if bankcard.Id == 0 {
		bankcard = srv.vendor.VerifyBankcard4(idno, realname, bankcardNo, phone)
		if bankcard != nil {
			_, err = srv.dbc.InsertBankcard(ctx, bankcard)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		} else {
			log.Println("验证失败:", idno, realname, bankcardNo, phone)
			return nil, fmt.Errorf("验证失败")
		}
	}

	if bankcard.IdcardNo == idno && bankcard.Realname == realname && bankcard.BankcardNo == bankcardNo && bankcard.Phone == phone {
		result.IdcardNo = idno
		result.Realname = realname
		result.BankcardNo = bankcardNo
		result.CardType = bankcard.CardType
		result.Bankname = bankcard.Bankname
		result.Phone = bankcard.Phone
	}
	return result, nil
}
