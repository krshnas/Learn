// package main

// import (
// 	"fmt"
// 	"time"
// )

// func greet(phrase string, doneChan chan bool) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// 	close(doneChan)
// }

// func main() {
// 	done := make(chan bool)
// 	// dones := make([]chan bool, 4)

// 	// dones[0] = make(chan bool)
// 	go greet("Nice to meet you!", done)

// 	// dones[1] = make(chan bool)
// 	go greet("How are you?", done)

// 	// dones[2] = make(chan bool)
// 	go slowGreet("How ... are ... you ...?", done)

// 	// dones[3] = make(chan bool)
// 	go greet("I hope you're liking the course!", done)

// 	// for _, done := range dones {
// 	// 	<-done
// 	// }

// 	for range done {
// 		// fmt.Println(doneChan)
// 	}
// }


package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// Load the default AWS SDK config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Call DescribeRegions to get the available AWS regions
	output, err := ec2Client.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{})
	if err != nil {
		log.Fatalf("unable to describe regions, %v", err)
	}

	// Print each region's name
	fmt.Println("Available AWS Regions:")
	for _, region := range output.Regions {
		fmt.Println(*region.RegionName)
	}
}
