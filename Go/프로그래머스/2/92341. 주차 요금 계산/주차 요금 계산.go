import (
    "strings"
    "strconv"
    "math"
    "sort"
)

type Record struct {
    In      int
    Acc     int
    Fee     int
}

func TimeConverter(time string) int {
    times := strings.Split(time, ":")
    hour, _ := strconv.Atoi(times[0])
    minute, _ := strconv.Atoi(times[1])
    intTime :=  hour * 60 + minute
    return intTime
}

func newRecord(time string) *Record {
    in := TimeConverter(time)
    return &Record{In: in, Acc: 0, Fee: 0} 
}

func (r *Record) Out(time string){
    out := TimeConverter(time)
    r.Acc += out - r.In
    r.In = 9999 
}

func (r *Record) SettleUp(fees []int) {
    fee := fees[1]
    if r.In != 9999 {        // 출차처리가 안되었을 경우
        r.Acc += 1439 - r.In // (23 * 60) + 59
    }
    toCalculate := r.Acc - fees[0]
    if toCalculate > 0 {
        unit := int(math.Ceil(float64(toCalculate) / float64(fees[2])))
        fee += unit * fees[3]
    }
    r.Fee += fee
}

func solution(fees []int, records []string) []int {
    recordMap := map[string]*Record{}
    for _, r := range records {
        chars := strings.Split(r, " ")
        time, number, status := chars[0], chars[1], chars[2]
        switch status {
            case "IN":
            if record, exist := recordMap[number]; exist{
                record.In = TimeConverter(time)
            } else {
                recordMap[number] = newRecord(time)
            }
            case "OUT":
            recordMap[number].Out(time)
        }
    }
    
    keys := make([]string, 0, len(recordMap))
    
    for k, _ := range recordMap {
        keys = append(keys, k)
    }
    
    sort.Strings(keys)
    
    result := []int{}
    
    for _, k := range keys {
        v := recordMap[k]
        v.SettleUp(fees)
        result = append(result, v.Fee)
    }
    
    return result
}