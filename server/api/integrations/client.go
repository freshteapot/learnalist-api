package integrations

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/freshteapot/learnalist-api/server/api/alist"
	"github.com/freshteapot/learnalist-api/server/api/client"
)

type Client struct {
	ApiClient client.Client
}

func (integrations *Client) RunIntegrationTests() {
	integrations.runV1Tests()
}

func (integrations *Client) runV1Tests() {
	fmt.Println("Running runV1Tests")
	aList := integrations.postListTypeV1NoData()
	integrations.postListTypeV1BadData()
	integrations.updateAlistV1WithData(*aList)
	integrations.getAlistV1(aList.Uuid)
	integrations.deleteAlist(aList.Uuid)
	shouldNotFindIt := integrations.getAlistV1(aList.Uuid)
	if shouldNotFindIt.Uuid != "" {
		log.Fatalln("The delete api might be broken.")
	}

	fmt.Println("Finished runV1Tests")
}

func (integrations *Client) postListTypeV1NoData() *alist.Alist {
	fmt.Println("integrations.postListTypeV1NoData start")
	body := strings.NewReader(`
{
		"data": [],
		"info": {
				"title": "Days of the Week",
				"type": "v1"
		}
}`)

	_, aList, err := integrations.ApiClient.PostAlist(body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("uuid is %s", aList.Uuid))
	fmt.Println("integrations.postListTypeV1NoData success")
	return aList
}

func (integrations *Client) postListTypeV1BadData() {
	fmt.Println("integrations.postListTypeV1BadData start")
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`
{
		"data": [""],
		"info": {
				"title": "Days of the Week",
				"type": "v1"
		}
}`)

	statusCode, _, _ := integrations.ApiClient.PostAlist(body)

	if statusCode != http.StatusBadRequest {
		log.Fatal("We were expecting a 400")
	}
	fmt.Println("integrations.postListTypeV1BadData success")
}

func (integrations *Client) updateAlistV1WithData(aList alist.Alist) *alist.Alist {
	fmt.Println("integrations.updateAlistV1WithData start")
	items := aList.Data.(alist.TypeV1)
	items = append(items, "apple")
	aList.Data = items

	byteData, _ := aList.MarshalJSON()
	body := strings.NewReader(string(byteData))

	statusCode, aListB, _ := integrations.ApiClient.PutAlist(aList.Uuid, body)
	fmt.Println(fmt.Sprintf("Status code is %d", statusCode))

	fmt.Println("integrations.updateAlistV1WithData success")
	return aListB
}

func (integrations *Client) getAlistV1(uuid string) alist.Alist {
	fmt.Println("integrations.getAlistV1 start")
	_, aList, _ := integrations.ApiClient.GetAlist(uuid)
	fmt.Println("integrations.getAlistV1 success")
	return aList
}

func (integrations *Client) deleteAlist(uuid string) {
	fmt.Println("integrations.deleteAlist start")
	integrations.ApiClient.DeleteAlist(uuid)
	fmt.Println("integrations.deleteAlist success")
}
