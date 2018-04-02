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

	valuesMA := make([]int8, durationMA, durationMA)

	for timer := 0; timer < 60; timer++ {

		d := GetCpuPercent()

		cpuPercent = int8(d[0])
		println("CPU Percent:", cpuPercent)

		valuesMA = append(valuesMA[1:], cpuPercent)

		totalMA = CalculateSMA(valuesMA)

		fmt.Println("SMA:", totalMA, "\n")
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
		//fmt.Printf("val: %v ", v)
		sum += float32(v)
		count++
	}

	average = sum / count
	return average
}
