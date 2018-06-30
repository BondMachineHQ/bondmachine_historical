package main

import (
	"bondmachine"
	"encoding/json"
	//"errors"
	"etherbond"
	"flag"
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	//"procbuilder"
	//"simbox"
	//"sort"
	//"strconv"
	"strings"
)

type string_slice []string

func (i *string_slice) String() string {
	return fmt.Sprint(*i)
}

func (i *string_slice) Set(value string) error {
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

type Redeployer struct {
	Cluster_file      string
	Bondmachine_ids   []int
	Bondmachine_files []string
	Bondmachine_maps  []string
}

type Transformation struct {
	Transformations []string
}

var debug = flag.Bool("d", false, "Debug")
var verbose = flag.Bool("v", false, "Verbose")

var redeployer_file = flag.String("redeployer-file", "", "Redeployer json file")

var transformations = flag.String("transformations", "", "Transformation json file")

var final_prefix = flag.String("final-prefix", "", "The filename prefix for all the outputs")

var list_objects = flag.Bool("list-objects", false, "List all possible objects")
var list_transformations = flag.Bool("list-transformations", false, "List transformations")
var emit_dot = flag.Bool("emit-dot", false, "Emit dot file on stdout")
var dot_detail = flag.Int("dot-detail", 1, "Detail of infos on dot file 1-5")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.Parse()
}

func main() {
	conf := new(bondmachine.Config)
	conf.Debug = *debug
	conf.Dotdetail = uint8(*dot_detail)

	rd_main := new(Redeployer)
	if *redeployer_file != "" {
		if rd_json, err := ioutil.ReadFile(*redeployer_file); err == nil {
			if err := json.Unmarshal([]byte(rd_json), rd_main); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}

		// Create the Redeployer data struct starting from the json redeployer file
		rd := new(bondmachine.Redeployer)

		rd.Init()

		econfig := new(etherbond.Config)
		// TODO Import register size from whitin BM internals
		econfig.Rsize = uint8(8)

		if rd_main.Cluster_file != "" {
			if cluster, err := etherbond.UnmarshallCluster(econfig, rd_main.Cluster_file); err != nil {
				panic(err)
			} else {
				rd.Cluster = cluster
			}
		} else {
			panic("Wrong Cluster file")
		}

		if len(rd_main.Bondmachine_ids) != len(rd_main.Bondmachine_files) {
			panic("Wrong number of files")
		}
		if len(rd_main.Bondmachine_ids) != len(rd_main.Bondmachine_maps) {
			panic("Wrong number of files")
		}

		for i, id := range rd_main.Bondmachine_ids {
			bmach := new(bondmachine.Bondmachine)
			if bondmachine_json, err := ioutil.ReadFile(rd_main.Bondmachine_files[i]); err == nil {
				var bmachj bondmachine.Bondmachine_json
				if err := json.Unmarshal([]byte(bondmachine_json), &bmachj); err == nil {
					bmach = (&bmachj).Dejsoner()
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}

			rd.Bondmachines[id] = bmach

			iomap := new(bondmachine.IOmap)
			if mapfile_json, err := ioutil.ReadFile(rd_main.Bondmachine_maps[i]); err == nil {
				if err := json.Unmarshal([]byte(mapfile_json), iomap); err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}

			rd.Maps[id] = iomap

		}

		// TODO Complete the importing

		// Start processing options

		if *list_transformations {
			tr := new(Transformation)
			if *transformations != "" {
				if tr_json, err := ioutil.ReadFile(*transformations); err == nil {
					if err := json.Unmarshal([]byte(tr_json), tr); err != nil {
						panic(err)
					}
				} else {
					panic(err)
				}
			}
			for _, trrule := range tr.Transformations {
				fmt.Println(trrule)
			}
		} else if *list_objects {
			// TODO
		} else if *emit_dot {
			fmt.Print(rd.Dot(conf))
		}

		f, err := os.Create(*redeployer_file)
		check(err)
		b, errj := json.Marshal(rd_main)
		check(errj)
		_, err = f.WriteString(string(b))
		check(err)
	}
}
