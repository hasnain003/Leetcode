/*
An underground railway system is keeping track of customer travel times between different stations.
They are using this data to calculate the average time it takes to travel from one station to another.
https://leetcode.com/problems/design-underground-system/description/?envType=study-plan-v2&envId=programming-skills
*/
class TravelInfo {
    public double totalTime,totalTravel;
    public TravelInfo(double totalTime, double totalTravel) {
        this.totalTime = totalTime;
        this.totalTravel = totalTravel;
    }
}

class Travel {
    public double inTime;
    public String station;
    public Travel(double inTime,String station) {
        this.inTime = inTime;
        this.station = station;
    }
}

class UndergroundSystem {
    private HashMap<Integer,Travel> checkInInfo;
    private HashMap<String,TravelInfo> totalInfo;

    public UndergroundSystem() {
        this.checkInInfo = new HashMap<>();
        this.totalInfo = new HashMap<>();
    }
    
    public void checkIn(int id, String stationName, int t) {
        this.checkInInfo.put(id, new Travel(t,stationName));
    }
    
    public void checkOut(int id, String stationName, int t) {
        if (!this.checkInInfo.containsKey(id)) {
            return;
        }
        String key = this.checkInInfo.get(id).station + ":" + stationName;
        if (this.totalInfo.containsKey(key)) {
            TravelInfo val = this.totalInfo.get(key);
            val.totalTime += (t - this.checkInInfo.get(id).inTime);
            val.totalTravel++;
            this.totalInfo.put(key,val);
        } else {
            this.totalInfo.put(key,new TravelInfo((t - this.checkInInfo.get(id).inTime) ,1.0));
        }
        this.checkInInfo.remove(id);
    }
    
    public double getAverageTime(String startStation, String endStation) {
        String key = startStation + ":" + endStation;
        if (!this.totalInfo.containsKey(key) || this.totalInfo.get(key).totalTravel == 0 ) {
            return 0.0;
        }
        TravelInfo val = this.totalInfo.get(key);
        return val.totalTime / val.totalTravel;
    }
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * UndergroundSystem obj = new UndergroundSystem();
 * obj.checkIn(id,stationName,t);
 * obj.checkOut(id,stationName,t);
 * double param_3 = obj.getAverageTime(startStation,endStation);
 */