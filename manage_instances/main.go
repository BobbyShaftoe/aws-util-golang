package main

import (
	"fmt"
	_ "fmt"
	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/s3"
	"github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/mem"
	_ "os"
	_ "os/exec"
	_ "strings"
	"time"
	"math"
)

type CPU struct {
	cpu  string
	user float64
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
	}

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", sess)

	// ec2Query := ec2.New(sess)
	// result, err := ec2Query.DescribeInstances(nil)
	//
	// if err  != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	var cpuPercent int8
	var durationMA int8 = 10
	var totalMA float32
	var standardDeviation float32
	var coefficientVar int8

	valuesMA := make([]int8, durationMA, durationMA)

	for timer := 0; timer < 60; timer++ {

		d := GetCpuPercent()

		cpuPercent = int8(d[0])
		valuesMA = append(valuesMA[1:], cpuPercent)
		totalMA = CalculateSMA(valuesMA)
		standardDeviation = CalculateStdDev(valuesMA, totalMA)
		coefficientVar = int8(standardDeviation / totalMA * 100)
		coefficientVar = int8(math.Abs(float64(coefficientVar)))

		fmt.Printf("Current CPU: %v%%\n", cpuPercent)
		fmt.Println("SMA:", totalMA)
		fmt.Println("STD:", standardDeviation)
		fmt.Printf("CV: %v%%\n\n", coefficientVar)

		time.Sleep(1000 * time.Millisecond)
	}
}

func GetCpuPercent() []float64 {
	c, _ := cpu.Percent(0, true)
	return c
}

func CalculateSMA(values []int8) float32 {
	var average float32
	var sum float32
	var count float32

	for _, v := range values {
		sum += float32(v)
		count++
	}
	average = sum / count
	return average
}

func CalculateStdDev(values []int8, average float32) float32 {
	var std float32
	var diff float64
	var count float64

	for _, v := range values {
		diff += math.Pow(float64(v) - float64(average), 2)
		count++
	}
	std = float32(math.Sqrt(diff / count))
	return std
}

