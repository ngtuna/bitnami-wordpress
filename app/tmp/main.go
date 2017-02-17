// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/ngtuna/bitnami-wordpress-aws/app"
	controllers "github.com/ngtuna/bitnami-wordpress-aws/app/controllers"
	tests "github.com/ngtuna/bitnami-wordpress-aws/tests"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "DescribeEC2State",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetIPAddress",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "data", Type: reflect.TypeOf((*[]byte)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Input",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					152: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "EC2",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					157: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RunEC2",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "accessKey", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "secretKey", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "StopEC2",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "data", Type: reflect.TypeOf((*[]byte)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					273: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "TerminateEC2",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "data", Type: reflect.TypeOf((*[]byte)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					305: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Fail",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					309: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
