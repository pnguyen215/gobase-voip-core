package ami

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/textproto"
	"reflect"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func NewActionWith(name string) *AMIMessage {
	a := NewMessage()
	a.AddField(config.AmiActionKey, name)
	return a
}

func NewMessage() *AMIMessage {
	header := make(textproto.MIMEHeader)
	return ofMessage(header)
}

func ofMessage(header textproto.MIMEHeader) *AMIMessage {
	m := &AMIMessage{}
	m.Header = header
	return m
}

// Login action by message
func LoginWith(username, password string) *AMIMessage {
	a := NewActionWith(config.AmiLoginKey)
	a.AddField(config.AmiUsernameField, username)
	a.AddField(config.AmiSecretField, password)
	return a
}

// Field return first value associated with then given key
func (k *AMIMessage) Field(key string) string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Get(key)
}

// Return first value associated of field by the given key
func (k *AMIMessage) GetFirstValueByField(key string) string {
	return k.Field(key)
}

// Field Values return all values associated with the given key
func (k *AMIMessage) FieldValues(key string) []string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Values(key)
}

// Return all value associated by the given key
func (k *AMIMessage) GetValuesByField(key string) []string {
	return k.FieldValues(key)
}

// Added new pair header as form key:value
func (k *AMIMessage) AddField(key, value string) {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	k.Header.Add(key, value)
}

// Added collection pair header as form key:value
func (k *AMIMessage) AddFields(fields map[string]string) {
	for key, value := range fields {
		k.AddField(key, value)
	}
}

// Added new Action-Id and insert into message
func (k *AMIMessage) AddActionIdWith(id string) {
	k.AddField(config.AmiActionIdKey, id)
}

// Added Action-id generated random
func (k *AMIMessage) AddActionId() {
	b := make([]byte, 12)

	_, err := rand.Read(b)

	if err == nil {
		k.AddActionIdWith(fmt.Sprintf("%x", b))
	}
}

// Return AMI message Action Id as string
func (k *AMIMessage) GetActionId() string {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	return k.Header.Get(strings.ToLower(config.AmiActionIdKey))
}

// Remove fields associated with the given key
func (k *AMIMessage) RemoveField(key string) {
	k.Mutex.Lock()
	defer k.Mutex.Unlock()
	k.Header.Del(key)
}

// Remove fields associated with the collection given keys
func (k *AMIMessage) RemoveFields(keys ...string) {
	for _, key := range keys {
		k.RemoveField(key)
	}
}

// Return new bytes buffer
func (k *AMIMessage) toBytesBuffer() bytes.Buffer {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()
	var buffer bytes.Buffer

	for key := range k.Header {
		for _, value := range k.Header.Values(key) {
			buffer.WriteString(key)
			buffer.WriteString(": ")
			buffer.WriteString(value)
			buffer.WriteString("\r\n")
		}
	}
	buffer.WriteString("\r\n")

	return buffer
}

// Return headers Asterisk Manager Interface (AMI) as string
func (k *AMIMessage) String() string {
	_buffer := k.toBytesBuffer()
	return _buffer.String()
}

// Return headers Asterisk Manager Interface (AMI) as byte array
func (k *AMIMessage) Bytes() []byte {
	_buffer := k.toBytesBuffer()
	return _buffer.Bytes()
}

// Return true if the AMI message is event type
func (k *AMIMessage) IsEvent() bool {
	return k.Field(config.AmiEventKey) != ""
}

// Return true if the AMI message is response type
func (k *AMIMessage) IsResponse() bool {
	return k.Field(config.AmiResponseKey) != ""
}

// Return true if the AMI message is response type with value success
func (k *AMIMessage) IsSuccess() bool {
	return strings.EqualFold(k.Field(config.AmiResponseKey), config.AmiStatusSuccessKey)
}

func (k *AMIMessage) PreVars() []string {
	vars := append(k.FieldValues("variable"), k.FieldValues("chanvariable")...)
	vars = append(vars, k.FieldValues("ParkeeChanVariable")...)
	vars = append(vars, k.FieldValues("OrigTransfererChanVariable")...)
	vars = append(vars, k.FieldValues("SecondTransfererChanVariable")...)
	vars = append(vars, k.FieldValues("TransfereeChanVariable")...)
	vars = append(vars, k.FieldValues("TransferTargetChanVariable")...)
	vars = append(vars, k.FieldValues("SpyeeChanVariable")...)
	vars = append(vars, k.FieldValues("SpyerChanVariable")...)

	return vars
}

// Var search in AMI message fields Variable and ChanVariable for a value
// of the type key=value or just key. If found, returns value as string
// and true. Variable name is case sensitive.
func (k *AMIMessage) Var(key string) (string, bool) {
	return k.VarWith(key, k.PreVars())
}

func (k *AMIMessage) VarWith(key string, vars []string) (string, bool) {
	if len(vars) == 0 || key == "" {
		return key, false
	}

	for _, value := range vars {
		e, v := utils.VarsSplit(value)
		if e == key || strings.EqualFold(e, key) {
			return v, true
		}
	}

	return "", false
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessage() map[string]interface{} {
	return k.ProduceMessageWith(false)
}

// Return AMI message as interface{}
func (k *AMIMessage) ProduceMessageWith(lowercaseField bool) map[string]interface{} {
	k.Mutex.RLock()
	defer k.Mutex.RUnlock()

	data := make(map[string]interface{})

	for key, value := range k.Header {
		var field string = key

		if lowercaseField {
			field = strings.ToLower(key)
		} else {
			field = key
		}

		if len(value) == 1 {
			data[field] = value[0]
		} else {
			data[field] = utils.VarsMap(value)
		}
	}

	return data
}

// Return AMI messages as Json string
func (k *AMIMessage) Json() string {
	result, err := json.Marshal(k.ProduceMessage())

	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return string(result)
}

// Create AMI message from json string
func FromJson(jsonString string) (*AMIMessage, error) {
	var builder interface{}

	message := ofMessage(textproto.MIMEHeader{})

	err := json.Unmarshal([]byte(jsonString), &builder)

	if err != nil {
		log.Fatal(err.Error())
		return message, err
	}

	for k, v := range builder.(map[string]interface{}) {
		refType := reflect.ValueOf(v)
		switch refType.Kind() {
		case reflect.Map:
			for v_name, v_val := range v.(map[string]interface{}) {
				message.AddField(k, fmt.Sprintf("%s=%v", v_name, v_val))
			}
		default:
			message.AddField(k, fmt.Sprintf("%v", v))
		}
	}

	return message, nil
}
