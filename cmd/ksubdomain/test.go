package main

import (
	"github.com/Tw1ps/ksubdomain/runner"
	"github.com/urfave/cli/v2"
)

var testCommand = &cli.Command{
	Name:  "test",
	Usage: "测试本地网卡的最大发送速度",
	Action: func(c *cli.Context) error {
		ether := runner.GetDeviceConfig()
		runner.TestSpeed(ether)
		return nil
	},
}
