package hbci

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

func MakeUnencryptedMessage300(msgNum int, dialogId string, segments []string) string {
	const HEAD_LEN int = 29  // the constant part, excluding the length of the message number + dialog ID
	const TRAIL_LEN int = 11 // the constant part, excluding the length of the message number

	content := strings.Join(segments, "")
	//The message number appears in header and trailer and the dialog ID in the header only.
	paddedMsgLen := fmt.Sprintf("%012d", HEAD_LEN+TRAIL_LEN+len(strconv.Itoa(msgNum))*2+len(dialogId)+len(content))

	header := fmt.Sprintf("HNHBK:1:3+%s+300+%s+%d'", paddedMsgLen, dialogId, msgNum)
	trailer := fmt.Sprintf("HNHBS:%d:1+%d'", len(segments)+2, msgNum)

	return fmt.Sprintf("%s%s%s", header, strings.Join(segments, ""), trailer)
}

func MakeMessage300(msgNum int, dialogId string, segments []string, bank *Bank, userId string) string {
  const HEAD_LEN int  = 29 // the constant part, excluding the length of the message number + dialog ID
  const TRAIL_LEN int = 11 // the constant part, excluding the length of the message number

  content     := strings.Join(segments, "")
  encHead     := fmt.Sprintf("HNVSK:998:3+PIN:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+%s:%s:%s:V:0:0+0'", string(bank.Country[:]), bank.Blz, userId)
  encEnvelope := fmt.Sprintf("HNVSD:999:1+@%v@%s'", len(content), content)

  //The message number appears in header and trailer and the dialog ID in the header only.
  paddedMsgLen := fmt.Sprintf("%012d", HEAD_LEN + TRAIL_LEN + len(strconv.Itoa(msgNum)) * 2 + len(dialogId) + len(encHead) + len(encEnvelope));

  header  := fmt.Sprintf("HNHBK:1:3+%s+300+%s+%d'", paddedMsgLen, dialogId, msgNum)
  trailer := fmt.Sprintf("HNHBS:%d:1+%d'", len(segments) + 2,msgNum)

  return fmt.Sprintf("%s%s%s%s", header, encHead, encEnvelope, trailer)
}

func MakeUnencryptedMessage220(msgNum int, dialogId string, segments []string) string {
	const HEAD_LEN int = 29  // the constant part, excluding the length of the message number + dialog ID
	const TRAIL_LEN int = 11 // the constant part, excluding the length of the message number

	content := strings.Join(segments, "")
	//The message number appears in header and trailer and the dialog ID in the header only.
	paddedMsgLen := fmt.Sprintf("%012d", HEAD_LEN+TRAIL_LEN+len(strconv.Itoa(msgNum))*2+len(dialogId)+len(content))

	header := fmt.Sprintf("HNHBK:1:3+%s+220+%s+%d'", paddedMsgLen, dialogId, msgNum)
	trailer := fmt.Sprintf("HNHBS:%d:1+%d'", len(segments)+2, msgNum)

	return fmt.Sprintf("%s%s%s", header, strings.Join(segments, ""), trailer)
}

func MakeMessage220(msgNum int, dialogId string, segments []string, bank *Bank, userId string) string {
	const HEAD_LEN int  = 29 // the constant part, excluding the length of the message number + dialog ID
	const TRAIL_LEN int = 11 // the constant part, excluding the length of the message number

	content     := strings.Join(segments, "")
	encHead     := fmt.Sprintf("HNVSK:998:2+998+1+1::0+1:20140714:142021+2:2:13:@8@\x00\x00\x00\x00\x00\x00\x00\x00:5:1+%s:%s:%s:V:0:0+0'", string(bank.Country[:]), bank.Blz, userId)
	encEnvelope := fmt.Sprintf("HNVSD:999:1+@%v@%s'", len(content), content)

	//The message number appears in header and trailer and the dialog ID in the header only.
	paddedMsgLen := fmt.Sprintf("%012d", HEAD_LEN + TRAIL_LEN + len(strconv.Itoa(msgNum)) * 2 + len(dialogId) + len(encHead) + len(encEnvelope));

	header  := fmt.Sprintf("HNHBK:1:3+%s+220+%s+%d'", paddedMsgLen, dialogId, msgNum)
	trailer := fmt.Sprintf("HNHBS:%d:1+%d'", len(segments) + 2,msgNum)

	return fmt.Sprintf("%s%s%s%s", header, encHead, encEnvelope, trailer)
}

func UnwrapEncryptedData(message string) string {
  regEx := regexp.MustCompile("(HNVSD:\\d+:\\d+)\\+@\\d+\\@(.*)''")
  return regEx.ReplaceAllString(message, "$2'")
}
