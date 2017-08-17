/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package message

import (
	"github.com/apache/incubator-rocketmq-externals/rocketmq-go/model/constant"
	"github.com/apache/incubator-rocketmq-externals/rocketmq-go/util"
	"math"
)




type MessageExtImpl struct {
	*MessageImpl
	QueueId                   int32
	StoreSize                 int32
	QueueOffset               int64
	SysFlag                   int32
	BornTimestamp             int64
	BornHost                  string
	StoreTimestamp            int64
	StoreHost                 string
	MsgId                     string
	CommitLogOffset           int64
	BodyCRC                   int32
	ReconsumeTimes            int32
	PreparedTransactionOffset int64

	propertyConsumeStartTimestamp string
}

//get message topic
func (m *MessageExtImpl)Topic2()(topic string){
	topic = m.Topic
	return
}
//get message tag
func (m *MessageExtImpl)Tag2() (tag string){
	tag = m.Tag()
	return
}

// get body
func (m *MessageExtImpl)Body2()(body []byte){
	body =  m.Body
	return
}
func (m *MessageExtImpl)MsgId2()(msgId string){
	msgId = m.MsgId
	return
}

func (m *MessageExtImpl) GetOriginMessageId() string {
	if m.Properties != nil {
		originMessageId := m.Properties[constant.PROPERTY_ORIGIN_MESSAGE_ID]
		if len(originMessageId) > 0 {
			return originMessageId
		}
	}
	return m.MsgId
}

func (m *MessageExtImpl) GetConsumeStartTime() int64 {
	if len(m.propertyConsumeStartTimestamp) > 0 {
		return util.StrToInt64WithDefaultValue(m.propertyConsumeStartTimestamp, -1)
	}
	return math.MaxInt64
}

func (m *MessageExtImpl) SetConsumeStartTime() {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	nowTime := util.CurrentTimeMillisStr()
	m.Properties[constant.PROPERTY_KEYS] = nowTime
	m.propertyConsumeStartTimestamp = nowTime
	return
}