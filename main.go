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

func shiftBinary(num int64, shift int64) int64 {
	return num << shift
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

func getTimeStampBinary() int64 {
	return shiftBinary(getUnixMill(), int64(machineIDLen) + int64(sequenceIDLen))
}

func getMachineIDBinary() (int64, error) {
	machineID, err := getRandomNum(0, int64(getMaxFromBitLen(machineIDLen)))
	if err != nil {
		return 0, err
	}
	return shiftBinary(machineID, int64(sequenceIDLen)), nil
}

func getSequenceIDBinary() (int64, error) {
	sequenceID, err := getRandomNum(0, int64(getMaxFromBitLen(sequenceIDLen)))
	if err != nil {
		return 0, err
	}
	return sequenceID, nil
}

func getSnowflakeIDBinary() (int64, error) {
	bTimeStamp := getTimeStampBinary()
	bMachineID, err := getMachineIDBinary()
	if err != nil {
		return 0, err
	}
	bSequenceID, err := getSequenceIDBinary()
	if err != nil {
		return 0, err
	}
	return bTimeStamp + bMachineID + bSequenceID, nil
}

func main()  {
	int64SnowFlakeID, err := getSnowflakeIDBinary()
	if err != nil {
		fmt.Println(err)
		return
	}
	bSnowFlakeID, err := padBinary(iToB(int64SnowFlakeID), 63)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("binary:", fmt.Sprintf("0%s", bSnowFlakeID))
	fmt.Println("int64 :", int64SnowFlakeID)
}