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
 
package inst_test

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/config"
	"configcenter/src/framework/core/output/module/client"
	"configcenter/src/framework/core/output/module/inst"
	"configcenter/src/framework/core/output/module/model"
	//"configcenter/src/framework/core/types"
	"testing"
)

func TestInstBusinessManager(t *testing.T) {

	client.NewForConfig(config.Config{"core.supplierAccount": "0", "core.user": "build_user", "core.ccaddress": "http://test.apiserver:8080"}, nil)

	clsItem, err := model.FindClassificationsByCondition(common.CreateCondition().Field("bk_classification_id").Eq("bk_organization"))
	if nil != err {
		t.Errorf("failed to find classifications, %s", err.Error())
		return
	}

	if nil == clsItem {
		t.Errorf("not found the host classification")
		return
	}

	clsItem.ForEach(func(item model.Classification) error {

		modelIter, err := item.FindModelsByCondition(common.CreateCondition().Field("bk_obj_id").Eq("biz"))
		if nil != err {
			t.Errorf("failed to search classification, %s", err.Error())
			return nil
		}

		if nil == modelIter {
			t.Log("not found the model")
			return nil
		}

		// deal common inst model
		modelIter.ForEach(func(modelItem model.Model) error {

			// create a common inst
			commonInst, err := inst.CreateInst(modelItem)
			if nil != err {
				t.Errorf("failed to create inst ")
				return nil
			}

			// Only test
			t.Logf("model name:%s %s", commonInst.GetModel().GetName(), commonInst.GetModel().GetID())

			// set inst value
			commonInst.SetValue("bk_biz_name", "testBiz2")
			commonInst.SetValue("bk_biz_maintainer", "build_user2")

			// save inst info
			err = commonInst.Save()

			if nil != err {
				t.Errorf("failed to save ,%s", err.Error())
			}
			return nil
		})
		return nil

	})

}
