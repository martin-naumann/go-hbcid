package hbci

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeUnencryptedMessage300(t *testing.T) {
	segments := []string{"HKIDN:2:2+280:10090000+9999999999+0+0'", "HKVVB:3:2+0+0+0+GoBanking+1.0'"}
	msg := MakeUnencryptedMessage300(15, "0", segments)
	expectedMsg := "HNHBK:1:3+000000000113+300+0+15'HKIDN:2:2+280:10090000+9999999999+0+0'HKVVB:3:2+0+0+0+GoBanking+1.0'HNHBS:4:1+15'"

  assert.Equal(t, msg, expectedMsg, "Message is invalid")
}

func TestMakeMessage300(t *testing.T) {
	segments := []string{"HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:12345678:USER:S:0:0'", "HKIDN:3:2+280:12345678+USER+0+1'", "HKVVB:4:3+0+0+0+GoBanking+1.0'", "HKSYN:5:3+0'", "HNSHA:6:2+1337++SECRET'"}
	bank := Bank{"http://not-used.local", "Testbank", [...]byte{'2', '8', '0'}, "12345678"}
	msg := MakeMessage300(15, "0", segments, &bank, "USER")
	expectedMsg := "HNHBK:1:3+000000000350+300+0+15'HNVSK:998:3+PIN:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+280:12345678:USER:V:0:0+0'HNVSD:999:1+@191@HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:12345678:USER:S:0:0'HKIDN:3:2+280:12345678+USER+0+1'HKVVB:4:3+0+0+0+GoBanking+1.0'HKSYN:5:3+0'HNSHA:6:2+1337++SECRET''HNHBS:7:1+15'"

  assert.Equal(t, msg, expectedMsg, "Message is invalid");
}

func TestMakeUnencryptedMessage220(t *testing.T) {
	segments := []string{"HKIDN:2:2+280:10090000+9999999999+0+0'", "HKVVB:3:2+0+0+0+GoBanking+1.0'"}
	msg := MakeUnencryptedMessage220(15, "0", segments)
	expectedMsg := "HNHBK:1:3+000000000113+220+0+15'HKIDN:2:2+280:10090000+9999999999+0+0'HKVVB:3:2+0+0+0+GoBanking+1.0'HNHBS:4:1+15'"

	assert.Equal(t, msg, expectedMsg, "Message is invalid")
}

func TestMakeMessage220(t *testing.T) {
	segments := []string{"HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:12345678:USER:S:0:0'", "HKIDN:3:2+280:12345678+USER+0+1'", "HKVVB:4:3+0+0+0+GoBanking+1.0'", "HKSYN:5:3+0'", "HNSHA:6:2+1337++SECRET'"}
	bank := Bank{"http://not-used.local", "Testbank", [...]byte{'2', '8', '0'}, "12345678"}
	msg := MakeMessage220(15, "0", segments, &bank, "USER")
	expectedMsg := "HNHBK:1:3+000000000344+220+0+15'HNVSK:998:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+280:12345678:USER:V:0:0+0'HNVSD:999:1+@191@HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:12345678:USER:S:0:0'HKIDN:3:2+280:12345678+USER+0+1'HKVVB:4:3+0+0+0+GoBanking+1.0'HKSYN:5:3+0'HNSHA:6:2+1337++SECRET''HNHBS:7:1+15'"

	assert.Equal(t, msg, expectedMsg, "Message is invalid");
}

func TestUnwrapEncryptedData(t *testing.T) {
	encryptedMsg := "HNHBK:1:3+000000000362+300+0+1'HNVSK:998:3+PIN:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+280:BLZ:USER:V:0:0+0'HNVSD:999:1+@200@HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:BLZ:USER:S:0:0'HKIDN:3:2+280:BLZ+USER+0+1'HKVVB:4:3+0+0+0+GoBanking+1.0'HKSYN:5:3+0'HNSHA:6:2+1337++PIN''HNHBS:7:1+1'"
	unwrappedMsg := "HNHBK:1:3+000000000362+300+0+1'HNVSK:998:3+PIN:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+280:BLZ:USER:V:0:0+0'HNSHK:2:4+PIN:2+900+1337+1+1+1::0+1+1:20140711:151800+1:999:1+6:10:16+280:BLZ:USER:S:0:0'HKIDN:3:2+280:BLZ+USER+0+1'HKVVB:4:3+0+0+0+GoBanking+1.0'HKSYN:5:3+0'HNSHA:6:2+1337++PIN'HNHBS:7:1+1'"
  result := UnwrapEncryptedData(encryptedMsg)
	assert.Equal(t, result, unwrappedMsg, "The message should not contain the HNVSD segment overhead")
	assert.Equal(t, result, unwrappedMsg, "UnwrapEncryptedData should not change a non-encrypted message")
}
