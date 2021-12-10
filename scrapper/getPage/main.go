package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println(len(jobs), "개 완료")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	// 새 csv 파일을 생성하는 명령을 w에 저장
	w := csv.NewWriter(file)
	// 함수가 끝나면 Flush로 데이터들을 파일에 입력하여 저장
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}

		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	// strconv 패키지의 Itoa()를 사용하면 int형을 string으로 바꿔준다
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println(pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)

	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs := append(jobs, job)
	}
	return jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := card.Find(".salary-snippet").Text()
	summary := card.Find(".job-snippet").Text()

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

// TrimSpace : 문자열 앞뒤의 공백 제거
// Fields : 문자열을 공백기준으로 잘라 배열로 반환
// Join : 배열을 문자열로 바꾸고 배열 구분 , 대신 " "로 바꿔서 반환
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// getPages 함수가 끝나면 닫아서 메모리가 새어나가는걸 막는다
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("상태로 요청 실패:", res.StatusCode)
	}
}
