package mq

import (
	"hk4e/common/config"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"

	"github.com/nats-io/nats.go"
	"github.com/vmihailenco/msgpack/v5"
	pb "google.golang.org/protobuf/proto"
)

type MessageQueue struct {
	natsConn     *nats.Conn
	natsMsgChan  chan *nats.Msg
	netMsgInput  chan *cmd.NetMsg
	netMsgOutput chan *cmd.NetMsg
	cmdProtoMap  *cmd.CmdProtoMap
}

func NewMessageQueue(netMsgInput chan *cmd.NetMsg, netMsgOutput chan *cmd.NetMsg) (r *MessageQueue) {
	r = new(MessageQueue)
	conn, err := nats.Connect(config.CONF.MQ.NatsUrl)
	if err != nil {
		logger.LOG.Error("connect nats error: %v", err)
		return nil
	}
	r.natsConn = conn
	r.natsMsgChan = make(chan *nats.Msg, 10000)
	_, err = r.natsConn.ChanSubscribe("GATE_HK4E", r.natsMsgChan)
	if err != nil {
		logger.LOG.Error("nats subscribe error: %v", err)
		return nil
	}
	r.netMsgInput = netMsgInput
	r.netMsgOutput = netMsgOutput
	r.cmdProtoMap = cmd.NewCmdProtoMap()
	return r
}

func (m *MessageQueue) Start() {
	go m.startRecvHandler()
	go m.startSendHandler()
}

func (m *MessageQueue) Close() {
	m.natsConn.Close()
}

func (m *MessageQueue) startRecvHandler() {
	for {
		natsMsg := <-m.natsMsgChan
		// msgpack NetMsg
		netMsg := new(cmd.NetMsg)
		err := msgpack.Unmarshal(natsMsg.Data, netMsg)
		if err != nil {
			logger.LOG.Error("parse bin to net msg error: %v", err)
			continue
		}
		if netMsg.EventId == cmd.NormalMsg {
			// protobuf PayloadMessage
			payloadMessage := m.cmdProtoMap.GetProtoObjByCmdId(netMsg.CmdId)
			err = pb.Unmarshal(netMsg.PayloadMessageData, payloadMessage)
			if err != nil {
				logger.LOG.Error("parse bin to payload msg error: %v", err)
				continue
			}
			netMsg.PayloadMessage = payloadMessage
		}
		m.netMsgOutput <- netMsg
	}
}

func (m *MessageQueue) startSendHandler() {
	for {
		netMsg := <-m.netMsgInput
		// protobuf PayloadMessage
		payloadMessageData, err := pb.Marshal(netMsg.PayloadMessage)
		if err != nil {
			logger.LOG.Error("parse payload msg to bin error: %v", err)
			continue
		}
		netMsg.PayloadMessageData = payloadMessageData
		// msgpack NetMsg
		netMsgData, err := msgpack.Marshal(netMsg)
		if err != nil {
			logger.LOG.Error("parse net msg to bin error: %v", err)
			continue
		}
		natsMsg := nats.NewMsg("GS_HK4E")
		natsMsg.Data = netMsgData
		err = m.natsConn.PublishMsg(natsMsg)
		if err != nil {
			logger.LOG.Error("nats publish msg error: %v", err)
			continue
		}
	}
}