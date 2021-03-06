// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"github.com/apache/servicecomb-service-center/server/alarm/model"
	pb "github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"
)

type AlarmListRequest struct {
}

type AlarmListResponse struct {
	Response *pb.Response        `json:"response,omitempty"`
	Alarms   []*model.AlarmEvent `json:"alarms,omitempty"`
}

type ClustersRequest struct {
}

type ClustersResponse struct {
	Response *pb.Response      `json:"response,omitempty"`
	Clusters registry.Clusters `json:"clusters,omitempty"`
}

type ClearAlarmRequest struct {
}

type ClearAlarmResponse struct {
	Response *pb.Response `json:"response,omitempty"`
}
