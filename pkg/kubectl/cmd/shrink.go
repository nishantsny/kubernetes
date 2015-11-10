/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	// "fmt"
	"io"
	// "strings"

	"github.com/spf13/cobra"

	// "k8s.io/kubernetes/pkg/api"
	// "k8s.io/kubernetes/pkg/kubectl"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	// "k8s.io/kubernetes/pkg/kubectl/resource"
	// "k8s.io/kubernetes/pkg/runtime"

	"os"
	"log"
	// "encoding/json"
	// "bytes"

)

// CreateOptions is the start of the data required to perform the operation.
// As new fields are added, add them here instead of referencing the cmd.Flags()
// type ShrinkOptions struct {
// 	// Filenames []string
// }

const (
	shrink_long = `This will shrink total pods, if possible, to a fewer number of nodes`
	shrink_example = `# Call Shrink.
$ kubectl shrink

# Do a dry run.
$ kubectl shrink --dry-run `
)

func NewCmdShrink(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	// options := &ShrinkOptions{}

	cmd := &cobra.Command{
		Use:     "shrink [--dry-run]",
		Short:   "Shrink pods to fit on fewer nodes",
		Long:    shrink_long,
		Example: shrink_example,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(ValidateArgs(cmd, args))
			cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunShrink(f, out, cmd, args))
		},
	}

	// usage := "just call it"
	// kubectl.AddJsonFilenameFlag(cmd, &options.Filenames, usage)
	cmdutil.AddOutputFlagsForMutation(cmd)
	cmd.Flags().BoolP("dry-run", "d", false, "Sets dry run flag")
	return cmd
}

func RunShrink(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, args []string) error {
	dry_run := cmdutil.GetFlagBool(cmd, "dry-run")
	cmdNamespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	// mapper, typer := f.Object()
	// r := resource.NewBuilder(mapper, typer, f.ClientMapperForCommand()).
	// 	ContinueOnError().
	// 	NamespaceParam(cmdNamespace).DefaultNamespace().
	// 	// FilenameParam(enforceNamespace, "").
	// 	SelectorParam("").
	// 	ResourceTypeOrNameArgs(true, args...).
	// 	Flatten().
	// 	Do()
	// err = r.Err()
	// if err != nil {
	// 	return err
	// }

	{

    //////////////////////////////////////////////////////////
		var logFileName2 string = "/home/nishant/test/kuberLogShrink"
		f2, err2 := os.OpenFile(logFileName2, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err2 != nil {
		    log.Printf("ErrorY: %+v\n",err2)
		}
		defer f2.Close()
		log.SetOutput(f2)
		log.Printf("InShrink: dry-run:%v DefaultNamespace:%v \n",dry_run,cmdNamespace)
		// b, err6 := json.Marshal(r)
	 //    if err6 != nil {
	 //        log.Printf("ErrorY :%s",err6) 
	 //    } else {       
	 //        var pretty_parsed_body bytes.Buffer
	 //        if(json.Indent(&pretty_parsed_body,b, "", "  ") != nil) {
	 //            log.Printf("ErrorY")
	 //        } else{
		// 		log.Printf("InShrink_Res_Pretty: %s\n",string(pretty_parsed_body.Bytes()))
	 //        }
	 //    }
    //////////////////////////////////////////////////////////
	}

	return nil
}