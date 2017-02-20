package controllers

import (
	"fmt"
	"os"

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

var imageID, instanceType, region, keyName, securityGroup, subnet string

func init() {
	imageID = os.Getenv("IMAGEID")
	if imageID == "" {
		fmt.Fprintf(os.Stderr, "AMI is required")
		os.Exit(1)
	}

	instanceType = os.Getenv("INSTANCETYPE")
	if instanceType == "" {
		fmt.Fprintf(os.Stderr, "Instance type is required.")
		os.Exit(1)
	}

	region = os.Getenv("AWSREGION")
	if region == "" {
		fmt.Fprintf(os.Stderr, "Region is required")
		os.Exit(1)
	}

	keyName = os.Getenv("KEYNAME")
	if keyName == "" {
		fmt.Fprintf(os.Stderr, "Keyname is required")
		os.Exit(1)
	}

	securityGroup = os.Getenv("SECURITYGROUP")
	if securityGroup == "" {
		fmt.Fprintf(os.Stderr, "Security group is required")
		os.Exit(1)
	}

	subnet = os.Getenv("SUBNET")
	if subnet == "" {
		fmt.Fprintf(os.Stderr, "Subnet is required")
		os.Exit(1)
	}
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
		Region:      aws.String(region),
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
		Region:      aws.String(region),
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
		Region:      aws.String(region),
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
		ImageId:  aws.String(imageID), // bitnami-wordpress-4.7.2-0-linux-ubuntu-14.04.3-x86_64-hvm
		MaxCount: aws.Int64(1),
		MinCount: aws.Int64(1),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					VolumeSize:          aws.Int64(10),
					VolumeType:          aws.String("gp2"),
				},
			},
		},
		DisableApiTermination:             aws.Bool(false),
		EbsOptimized:                      aws.Bool(false),
		InstanceInitiatedShutdownBehavior: aws.String("stop"),
		InstanceType:                      aws.String(instanceType),
		KeyName:                           aws.String(keyName),
		Monitoring: &ec2.RunInstancesMonitoringEnabled{
			Enabled: aws.Bool(true),
		},
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination:      aws.Bool(true),
				DeviceIndex:              aws.Int64(0),
				Groups: []*string{
					aws.String(securityGroup),
				},
				SubnetId: aws.String(subnet),
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
		Region:      aws.String(region),
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
		Region:      aws.String(region),
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
