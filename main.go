package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	timestampLen = 41
	machineIDLen = 10
	sequenceIDLen = 12
)

func padBinary(binaryStr string, binaryLen int) (string, error) {
	if len(binaryStr) > binaryLen {
		return "", fmt.Errorf("padBinary Error: overflow")
	} else {
		return strings.Repeat("0", binaryLen - len(binaryStr)) + binaryStr, nil
	}
}

func iToB(num int64) string {
	return strconv.FormatInt(num, 2)
}

func getUnixMill() int64  {
	now := time.Now()
	unixNanoTime := now.UnixNano()
	unixMilliTime := unixNanoTime / int64(time.Millisecond)
	return unixMilliTime
}

func getRandomNum(min, max int64) (int64, error) {
	bigIntMax := big.NewInt(max)
	n, err := rand.Int(rand.Reader, bigIntMax)
	if err != nil {
		return 0, err
	}
	randomNum := n.Int64() + min
	return randomNum, nil
}

func getMaxFromBitLen(len int) int {
	max := 0
	work := 1;
	for i:=0; i<len; i++ {
		max += work
		work *= 2
	}
	return max
}

func getTimeStampBinary() (string, error) {
	timestamp := getUnixMill()
	bTimeStamp := iToB(timestamp)
	bTimeStamp, err := padBinary(bTimeStamp, timestampLen)
	if err != nil {
		return "", err
	}
	return bTimeStamp, nil
}

func getMachineIDBinary() (string, error) {
	machineID, err := getRandomNum(0, int64(getMaxFromBitLen(machineIDLen)))
	if err != nil {
		return "", err
	}
	bMachineID := iToB(machineID)
	bMachineID, err = padBinary(bMachineID, machineIDLen)
	if err != nil {
		return "", err
	}
	return bMachineID, err
}

func getSequenceIDBinary() (string, error) {
	sequenceID, err := getRandomNum(0, int64(getMaxFromBitLen(sequenceIDLen)))
	if err != nil {
		return "", err
	}
	bSequenceID := iToB(sequenceID)
	bSequenceID, err = padBinary(bSequenceID, sequenceIDLen)
	if err != nil {
		return "", err
	}
	return bSequenceID, err
}

func getSnowflakeIDBinary() (string, error) {
	bTimeStamp, err := getTimeStampBinary()
	if err != nil {
		return "", err
	}
	machineID, err := getMachineIDBinary()
	if err != nil {
		return "", err
	}
	sequenceID, err := getSequenceIDBinary()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0%s%s%s", bTimeStamp, machineID, sequenceID), nil
}

func main()  {
	bSnowFrakeID, err := getSnowflakeIDBinary()
	if err != nil {
		fmt.Println(err)
		return
	}
	int64SnowFrakeID, err := strconv.ParseInt(bSnowFrakeID, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("binary:",bSnowFrakeID)
	fmt.Println("int64 :",int64SnowFrakeID)
}