package controller

import (
	// "encoding/json"
	"fmt"
	"github.com/cloud-barista/cb-webtool/src/model"
	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// deprecated
// func NsRegController(c echo.Context) error {
// 	username := c.FormValue("username")
// 	description := c.FormValue("description")

// 	fmt.Println("NSRegController : ", username, description)
// 	return nil
// }

// func NsRegForm(c echo.Context) error {
func NameSpaceRegForm(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// 	comURL := service.GetCommonURL()
	// 	apiInfo := service.AuthenticationHandler()
	// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
	// return c.Render(http.StatusOK, "NSRegister.html", map[string]interface{}{
	return echotemplate.Render(c, http.StatusOK, "setting/namespaces/NSRegister.html", map[string]interface{}{
		"LoginInfo": loginInfo,
		// 			"comURL":    comURL,
		// 			"apiInfo":   apiInfo,
	})
	// 	}
	// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// namespace 등록 처리
func NameSpaceRegProc(c echo.Context) error {

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// namespace := c.FormValue("name")
	// description := c.FormValue("description")
	// fmt.Println("namespace : " + namespace + " , description :" + description)
	// nameSpaceInfo := new(model.NameSpaceInfo)
	// nameSpaceInfo.Name = namespace
	// nameSpaceInfo.Description = description

	nameSpaceInfo := new(model.NameSpaceInfo)
	if err := c.Bind(nameSpaceInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// if err = c.Validate(nameSpaceInfo); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	fmt.Println("nameSpaceInfo : ", nameSpaceInfo)

	// Tubblebug 호출하여 namespace 생성

	// person := Person{"Alex", 10}
	// pbytes, _ := json.Marshal(person)
	respBody, nsErr := service.RegNameSpace(nameSpaceInfo)
	fmt.Println("=============respBody =============", respBody)
	if nsErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, err := service.GetNameSpaceList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       "fail",
			"status":        "403",
			"NameSpaceList": nil,
		})
	}
	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        "200",
		"NameSpaceList": nameSpaceList,
	})
}

// Namespace 수정
func NameSpaceUpdateProc(c echo.Context) error {
	log.Println("NameSpaceUpdateProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nameSpaceInfo := new(model.NameSpaceInfo)
	if err := c.Bind(nameSpaceInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	respBody, respStatus := service.UpdateNameSpace(nameSpaceInfo)
	fmt.Println("=============respBody =============", respBody)
	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, err := service.GetNameSpaceList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       "fail",
			"status":        err,
			"NameSpaceList": nil,
		})
	}
	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"NameSpaceList": nameSpaceList,
	})

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"status":  "200",
	// })
}

// NameSpace 삭제
func NameSpaceDelProc(c echo.Context) error {
	log.Println("NameSpaceDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	paramNameSpaceID := c.Param("nameSpaceID")
	log.Println(paramNameSpaceID)

	respBody, reErr := service.DelNameSpace(paramNameSpaceID)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, err := service.GetNameSpaceList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       "fail",
			"status":        "403",
			"NameSpaceList": nil,
		})
	}
	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        "200",
		"NameSpaceList": nameSpaceList,
	})

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"status":  "200",
	// })
}

// NsListForm -> NameSpaceMngForm으로 변경
//func NsListForm(c echo.Context) error {
func NameSpaceMngForm(c echo.Context) error {
	fmt.Println("=============start NameSpaceMngForm =============")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	// comURL := service.GetCommonURL()
	// apiInfo := service.AuthenticationHandler()
	// loginInfo := service.CallLoginInfo(c)
	// if loginInfo.Username != "" {
	fmt.Println("=============start GetNSList =============")
	// nsList := service.GetNSList()
	nameSpaceList, _ := service.GetNameSpaceList()
	fmt.Println("=============start GetNSList =============", nameSpaceList)
	// if nsList != nil {
	// return c.Render(http.StatusOK, "NameSpace.html", map[string]interface{}{
	// "LoginInfo": loginInfo,
	// "NSList": nsList,
	// "comURL":    comURL,
	// "apiInfo":   apiInfo,
	// })
	// } else {
	// 	return c.Redirect(http.StatusTemporaryRedirect, "/NS/reg")
	// }

	// status, filepath(vies기준), return params
	return echotemplate.Render(c, http.StatusOK,
		"setting/namespaces/NameSpaceMng", // 파일명
		map[string]interface{}{
			"LoginInfo":     loginInfo,
			"NameSpaceList": nameSpaceList,
		})

}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	//return c.Redirect(http.StatusPermanentRedirect, "/login")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// 사용자의 namespace 목록 조회
func GetNameSpaceList(c echo.Context) error {
	fmt.Println("====== GET NAMESPACE LIST ========")
	store := echosession.FromContext(c)
	nameSpaceInfoList, nsErr := service.GetNameSpaceList()
	if nsErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	} else {
		store.Set("namespacelist", nameSpaceInfoList)
		store.Save()
	}

	return c.JSON(http.StatusOK, nameSpaceInfoList)

}

// 기본 namespace set. set default Namespace
func SetNameSpace(c echo.Context) error {
	fmt.Println("====== SET SELECTED NAME SPACE ========")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"message": "Not Login",
			"status":  "403",
		})
	}
	store := echosession.FromContext(c)

	result, ok := store.Get(loginInfo.Username)
	if !ok {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"message": "Not Login",
			"status":  "403",
		})
	}
	storedUser := result.(map[string]string)

	nameSpaceID := c.Param("nameSpaceID")
	loginInfo.DefaultNameSpaceID = nameSpaceID

	nsResult, nsOk := store.Get("namespacelist")
	fmt.Println("nsResult : ", nsResult)
	if !nsOk {
		fmt.Println("nsOk : ", nsOk)
	} else {
		fmt.Println("______________")

		//interface conversion: interface {} is []model.NameSpaceInfo
		nsList := nsResult.([]model.NameSpaceInfo)
		fmt.Println("nsList ", nsList)
		for _, nsInfo := range nsList {
			fmt.Println(nsInfo.ID + " :  " + nameSpaceID)
			if nsInfo.ID == nameSpaceID {
				loginInfo.DefaultNameSpaceID = nsInfo.ID
				loginInfo.DefaultNameSpaceName = nsInfo.Name
				fmt.Println(nsInfo.ID + " :  " + nameSpaceID + " found " + nsInfo.Name)
				storedUser["defaultnameSpacename"] = nsInfo.ID
				storedUser["defaultnamespaceid"] = nsInfo.Name
				break
			}
		}
	}

	// storedUser["defaultnamespaceid"] = nameSpaceID
	fmt.Println("storedUser : ", storedUser)

	store.Set(loginInfo.Username, storedUser)

	storeErr := store.Save()
	if storeErr != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"message": storeErr.Error(),
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success",
		"status":    "200",
		"LoginInfo": loginInfo,
	})
}

// 기본 namespace get. get default Namespace
func GetNameSpace(c echo.Context) error {
	fmt.Println("====== GET SELECTED NAME SPACE ========")
	store := echosession.FromContext(c)

	getInfo, ok := store.Get("namespace")
	if !ok {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"message": "Not Exist",
		})
	}
	nsId := getInfo.(string)

	res := map[string]string{
		"message": "success",
		"nsID":    nsId,
	}

	return c.JSON(http.StatusOK, res)
}
