package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"reflect"
	//"github.com/davecgh/go-spew/spew"
	
	//model "github.com/cloud-barista/cb-webtool/src/model"
	model "github.com/dogfootman/cb-webtool/src/model"
	// model "./src/model"
	
)

// var NameSpaceUrl = "http://15.165.16.67:1323"
var NameSpaceUrl = os.Getenv("TUMBLE_URL")

type NSInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetNS(nsID string) NSInfo {
	url := NameSpaceUrl + "ns" + nsID

	body := HttpGetHandler(url)
	defer body.Close()
	nsInfo := NSInfo{}
	json.NewDecoder(body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo.ID)
	return nsInfo

}

func GetNSList() []NSInfo {
	url := NameSpaceUrl + "/ns"
	fmt.Println("============= NameSpace URL =============", url)
	// authInfo := controller.AuthenticationHandler()
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {

	// }
	// req.Header.Add("Authorization", authInfo)
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// fmt.Println("=============result GetNSList =============", resp)
	// //spew.Dump(resp)
	// if err != nil {
	// 	fmt.Println("========= GetNSList Error : ", err)
	// 	fmt.Println("request URL : ", url)
	// 	return nil
	// }

	// defer resp.Body.Close()
	body := HttpGetHandler(url)
	nsInfo := map[string][]NSInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&nsInfo)
	//spew.Dump(body)
	return nsInfo["ns"]

}

func GetNSCnt() int {
	url := NameSpaceUrl + "/ns"
	fmt.Println("============= NameSpace URL =============", url)

	// defer resp.Body.Close()
	body := HttpGetHandler(url)
	nsInfo := map[string][]NSInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&nsInfo)
	//spew.Dump(body)
	if nsInfo["ns"] == nil {
		return 0
	} else {
		return len(nsInfo["ns"])

	}

}

// func RegNS() error {

// }

func RequestGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]NSInfo{}
	fmt.Println("nsInfo type : ", reflect.TypeOf(nsInfo))
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Get Data Error")
	// }
	// fmt.Println("GetData : ", string(data))

}

func HttpGetHandler(url string) io.ReadCloser {
	authInfo := AuthenticationHandler()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	client := &http.Client{}
	resp, _ := client.Do(req)

	//defer resp.Body.Close()

	return resp.Body
}
func AuthenticationHandler() string {

	api_username := os.Getenv("API_USERNAME")
	api_password := os.Getenv("API_PASSWORD")

	//The header "KEY: VAL" is "Authorization: Basic {base64 encoded $USERNAME:$PASSWORD}".
	apiUserInfo := api_username + ":" + api_password
	encA := base64.StdEncoding.EncodeToString([]byte(apiUserInfo))
	//req.Header.Add("Authorization", "Basic"+encA)
	return "Basic " + encA

}


// 사용자의 namespace 목록 조회
func GetNameSpaceList() ([]model.NSInfo, error) {
	fmt.Println("GetNameSpaceList start")
	url := NameSpaceUrl + "/ns"

	body, err := HttpGetHandler(url)

	fmt.Println("called HttpGetHandler")
	fmt.Println(body)
	if err != nil {
		// Tumblebug 접속 확인하라고
		return nil, err
	}

	//nsInfoList := map[string][]model.NSInfo{}
	// nsInfoList := map[string][]model.NSInfo{}
	// nserr := json.NewDecoder(body).Decode(&nsInfoList)
	// if nserr != nil {
	// 	return nil, nserr
	// }
	// fmt.Println(nsInfoList["ns"])

	nsInfoList := map[string][]model.NSInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&nsInfoList)
	//spew.Dump(body)
	fmt.Println(nsInfoList["ns"])

	// nsInfo := map[string][]model.NSInfo{}
	// defer body.Close()
	// json.NewDecoder(body).Decode(&nsInfo)
	// //spew.Dump(body)
	// fmt.Println(nsInfo["ns"])
	// return nsInfo["ns"]

	//nsInfoList := map[string][]model.NSInfo{}
	// nsInfoMap := map[string][]model.NSInfo{}

	// nsInfoList := model.NSInfoList{}
	// keys := []string{}
	// for key, _ := range nsInfoMap {
	// 	fmt.Println(key)
	// 	// keys = append(keys, key)
	// 	nsInfoList = append(nsInfoList, key)
	// }
	// //nsList := model.NsInfoList{}
	// defer body.Close()

	// err := json.NewDecoder(body).Decode(&nsInfo)
	// if err != nil {
	// 	return nil, err
	// }
	// //spew.Dump(body)
	// fmt.Println(nsInfo)
	//return nsInfo["ns"]
	// 접속오류인지
	// namespace가 생성되지 않았는지
	// return nil, nil
	return nsInfoList["ns"], nil
}