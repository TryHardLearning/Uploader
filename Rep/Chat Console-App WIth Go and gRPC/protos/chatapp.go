package app

import (
	"log"
	"sync"
	"math/rand"
	"time"
)
type messageUnit struct{
	ID int
	Username string
	Msg string
	MsgCode int
}
type messageHandler struct{
	AllMsg []messageUnit
	Mute sync.Mutex
}
var messageHandlerObect = MessageHandler{}

type ChatServer struct{

} 
func (is *ChatServer) ChatService(service Services_ChatServiceServer) error {
	ID := rand.Intn(1e6)
	e := make(chan error)

	//Receive Messages - Routine 
	go reciveFromStream(service, ID,e)

	//Send messages - Routine
	go sendToStream(service, ID,e)


	return <- e
}

func reciveFromStream(service servicesChatServiceServer,ID int, e chan error){
	for{
		msg, err : service.Recv()
		if(err != nil){
			log.Printf("Erro recive msg from Client")
			e <- err
		}else{
			messageHandlerObect.Mute.Lock()
			messageHandlerObect.AllMsg = append(messageHandlerObect.AllMsg, messageUnit{
				ID := ID,
				Username := msg.myName,
				Msg := msg.myMsg,
				MsgCode := rand.Intn(1e8)
			})
			messageHandlerObect.Mute.Unlock()

			log.Printf("%v Messages", messageHandlerObect.AllMsg[len(messageHandlerObect.AllMsg)-1])

		}

	}
}
func sendToStream(service servicesChatServiceServer,ID int, e chan error){
	for{
		for{
			time.Sleep(500 * time.Millisecond)
			messageHandlerObect.Mute.Lock()

			if(len (messageHandlerObect.AllMsg) == 0){
				messageHandlerObect.Mute.Unlock()
				break
			}
			senderID := messageHandlerObect.AllMsg[0].ID
			senderUsername := messageHandlerObect.AllMsg[0].Username
			SenderMsg := messageHandlerObect.AllMsg[0].Msg

			messageHandlerObect.Mute.Unlock()
			
			if(senderID != ID){
				err := service.Send(&FromServer{myName: senderUsername, myMsg: SenderMsg})
				if (err != nil){
					e <- err
				}

				messageHandlerObect.Mute.Lock()

				if(len(messageHandlerObect.AllMsg) > 1){
					messageHandlerObect = messageHandlerObect.AllMsg[1:]
					//WtF
				}else{
					messageHandlerObect.AllMsg = []messageUnit{}
					//Yes he is sick
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
/*
MQue = AllMsg
errch = e 
csi = service
mu = Mute
UniqueCode = ID
NameCLient =  Username
message4CLient = SenderMsg
*/