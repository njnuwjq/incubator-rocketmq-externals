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
	"strings"
)



type MessageImpl struct {
	Topic      string
	Flag       int
	Properties map[string]string
	Body       []byte
}

func NewMessageImpl()(message *MessageImpl) {
	message = &MessageImpl{}
	return
}

func (m *MessageImpl) SetTopic(topic string) {
	m.Topic = topic
}

func (m *MessageImpl) SetBody(body []byte) {
	m.Body = body
}

//set message tag
func (m *MessageImpl) SetTag(tag string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_TAGS] = tag
}

//get message tag from Properties
func (m *MessageImpl) Tag() (tag string) {
	if m.Properties != nil {
		tag = m.Properties[constant.PROPERTY_TAGS]
	}
	return
}

//set message key
func (m *MessageImpl) SetKeys(keys []string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_KEYS] = strings.Join(keys, " ")
}

//SetDelayTimeLevel
func (m *MessageImpl) SetDelayTimeLevel(delayTimeLevel int) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_DELAY_TIME_LEVEL] = util.IntToString(delayTimeLevel)
}

////SetWaitStoreMsgOK
//func (m *MessageImpl) SetWaitStoreMsgOK(waitStoreMsgOK bool) {
//	if m.Properties == nil {
//		m.Properties = make(map[string]string)
//	}
//	m.Properties[constant.PROPERTY_WAIT_STORE_MSG_OK] = strconv.FormatBool(waitStoreMsgOK)
//}
//GeneratorMsgUniqueKey only use by system
func (m *MessageImpl) GeneratorMsgUniqueKey() {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	if len(m.Properties[constant.PROPERTY_UNIQ_CLIENT_MESSAGE_ID_KEYIDX]) > 0 {
		return
	}
	m.Properties[constant.PROPERTY_UNIQ_CLIENT_MESSAGE_ID_KEYIDX] = util.GeneratorMessageClientId()
}

//GetMsgUniqueKey only use by system
func (m *MessageExtImpl) GetMsgUniqueKey() string {
	if m.Properties != nil {
		originMessageId := m.Properties[constant.PROPERTY_UNIQ_CLIENT_MESSAGE_ID_KEYIDX]
		if len(originMessageId) > 0 {
			return originMessageId
		}
	}
	return m.MsgId
}

//only use by system
func (m *MessageImpl) SetOriginMessageId(messageId string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_ORIGIN_MESSAGE_ID] = messageId
}

//only use by system
func (m *MessageImpl) SetRetryTopic(retryTopic string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_RETRY_TOPIC] = retryTopic
}

//only use by system
func (m *MessageImpl) SetReconsumeTime(reConsumeTime int) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_RECONSUME_TIME] = util.IntToString(reConsumeTime)
}

//only use by system
func (m *MessageImpl) GetReconsumeTimes() (reConsumeTime int) {
	reConsumeTime = 0
	if m.Properties != nil {
		reConsumeTimeStr := m.Properties[constant.PROPERTY_RECONSUME_TIME]
		if len(reConsumeTimeStr) > 0 {
			reConsumeTime = util.StrToIntWithDefaultValue(reConsumeTimeStr, 0)
		}
	}
	return
}

//only use by system
func (m *MessageImpl) SetMaxReconsumeTimes(maxConsumeTime int) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[constant.PROPERTY_MAX_RECONSUME_TIMES] = util.IntToString(maxConsumeTime)
}

//only use by system
func (m *MessageImpl) GetMaxReconsumeTimes() (maxConsumeTime int) {
	maxConsumeTime = 0
	if m.Properties != nil {
		reConsumeTimeStr := m.Properties[constant.PROPERTY_MAX_RECONSUME_TIMES]
		if len(reConsumeTimeStr) > 0 {
			maxConsumeTime = util.StrToIntWithDefaultValue(reConsumeTimeStr, 0)
		}
	}
	return
}