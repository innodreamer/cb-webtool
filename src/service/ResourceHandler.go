package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	// "math"
	"net/http"
	// "strconv"
	// "sync"
	"io/ioutil"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// 해당 namespace의 vpc 목록 조회
//func GetVnetList(nameSpaceID string) (io.ReadCloser, error) {
func GetVnetList(nameSpaceID string) ([]model.VNetInfo, model.WebStatus) {
	fmt.Println("GetVnetList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	vNetInfoList := map[string][]model.VNetInfo{}
	json.NewDecoder(respBody).Decode(&vNetInfoList)
	//spew.Dump(body)
	fmt.Println(vNetInfoList["vNet"])

	return vNetInfoList["vNet"], model.WebStatus{StatusCode: respStatus}

}

// vpc 상세 조회-> ResourceHandler로 이동
func GetVpcData(nameSpaceID string, vNetID string) (*model.VNetInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet/" + vNetID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	vNetInfo := model.VNetInfo{}
	if err != nil {
		fmt.Println(err)
		return &vNetInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	// json.NewDecoder(body).Decode(&vNetInfo)
	json.NewDecoder(respBody).Decode(&vNetInfo)
	fmt.Println(vNetInfo)

	// return vNetInfo, err
	return &vNetInfo, model.WebStatus{StatusCode: respStatus}
}

// vpc 등록
// func RegVpc(nameSpaceID string, vnetRegInfo *model.VNetRegInfo) (io.ReadCloser, int) {
func RegVpc(nameSpaceID string, vnetRegInfo *model.VNetRegInfo) (*model.VNetInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	fmt.Println("vnetRegInfo : ", vnetRegInfo)

	pbytes, _ := json.Marshal(vnetRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	vNetInfo := model.VNetInfo{}
	if err != nil {
		fmt.Println(err)
		return &vNetInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("Reg respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(&vNetInfo)
	fmt.Println(vNetInfo)

	util.DisplayResponse(resp)

	// return respBody, respStatusCode
	return &vNetInfo, model.WebStatus{StatusCode: respStatus}
}

// vpc 삭제
func DelVpc(nameSpaceID string, vNetID string) (io.ReadCloser, model.WebStatus) {
	// if ValidateString(vNetID) != nil {
	if len(vNetID) == 0 {
		log.Println("vNetID 가 없으면 해당 namespace의 모든 vpc가 삭제되므로 처리할 수 없습니다.")
		return nil, model.WebStatus{StatusCode: 4040, Message: "vNetID 가 없으면 해당 namespace의 모든 vpc가 삭제되므로 처리할 수 없습니다."}
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet/" + vNetID

	fmt.Println("vNetID : ", vNetID)

	pbytes, _ := json.Marshal(vNetID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

	return respBody, model.WebStatus{StatusCode: respStatus}

}

// 해당 namespace의 SecurityGroup 목록 조회
func GetSecurityGroupList(nameSpaceID string) ([]model.SecurityGroupInfo, model.WebStatus) {
	fmt.Println("GetSecurityGroupList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	securityGroupList := map[string][]model.SecurityGroupInfo{}

	json.NewDecoder(respBody).Decode(&securityGroupList)
	//spew.Dump(body)
	fmt.Println(securityGroupList["securityGroup"])

	return securityGroupList["securityGroup"], model.WebStatus{StatusCode: respStatus}

}

// SecurityGroup 상세 조회
func GetSecurityGroupData(nameSpaceID string, securityGroupID string) (*model.SecurityGroupInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup/" + securityGroupID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	securityGroupInfo := model.SecurityGroupInfo{}
	if err != nil {
		fmt.Println(err)
		return &securityGroupInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&securityGroupInfo)
	fmt.Println(securityGroupInfo)

	return &securityGroupInfo, model.WebStatus{StatusCode: respStatus}
}

// SecurityGroup 등록
func RegSecurityGroup(nameSpaceID string, securityGroupRegInfo *model.SecurityGroupRegInfo) (*model.SecurityGroupInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup"

	fmt.Println("RegSecurityGroup : ")

	pbytes, _ := json.Marshal(securityGroupRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	securityGroupInfo := model.SecurityGroupInfo{}
	if err != nil {
		log.Println("-----")
		fmt.Println(err)
		log.Println("-----1111")
		fmt.Println(err.Error())
		log.Println("-----222")
		return &securityGroupInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	log.Println("respStatusCode = ", resp.StatusCode)
	log.Println("respStatus = ", resp.Status)
	if respStatus != 200 {
		// b, _ := ioutil.ReadAll(respBody)
		// log.Fatal(string(b))
		
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println(errorInfo)
		return nil, model.WebStatus{StatusCode: 500, Message: errorInfo.Message}
	  }

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(&securityGroupInfo)
	fmt.Println(securityGroupInfo)
	// return respBody, respStatusCode
	return &securityGroupInfo, model.WebStatus{StatusCode: respStatus}
}

// SecurityGroup 삭제
func DelSecurityGroup(nameSpaceID string, securityGroupID string) (io.ReadCloser, model.WebStatus) {
	// if ValidateString(vNetID) != nil {
	if len(securityGroupID) == 0 {
		log.Println("securityGroupID 가 없으면 해당 namespace의 모든 securityGroup이 삭제되므로 처리할 수 없습니다.")
		return nil, model.WebStatus{StatusCode: 4040, Message: "securityGroupID 가 없으면 해당 namespace의 모든 securityGroup이 삭제되므로 처리할 수 없습니다."}
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup/" + securityGroupID

	fmt.Println("securityGroupID : ", securityGroupID)

	pbytes, _ := json.Marshal(securityGroupID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return respBody, model.WebStatus{StatusCode: respStatus}

}

// SSHKey 목록 조회 : /ns/{nsId}/resources/sshKey
func GetSshKeyInfoList(nameSpaceID string) ([]model.SshKeyInfo, model.WebStatus) {
	fmt.Println("GetSshKeyInfoList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/sshKey"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	sshKeyList := map[string][]model.SshKeyInfo{}

	json.NewDecoder(respBody).Decode(&sshKeyList)
	//spew.Dump(body)
	fmt.Println(sshKeyList["sshKey"])

	return sshKeyList["sshKey"], model.WebStatus{StatusCode: respStatus}

}

// sshKey 상세 조회
func GetSshKeyData(nameSpaceID string, sshKeyID string) (*model.SshKeyInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/sshKey/" + sshKeyID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	sshKeyInfo := model.SshKeyInfo{}
	if err != nil {
		fmt.Println(err)
		return &sshKeyInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(&sshKeyInfo)
	fmt.Println(sshKeyInfo)

	return &sshKeyInfo, model.WebStatus{StatusCode: respStatus}
}

// sshKey 등록
func RegSshKey(nameSpaceID string, sshKeyRegInfo *model.SshKeyRegInfo) (*model.SshKeyInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/sshKey"

	// fmt.Println("vnetInfo : ", vnetInfo)

	pbytes, _ := json.Marshal(sshKeyRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	sshKeyInfo := model.SshKeyInfo{}
	if err != nil {
		fmt.Println(err)
		return &sshKeyInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(&sshKeyInfo)
	fmt.Println(sshKeyInfo)
	// return respBody, respStatusCode
	return &sshKeyInfo, model.WebStatus{StatusCode: respStatus}
}

// sshKey 삭제
func DelSshKey(nameSpaceID string, sshKeyID string) (io.ReadCloser, model.WebStatus) {
	// if ValidateString(sshKeyID) != nil {
	if len(sshKeyID) == 0 {
		log.Println("securityGroupID 가 없으면 해당 namespace의 모든 securityGroup이 삭제되므로 처리할 수 없습니다.")
		return nil, model.WebStatus{StatusCode: 4040, Message: "securityGroupID 가 없으면 해당 namespace의 모든 securityGroup이 삭제되므로 처리할 수 없습니다."}
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/sshKey/" + sshKeyID

	fmt.Println("sshKeyID : ", sshKeyID)

	pbytes, _ := json.Marshal(sshKeyID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return respBody, model.WebStatus{StatusCode: respStatus}

}

// VirtualMachineImage 목록 조회
func GetVirtualMachineImageInfoList(nameSpaceID string) ([]model.VirtualMachineImageInfo, model.WebStatus) {
	fmt.Println("GetVirtualMachineImageInfoList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/image"

	pbytes, _ := json.Marshal(nameSpaceID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// TODO : defer를 넣어줘야 할 듯. defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	virtualMachineImageList := map[string][]model.VirtualMachineImageInfo{}

	json.NewDecoder(respBody).Decode(&virtualMachineImageList)
	//spew.Dump(body)
	fmt.Println(virtualMachineImageList["image"])

	robots, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	return virtualMachineImageList["image"], model.WebStatus{StatusCode: respStatus}

}

// VirtualMachineImage 상세 조회
func GetVirtualMachineImageData(nameSpaceID string, virtualMachineImageID string) (*model.VirtualMachineImageInfo, model.WebStatus) {
	fmt.Println("GetVirtualMachineImageData ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/image/" + virtualMachineImageID

	// fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	virtualMachineImageInfo := model.VirtualMachineImageInfo{}
	if err != nil {
		fmt.Println(err)
		return &virtualMachineImageInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(&virtualMachineImageInfo)
	fmt.Println(virtualMachineImageInfo)

	return &virtualMachineImageInfo, model.WebStatus{StatusCode: respStatus}
}

// VirtualMachineImage 등록
func RegVirtualMachineImage(nameSpaceID string, virtualMachineImageRegInfo *model.VirtualMachineImageRegInfo) (*model.VirtualMachineImageInfo, model.WebStatus) {
	fmt.Println("RegVirtualMachineImage ************ : ")

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/image?action=registerWithInfo" //
	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/image?action=registerWithId"//

	// fmt.Println("vnetInfo : ", vnetInfo)

	pbytes, _ := json.Marshal(virtualMachineImageRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	virtualMachineImageInfo := model.VirtualMachineImageInfo{}
	if err != nil {
		fmt.Println(err)
		return &virtualMachineImageInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(&virtualMachineImageInfo)
	fmt.Println(virtualMachineImageInfo)
	// return respBody, respStatusCode
	return &virtualMachineImageInfo, model.WebStatus{StatusCode: respStatus}
}

// VirtualMachineImage 삭제
func DelVirtualMachineImage(nameSpaceID string, virtualMachineImageID string) (io.ReadCloser, model.WebStatus) {
	// if ValidateString(VirtualMachineImageID) != nil {
	if len(virtualMachineImageID) == 0 {
		log.Println("ImageID 가 없으면 해당 namespace의 모든 image가 삭제되므로 처리할 수 없습니다.")
		return nil, model.WebStatus{StatusCode: 4040, Message: "ImageID 가 없으면 해당 namespace의 모든 image가 삭제되므로 처리할 수 없습니다."}
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/image/" + virtualMachineImageID

	fmt.Println("virtualMachineImageID : ", virtualMachineImageID)

	pbytes, _ := json.Marshal(virtualMachineImageID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return respBody, model.WebStatus{StatusCode: respStatus}

}

func LookupVirtualMachineImageList(connectionName string) ([]model.VirtualMachineLookupImageInfo, model.WebStatus) {
	fmt.Println("LookupVirtualMachineImageList ************ : ", connectionName)
	url := util.TUMBLEBUG + "/lookupImage"

	// body, err := util.CommonHttpGet(url)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	pbytes, _ := json.Marshal(connectionName)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)
	log.Println("LookupVirtualMachineImageList called 1 ")
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode
	log.Println("LookupVirtualMachineImageList called 2 ", respStatus)
	// return respBody, respStatus
	// log.Println(respBody)
	lookupImageList := map[string][]model.VirtualMachineLookupImageInfo{}

	json.NewDecoder(respBody).Decode(&lookupImageList)
	log.Println("LookupVirtualMachineImageList called 3 ")
	// //spew.Dump(body)
	// // fmt.Println(lookupImageList["image"])

	// util.DisplayResponse(resp)
	// log.Println("LookupVirtualMachineImageList called 4 ")


	outputAsBytes, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	log.Println("LookupVirtualMachineImageList called 4 ")
	fmt.Println(string(outputAsBytes))
	var output apiResponse
	err = json.Unmarshal(outputAsBytes, &output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)

	log.Println("LookupVirtualMachineImageList called 5 ")
	return lookupImageList["image"], model.WebStatus{StatusCode: respStatus}

}

type apiResponse struct {
	Id, Kind, LongURL string
}

func LookupVirtualMachineImageData(virtualMachineImageID string) (*model.VirtualMachineImageInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/lookupImage/" + virtualMachineImageID

	fmt.Println("virtualMachineImageID : ", virtualMachineImageID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	virtualMachineImageInfo := model.VirtualMachineImageInfo{}
	if err != nil {
		fmt.Println(err)
		return &virtualMachineImageInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(virtualMachineImageInfo)
	fmt.Println(virtualMachineImageInfo)

	return &virtualMachineImageInfo, model.WebStatus{StatusCode: respStatus}
}

// VirtualMachineImage 상세 조회
func SearchVirtualMachineImageList(nameSpaceID string, virtualMachineImageID string) (*model.VirtualMachineImageInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/searchImage/"

	fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	virtualMachineImageInfo := model.VirtualMachineImageInfo{}
	if err != nil {
		fmt.Println(err)
		return &virtualMachineImageInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(&virtualMachineImageInfo)
	fmt.Println(virtualMachineImageInfo)

	return &virtualMachineImageInfo, model.WebStatus{StatusCode: respStatus}
}

// InstanceSpec 목록 조회
func GetInstanceSpecInfoList(nameSpaceID string) ([]model.InstanceSpecInfo, model.WebStatus) {
	fmt.Println("GetInstanceSpecInfoList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/spec"

	// pbytes, _ := json.Marshal(nameSpaceID)
	// resp, err := util.CommonHttp(url, pbytes, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// TODO : defer를 넣어줘야 할 듯. defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	instanceSpecList := map[string][]model.InstanceSpecInfo{}

	json.NewDecoder(respBody).Decode(&instanceSpecList)
	//spew.Dump(body)
	fmt.Println(instanceSpecList["spec"])

	return instanceSpecList["spec"], model.WebStatus{StatusCode: respStatus}

}

// InstanceSpec 상세 조회
func GetInstanceSpecData(nameSpaceID string, instanceSpecID string) (*model.InstanceSpecInfo, model.WebStatus) {
	fmt.Println("GetInstanceSpecData ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/spec/" + instanceSpecID

	// fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	instanceSpecInfo := model.InstanceSpecInfo{}
	if err != nil {
		fmt.Println(err)
		return &instanceSpecInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(&instanceSpecInfo)
	fmt.Println(instanceSpecInfo)

	return &instanceSpecInfo, model.WebStatus{StatusCode: respStatus}
}

// InstanceSpec 등록
func RegInstanceSpec(nameSpaceID string, instanceSpecRegInfo *model.InstanceSpecRegInfo) (*model.InstanceSpecInfo, model.WebStatus) {
	fmt.Println("RegInstanceSpec ************ : ")

	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/spec?action=registerWithInfo"// parameter를 모두 받지않기 때문에 param의 data type이 틀려 오류남.
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/spec" // action 인자없이 전송

	// fmt.Println("vnetInfo : ", vnetInfo)

	pbytes, _ := json.Marshal(instanceSpecRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	instanceSpecInfo := model.InstanceSpecInfo{}
	if err != nil {
		fmt.Println(err)
		return &instanceSpecInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(&instanceSpecInfo)
	fmt.Println(instanceSpecInfo)
	// return respBody, respStatusCode
	return &instanceSpecInfo, model.WebStatus{StatusCode: respStatus}
}

// InstanceSpec 삭제
func DelInstanceSpec(nameSpaceID string, instanceSpecID string) (io.ReadCloser, model.WebStatus) {
	// if ValidateString(InstanceSpecID) != nil {
	if len(instanceSpecID) == 0 {
		log.Println("specID 가 없으면 해당 namespace의 모든 image가 삭제되므로 처리할 수 없습니다.")
		return nil, model.WebStatus{StatusCode: 4040, Message: "specID 가 없으면 해당 namespace의 모든 image가 삭제되므로 처리할 수 없습니다."}
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/spec/" + instanceSpecID

	fmt.Println("instanceSpecID : ", instanceSpecID)

	pbytes, _ := json.Marshal(instanceSpecID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return respBody, model.WebStatus{StatusCode: respStatus}

}

func LookupInstanceSpecList() ([]model.InstanceSpecInfo, model.WebStatus) {
	fmt.Println("LookupInstanceSpecList ************ : ")
	url := util.TUMBLEBUG + "/lookupSpec"

	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	instanceSpecList := map[string][]model.InstanceSpecInfo{}

	json.NewDecoder(respBody).Decode(&instanceSpecList)
	//spew.Dump(body)
	fmt.Println(instanceSpecList["vmspec"])

	return instanceSpecList["vmspec"], model.WebStatus{StatusCode: respStatus}

}

func LookupInstanceSpecData(instanceSpecName string) (*model.InstanceSpecInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/lookupSpec/" + instanceSpecName

	fmt.Println("instanceSpecName : ", instanceSpecName)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	instanceSpecInfo := model.InstanceSpecInfo{}
	if err != nil {
		fmt.Println(err)
		return &instanceSpecInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	
	json.NewDecoder(respBody).Decode(&instanceSpecInfo)
	fmt.Println(instanceSpecInfo)

	return &instanceSpecInfo, model.WebStatus{StatusCode: respStatus}
}

// resourcesGroup.PUT("/instancespec/put/:specID", controller.InstanceSpecPutProc)	// RegProc _ SshKey 같이 앞으로 넘길까
// resourcesGroup.POST("/instancespec/filterspecs", controller.FilterInstanceSpecList)
// resourcesGroup.POST("/instancespec/filterspecsbyrange", controller.FilterInstanceSpecListByRange)
