// Copyright 2016 R. Stanley Hum
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cuxlsx

import (
	"fmt"
	"github.com/humrs/cumed/clerkship"
	"github.com/tealeg/xlsx"
	"log"
)

func parseEvalRow(cells []*xlsx.Cell) (clerkship.StudentEval, error) {
	var result clerkship.StudentEval
	var err error
	result.ID, err = cells[0].String()
	if err != nil {
		fmt.Printf("Problem with Id: %v\n", err)
	}
	result.Year, err = cells[1].Int()
	if err != nil {
		fmt.Printf("Problem with Year: %v:%v\n", result.ID, err)
	}
	result.Block, err = cells[2].Int()
	if err != nil {
		fmt.Printf("Problem with Block: %v:%v\n", result.ID, err)
	}
	result.Rotation, err = cells[3].String()
	if err != nil {
		fmt.Printf("Problem with Rotation: %v\n", err)
	}
	result.Site, err = cells[4].String()
	if err != nil {
		fmt.Printf("Problem with Site: %v\n", err)
	}
	result.Overall, err = cells[5].Int()
	if err != nil {
		fmt.Printf("Problem with Overall: %v\n", err)
	}
	result.Strength, err = cells[6].String()
	if err != nil {
		fmt.Printf("Problem with Strength: %v\n", err)
	}
	result.Weakness, err = cells[7].String()
	if err != nil {
		fmt.Printf("Problem with Weakness: %v\n", err)
	}
	return result, err
}

// OpenCEREFile - opens data file given to us by CERE
func OpenCEREFile(fn string) ([]clerkship.StudentEval, error) {
	var data []clerkship.StudentEval
	var myval clerkship.StudentEval
	var err error

	data = make([]clerkship.StudentEval, 0)

	xlFile, err := xlsx.OpenFile(fn)
	if err != nil {
		log.Fatalf("Problem reading excelfile: %v", err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			myval, err = parseEvalRow(row.Cells)
			if err != nil {
				fmt.Printf("error %v\n", err)
				continue
			} else {
				fmt.Printf("appending: %v\n", myval.ID)
				data = append(data, myval)
			}
		}
	}

	return data, err
}
