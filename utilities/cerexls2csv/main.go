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
package main

import (
	//"bufio"
	//"bytes"
	"encoding/csv"
	//"encoding/json"
	"fmt"
	"github.com/humrs/cumed/cuxlsx"
	"github.com/linkedin/goavro"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error

	excelFilename := "/Users/stanley/Desktop/analysis/qry_ForStanleyFinal-2015-11-20.xlsx"
	avroFilename := "/Users/stanley/Desktop/analysis/data.avro"
	avroSchemaFilename := "/Users/stanley/Desktop/analysis/data.avsc"
	csvFilename := "/Users/stanley/Desktop/analysis/data.go.csv"
	csv2Filename := "/Users/stanley/Desktop/analysis/data2.go.csv"

	studentEvals, err := cuxlsx.OpenCEREFile(excelFilename)
	if err != nil {
		log.Fatal(err)
	}

	avscFile, err := ioutil.ReadFile(avroSchemaFilename)
	if err != nil {
		log.Fatal(err)
	}
	avsc := string(avscFile)

	f, err := os.Create(avroFilename)
	if err != nil {
		log.Fatal(err)
	}

	csvf, err := os.Create(csvFilename)
	if err != nil {
		log.Fatal(err)
	}

	csv2f, err := os.Create(csv2Filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		f.Close()
		csvf.Close()
		csv2f.Close()
	}()
	
	fw, err := goavro.NewWriter(
        goavro.BlockSize(10), // example; default is 10
		goavro.Compression(goavro.CompressionSnappy),
        goavro.WriterSchema(avsc),
        goavro.ToWriter(f))
    if err != nil {
        log.Fatal("cannot create Writer: ", err)
    }
	
	for _, eval := range studentEvals {
		record, err := goavro.NewRecord(goavro.RecordSchema(avsc))
		if err != nil {
			log.Fatal(err)
		}
		record.Set("id", eval.ID)
		record.Set("year", int32(eval.Year))
		record.Set("block", int32(eval.Block))
		record.Set("rotation", eval.Rotation)
		record.Set("site", eval.Site)
		record.Set("overall", int32(eval.Overall))
		record.Set("strength", eval.Strength)
		record.Set("weakness", eval.Weakness)
		fw.Write(record)
	}
	fw.Close()
	
	// myoutput, err := json.Marshal(studentEvals)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	csvw := csv.NewWriter(csvf)
	csv2w := csv.NewWriter(csv2f)

	csvrec := []string{"He\nllo", "The,re", "Ag\"ain"}
	fmt.Printf("%v\n", csvrec)
	if err = csv2w.Write(csvrec); err != nil {
		log.Fatal(err)
	}

	var newcsvrec []string

	for _, s := range csvrec {
		newcsvrec = append(newcsvrec, strconv.QuoteToASCII(s))
	}
	fmt.Printf("%v\n", newcsvrec)
	if err = csvw.Write(newcsvrec); err != nil {
		log.Fatal(err)
	}

	csvw.Flush()
	csv2w.Flush()

	if err = csvw.Error(); err != nil {
		log.Fatal(err)
	}

	if err = csv2w.Error(); err != nil {
		log.Fatal(err)
	}

	// var myout bytes.Buffer
	// json.Indent(&myout, myoutput, "", "    ")
	// fmt.Printf("%v\n", myout.String())
}
