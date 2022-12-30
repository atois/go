package demo

import (
	"fmt"
	"testing"

	//"time"

	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	//"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	//vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
	//vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

func initVmClient() *VmClient {
	accessKey := "AB10BD3D9D7C8CB90931B72B74507ABA"
	secretKey := "33177936DE9047C7074007F5B7DBAFA9"
	credentials := NewCredentials(accessKey, secretKey)

	return NewVmClient(credentials)
}

func TestVmGetInstances(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequest("cn-east-2")
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstancesByPage(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequest("cn-east-2")
	req.SetPageNumber(2)
	req.SetPageSize(10)
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}

	if len(resp.Result.Instances) != 10 {
		t.Error("describe instances by paging failed.", len(resp.Result.Instances))
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstanceTypes(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstanceTypesRequest("cn-east-2")
	resp, err := client.DescribeInstanceTypes(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance types failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.InstanceTypes)
}
