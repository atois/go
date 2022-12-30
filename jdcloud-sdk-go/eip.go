package demo

import (
	"encoding/json"
	"errors"

	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
)

type VpcClient struct {
	core.JDCloudClient
}

func NewVpcClient(credential *core.Credential) *VpcClient {
	if credential == nil {
		return nil
	}

	config := core.NewConfig()
	config.SetEndpoint("vpc.jdcloud-api.com")

	return &VpcClient{
		core.JDCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "vpc",
			Revision:    "1.0.1",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *VpcClient) SetConfig(config *core.Config) {
	c.Config = *config
}

func (c *VpcClient) SetLogger(logger core.Logger) {
	c.Logger = logger
}

func (c *VpcClient) DisableLogger() {
	c.Logger = core.NewDummyLogger()
}

func (c *VpcClient) DescribeNetworkSecurityGroup(request *vpc.DescribeNetworkSecurityGroupRequest) (*vpc.DescribeNetworkSecurityGroupResponse, error) {
	if request == nil {
		return nil, errors.New("request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkSecurityGroupResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}
