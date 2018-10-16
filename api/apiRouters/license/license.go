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

package license

import (
	"github.com/goodrain/rainbond/api/controller"
	"github.com/goodrain/rainbond/api/util"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Sirupsen/logrus"
)

//License license struct
type License struct{}

//Routes routes
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", controller.GetLicenseManager().AnalystLicense)
	return r
}

//CheckLicense 验证
func CheckLicense(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		licenseCome := r.Header.Get("License")
		logrus.Debug("License is :" + licenseCome)
		if 1 == 1 {
			next.ServeHTTP(w, r)
			return
		}
		util.CloseRequest(r)
		w.WriteHeader(http.StatusUnauthorized)
	}
	return http.HandlerFunc(fn)
}
