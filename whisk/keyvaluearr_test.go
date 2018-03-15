// +build unit

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package whisk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValueArrReplaceOrAdd(t *testing.T) {
	kvArr := make(KeyValueArr, 3)
	kvArr[0] = KeyValue{"key0", "value0"}
	kvArr[1] = KeyValue{"key1", "value1"}
	kvArr[2] = KeyValue{"key2", "value2"}

	kvNew := &KeyValue{"keyAdd", "valueAdd"}
	kvArrNew := kvArr.AddOrReplace(kvNew)
	assert.Equal(t, len(kvArrNew), 4)
	assert.Equal(t, kvArrNew.GetValue(kvNew.Key), kvNew.Value)

	kvAdd := &KeyValue{"key3", "valueReplace"}
	kvArrNew = kvArr.AddOrReplace(kvAdd)
	assert.Equal(t, len(kvArrNew), 4)
	assert.Equal(t, kvArrNew.GetValue(kvAdd.Key), kvAdd.Value)
}
