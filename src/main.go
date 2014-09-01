package main

import (
  "fmt"
  "flag"
  "./hbci"
  "strconv"
  "strings"
  "net/http"
  "io/ioutil"
  "./persistence"
  "encoding/base64"
  "github.com/golang/glog"
  "github.com/dchest/uniuri"
)

type MessageContext struct {
  NextMsgNum int
  DialogId string
  SystemId string
  Segments []string
  UserId string
  Pin string
  Bank hbci.Bank
}

func ParseIncomingMessage(message string) (MessageContext, error) {
  var msgContext MessageContext
  var responseSegments []string

  segments := strings.Split(message, "'")

  for _, segment := range(segments) {
    if len(segment) < 6 || strings.Index(segment, ":") < 0 {
      continue
    }

    dataElems := strings.Split(segment, "+")
    segHead := strings.Split(dataElems[0], ":")
    switch segHead[0] {

      case "HNHBK":
      if len(dataElems) < 5 {
        seg := fmt.Sprintf("HIRMS:%d:2+9010:%s:Syntax error'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }
      if dataElems[3] == "0" {
        msgContext.DialogId = uniuri.New()
        glog.Infof("Connection needs Dialog ID... setting %s", msgContext.DialogId)
      } else {
        msgContext.DialogId = dataElems[3]
        glog.Infof("Dialog ID found: %s", msgContext.DialogId)
      }
      parsedMsgNum, err := strconv.Atoi(dataElems[4])
      if err != nil {
        msgContext.NextMsgNum = 2 // 1 is usually the DialogInitialisation message number, so 2 should be good for the next message number.
      } else {
        msgContext.NextMsgNum = parsedMsgNum + 1
      }
      break

      case "HKIDN":
      if len(dataElems) < 5 {
        seg := fmt.Sprintf("HIRMS:%d:2+9010:%s:Syntax error'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }
      bankInfo := strings.Split(dataElems[1], ":")
      msgContext.UserId = dataElems[2]
      msgContext.Bank.Blz = bankInfo[1]
      msgContext.Bank.Country = [...]byte{'2', '8', '0'}
      glog.Infof("Identification for User %s at BLZ %s", msgContext.UserId, msgContext.Bank.Blz)
      break

      case "HKVVB":
      if len(dataElems) < 6 {
        seg := fmt.Sprintf("HIRMS:%d:2+9010:%s:Syntax error'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }
      glog.Infof("The connection is initiated by \"%s\" version %s", dataElems[4], dataElems[5])
      break

      case "HKSYN":
      if len(dataElems) < 2 {
        seg := fmt.Sprintf("HIRMS:%d:2+9010:%s:Syntax error'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }
      if dataElems[1] == "0" {
        msgContext.SystemId = uniuri.New()
        glog.Infof("Connection needs System ID... setting %s", msgContext.SystemId)
        seg := fmt.Sprintf("HISYN:%d:2:%d+%s'", len(responseSegments)+2, segHead[1], msgContext.SystemId)
        responseSegments = append(responseSegments, seg)

      }
      break

      case "HNSHA":
      if len(dataElems) < 4 {
        seg := fmt.Sprintf("HIRMS:%d:2+9010:%s:Syntax error'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }
      msgContext.Pin = dataElems[3]
      glog.Infof("Signed with PIN %s", msgContext.Pin)

      result, err := persistence.Load(msgContext.UserId)
      if err != nil {
        glog.Warningf("(Auth) Cannot read from redis: %s", err)
        seg := fmt.Sprintf("HIRMS:%d:2+9931:%s:Invalid Login'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
        break
      }

      if msgContext.Pin != result {
        glog.Warningf("Invalid login detected")
        seg := fmt.Sprintf("HIRMS:%d:2+9931:%s:Invalid Login'", len(responseSegments)+2, segHead[1])
        responseSegments = append(responseSegments, seg)
      } else {
        seg := fmt.Sprintf("HIRMG:%d:2+0000::OHAI :)'", len(responseSegments)+2)
        responseSegments = append(responseSegments, seg)
        glog.Infof("Initialisation successful")
      }
      break
    }
  }

  msgContext.Segments = responseSegments

  return msgContext, nil
}

func MakeResponseMessage(message string) (string, error) {
  unwrappedMsg := hbci.UnwrapEncryptedData(message)

  msgContext, err := ParseIncomingMessage(unwrappedMsg)
  if err != nil {
    glog.Errorf("Error in parsing: %s", err)
    return "", err
  }

  return hbci.MakeMessage300(msgContext.NextMsgNum, msgContext.DialogId, msgContext.Segments, &msgContext.Bank, msgContext.UserId), nil

}

func hbciHandler(w http.ResponseWriter, r *http.Request) {
  message, _ := ioutil.ReadAll(r.Body)
  decodedBytes, err := base64.StdEncoding.DecodeString(string(message))
  if err != nil {
    glog.Errorf("Decoding error: %s", err)
    return
  }

  glog.Infof("Decoded message: %s", decodedBytes)
  response, _ := MakeResponseMessage(string(decodedBytes))
  fmt.Fprintf(w,"%s", base64.StdEncoding.EncodeToString([]byte(response)))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()

  loginId := r.PostFormValue("login_id")
  pin := r.PostFormValue("pin")

  glog.Infof("Storing new user: %s (PIN=%s)", loginId, pin)

  err := persistence.Save(loginId, pin)
  if err != nil {
    glog.Errorf("Cannot save to redis: %s", err)
    fmt.Fprintf(w, "Failed to save to persistence backend")
    return
  }

  fmt.Fprintf(w, "OK")
}

func main() {
  addr := flag.String("addr", ":8080", "Host & Port the HTTPS server will listen on (defaults to :8080)")

  flag.Parse() // glog needs this for the loglevels!
  glog.Infof("Server starting...")
  http.HandleFunc("/hbci", hbciHandler)
  http.HandleFunc("/users", userHandler)
  http.ListenAndServeTLS(*addr, "server.crt", "server.key", nil)
}
