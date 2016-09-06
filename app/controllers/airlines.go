package controllers
import (
	"github.com/revel/revel"
)

type AirlineController struct {
	*revel.Controller
}

func (c AirlineController) Get() revel.Result  {
//	var userRequest = c.Request.Body
//	data , _:= ioutil.ReadAll(userRequest);
//	var userInfo interface{}  ;
//	if err :=  json.Unmarshal(data,&userInfo); err != nil {
//		log.Println(err);
//		log.Println("there is a problem on data binding")
//	}
 var userResponse =` {

  "search_id": "T5dJTxtxQv-uQ6OuYtaWkw",
  "trips": [
    {
      "id": "abc123",
      "departure_city": "Mogadishu",
      "arrival_city": "Dubai",
      "departure_date": "07/31/2019",
      "return_date": "08/29/2016" ,
      "price": 800,
      "airline_code": "0B"
    },  {
      "id": "abc123",
      "departure_city": "Mogadishu",
      "arrival_city": "Dubai",
      "departure_date": "07/31/2019",
      "return_date": "08/29/2016" ,
      "price": 200,
      "airline_code": "0C"
    },
    {
      "id": "def456",
      "departure_city": "Mogadishu",
      "arrival_city": "Dubai",
      "departure_date": "07/31/2019",
      "return_date": "08/29/2016",
      "price": 900,
      "airline_code": "0A"
    }
  ],
  "number_adults": 1,
  "number_children": 0,
  "cabin": "economy",
  "sort": "price",
  "order": "asc",
  "currency_code": "USD",
  "price_min_usd": 515,
  "price_max_usd": 1700,
  "stop_types": ["none"],
  "distance": 1500,
  "meal": ["Breakfast", "Lunch"]
}`

//	log.Println(userInfo);
	return c.RenderText(userResponse)
}

func (c AirlineController) List() revel.Result  {
 var AirlineList =`[
		  {
			"code":"0A",
			"name":"Jubba Airways",
			"image":["ABCADAFASFEADFASDF"]
		  },
		  {
			"code":"0B",
			"name":"Blue Air",
			"image":["ABCADAFASFEADFASDF"]
		  },
		  {
			"code":"0C",
			"name":"Turkis Air",
			"image":["ABCADAFASFEADFASDF"]
		  }
		]`


	return c.RenderText(AirlineList)
}