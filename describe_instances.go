// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Describes your Amazon EC2 instances.]
// snippet-keyword:[Amazon Elastic Compute Cloud]
// snippet-keyword:[DescribeInstances function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[ec2]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
   Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	//"reflect"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Instance-ID", "Status"})
	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		//fmt.Println("Success", result.Reservations)
		for _, r := range result.Reservations {
			//fmt.Println("Reservation ID: " + *r.ReservationId)
			//fmt.Println("Instance IDs:")
			for _, i := range r.Instances {
				//fmt.Println("Instance IDs:" + *i.InstanceId)
				//fmt.Println("Public IP:" + *i.PublicIpAddress)
				//fmt.Println(*i.State.Name)
				t.AppendRows([]table.Row{
					{1, *i.InstanceId, *i.State.Name},
				})
				t.AppendSeparator()
			}

			//fmt.Println("")
		}
	}
	//fmt.Println(reflect.TypeOf(result))

	t.Render()
}
