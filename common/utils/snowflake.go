package utils

import (
	"sync"
	"time"
)

const (
	StartTimestamp   int64 = 1480166465631
	SequenceBits     uint  = 12 // 序号
	MachineBits      uint  = 5  // 机器id
	DataCenterBits   uint  = 5  // 数据中心
	MaxSequence      int64 = -1 ^ (-1 << SequenceBits)
	MaxMachineNum    int64 = -1 ^ (-1 << MachineBits)
	MaxDataCenterNum int64 = -1 ^ (-1 << DataCenterBits)
	MachineLeft      uint  = SequenceBits
	DataCenterLeft   uint  = SequenceBits + MachineBits
	TimestampLeft    uint  = DataCenterLeft + DataCenterBits
)

// SnowFlake ...
type SnowFlake struct {
	mu            sync.Mutex
	dataCenterId  int64
	machineId     int64
	sequence      int64
	lastTimestamp int64
}

// SnowFlakeConf .
type SnowFlakeConf struct {
	DataCenterId int64
	MachineId    int64
}

// NewSnowFlake ...
func NewSnowFlake(c *SnowFlakeConf) *SnowFlake {
	if c.DataCenterId > MaxDataCenterNum || c.DataCenterId < 0 {
		panic("dataCenterId must greater than 0 and less than MaxDataCenterNum")
	}
	if c.MachineId > MaxMachineNum || c.MachineId < 0 {
		panic("machineId must greater than 0 and less than MaxMachineNum")
	}
	return &SnowFlake{
		dataCenterId: c.DataCenterId,
		machineId:    c.MachineId,
	}
}

// NextId generates the next unique ID
func (s *SnowFlake) NextId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	currTimestamp := s.getNewTimestamp()
	if currTimestamp < s.lastTimestamp {
		// 时钟回拨
		return 0
	}

	if currTimestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & MaxSequence
		if s.sequence == 0 {
			currTimestamp = s.getNextMill()
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = currTimestamp

	return (currTimestamp-StartTimestamp)<<TimestampLeft |
		(s.dataCenterId << DataCenterLeft) |
		(s.machineId << MachineLeft) |
		s.sequence
}

func (s *SnowFlake) getNextMill() int64 {
	mill := s.getNewTimestamp()
	for mill <= s.lastTimestamp {
		mill = s.getNewTimestamp()
	}
	return mill
}

func (s *SnowFlake) getNewTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
