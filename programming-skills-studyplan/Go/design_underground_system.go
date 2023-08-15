package main

/*
An underground railway system is keeping track of customer travel times between different stations.
They are using this data to calculate the average time it takes to travel from one station to another.
https://leetcode.com/problems/design-underground-system/description/?envType=study-plan-v2&envId=programming-skills
*/

type UndergroundSystem struct {
	Journey map[int]CheckInData
	Time    map[string]StationAToB
}

type StationAToB struct {
	TotalTime    float64
	TotalJourney float64
}

type CheckInData struct {
	CheckIn string
	InTime  int
}

func UndergroundSystemConstructor() UndergroundSystem {
	Journey := make(map[int]CheckInData)
	TotalTime := make(map[string]StationAToB)
	return UndergroundSystem{
		Journey: Journey,
		Time:    TotalTime,
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.Journey[id] = CheckInData{CheckIn: stationName, InTime: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, exists := this.Journey[id]; !exists {
		return
	}
	stationAToB := this.Journey[id].CheckIn + ":" + stationName
	if info, exists := this.Time[stationAToB]; exists {
		info.TotalTime += float64(t - this.Journey[id].InTime)
		info.TotalJourney++
		this.Time[stationAToB] = info
	} else {
		this.Time[stationAToB] = StationAToB{
			TotalTime:    float64(t - this.Journey[id].InTime),
			TotalJourney: 1.0,
		}
	}
	delete(this.Journey, id)

}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	AtoB := startStation + ":" + endStation
	if _, exists := this.Time[AtoB]; !exists || startStation == "" || endStation == "" {
		return 0.0
	}
	return (this.Time[AtoB].TotalTime / this.Time[AtoB].TotalJourney)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
