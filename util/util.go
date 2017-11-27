package util

import (
	"net/http"
	"io/ioutil"
	"github.com/siddontang/go/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/hex"
)

var mgoSessioon *mgo.Session

const idsURL = "http://utils.instreet.cc:1273/id"
const mongoURL = "10.31.155.86:27017"
const remote = false

func GetObjectId() string {
	if remote {
		r, _ := http.Get(idsURL)
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if nil != err {
			log.Error(err)
		}
		return string(body)
	} else {
		return getID()
	}
}

func getSession() *mgo.Session {
	if nil == mgoSessioon {
		var err error
		mgoSessioon, err = mgo.Dial(mongoURL)
		if nil != err {
			log.Error(err)
		}
	}
	return mgoSessioon.Clone()
}

func getID() string {
	objectId := bson.NewObjectId()
	objIdStr := hex.EncodeToString([]byte(objectId))
	log.Info(objIdStr)
	return objIdStr
}
