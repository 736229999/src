package admin

import (
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"log"
)

func (agt *DbAdminAgent) QueryLeagueByName(league_name string, league_id *int) error {
	sql := fmt.Sprintf("SELECT n.id FROM %s AS n WHERE n.name = $1;", TABLE_LEAGUE)
	err := agt.footballDbConn.QueryRow(sql, league_name).Scan(league_id)
	if err != nil {
		log.Println(sql)
		return err
	}
	return nil
}

func (agt *DbAdminAgent) QueryTeamByName(home_name string, guest_name string, home_id *int, guest_id *int) error {
	sql := fmt.Sprintf("SELECT n.id FROM %s AS n WHERE n.name = $1 OR n.name=$2;", TABLE_TEAMS)
	teamResult, err := agt.footballDbConn.Query(sql, home_name, guest_name)
	teamResult.Next()
	err = teamResult.Scan(home_id)
	teamResult.Next()
	err = teamResult.Scan(guest_id)
	if err != nil {
		log.Println(sql)
		return err
	}
	return nil
}

func (agt *DbAdminAgent) QueryLeagueById(league_id int, arg *dbproto.GameInfo) error {
	sql := fmt.Sprintf("SELECT n.name FROM %s AS n WHERE n.id = $1;", TABLE_LEAGUE)
	err := agt.footballDbConn.QueryRow(sql, league_id).Scan(&arg.GameType)
	if err != nil {
		log.Println(sql)
		return err
	}
	return nil
}

func (agt *DbAdminAgent) QueryTeamById(home_id int, guest_id int, arg *dbproto.GameInfo) error {
	sql := fmt.Sprintf("SELECT n.name FROM %s AS n WHERE n.id = $1;", TABLE_TEAMS)
	err := agt.footballDbConn.QueryRow(sql, home_id).Scan(&arg.HomeTeam)
	err = agt.footballDbConn.QueryRow(sql, guest_id).Scan(&arg.GuestTeam)
	if err != nil {
		log.Println(sql, arg.HomeTeam, arg.GuestTeam)
		return err
	}
	return nil
}

func (agt *DbAdminAgent) IsIdExist(id int64, tableName string) int {
	var testId int
	sql := fmt.Sprintf("SELECT id FROM %s where id=$1", tableName)
	err := agt.footballDbConn.QueryRow(sql, id).Scan(&testId)
	if err != nil {
		return 0
	} else {
		return testId
	}
}
