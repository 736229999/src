package options

import (
	"log"
	"time"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

//添加反馈.
func (agt *DbOptionsAgent) InsertFeedback(ctx context.Context, arg *dbproto.Feedback) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `INSERT INTO feedback (email, name, content, status, create_time) VALUES ($1, $2, $3, $4, $5)`
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetEmail(), arg.GetName(), arg.GetContent(), dbproto.FeedbackStatus_wait, time.Now().Unix())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}
