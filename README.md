# delighted-go
[![Circle CI](https://circleci.com/gh/jonnonz/delighted-go/tree/master.svg?style=svg)](https://circleci.com/gh/jonnonz/delighted-go/tree/master)

___
Simple Golang API Client for [Delighted](https://delighted.com). Currently supports the following. 

- People 
	- Adding a person to a survey.
	- Unsubscribing a person from a survey.
	- Deleting pending survey requests.
	
- Survey
	- Listing all survey responses.

- Metrics
	- Getting metrics.


### Installation 

```
go get github.com/jonnonz/delighted-go
```


```
import github.com/jonnonz/delighted-go
```


### Examples

#### Create a client

```
client, err := delighted.NewClient("DELIGHTED_API_KEY", nil)
if err != nil {
	fmt.Println("Doh!", err)
}
```
Note: You can specify a http.Client instead of nil, this is useful for mocking during testing.

#### Add a person to a survey
```
p := delighted.Person{Email: "test@test.com", Name: "John Doe"}
r, err := client.PersonService.Create(&p)
if err != nil {
	fmt.Println("Doh!", err)
}
fmt.Printf("#%v",r)
```
Note: Email is required when adding a person to a survey.

#### Unsubscribing a person from a survey

```
p := delighted.Person{Email: "test@test.com"}
_, err := client.PersonService.Unsubscribe(&p) // returns a bool.
if err != nil {
	fmt.Println("Doh!", err)
}

```

#### Listing all survey responses

```
sr := delighted.SurveyResponses{PerPage: 100, Since: 12345678910}
r, err := client.SurveyService.GetAll(&sr)
if err != nil {
	fmt.Println("Doh!", err)
}
fmt.Printf("#%v",r)
```


#### Get metrics

```
m := delighted.Metrics{Since: 12345678910, Until: 10987654321}
r, err := client.MetricService.Get(&m)
if err != nil {
	fmt.Println("Jar Jar Binks!", err)
}
fmt.Printf("#%v",r)
```

### Backwards Compatibility

As of 01/02/2016, This is currently under active development. Although I will try my best, there is definitely no garantee of BC breaks during the early stages of my work. Use fully at your own risk.


