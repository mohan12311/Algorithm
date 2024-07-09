import (
    "sort"
)

type Detail struct {
    first   Album
    second  Album
    total   int
}

type Album struct {
    play  int
    index int
}

const NULL = -1

func solution(genres []string, plays []int) []int {
    genresMap := map[string]Detail{}
    for i := 0; i < len(genres); i++ {
        genre := genres[i]
        play := plays[i]
        target, exist := genresMap[genre]
        if exist {
            if target.first.play < play {
                target.second = target.first
                target.first = Album{play, i}
            } else if target.second.play < play {
                target.second = Album{play, i}
            }
            target.total += play
        } else {
            target = Detail{ Album{play, i}, Album{NULL, NULL}, play}
        }
        genresMap[genre] = target
    }
    
    keys := []string{}
    for k := range genresMap {
        keys = append(keys, k)
    }
    
    sort.Slice(keys, func(a, b int) bool {
        return genresMap[keys[a]].total > genresMap[keys[b]].total
    })
    
    result := []int{}
    for _, k := range keys {
        fst := genresMap[k].first.index
        scd := genresMap[k].second.index
        result = append(result, fst)
        if scd != NULL {
            result = append(result, scd)
        }
    }
    
    return result
}