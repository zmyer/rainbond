// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"net/http"
	"github.com/goodrain/rainbond/db"
	dbmodel "github.com/goodrain/rainbond/db/model"
	httputil "github.com/goodrain/rainbond/util/http"
	"github.com/Sirupsen/logrus"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
)

func GetEventsByIds(w http.ResponseWriter, r *http.Request) {
	b,_:=ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	j,err:=simplejson.NewJson(b)
	if err != nil {
		logrus.Errorf("error decode json,details %s",err.Error())
		httputil.ReturnError(r,w,400,"bad request")
		return
	}
	eventIDS,err:=j.Get("event_ids").StringArray()
	if err != nil {
		logrus.Errorf("error get event_id in json,details %s",err.Error())
		httputil.ReturnError(r,w,400,"bad request")
		return
	}
	result:=[]*dbmodel.ServiceEvent{}
	for _,v:=range eventIDS{
		serviceEvent,err:=db.GetManager().ServiceEventDao().GetEventByEventID(v)
		if err != nil {
			logrus.Warnf("can't find event by given id %s ,details %s",v,err.Error())
			continue
		}
		result=append(result,serviceEvent)
	}
	httputil.ReturnSuccess(r, w, result)
}