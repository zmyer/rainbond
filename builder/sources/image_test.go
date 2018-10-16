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

package sources

import (
	"fmt"
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

func TestImageName(t *testing.T) {
	imageName := []string{
		"hub.goodrain.com/nginx:v1",
		"hub.goodrain.cn/nginx",
		"nginx:v2",
		"tomcat",
	}
	for _, i := range imageName {
		in := ImageNameHandle(i)
		fmt.Printf("host: %s, name: %s, tag: %s\n", in.Host, in.Name, in.Tag)
	}
}

func TestBuildImage(t *testing.T) {
	dc, _ := client.NewEnvClient()
	buildOptions := types.ImageBuildOptions{
		Tags:   []string{"goodrain.me/gr1e1a6c_goodrain-apps_mysql:20180307135753"},
		Remove: true,
	}
	if err := ImageBuild(dc, "/Users/qingguo/tmp/nginx", buildOptions, nil, 1); err != nil {
		t.Fatal(err)
	}
}

func TestPushImage(t *testing.T) {
	dc, _ := client.NewEnvClient()
	if err := ImagePush(dc, "hub.goodrain.com/zengqg-test/etcd:v2.2.0", "zengqg-test", "zengqg-test", nil, 2); err != nil {
		t.Fatal(err)
	}
}

func TestTrustedImagePush(t *testing.T) {
	dc, _ := client.NewEnvClient()
	if err := TrustedImagePush(dc, "hub.goodrain.com/zengqg-test/etcd:v2.2.0", "zengqg-test", "zengqg-test", nil, 2); err != nil {
		t.Fatal(err)
	}
}

func TestCheckTrustedRepositories(t *testing.T) {
	err := CheckTrustedRepositories("hub.goodrain.com/zengqg-test/etcd2:v2.2.0", "zengqg-test", "zengqg-test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestImageSave(t *testing.T) {
	dc, _ := client.NewEnvClient()
	if err := ImageSave(dc, "hub.goodrain.com/zengqg-test/etcd:v2.2.0", "/tmp/testsaveimage.tar", nil); err != nil {
		t.Fatal(err)
	}
}

func TestImageImport(t *testing.T) {
	dc, _ := client.NewEnvClient()
	if err := ImageImport(dc, "hub.goodrain.com/zengqg-test/etcd:v2.2.0", "/tmp/testsaveimage.tar", nil); err != nil {
		t.Fatal(err)
	}
}
