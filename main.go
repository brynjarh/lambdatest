package main
import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/epsagon/epsagon-go/epsagon"
)
type response struct {
	UTC time.Time `json:"utc"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("In myHandler, received body: ", request.Body)
	now := time.Now()
	resp := &response{
		UTC: now.UTC(),
	}
	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}
func main() {
	log.Println("enter main")
	token := os.Getenv("EPSAGON_TOKEN")
	lambda.Start(epsagon.WrapLambdaHandler(
		epsagon.NewTracerConfig("Epsagon Application",token),
		handleRequest))
}
