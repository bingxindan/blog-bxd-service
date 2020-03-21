package database

import (
	"bxd-middleware-service/config"
	"bxd-middleware-service/utils/log"
	"github.com/go-xorm/xorm"
	"math/rand"
	"os"
	"time"
)

var (
	dbm        *xorm.Engine
	dbs        []*xorm.Engine
	checkCount int
)

func Instance(isMaster, isSlave bool, mConn, sConn string) {
	var err error

	// 主库添加
	if isMaster {
		if mConn == "" || len(mConn) == 0 {
			return
		}
		mDsn := config.Get(mConn)
		dbm, err = xorm.NewEngine("mysql", mDsn)
		dbm.SetMaxIdleConns(10)
		dbm.SetMaxOpenConns(200)
		dbm.ShowSQL(true)
		dbm.ShowExecTime(true)

		if err != nil {
			log.Errorf("Instance master err [%s]", err)
			os.Exit(1)
		}
	}

	// 从库添加
	if isSlave {
		if sConn == "" || len(sConn) == 0 {
			return
		}
		slaves := config.Get(sConn)

		for _, sDsn := range slaves {
			_dbs, err := xorm.NewEngine("mysql", sDsn)
			_dbs.SetMaxIdleConns(10)
			_dbs.SetMaxOpenConns(200)
			_dbs.ShowSQL(true)
			_dbs.ShowExecTime(true)

			if err != nil {
				log.Errorf("Instance slave err [%s]", err)
			} else {
				dbs = append(dbs, _dbs)
			}
		}
	}
}

func Master() *xorm.Engine {
	return dbm
}

func Slave() *xorm.Engine {
	rand.Seed(time.Now().Unix())
	rn := rand.Intn(len(dbs) - 1)
	return dbs[rn]
}