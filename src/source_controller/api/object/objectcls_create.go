/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package object

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"encoding/json"
	"fmt"
)

// CreateMetaObjectCls 创建元对象分类, 如果成功则返回 新数据的ID
func (cli *Client) CreateMetaObjectCls(forward *ForwardParam, data []byte) (int, error) {

	if len(data) == 0 {
		return 0, Err_Not_Set_Input
	}

	rst, err := cli.base.HttpCli.POST(fmt.Sprintf("%s/object/v1/meta/object/classification", cli.address), forward.Header, data)
	if nil != err {
		blog.Error("request failed, error:%v", err)
		return 0, Err_Request_Object
	}

	var rstRes ObjClsRsp
	if jserr := json.Unmarshal(rst, &rstRes); nil != jserr {
		blog.Error("can not unmarshal the result , error information is %v, data:%s", jserr, string(rst))
		return 0, jserr
	}

	if rstRes.Code != common.CCSuccess {
		return 0, fmt.Errorf("%v", rstRes.Message)
	}

	return rstRes.Data[0].ID, nil
}
