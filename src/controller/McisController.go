package controller

import (
	"fmt"
	model "github.com/cloud-barista/cb-webtool/src/model"
	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"
	"log"
	"net/http"

	echotemplate "github.com/foolin/echo-template"
	"github.com/labstack/echo"
	// echosession "github.com/go-session/echo-session"
)

// type SecurityGroup struct {
// 	Id []string `form:"sg"`
// }

func McisRegForm(c echo.Context) error {
	fmt.Println("McisRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manage/McisRegister", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// MCIS 관리 화면 McisListForm 에서 이름 변경 McisMngForm으로
// func McisListForm(c echo.Context) error {
func McisMngForm(c echo.Context) error {
	fmt.Println("McisMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// provider 별 연결정보 count(MCIS 무관)
	cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList()
	connectionConfigCountMap, providerCount := service.GetCloudConnectionCountMap(cloudConnectionConfigInfoList)
	totalConnectionCount := len(cloudConnectionConfigInfoList)

	// 모든 MCIS 조회
	mcisList, _ := service.GetMcisList(defaultNameSpaceID)
	log.Println(" mcisList  ", mcisList)

	totalMcisCount := len(mcisList) // mcis 갯수
	totalVmCount := 0               // 모든 vm 갯수

	totalMcisStatusCountMap := make(map[string]int)             // 모든 MCIS의 상태 Map
	mcisStatusCountMapByMcis := make(map[string]map[string]int) // MCIS ID별 mcis status
	totalVmStatusCountMap := make(map[string]int)               // 모든 VM의 상태 Map
	vmStatusCountMapByMcis := make(map[string]map[string]int)   // MCIS ID 별 vmStatusMap
	mcisSimpleInfoList := []model.MCISSimpleInfo{}              // 표에 뿌려줄 mics summary 정보

	for _, mcisInfo := range mcisList {
		resultMcisStatusCountMap := service.GetMcisStatusCountMap(mcisInfo)
		// mcisStatusMap["RUNNING"] = mcisStatusRunning
		// mcisStatusMap["STOPPED"] = mcisStatusStopped
		// mcisStatusMap["TERMINATED"] = mcisStatusTerminated
		// mcisStatusMap["TOTAL"] = mcisStatusRunning + mcisStatusStop

		for mcisStatusKey, mcisStatusCountVal := range resultMcisStatusCountMap {
			if mcisStatusKey == "TOTAL" { // Total까지 오므로 Total은 제외
				continue
			}

			val, exists := totalMcisStatusCountMap[mcisStatusKey]
			if exists {
				totalMcisStatusCountMap[mcisStatusKey] = val + mcisStatusCountVal
			} else {
				totalMcisStatusCountMap[mcisStatusKey] = mcisStatusCountVal
			}
		}

		mcisStatusCountMapByMcis[mcisInfo.ID] = resultMcisStatusCountMap // 각 MCIS의 status별 cnt
		// connectionConfigCountMap[util.GetProviderName(connectionInfo.ProviderName)] = count

		//////////// vm status area
		resultVmStatusCountMap := service.GetVMStatusCountMap(mcisInfo)
		for i, _ := range util.STATUS_ARRAY {
			// status_array는 고정값이므로 없는 경우 default로 '0'으로 set
			_, exists := resultVmStatusCountMap[util.STATUS_ARRAY[i]]
			if !exists {
				resultVmStatusCountMap[util.STATUS_ARRAY[i]] = 0
			}
			totalVmStatusCountMap[util.STATUS_ARRAY[i]] += resultVmStatusCountMap[util.STATUS_ARRAY[i]]
		}
		// UI manage mcis > server 영역에서는 run/stopped/terminated 만 있음. etc를 stopped에 추가한다.
		totalVmStatusCountMap["stopped"] = totalVmStatusCountMap["stopped"] + resultVmStatusCountMap[util.VM_STATUS_ETC]

		totalVmCount += resultVmStatusCountMap["TOTAL"] // 모든 vm의 갯수

		totalVmCountByMcis := resultVmStatusCountMap["TOTAL"]        // 모든 vm의 갯수
		vmStatusCountMapByMcis[mcisInfo.ID] = resultVmStatusCountMap // MCIS 내 vm 상태별 cnt

		// Provider 별 connection count (Location 내에 있는 provider로 갯수 셀 것.)
		mcisConnectionMap := service.GetVMConnectionCountByMcis(mcisInfo) // 해당 MCIS의 각 provider별 connection count
		log.Println(mcisConnectionMap)
		////////////// return value 에 set
		mcisSimpleInfo := model.MCISSimpleInfo{}
		mcisSimpleInfo.ID = mcisInfo.ID
		mcisSimpleInfo.Status = mcisInfo.Status
		mcisSimpleInfo.McisStatus = util.GetMcisStatus(mcisInfo.Status)
		mcisSimpleInfo.Name = mcisInfo.Name
		mcisSimpleInfo.Description = mcisInfo.Description

		mcisSimpleInfo.VmCount = totalVmCountByMcis // 해당 mcis의 모든 vm 갯수
		mcisSimpleInfo.VmStatusCountMap = resultVmStatusCountMap
		// mcisSimpleInfo.VmRunningCount = vmStatusCountMap[util.STATUS_ARRAY[0]]    //running
		// mcisSimpleInfo.VmStoppedCount = vmStatusCountMap[util.STATUS_ARRAY[1]]    //stopped
		// mcisSimpleInfo.VmTerminatedCount = vmStatusCountMap[util.STATUS_ARRAY[2]] //terminated

		mcisSimpleInfo.ConnectionConfigProviderMap = mcisConnectionMap // 해당 MCIS 등록된 connection의 provider 목록
		// mcisSimpleInfo.ConnectionConfigProviderNames =
		mcisSimpleInfo.ConnectionConfigProviderCount = len(mcisConnectionMap)
		// mcisConnectionMap.ConnectionCount = mcisConnectionMap

		mcisSimpleInfoList = append(mcisSimpleInfoList, mcisSimpleInfo)

	}

	// log.Println(" totoalMcisCount  ", totoalMcisCount)
	// log.Println(" totoalVmCount  ", totoalVmCount)

	// // mcis 별 vmCnt
	// // mcisSimpleInfos = model.McisSimpleInfos{}
	// connectionCountTotal := 0
	// connectionCountByMcis := 0
	// vmCountTotal := 0
	// vmRunningCountByMcis := 0
	// vmStoppedCountByMcis := 0
	// vmTerminatedCountByMcis := 0
	// for mcisIndex, mcisInfo := range mcisList {
	// 	// mcis.ID, mcis.status, mcis.name, mcis.description
	// 	// csp : 해당 mcis의 connection cnt
	// 	// vm_cnt : 해당 mcis의 vm cnt
	// 	// vm_run_cnt, vm_stop_cnt
	// 	vmList := mcisInfo.VMs
	// 	mcisConnectionCountMap := make(map[string]int)
	// 	mcisVmStatusCountMap := make(map[string]int)
	// 	for vmIndex, vmInfo := range vmList {
	// 		locationInfo := vmInfo.LocationInfo
	// 		cloudType := locationInfo.CloudType // CloudConnection
	// 		providerCount := 0
	// 		val, exists := mcisConnectionCountMap[util.GetProviderName(locationInfo.CloudType)]
	// 		if !exists {
	// 			providerCount = 1
	// 		} else {
	// 			providerCount = val + 1
	// 		}
	// 		mcisConnectionCountMap[util.GetProviderName(locationInfo.CloudType)] = providerCount

	// 		vmStatus := vmInfo.Status
	// 		vnStatusCount := 0
	// 		val2, exists2 := mcisVmStatusCountMap[util.GetVmStatus(vmInfo.Status)]
	// 		if !exists2 {
	// 			vnStatusCount = 1
	// 		} else {
	// 			vnStatusCount = val2 + 1
	// 		}
	// 		mcisVmStatusCountMap[util.GetVmStatus(vmInfo.Status)] = vnStatusCount
	// 	}

	// 	mcisSimpleInfo := model.McisSimpleInfo{}
	// 	mcisSimpleInfo.ID = mcisInfo.ID
	// 	mcisSimpleInfo.Status = mcisInfo.Status
	// 	mcisSimpleInfo.Name = mcisInfo.Name
	// 	mcisSimpleInfo.Description = mcisInfo.Description

	// 	mcisSimpleInfo.VmCount = len(vmList)
	// 	mcisSimpleInfo.VmRunningCount = mcisVmStatusCountMap[util.VM_STATUS_RUNNING]
	// 	mcisSimpleInfo.VmStoppedCount = mcisVmStatusCountMap[util.VM_STATUS_RUNNING]
	// 	mcisSimpleInfo.VmTerminatedCount = mcisVmStatusCountMap[util.VM_STATUS_RUNNING]
	// }

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manage/McisMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"McisList":           mcisList,
			// "McisIDList":         mcisIdArr,
			// "VmIDList":           vmIdArr,
			// "VMStatusList":  vmStatusArr,
			// "MCISStatusMap":            mcisStatusMap,
			// "MCISCount":                totoalMcisCount,
			// "VMStatusMap":              vmStatusMap,
			// "VMCount":                  totoalVmCount,
			// "ConnectionConfigCountMap": connectionConfigCountMap,
			// "ConnectionCount":          totalConnectionCount,
			// "ProviderCount":            providerCount,

			// mcis count 영역
			"TotalMCISCount":          totalMcisCount,
			"TotalMCISStatusCountMap": totalMcisStatusCountMap, // 모든 MCIS의 상태 Map

			// server count 영역
			"TotalVmCount":          totalVmCount,
			"TotalVMStatusCountMap": totalVmStatusCountMap, // 모든 VmStatus 별 count Map(MCIS 무관)

			// cp count 영역
			"TotalProviderCount":         providerCount,            // VM이 등록 된 provider 목록
			"TotalConnectionConfigCount": totalConnectionCount,     // 총 connection 갯수
			"ConnectionConfigCountMap":   connectionConfigCountMap, // provider별 connection 수
			// mcis list
			"MCISList":               mcisSimpleInfoList,     // 표에 뿌려줄 mics summary 정보
			"VmStatusCountMapByMCIS": vmStatusCountMapByMcis, // MCIS ID 별 vmStatusMap

		})
}

// MCIS 목록 조회
func GetMcisList(c echo.Context) error {
	log.Println("GetMcisList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	mcisList, respStatus := service.GetMcisList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"McisList":           mcisList,
	})
}

// func McisListFormWithParam(c echo.Context) error {
// 	mcis_id := c.Param("mcis_id")
// 	mcis_name := c.Param("mcis_name")
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if mcis_id == "" && mcis_name == "" {
// 		mcis_id = ""
// 		mcis_name = ""
// 	}
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Mcis.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"McisID":    mcis_id,
// 			"McisName":  mcis_name,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	//return c.Render(http.StatusOK, "MCISlist.html", nil)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func VMAddForm(c echo.Context) error {
// 	mcis_id := c.Param("mcis_id")
// 	mcis_name := c.Param("mcis_name")
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if mcis_id == "" && mcis_name == "" {
// 		mcis_id = ""
// 		mcis_name = ""
// 	}
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Create_VM.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"McisID":    mcis_id,
// 			"McisName":  mcis_name,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	//return c.Render(http.StatusOK, "MCISlist.html", nil)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func McisRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Create_Mcis.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func McisRegController(c echo.Context) error {
// 	m := new(model.MCISRequest)

// 	vmspec := c.FormValue("vmspec")
// 	namespace := c.FormValue("namespace")
// 	mcis_name := c.FormValue("mcis_name")
// 	provider := c.FormValue("provider")
// 	sg := c.FormValue("sg")

// 	fmt.Println("namespace : ", namespace)
// 	fmt.Println("mcis_name : ", mcis_name)
// 	fmt.Println("vmSpec : ", vmspec)
// 	fmt.Println("provider : ", provider)
// 	fmt.Println("sg : ", sg)

// 	if err := c.Bind(m); err != nil {
// 		fmt.Println("bind Error")
// 		return err
// 	}
// 	fmt.Println("Bind Form : ", m)
// 	fmt.Println("nameSPace:", m.NameSpace)
// 	fmt.Println("vmName 0 : ", m.VMName[0])
// 	fmt.Println("vmName 1 : ", m.VMName[1])
// 	fmt.Println("vmSpec 0 : ", m.VMSpec[0])
// 	fmt.Println("vmspec 1 : ", m.VMSpec[1])

// 	//spew.Dump(m)
// 	//return c.Redirect(http.StatusTemporaryRedirect, "/MCIS/list")
// 	return nil
// }
