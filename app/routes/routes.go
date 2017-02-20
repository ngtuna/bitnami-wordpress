// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) DescribeEC2State(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.DescribeEC2State", args).Url
}

func (_ tApp) GetIPAddress(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetIPAddress", args).Url
}

func (_ tApp) Input(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Input", args).Url
}

func (_ tApp) EC2(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EC2", args).Url
}

func (_ tApp) RunEC2(
		accessKey string,
		secretKey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accessKey", accessKey)
	revel.Unbind(args, "secretKey", secretKey)
	return revel.MainRouter.Reverse("App.RunEC2", args).Url
}

func (_ tApp) StopEC2(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.StopEC2", args).Url
}

func (_ tApp) TerminateEC2(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.TerminateEC2", args).Url
}

func (_ tApp) Fail(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Fail", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


