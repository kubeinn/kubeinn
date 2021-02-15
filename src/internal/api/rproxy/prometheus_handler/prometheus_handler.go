package prometheus_handler

import (
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"strings"

	"encoding/json"

	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/kubeinn/src/internal/global"
)

type ReverseProxyResponseObject struct {
	Pod                           string `json:"pod"`
	Namespace                     string `json:"namespace"`
	CreatedByName                 string `json:"created_by_name"`
	Node                          string `json:"node"`
	KubePodCreated                string `json:"kube_pod_created"`
	KubePodCompleted              string `json:"kube_pod_completed"`
	ContainerCPUUsageSecondsTotal string `json:"container_cpu_usage_seconds_total"`
	ContainerMemoryUsageBytes     string `json:"container_memory_usage_bytes"`
	KubePodStatusPhase            string `json:"kube_pod_status_phase"`
}

func newReverseProxyResponseObject() *ReverseProxyResponseObject {
	o := ReverseProxyResponseObject{}
	o.Pod = ""
	o.Namespace = ""
	o.CreatedByName = ""
	o.Node = ""
	o.KubePodCreated = ""
	o.KubePodCompleted = ""
	o.ContainerCPUUsageSecondsTotal = ""
	o.ContainerMemoryUsageBytes = ""
	o.KubePodStatusPhase = ""
	return &o
}

type PrometheusResponse struct {
	Status string              `json:"status"`
	Data   PrometheusDataField `json:"data"`
}

type PrometheusDataField struct {
	ResultType string                      `json:"resultType"`
	Result     []PrometheusDataResultField `json:"result"`
}

type PrometheusDataResultField struct {
	Metric map[string]interface{} `json:"metric"`
	Value  []interface{}          `json:"value"`
}

func unmarshalPrometheusResponse(b []byte) (PrometheusResponse, error) {
	prometheusResponse := PrometheusResponse{}
	err := json.Unmarshal(b, &prometheusResponse)
	if err != nil {
		return prometheusResponse, err
	}
	return prometheusResponse, nil
}

func PrometheusHandler(c *gin.Context) {
	// Parse context request
	var audience string
	var subject string

	// Identify source
	if strings.HasPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.INNKEEPER_ROUTE_PREFIX) {
		audience = global.JWT_AUDIENCE_INNKEEPER
	} else if strings.HasPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.PILGRIM_ROUTE_PREFIX) {
		audience = global.JWT_AUDIENCE_PILGRIM
	}
	subject = c.Request.Header.Get("subject")

	// Get all projects from database
	projectsMap, err := global.PG_CONTROLLER.SelectProjects()
	for k, v := range projectsMap {
		log.Printf("key[%s] value[%s]\n", k, v)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Get slice of projects to query
	var projectsSlice []string
	for k, v := range projectsMap {
		log.Printf("key[%s] value[%s]\n", k, v)
		// if requester is admin or if user queries for his/her own usage
		if audience == global.JWT_AUDIENCE_INNKEEPER || subject == v {
			// add namespace of project to slice
			projectsSlice = append(projectsSlice, k)
		}
	}

	// Instantiate reverse proxy response structure
	reverseProxyResponseMap := make(map[string]ReverseProxyResponseObject)

	// Query PromQL for kube_pod_info
	err = getKubePodInfo(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Query PromQL for kube_pod_created
	err = getKubePodCreated(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Query PromQL for kube_pod_completed
	err = getKubePodCompleted(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Query PromQL for container_cpu_usage_seconds_total
	err = getContainerCPUUsageSecondsTotal(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Query PromQL for container_memory_usage_bytes
	err = getContainerMemoryUsageBytes(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Query PromQL for kube_pod_status_phase
	err = getKubePodStatusPhase(reverseProxyResponseMap, projectsSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Convert reverseProxyResponseMap to slice
	var reverseProxyResponseSlice []ReverseProxyResponseObject
	for _, v := range reverseProxyResponseMap {
		reverseProxyResponseSlice = append(reverseProxyResponseSlice, v)
	}
	b, err := json.Marshal(reverseProxyResponseSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	c.Data(http.StatusOK, "application/json", b)

	// //  Write response headers
	// for header, values := range proxyRes.Header {
	// 	for _, value := range values {
	// 		log.Println(header + ": " + value)
	// 		c.Writer.Header().Set(header, value)
	// 	}
	// }
	// c.Data(proxyRes.StatusCode, "application/json", body)
}

func queryPromQL(function string, metric string, labels map[string]string, time string) ([]byte, error) {
	// Format request for Prometheus API
	var labelsSlice []string
	for k, v := range labels {
		labelsSlice = append(labelsSlice, k+"="+v)
	}

	// Create query string
	query := ""
	if time == "" {
		query = metric + "{" + strings.Join(labelsSlice[:], ",") + "}"
	} else {
		query = metric + "{" + strings.Join(labelsSlice[:], ",") + "}[" + time + "]"
	}
	if function != "" {
		query = function + "(" + query + ")"
	}

	// Format remaining request params
	url := "http://" + global.PROMETHEUS_URL + "/api/v1/query?query=" + query
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil

}

func getKubePodInfo(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	body, err := queryPromQL("", "kube_pod_info", labels, "")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// instantiate reverseProxyResponseObject
		reverseProxyResponseObject := newReverseProxyResponseObject()

		// fill in pod value
		if str, ok := result.Metric["pod"].(string); ok {
			reverseProxyResponseObject.Pod = str
		} else {
			fmt.Println(err)
			return err
		}

		// fill in namespace value
		if str, ok := result.Metric["namespace"].(string); ok {
			reverseProxyResponseObject.Namespace = str
		} else {
			fmt.Println(err)
			return err
		}

		// fill in created_by_name value
		if str, ok := result.Metric["created_by_name"].(string); ok {
			reverseProxyResponseObject.CreatedByName = str
		} else {
			fmt.Println(err)
			return err
		}

		// fill in node value
		if str, ok := result.Metric["node"].(string); ok {
			reverseProxyResponseObject.Node = str
		} else {
			fmt.Println(err)
			return err
		}

		// save struct as object
		reverseProxyResponseMap[reverseProxyResponseObject.Pod] = *reverseProxyResponseObject
	}
	return nil
}

func getKubePodCreated(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	body, err := queryPromQL("", "kube_pod_created", labels, "")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// get pod value
		pod, ok := result.Metric["pod"].(string)
		if !ok {
			fmt.Println(err)
			return err
		}

		// Check if key exists
		if _, ok := reverseProxyResponseMap[pod]; ok {
			// get reverseProxyResponseObject
			reverseProxyResponseObject := reverseProxyResponseMap[pod]

			// fill in kube_pod_created value
			if str, ok := result.Value[1].(string); ok {
				reverseProxyResponseObject.KubePodCreated = str
			} else {
				fmt.Println(err)
				return err
			}

			// save struct as object
			reverseProxyResponseMap[pod] = reverseProxyResponseObject
		}
	}
	return nil
}

func getKubePodCompleted(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	body, err := queryPromQL("", "kube_pod_completion_time", labels, "")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// get pod value
		pod, ok := result.Metric["pod"].(string)
		if !ok {
			fmt.Println(err)
			return err
		}

		// Check if key exists
		if _, ok := reverseProxyResponseMap[pod]; ok {
			// get reverseProxyResponseObject
			reverseProxyResponseObject := reverseProxyResponseMap[pod]

			// fill in kube_pod_created value
			if str, ok := result.Value[1].(string); ok {
				reverseProxyResponseObject.KubePodCompleted = str
			} else {
				fmt.Println(err)
				return err
			}

			// save struct as object
			reverseProxyResponseMap[pod] = reverseProxyResponseObject
		}
	}
	return nil
}

func getContainerCPUUsageSecondsTotal(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	labels["container"] = "\"\""
	body, err := queryPromQL("max_over_time", "container_cpu_usage_seconds_total", labels, "4w")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// get pod value
		pod, ok := result.Metric["pod"].(string)
		if !ok {
			fmt.Println(err)
			return err
		}

		// Check if key exists
		if _, ok := reverseProxyResponseMap[pod]; ok {
			// get reverseProxyResponseObject
			reverseProxyResponseObject := reverseProxyResponseMap[pod]

			// fill in kube_pod_created value
			if str, ok := result.Value[1].(string); ok {
				reverseProxyResponseObject.ContainerCPUUsageSecondsTotal = str
			} else {
				fmt.Println(err)
				return err
			}

			// save struct as object
			reverseProxyResponseMap[pod] = reverseProxyResponseObject
		}
	}
	return nil
}

func getContainerMemoryUsageBytes(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	labels["container"] = "\"\""
	body, err := queryPromQL("max_over_time", "container_memory_usage_bytes", labels, "4w")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// get pod value
		pod, ok := result.Metric["pod"].(string)
		if !ok {
			fmt.Println(err)
			return err
		}

		// Check if key exists
		if _, ok := reverseProxyResponseMap[pod]; ok {
			// get reverseProxyResponseObject
			reverseProxyResponseObject := reverseProxyResponseMap[pod]

			// fill in kube_pod_created value
			if str, ok := result.Value[1].(string); ok {
				reverseProxyResponseObject.ContainerMemoryUsageBytes = str
			} else {
				fmt.Println(err)
				return err
			}

			// save struct as object
			reverseProxyResponseMap[pod] = reverseProxyResponseObject
		}
	}
	return nil
}

func getKubePodStatusPhase(reverseProxyResponseMap map[string]ReverseProxyResponseObject, projects []string) error {
	// Get list of containers for all projects in projectsSlice
	labels := make(map[string]string)
	labels["namespace"] = "~\"" + strings.Join(projects[:], "|") + "\""
	body, err := queryPromQL("", "kube_pod_status_phase", labels, "")

	// Unmarshal response body
	prometheusResponse, err := unmarshalPrometheusResponse(body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fill response map with unmarshaled data
	for _, result := range prometheusResponse.Data.Result {
		// get pod value
		pod, ok := result.Metric["pod"].(string)
		if !ok {
			fmt.Println(err)
			return err
		}

		// fill in kube_pod_created value
		if str, ok := result.Value[1].(string); ok {
			if str == "1" {
				// Check if key exists
				if _, ok := reverseProxyResponseMap[pod]; ok {
					// get reverseProxyResponseObject
					reverseProxyResponseObject := reverseProxyResponseMap[pod]
					if str, ok := result.Metric["phase"].(string); ok {
						fmt.Println("phase: " + str)
						reverseProxyResponseObject.KubePodStatusPhase = str
						// save struct as object
						reverseProxyResponseMap[pod] = reverseProxyResponseObject
					}
				} else {
					fmt.Println(err)
					return err
				}
			}
		} else {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
