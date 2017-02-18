package controllers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

//Define state struct for instance
type InstanceState struct {
	InstanceId string `json:"instanceId"`
	Code       int    `json:"code"`
	State      string `json:"state"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) DescribeEC2State() revel.Result {

	cred := credentials.NewStaticCredentials(c.Session["AccessKey"], c.Session["SecretKey"], "")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.Redirect(App.Fail)
	}

	svc := ec2.New(sess, &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: cred,
	})

	// Call the DescribeInstances Operation
	var ins []*string
	instanceId := c.Session["InstanceID"]
	ins = append(ins, &instanceId)
	ipt := &ec2.DescribeInstancesInput{}
	ipt.SetInstanceIds(ins)
	resp, err := svc.DescribeInstances(ipt)
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.RenderError(err)
	}

	switch {
	case len(resp.Reservations) == 0:
		return c.RenderText("No instance found")
	case len(resp.Reservations) > 1:
		return c.RenderText("It should not have more than one instance")
	}
	// INSTANCE CODE
	//    * 0 : pending
	//    * 16 : running
	//    * 32 : shutting-down
	//    * 48 : terminated
	//    * 64 : stopping
	//    * 80 : stopped
	instanceStateCode := int(*resp.Reservations[0].Instances[0].State.Code)
	instanceStateString := *resp.Reservations[0].Instances[0].State.Name

	instanceState := InstanceState{InstanceId: c.Session["InstanceID"], Code: instanceStateCode, State: instanceStateString}
	return c.RenderJson(instanceState)
}

func (c App) GetIPAddress() revel.Result {

	cred := credentials.NewStaticCredentials(c.Session["AccessKey"], c.Session["SecretKey"], "")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.Redirect(App.Fail)
	}

	svc := ec2.New(sess, &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: cred,
	})

	// Call the DescribeInstances Operation
	var ins []*string
	instanceId := c.Session["InstanceID"]
	ins = append(ins, &instanceId)
	ipt := &ec2.DescribeInstancesInput{}
	ipt.SetInstanceIds(ins)

	resp, err := svc.DescribeInstances(ipt)
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.RenderError(err)
	}

	switch {
	case len(resp.Reservations) == 0:
		return c.RenderText("No instance found")
	case len(resp.Reservations) > 1:
		return c.RenderText("It should not have more than one instance")
	}

	return c.RenderText(*resp.Reservations[0].Instances[0].PublicIpAddress)
}

func (c App) Input() revel.Result {
	return c.Render()
}

func (c App) EC2() revel.Result {
	return c.Render()
}

func (c App) RunEC2(accessKey, secretKey string) revel.Result {

	//Set data into cookie
	c.Session["AccessKey"] = accessKey
	c.Session["SecretKey"] = secretKey

	cred := credentials.NewStaticCredentials(accessKey, secretKey, "")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session, ", err)
		return c.Redirect(App.Fail)
	}

	svc := ec2.New(sess, &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: cred,
	})

	instanceID, err := c.launchInstance(svc)
	if err != nil {
		return c.RenderError(err)
	}

	c.Session["InstanceID"] = instanceID
	return c.Redirect(App.EC2)
}

func (c App) launchInstance(svc *ec2.EC2) (string, error) {
	params := &ec2.RunInstancesInput{
		ImageId:  aws.String("ami-71c57212"), // bitnami-wordpress-4.7.2-0-linux-ubuntu-14.04.3-x86_64-hvm
		MaxCount: aws.Int64(1),               // Required
		MinCount: aws.Int64(1),               // Required
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{ // Required
				DeviceName: aws.String("/dev/sda1"), // default: /dev/sda1
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					VolumeSize:          aws.Int64(10),     // default: 10 GiB
					VolumeType:          aws.String("gp2"), // default: gp2
				},
			},
		},
		DisableApiTermination:             aws.Bool(false),
		EbsOptimized:                      aws.Bool(false), // default to false
		InstanceInitiatedShutdownBehavior: aws.String("stop"),
		InstanceType:                      aws.String("t2.micro"),       //FIXME: t2.micro
		KeyName:                           aws.String("tuna-singapore"), //FIXME: select keyname
		Monitoring: &ec2.RunInstancesMonitoringEnabled{
			Enabled: aws.Bool(true), // Required
		},
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{ // Required
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination:      aws.Bool(true),
				DeviceIndex:              aws.Int64(0),
				Groups: []*string{
					aws.String("sg-1a36d07e"), // FIXME - Security Group: default
				},
				SubnetId: aws.String("subnet-bebb9cc9"), //FIXME: fixed subnet
			},
		},
		//UserData: aws.String("String"),
	}

	// Call the RunInstances Operation
	resp, err := svc.RunInstances(params)

	if err != nil {
		return "", err
	}

	return *resp.Instances[0].InstanceId, nil
}

func (c App) StopEC2() revel.Result {

	cred := credentials.NewStaticCredentials(c.Session["AccessKey"], c.Session["SecretKey"], "")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.Redirect(App.Fail)
	}

	svc := ec2.New(sess, &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: cred,
	})

	var ins []*string
	instanceId := c.Session["InstanceID"]
	ins = append(ins, &instanceId)
	ipt := &ec2.StopInstancesInput{}
	ipt.SetInstanceIds(ins)

	resp, err := svc.StopInstances(ipt)

	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(resp.StoppingInstances[0].CurrentState.Code)
}

func (c App) TerminateEC2() revel.Result {

	cred := credentials.NewStaticCredentials(c.Session["AccessKey"], c.Session["SecretKey"], "")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return c.Redirect(App.Fail)
	}

	svc := ec2.New(sess, &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: cred,
	})

	var ins []*string
	instanceId := c.Session["InstanceID"]
	ins = append(ins, &instanceId)
	ipt := &ec2.TerminateInstancesInput{}
	ipt.SetInstanceIds(ins)

	resp, err := svc.TerminateInstances(ipt)

	if err != nil {
		return c.RenderError(err)
	}

	return c.Render(resp.TerminatingInstances[0].CurrentState.Code)
}

func (c App) Fail() revel.Result {
	return c.Render()
}
