package server

import (
	"encoding/json"
	"time"

	"gitlab.com/btlike/api/utils"
	"gopkg.in/olivere/elastic.v3"
)

//format
const (
	TIME = "2006-01-02T15:04:05Z07:00"
)

var (
	monthTrends []trend
	weekTrends  []trend
)

func getMonthTrend(latest time.Time) {
	for {
		section := elastic.NewRangeQuery("CreateTime").Gte(latest.Add(-time.Hour * 24 * 30).Format(TIME))
		var trends []trend
		result, err := utils.ES.Search().Index("torrent").Query(section).Sort("Heat", false).Size(1000).Do()
		if err != nil {
			utils.Log.Println(err)
			time.Sleep(time.Hour)
			continue
		}
		if result != nil && result.Hits != nil {
			for _, v := range result.Hits.Hits {
				var esdata estorrent
				err = json.Unmarshal(*v.Source, &esdata)
				if err != nil {
					utils.Log.Println(err)
					continue
				}
				trends = append(trends, trend{
					ID:         esdata.Infohash,
					Name:       esdata.Name,
					CreateTime: esdata.CreateTime,
					Length:     esdata.Length,
					Heat:       esdata.Heat,
				})
				if len(trends) >= 100 {
					monthTrends = make([]trend, 0)
					for _, v := range trends {
						monthTrends = append(monthTrends, v)
					}
					goto done
				}
			}
			monthTrends = make([]trend, 0)
			for _, v := range trends {
				monthTrends = append(monthTrends, v)
			}
			goto done
		}
	done:
		time.Sleep(time.Hour)
	}
}

func getWeekTrend(latest time.Time) {
	for {
		section := elastic.NewRangeQuery("CreateTime").Gte(latest.Add(-time.Hour * 24 * 7).Format(TIME))
		var trends []trend
		result, err := utils.ES.Search().Index("torrent").Query(section).Sort("Heat", false).Size(1000).Do()
		if err != nil {
			utils.Log.Println(err)
			time.Sleep(time.Hour)
			continue
		}
		if result != nil && result.Hits != nil {
			for _, v := range result.Hits.Hits {
				var esdata estorrent
				err = json.Unmarshal(*v.Source, &esdata)
				if err != nil {
					utils.Log.Println(err)
					continue
				}
				trends = append(trends, trend{
					ID:         esdata.Infohash,
					Name:       esdata.Name,
					CreateTime: esdata.CreateTime,
					Length:     esdata.Length,
					Heat:       esdata.Heat,
				})
				if len(trends) >= 100 {
					weekTrends = make([]trend, 0)
					for _, v := range trends {
						weekTrends = append(weekTrends, v)
					}
					goto done
				}
			}
			weekTrends = make([]trend, 0)
			for _, v := range trends {
				weekTrends = append(weekTrends, v)
			}
			goto done
		}
	done:
		time.Sleep(time.Hour)
	}
}

func getTrend() (err error) {
	result, err := utils.ES.Search().Index("torrent").Sort("CreateTime", false).Size(1).Do()
	if err != nil {
		utils.Log.Printf("如果未抓取数据，此报错不予理会，先运行crawl即可。原始错误:%v\n", err)
		return
	}

	var latest time.Time
	if result != nil && result.Hits != nil {
		for _, v := range result.Hits.Hits {
			var esdata estorrent
			json.Unmarshal(*v.Source, &esdata)
			latest = esdata.CreateTime
			if latest.IsZero() {
				latest = time.Now()
			}
		}
	}
	go getMonthTrend(latest)
	go getWeekTrend(latest)
	return
}
