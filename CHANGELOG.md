<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# Apache OpenWhisk Client Go

## 1.2.0

- Update for travis migration (#140)
- Add omit tag to ErrorResponse.Response field (#121)
- update status code if action returns error (#142)
- Migrate to Go Modules (#143)
- Update docs & provide example code (#144)
- Update testify dep. version and clarify use of go get (#145)
- Update go.mod to use go v1.15 (#146)

## 1.1.0

- Add DelAnnotations field to support del annotation (#137)
- Bump up go version to 1.13.14 for travis (#138)
- Add `updated` field on package, rule, and trigger entity (#135)
- Add updated field on Action struct (#129)
- Parse an action's application error (#133)
- Support alt. namespace resource uuid as tenant id to API gateway service (#130)

## 1.0.0

- Handle err return from url.Parse in GetUrlBase (#123)
- Add dynamic column sizing to wsk activation list command (#120)

## 0.10.0-incubating

- Added extra columns to activation records summary rows (#116)
- Replace godep with govendor (#113)
- Load X509 cert on client creation (#112)
- Add Concurrency to Limits (#94)
- Fix invalid warning message (#91)
- Allow NewClient to run concurrently (#103)
- Update Go Version (#104)
- Allow additional headers to override auth header val (#100)
- Replace trigger service type with interface (#99)

## 0.9.0-incubating

- Initial Apache Release
