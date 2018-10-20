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
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aqua/raspberrypi/onewire"
	"github.com/brnsampson/MrFreeze/controller"
	"github.com/felixge/pidctrl"
	"github.com/spf13/cobra"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/bcm283x"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: startRunCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startRunCmd(cmd *cobra.Command, args []string) {
	s, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initialized host:")
	fmt.Println(s)

	switchPin := bcm283x.GPIO25
	ps := controller.PinShim{switchPin}

	dev, err := onewire.NewDS18S20("28-00000854dafe")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Opened temperature device")
	}

	pid := pidctrl.NewPIDController(5, .001, 1)

	p := controller.PIDShim{Ctrl: pid}
	t := controller.TempShim{Dev: dev}

	f := controller.FreezerController{&p, &t, &ps}

	f.Set(69.5)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			f.Off()
			t := time.Now()
			log.Printf("%v Encountered SIGINT. Exiting...", t.Format(time.RFC822))
			os.Exit(0)
		}
	}()

	for {
		err := f.TestTemp()
		if err != nil {
			t := time.Now()
			log.Printf("%v Encountered error: %v", t.Format(time.RFC822), err)
			fmt.Printf("\n")
		}

		time.Sleep(10 * time.Second)
	}
}
