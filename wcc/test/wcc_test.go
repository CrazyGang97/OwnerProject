package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"code.byted.org/webcast/rpcv2_webcast_platform_wcc/kitex_gen/webcast/platform_wcc"
)

const (
	DefaultDomain     string = "https://webcast-arch.byted.org"
	OPERATOR          string = "donggang.crazy"
	ServiceId         int64  = 19
	NamespaceIDNorm   int64  = 40
	NamespaceNameNorm string = "normal_space"
)

type CreateConfigRequest struct {
	ConfigMetaList []*ConfigMeta `json:"configMetaList"`
	Operator       string        `json:"operator"`
}
type ConfigMeta struct {
	ID          int64
	ServiceID   int64
	Key         string
	Description string
	Region      platform_wcc.Region
	Env         string
	ProductID   int64
	Kind        platform_wcc.ConfigKind
	Type        platform_wcc.ConfigType
	NamespaceID int64
	SchemaID    int64
}

func Test_CreateConfig(t *testing.T) {

	configMetaList := []*ConfigMeta{
		{
			ServiceID:   ServiceId,
			Key:         "1",
			Region:      platform_wcc.Region_CN,
			Env:         "prod",
			ProductID:   99999,
			Kind:        platform_wcc.ConfigKind_Normal,
			Type:        platform_wcc.ConfigType_String,
			NamespaceID: 40,
		},
	}

	createConfigRequest := CreateConfigRequest{
		ConfigMetaList: configMetaList,
		Operator:       OPERATOR,
	}

	bodyBytes, err := json.Marshal(createConfigRequest)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	requestBody := bytes.NewReader(bodyBytes)

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", DefaultDomain+"/api/wcc/v1/configs", requestBody)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	request.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-USE-PPE":    {"1"},
		"X-TT-ENV":     {"ppe_biz_config"},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	createConfigReply := platform_wcc.CreateConfigReply{}
	err = json.Unmarshal(responseBody, &createConfigReply)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

type KeyValue struct {
	Type      platform_wcc.ConfigType
	Key       string
	ProductID int64
	Value     string
}
type Deployment struct {
	ServiceID int64
	Namespace string
	Region    platform_wcc.Region
	Env       string
	KeyValues []*KeyValue
	Operator  string
}
type CreateUnattendedDeploymentRequest struct {
	Deployment *Deployment `json:"deployment"`
	Async      bool        `json:"async"`
}

func Test_InsertDataNormal(t *testing.T) {
	keyValue := &KeyValue{
		Type:      platform_wcc.ConfigType_String,
		Key:       "1",
		ProductID: 99999,
		Value:     "test_value1",
	}

	keyValues := make([]*KeyValue, 0)
	keyValues = append(keyValues, keyValue)

	deployment := &Deployment{
		ServiceID: 19,
		Namespace: "normal_space",
		Region:    platform_wcc.Region_CN,
		Env:       "prod",
		KeyValues: keyValues,
		Operator:  "donggang.crazy",
	}

	createUnattendedDeploymentRequest := CreateUnattendedDeploymentRequest{
		Deployment: deployment,
	}

	bodyBytes, err := json.Marshal(createUnattendedDeploymentRequest)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	requestBody := bytes.NewReader(bodyBytes)

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", DefaultDomain+"/api/wcc/v1/deployments/unattended", requestBody)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	request.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-USE-PPE":    {"1"},
		"X-TT-ENV":     {"ppe_biz_config"},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(response.Header.Get("X-Tt-Logid"))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	createUnattendedDeploymentReply := platform_wcc.CreateUnattendedDeploymentReply{}
	err = json.Unmarshal(responseBody, &createUnattendedDeploymentReply)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func Test_InsertDataBig(t *testing.T) {
	keyValue := &KeyValue{
		Type:      platform_wcc.ConfigType_String,
		Key:       "1",
		ProductID: 99999,
		Value:     "test_value",
	}

	keyValues := make([]*KeyValue, 0)
	keyValues = append(keyValues, keyValue)

	deployment := &Deployment{
		ServiceID: 19,
		Namespace: "big_value_space",
		Region:    platform_wcc.Region_CN,
		Env:       "prod",
		KeyValues: keyValues,
		Operator:  "donggang.crazy",
	}

	createUnattendedDeploymentRequest := CreateUnattendedDeploymentRequest{
		Deployment: deployment,
	}

	bodyBytes, err := json.Marshal(createUnattendedDeploymentRequest)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	requestBody := bytes.NewReader(bodyBytes)

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", DefaultDomain+"/api/wcc/v1/deployments/unattended", requestBody)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	request.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-USE-PPE":    {"1"},
		"X-TT-ENV":     {"ppe_biz_config"},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	createUnattendedDeploymentReply := platform_wcc.CreateUnattendedDeploymentReply{}
	err = json.Unmarshal(responseBody, &createUnattendedDeploymentReply)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func createConfig(key string, namespaceID int64) (err error) {
	configMetaList := []*ConfigMeta{
		{
			ServiceID:   ServiceId,
			Key:         key,
			Region:      platform_wcc.Region_CN,
			Env:         "prod",
			ProductID:   99999,
			Kind:        platform_wcc.ConfigKind_Normal,
			Type:        platform_wcc.ConfigType_String,
			NamespaceID: namespaceID,
		},
	}

	createConfigRequest := CreateConfigRequest{
		ConfigMetaList: configMetaList,
		Operator:       OPERATOR,
	}

	bodyBytes, err := json.Marshal(createConfigRequest)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	requestBody := bytes.NewReader(bodyBytes)

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", DefaultDomain+"/api/wcc/v1/configs", requestBody)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	request.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-USE-PPE":    {"1"},
		"X-TT-ENV":     {"ppe_biz_config"},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	fmt.Println(response)
	fmt.Println(responseBody)
	return
}

func insertNormal(key string, value string, namespace string) (err error) {
	keyValue := &KeyValue{
		Type:      platform_wcc.ConfigType_String,
		Key:       key,
		ProductID: 99999,
		Value:     value,
	}

	keyValues := make([]*KeyValue, 0)
	keyValues = append(keyValues, keyValue)

	deployment := &Deployment{
		ServiceID: 19,
		Namespace: namespace,
		Region:    platform_wcc.Region_CN,
		Env:       "prod",
		KeyValues: keyValues,
		Operator:  "donggang.crazy",
	}

	createUnattendedDeploymentRequest := CreateUnattendedDeploymentRequest{
		Deployment: deployment,
	}

	bodyBytes, err := json.Marshal(createUnattendedDeploymentRequest)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	requestBody := bytes.NewReader(bodyBytes)

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", DefaultDomain+"/api/wcc/v1/deployments/unattended", requestBody)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	request.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-USE-PPE":    {"1"},
		"X-TT-ENV":     {"ppe_biz_config"},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	fmt.Println(response.Header.Get("X-Tt-Logid"))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	fmt.Println(response)
	fmt.Println(responseBody)
	return
}

func Producer(indexChan chan int) {
	for i := 10; i < 100; i++ {
		indexChan <- i
		fmt.Printf("index = %d \n", i)
	}
	close(indexChan)
}

func Consumer(indexChan chan int, done chan bool) {
	for {
		i, ok := <-indexChan
		if ok {
			key := "normalKey" + strconv.Itoa(i)
			err := createConfig(key, NamespaceIDNorm)
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}
			time.Sleep(2 * time.Second)
			err = insertNormal(key, randStr(10), NamespaceNameNorm)
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}
			time.Sleep(2 * time.Second)

		} else {
			fmt.Println("closed...")
			break
		}
	}
	done <- true
}

func Test_InsertNorm(t *testing.T) {

	indexChan := make(chan int, 50)
	num := 10
	done := make(chan bool, num)

	go Producer(indexChan)

	for i := 0; i < num; i++ {
		go Consumer(indexChan, done)
	}

	for i := 0; i < num; i++ {
		<-done
	}

	for i := 0; i < 10; i++ {
		key := "normalKey" + strconv.Itoa(i)
		err := createConfig(key, NamespaceIDNorm)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		time.Sleep(2 * time.Second)
		err = insertNormal(key, randStr(10), NamespaceNameNorm)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		time.Sleep(2 * time.Second)
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func randStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}