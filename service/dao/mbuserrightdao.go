package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"strings"
	"time"
)

const (
	tbNameMbUserRight  = "mb_user_right"
	tbFieldMbUserRight = "id,email,tg_id,right_status,right_status," +
		"ext_info,remark,del_flag,expire_time,create_time,update_time"
)

// MbUserRight  用户权益表
type MbUserRightEntity struct {
	RightId     uint64    `db:"id" json:"id"`                     //  用户权益id,自增主键
	Email       string    `db:"email" json:"email"`               //  用户邮箱
	TgId        string    `db:"tg_id" json:"tg_id"`               //  telegram id
	RightStatus int64     `db:"right_status" json:"right_status"` //  权益状态(0不可用,1可用)
	ExtInfo     string    `db:"ext_info" json:"ext_info"`         //  扩展信息
	Remark      string    `db:"remark" json:"remark"`             //  备注信息
	DelFlag     int64     `db:"del_flag" json:"del_flag"`         //  删除标识(0未删除,1已删除)
	ExpireTime  time.Time `db:"expire_time" json:"expire_time"`   //  权益过期时间
	CreateTime  time.Time `db:"create_time" json:"create_time"`   //  创建时间
	UpdateTime  time.Time `db:"update_time" json:"update_time"`   //  最近更新时间
}

func InsertMbUserRight(ctx context.Context, ens []*MbUserRightEntity) (int64, error) {
	st, traceId := time.Now(), commonlib.GetTrace(ctx)
	conn := commonlib.GetMysqlConn()
	if conn == nil {
		dlog.Errorf(`%s||InsertMbUserRight->GetMysqlConn=nil`, traceId)
		return int64(0), errors.New(`InsertMbUserRight->GetMysqlConn=nil`)
	}

	qs := "INSERT INTO " + tbNameMbUserRight + " (" + tbFieldMbUserRight + ") VALUES (" +
		":id," +
		":email," +
		":tg_id," +
		":right_status," +
		":ext_info," +
		":remark," +
		":del_flag," +
		":expire_time," +
		":create_time," +
		":update_time)"
	ret, err := conn.NamedExec(qs, ens)
	dlog.Infof(`%s||param=%+v,latency=%dms,err=%+v`, traceId, len(ens), time.Since(st).Milliseconds(), err)
	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	dlog.Infof(`%s||InsertMbUserRight->LastInsertId id=%v,err=%+v`, traceId, id, err)

	return id, err
}

func QueryMbUserRightList(ctx context.Context, where map[string]interface{}) (res []*MbCoinEntity, err error) {
	st, traceId := time.Now(), commonlib.GetTrace(ctx)
	conn := commonlib.GetMysqlConn()
	if conn == nil {
		dlog.Errorf(`%s||QueryMbUserRightList->GetMysqlConn=nil`, traceId)
		err = errors.New(`QueryMbUserRightList->GetMysqlConn=nil`)
		return
	}

	whereStr := "id > 0 AND "
	for k, _ := range where {
		whereStr += fmt.Sprintf("%v = :%v AND ", k, k)
	}
	qs := "SELECT " + tbFieldMbCoin + " FROM " + tbNameMbCoin +
		" WHERE " + strings.TrimRight(whereStr, " AND ") +
		" LIMIT 500"

	nstmt, err := conn.PrepareNamed(qs)
	defer commonlib.ReleaseStmt(nstmt)

	if err != nil {
		dlog.Errorf(`%s||QueryMbUserRightList conn.PrepareNamed err=%+v`, traceId, err)
		return
	}

	err = nstmt.Select(&res, where)
	dlog.Infof(`%s||where=%+v,cnt=%d,latency=%dms,err=%v`, traceId, where, len(res), time.Since(st).Milliseconds(), err)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return
}
