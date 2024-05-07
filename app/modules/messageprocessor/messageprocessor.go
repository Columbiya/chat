package messageprocessor

import (
	"encoding/json"
	"zbchat/app/entities"
)

func ProcessPayload(payload *[]byte, fromUser *entities.User) {
	inmes := entities.IncomingMessage{}
	err := json.Unmarshal(*payload, &inmes)

	if err != nil {
		panic(err)
	}

	// есть одна проблема
	// если оба сокета на линии, но оба не отправили хотя бы одного сообщения в комнату, то сообщение не придет

	// TODO:
	// надо хранить мапу с юзер айдишниками и их коннектами

	fromUser.SendMessage(&inmes.Text, &inmes.ToUserId)
}
