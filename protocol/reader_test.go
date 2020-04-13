package protocol_test

import (
	"chat.server.com/protocol"
	"reflect"
	"strings"
	"testing"
)

func TestReadCommand(t *testing.T) {
	tests := []struct{
		input 	string
		result 	interface{}
	}{
		{
			input: 	"SEND Test\n",
			result: protocol.SendCommand{Message: "Test"},
		},
		{
			input: 	"MESSAGE User Test\n",
			result: protocol.MessageCommand{Name: "User", Message: "Test"},
		},
		{
			input: 	"NAME Test\n",
			result: protocol.NameCommand{Name: "Test"},
		},
		{
			input:  "UNDEFINED Test\n",
			result: protocol.UndefinedCommand,
		},
	}

	for _, test := range tests {
		reader := protocol.NewCommandReader(strings.NewReader(test.input))
		// reader.Read 메서드를 이용하면 위에서 넘긴 문자열 명령어가 커맨드 타입으로 변환되어 반환된다.
		result, err := reader.Read()

		if err != nil { // 잘못된 프로토콜의 명령어가 들어올 경우, 파싱하는데 문제가 생겨 EOF 에러를 발생시킨다.
			t.Errorf("Unable to read command %v", err)
		} else if !reflect.DeepEqual(result, test.result) { // reflect.DeepEqual 함수를 이용하여 두 객체가 동일한지 알 수 있다.
			t.Errorf("Command output is not same: %v, %v", test.result, result)
		}
	}
}