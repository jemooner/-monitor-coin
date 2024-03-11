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
	tbNameMbCoin  = "mb_coin"
	tbFieldMbCoin = "id,coin_name,coin_price,market_tag,ext_info,remark,del_flag,list_time,create_time,update_time"
)

// MbCoin  币种表
type MbCoinEntity struct {
	CoinId     uint64 `db:"id" json:"id"`                  //  用户权益id,自增主键
	CoinName   string `db:"coin_name" json:"coinName"`     //  币名称
	CoinPrice  string `db:"coin_price" json:"coinPrice"`   //  最近价格
	MarketTag  string `db:"market_tag" json:"marketTag"`   //  交易所标识
	ExtInfo    string `db:"ext_info" json:"extInfo"`       //  扩展信息
	Remark     string `db:"remark" json:"remark"`          //  备注信息
	DelFlag    int    `db:"del_flag" json:"delFlag"`       //  删除标识(0未删除,1已删除)
	ListTime   string `db:"list_time" json:"listTime"`     //  发币时间
	CreateTime string `db:"create_time" json:"createTime"` //  创建时间
	UpdateTime string `db:"update_time" json:"updateTime"` //  最近更新时间
}

func InsertMbCoin(ctx context.Context, ens []*MbCoinEntity) (int64, error) {
	st, traceId := time.Now(), commonlib.GetTrace(ctx)
	conn := commonlib.GetMysqlConn()
	if conn == nil {
		dlog.Errorf(`%s||InsertMbCoin->GetMysqlConn=nil`, traceId)
		return int64(0), errors.New(`InsertMbCoin->GetMysqlConn=nil`)
	}

	qs := "INSERT INTO " + tbNameMbCoin + " (" + tbFieldMbCoin + ") VALUES (" +
		":id," +
		":coin_name," +
		":coin_price," +
		":market_tag," +
		":ext_info," +
		":remark," +
		":del_flag," +
		":list_time," +
		":create_time," +
		":update_time)"
	ret, err := conn.NamedExec(qs, ens)
	dlog.Infof(`%s||param=%+v,latency=%dms,err=%+v`, traceId, len(ens), time.Since(st).Milliseconds(), err)
	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	dlog.Infof(`%s||InsertMbCoin->LastInsertId id=%v,err=%+v`, traceId, id, err)

	return id, err
}

func QueryMbCoinList(ctx context.Context, where map[string]interface{}) (res []*MbCoinEntity, err error) {
	st, traceId := time.Now(), commonlib.GetTrace(ctx)
	conn := commonlib.GetMysqlConn()
	if conn == nil {
		dlog.Errorf(`%s||QueryMbCoinList->GetMysqlConn=nil`, traceId)
		err = errors.New(`QueryMbCoinList->GetMysqlConn=nil`)
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
		dlog.Errorf(`%s||QueryMbCoinList conn.PrepareNamed err=%+v`, traceId, err)
		return
	}

	err = nstmt.Select(&res, where)
	dlog.Infof(`%s||where=%+v,cnt=%d,latency=%dms,err=%v`, traceId, where, len(res), time.Since(st).Milliseconds(), err)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return
}

func UpdateMbCoin(ctx context.Context, update map[string]interface{}, where map[string]interface{}) (int64, error) {
	st, traceId := time.Now(), commonlib.GetTrace(ctx)
	conn := commonlib.GetMysqlConn()
	if conn == nil {
		dlog.Errorf(`%s||UpdateMbCoin->GetMysqlConn=nil`, traceId)
		return 0, errors.New(`UpdateMbCoin->GetMysqlConn=nil`)
	}

	argsMap := make(map[string]interface{})
	setStr := ""
	update[`update_time`] = time.Now().Format("2006-01-02 15:04:05")
	for k, v := range update {
		setStr += fmt.Sprintf("%v = :%v, ", strings.Replace(k, "pre_", "", 1), k)
		argsMap[k] = v
	}

	whereStr := ""
	for k, v := range where {
		whereStr += fmt.Sprintf("%v = :%v AND ", strings.Replace(k, "pre_", "", 1), k)
		argsMap[k] = v
	}

	qs := "UPDATE " + tbNameMbCoin +
		" SET " + strings.TrimRight(setStr, ", ") +
		" WHERE " + strings.TrimRight(whereStr, " AND ")

	res, err := conn.NamedExec(qs, argsMap)
	dlog.Infof(`%s||sql=%s;param=%+v,latency=%dms,err=%v`, traceId, qs, argsMap, time.Since(st).Milliseconds(), err)
	if err != nil {
		dlog.Errorf(`%s||UpdateMbCoin->NamedExec fail,err=%s`, traceId, err.Error())
		return 0, err
	}

	row, err := res.RowsAffected()
	dlog.Infof(`%s||UpdateMbCoin->RowsAffected,err=%+v,row=%v`, traceId, err, row)
	return row, err
}
