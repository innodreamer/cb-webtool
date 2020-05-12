package controller

import (
	"fmt"

	"net/http"

	"github.com/cloud-barista/cb-webtool/src/service"
	"github.com/labstack/echo"
	//"github.com/davecgh/go-spew/spew"
)

// Driver Contorller
func DriverRegController(c echo.Context) error {

	username := c.FormValue("username")
	description := c.FormValue("description")

	fmt.Println("NSRegController : ", username, description)
	return nil
}

func DriverRegForm(c echo.Context) error {
	comURL := GetCommonURL()
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "DriverRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func DriverListForm(c echo.Context) error {
	fmt.Println("=============start NsListForm =============")
	comURL := GetCommonURL()
	loginInfo := CallLoginInfo(c)
	if loginInfo.Username != "" {
		//nsList := service.GetDriverList()
		return c.Render(http.StatusOK, "DriverList.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			//"NSList": nsList,
		})
	}

	fmt.Println("LoginInfo : ", loginInfo)

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//Credential Controller
func CredertialRegForm(c echo.Context) error {
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "CredentialRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func CredertialListForm(c echo.Context) error {

	fmt.Println("=============start CredertialRegForm =============")
	loginInfo := CallLoginInfo(c)
	comURL := GetCommonURL()
	if loginInfo.Username != "" {
		//nsList := service.GetCredentialList()
		return c.Render(http.StatusOK, "CredentialList.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			// "NSList": nsList,
		})
	}

	fmt.Println("LoginInfo : ", loginInfo)
	return c.Redirect(http.StatusTemporaryRedirect, "/login")

}

//Region Controller
func RegionRegForm(c echo.Context) error {
	comURL := GetCommonURL()
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "RegionRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
		})
	}
	// return c.Redirect(http.StatusPermanentRedirect, "/login")
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func RegionListForm(c echo.Context) error {
	comURL := GetCommonURL()
	loginInfo := CallLoginInfo(c)
	if loginInfo.Username != "" {
		nsList := service.GetRegionList()
		fmt.Println("REGION List : ",nsList)
		
		//spew.Dump(nsList)
		return c.Render(http.StatusOK, "RegionList.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			"NSList":    nsList,
		})
	}

	fmt.Println("LoginInfo : ", loginInfo)
	return c.Redirect(http.StatusTemporaryRedirect, "/login")

}

//Connection Controller
func ConnectionRegForm(c echo.Context) error {
	comURL := GetCommonURL()
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {

		return c.Render(http.StatusOK, "ConnectionRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
	//return c.Render(http.StatusOK, "RegionRegister.html", nil)
}

func ConnectionListForm(c echo.Context) error {
	comURL := GetCommonURL()
	loginInfo := CallLoginInfo(c)
	if loginInfo.Username != "" {
		//cList := service.GetConnectionList()
		return c.Render(http.StatusOK, "ConnectionList.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
		})
	}

	fmt.Println("LoginInfo : ", loginInfo)

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
