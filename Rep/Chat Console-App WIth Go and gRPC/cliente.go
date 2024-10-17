import (
	"bufio"
	"context"
	"fmt"
	"grpcChatServer/chatserver"
	"log"
	"os"
	"strings"
	"google.golang.org/grpc"
)
func main(){
	fmt.Println("Enter Server Ip and Port")
	reader := bufio.NewReader(os.Stdin)
	serverID, err := reader.ReadString('\n')
	if (err != nil) {
		log.Println(Failed to Read the console)
	}
	serverID = strings.TrimSpace(serverID,"/r/n")

	log.Println("Connecting: ",serverID)

	//call CHatService
	con, err := net.Dial(serverID, grpc.WithInsecure())

	if(err != nil) {
		log.Fatalf("Failed To Connect TO gRPC Server: %v"err)

	}
	defer con.Close()

	client : chatserver.NewServicesClient(con)

	stream, err := client.ChatService(context.Background())
	if(err != nil) {
		log.Fatalf("Failed to connect whit Chate Service: %v "err)
	}
	//Implement Communication with gRPC Server

	type clienthandle struct{
		stream chatserver.Services_ChatServiceClient
		clientName string
	}
	func (ch *clienthandle) clientConfig(){
		reader := bufio.NewReader()
		fmt.Printf("Your Nme Please: ")
		name, err := reader.ReadString('\n')
		if(err != nil){
			log.Fatalf("Failed to read your Name")
		}
		ch.clientName = strings.Trim(name,"\r\n")
	}
	func (ch *clienthandle) sendMessage(){

		for{
			reader := bufio.NewReader(os.Stdin)

			clientMessage, err := reader.ReadString('\n')
			if(err != nil){
				log.Fatalf("Failed to read Msg from Console")
			}
			clientMessage = strings.Trim(clientMessage,"\r\t")

			clientMessageBox := &chatserver.FromClient{
				myName: ch.clientName,
				myMsg: clientMessage,
			}

			err = ch.stream.Send(clientMessageBox)

			if (err != nil){
				log.Println("Failed to send Message to te Server")
			}
		}
	}
	func (ch *clienthandle) reciveMessage(){

		for{
			msg, err := ch.stream.Recv()
			if(err != nil){
				log.Println("Failed to recive Message the message from the Server")
			}

			fmt.Printf("%s : %s\n", msg.myName, msg.myMsg)
		}
	}
}