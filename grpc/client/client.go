package Client

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"log"
	"strconv"
	pb "test/grpc/proto"
	"time"

	"google.golang.org/grpc"
)

type BeefRespone struct {
	Beef map[string]int32
}
type TestCase struct {
	request        string
	resultExpected BeefRespone
}

func StartSendRequest() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBeefClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dataTest := []TestCase{
		{
			request: "testtest",
			resultExpected: BeefRespone{
				Beef: map[string]int32{"bresaola": 1, "enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4},
			},
		},
	}

	for _,item := range dataTest{
		result, err := c.BeefSummary(ctx, &pb.Request{Text: item.request})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println("send text : " , string(item.request))
		resultJSON, _ := json.Marshal(result)
		fmt.Println("service grpc response  :" , string(resultJSON))

		var resultMap BeefRespone
		json.Unmarshal(resultJSON, &resultMap)
		fmt.Println("expected value : " ,  item.resultExpected)
		fmt.Println("is correct : " , strconv.FormatBool(reflect.DeepEqual(resultMap.Beef,item.resultExpected.Beef)))
	}
}
