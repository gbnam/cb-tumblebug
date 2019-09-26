package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type specReq struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Csp         string `json:"csp"`
	Os_type     string `json:"os_type"`
	Num_vCPU    string `json:"num_vCPU"`
	Num_core    string `json:"num_core"`
	Mem_GiB     string `json:"mem_GiB"`
	Storage_GiB string `json:"storage_GiB"`
	Description string `json:"description"`
}

type specInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Csp         string `json:"csp"`
	Os_type     string `json:"os_type"`
	Num_vCPU    string `json:"num_vCPU"`
	Num_core    string `json:"num_core"`
	Mem_GiB     string `json:"mem_GiB"`
	Storage_GiB string `json:"storage_GiB"`
	Description string `json:"description"`

	Cost_per_hour         string `json:"cost_per_hour"`
	Num_storage           string `json:"num_storage"`
	Max_num_storage       string `json:"max_num_storage"`
	Max_total_storage_TiB string `json:"max_total_storage_TiB"`
	Net_bw_Gbps           string `json:"net_bw_Gbps"`
	Ebs_bw_Mbps           string `json:"ebs_bw_Mbps"`
	Gpu_model             string `json:"gpu_model"`
	Num_gpu               string `json:"num_gpu"`
	Gpumem_GiB            string `json:"gpumem_GiB"`
	Gpu_p2p               string `json:"gpu_p2p"`
}

/* FYI
g.POST("/:nsId/resources/spec", restPostSpec)
g.GET("/:nsId/resources/spec/:specId", restGetSpec)
g.GET("/:nsId/resources/spec", restGetAllSpec)
g.PUT("/:nsId/resources/spec/:specId", restPutSpec)
g.DELETE("/:nsId/resources/spec/:specId", restDelSpec)
g.DELETE("/:nsId/resources/spec", restDelAllSpec)
*/

// MCIS API Proxy: Spec
func restPostSpec(c echo.Context) error {

	nsId := c.Param("nsId")

	u := &specReq{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if true { // if error not occurred
		fmt.Println("[Registering Spec]")
		content := registerSpec(nsId, u)
		return c.JSON(http.StatusCreated, content)

	} else { // if error occurred
		mapA := map[string]string{"message": "lookupSpec(specRequest) failed."}
		return c.JSON(http.StatusFailedDependency, &mapA)
	}

}

func restGetSpec(c echo.Context) error {

	nsId := c.Param("nsId")

	id := c.Param("specId")

	content := specInfo{}
	/*
		var content struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Csp         string `json:"csp"`
			Os_type     string `json:"os_type"`
			Num_vCPU    string `json:"num_vCPU"`
			Num_core    string `json:"num_core"`
			Mem_GiB     string `json:"mem_GiB"`
			Storage_GiB string `json:"storage_GiB"`
			Description string `json:"description"`

			Cost_per_hour         string `json:"cost_per_hour"`
			Num_storage           string `json:"num_storage"`
			Max_num_storage       string `json:"max_num_storage"`
			Max_total_storage_TiB string `json:"max_total_storage_TiB"`
			Net_bw_Gbps           string `json:"net_bw_Gbps"`
			Ebs_bw_Mbps           string `json:"ebs_bw_Mbps"`
			Gpu_model             string `json:"gpu_model"`
			Num_gpu               string `json:"num_gpu"`
			Gpumem_GiB            string `json:"gpumem_GiB"`
			Gpu_p2p               string `json:"gpu_p2p"`
		}
	*/

	fmt.Println("[Get spec for id]" + id)
	key := "/ns/" + nsId + "/resources/spec/" + id
	fmt.Println(key)

	keyValue, _ := store.Get(key)
	fmt.Println("<" + keyValue.Key + "> \n" + keyValue.Value)
	fmt.Println("===============================================")

	json.Unmarshal([]byte(keyValue.Value), &content)
	content.Id = id // Optional. Can be omitted.

	return c.JSON(http.StatusOK, &content)

}

func restGetAllSpec(c echo.Context) error {

	nsId := c.Param("nsId")

	var content struct {
		//Name string     `json:"name"`
		Spec []specInfo `json:"spec"`
	}

	specList := getSpecList(nsId)

	for _, v := range specList {

		key := "/ns/" + nsId + "/resources/spec/" + v
		fmt.Println(key)
		keyValue, _ := store.Get(key)
		fmt.Println("<" + keyValue.Key + "> \n" + keyValue.Value)
		specTmp := specInfo{}
		json.Unmarshal([]byte(keyValue.Value), &specTmp)
		specTmp.Id = v
		content.Spec = append(content.Spec, specTmp)

	}
	fmt.Printf("content %+v\n", content)

	return c.JSON(http.StatusOK, &content)

}

func restPutSpec(c echo.Context) error {
	//nsId := c.Param("nsId")

	return nil
}

func restDelSpec(c echo.Context) error {

	nsId := c.Param("nsId")
	id := c.Param("specId")

	err := delSpec(nsId, id)
	if err != nil {
		cblog.Error(err)
		mapA := map[string]string{"message": "Failed to delete the spec"}
		return c.JSON(http.StatusFailedDependency, &mapA)
	}

	mapA := map[string]string{"message": "The spec has been deleted"}
	return c.JSON(http.StatusOK, &mapA)
}

func restDelAllSpec(c echo.Context) error {

	nsId := c.Param("nsId")

	specList := getSpecList(nsId)

	for _, v := range specList {
		err := delSpec(nsId, v)
		if err != nil {
			cblog.Error(err)
			mapA := map[string]string{"message": "Failed to delete All specs"}
			return c.JSON(http.StatusFailedDependency, &mapA)
		}
	}

	mapA := map[string]string{"message": "All specs has been deleted"}
	return c.JSON(http.StatusOK, &mapA)

}

func registerSpec(nsId string, u *specReq) specInfo {

	// TODO: Implement error check logic
	// TODO: Implement spec retrieving logic

	content := specInfo{}

	/* TODO: Implement the code below
	content, err := lookupSpec(u)

	if 1 { // if lookupSpec(u) succeeds
		content.Id = genUuid()
		...
	} else { // if lookupSpec(u) fails

	}
	*/

	// Temporary code
	content.Id = genUuid()
	content.Name = u.Name
	content.Csp = u.Csp
	content.Os_type = u.Os_type
	content.Num_vCPU = u.Num_vCPU
	content.Num_core = u.Num_core
	content.Mem_GiB = u.Mem_GiB
	content.Storage_GiB = u.Storage_GiB
	content.Description = u.Description

	/* FYI
	type specInfo struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Csp         string `json:"csp"`
		Os_type     string `json:"os_type"`
		Num_vCPU    string `json:"num_vCPU"`
		Num_core    string `json:"num_core"`
		Mem_GiB     string `json:"mem_GiB"`
		Storage_GiB string `json:"storage_GiB"`
		Description string `json:"description"`

		Cost_per_hour         string `json:"cost_per_hour"`
		Num_storage           string `json:"num_storage"`
		Max_num_storage       string `json:"max_num_storage"`
		Max_total_storage_TiB string `json:"max_total_storage_TiB"`
		Net_bw_Gbps           string `json:"net_bw_Gbps"`
		Ebs_bw_Mbps           string `json:"ebs_bw_Mbps"`
		Gpu_model             string `json:"gpu_model"`
		Num_gpu               string `json:"num_gpu"`
		Gpumem_GiB            string `json:"gpumem_GiB"`
		Gpu_p2p               string `json:"gpu_p2p"`
	}
	*/

	// cb-store
	fmt.Println("=========================== PUT registerSpec")
	Key := "/ns/" + nsId + "/resources/spec/" + content.Id
	mapA := map[string]string{
		"name":        content.Name,
		"csp":         content.Csp,
		"os_type":     content.Os_type,
		"Num_vCPU":    content.Num_vCPU,
		"Num_core":    content.Num_core,
		"mem_GiB":     content.Mem_GiB,
		"storage_GiB": content.Storage_GiB,
		"description": content.Description,

		"cost_per_hour":         content.Cost_per_hour,
		"num_storage":           content.Num_storage,
		"max_num_storage":       content.Max_num_storage,
		"max_total_storage_TiB": content.Max_total_storage_TiB,
		"net_bw_Gbps":           content.Net_bw_Gbps,
		"ebs_bw_Mbps":           content.Ebs_bw_Mbps,
		"gpu_model":             content.Gpu_model,
		"num_gpu":               content.Num_gpu,
		"gpumem_GiB":            content.Gpumem_GiB,
		"gpu_p2p":               content.Gpu_p2p,
	}
	Val, _ := json.Marshal(mapA)
	err := store.Put(string(Key), string(Val))
	if err != nil {
		cblog.Error(err)
	}
	keyValue, _ := store.Get(string(Key))
	fmt.Println("<" + keyValue.Key + "> \n" + keyValue.Value)
	fmt.Println("===========================")

	return content
}

func getSpecList(nsId string) []string {

	fmt.Println("[Get specs")
	key := "/ns/" + nsId + "/resources/spec"
	fmt.Println(key)

	keyValue, _ := store.GetList(key, true)
	var specList []string
	for _, v := range keyValue {
		//if !strings.Contains(v.Key, "vm") {
		specList = append(specList, strings.TrimPrefix(v.Key, "/ns/"+nsId+"/resources/spec/"))
		//}
	}
	for _, v := range specList {
		fmt.Println("<" + v + "> \n")
	}
	fmt.Println("===============================================")
	return specList

}

func delSpec(nsId string, Id string) error {

	fmt.Println("[Delete spec] " + Id)

	key := "/ns/" + nsId + "/resources/spec/" + Id
	fmt.Println(key)

	err := store.Delete(key)
	if err != nil {
		cblog.Error(err)
		return err
	}

	return nil
}