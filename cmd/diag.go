// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// diagCmd represents the diag command
var diagCmd = &cobra.Command{
	Use:   "diag",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: diagRunCmd,
}

func init() {
	rootCmd.AddCommand(diagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func diagRunCmd(cmd *cobra.Command, args []string) {
	//s, err := host.Init()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println("Initialized host:")
	//fmt.Println(s)

	//// Enumerate all I²C buses available and the corresponding pins.
	//fmt.Print("I²C buses available:\n")
	//for _, ref := range i2creg.All() {
	//	fmt.Printf("- %s\n", ref.Name)
	//	if ref.Number != -1 {
	//		fmt.Printf("  %d\n", ref.Number)
	//	}
	//	if len(ref.Aliases) != 0 {
	//		fmt.Printf("  %s\n", strings.Join(ref.Aliases, " "))
	//	}

	//	b, err := ref.Open()
	//	if err != nil {
	//		fmt.Printf("  Failed to open: %v", err)
	//	}
	//	if p, ok := b.(i2c.Pins); ok {
	//		fmt.Printf("  SDA: %s", p.SDA())
	//		fmt.Printf("  SCL: %s", p.SCL())
	//	}
	//	if err := b.Close(); err != nil {
	//		fmt.Printf("  Failed to close: %v", err)
	//	}
	//}

	//fmt.Println("Attempting to open default i2c bus...")
	//b, err := i2creg.Open("")
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Successfully opened default i2c bus.")
	//}
	//defer b.Close()

	//// Open the DS248x to get a 1-wire bus.
	//fmt.Println("Attempting to open 1-wire bus via ds248x...")
	//ob, err := ds248x.New(b, 0x18, &ds248x.DefaultOpts)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Successfully opened 1-wire bus via ds248x.")
	//}
	//// Search devices on the bus
	//devices, err := ob.Search(false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Found %d 1-wire devices: ", len(devices))
	//for _, d := range devices {
	//	fmt.Printf(" %#16x", uint64(d))
	//}
	//fmt.Print("\n")

	//fmt.Println("Attempting to create ds18b20 handle...")
	//dev, err := ds18b20.New(ob, 0x2800000854dafe, 9)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Successfully created ds18b20 handle.")
	//}

	//fmt.Println("Attempting to read ds18b20...")
	//t := &controller.TempShim{dev}
	//p := &controller.PIDShim{}
	//f := controller.FreezerController{p, t}
	//temp, err := f.GetTemp()

	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(temp)

	//fmt.Println("All functionality succeeded.")
}
