function solution(participant, completion) {
    const completionMap = new Map();

    // 완주자 이름을 키로 하고, 등장 횟수를 값으로 하는 맵을 생성
    for (const name of completion) {
        const count = completionMap.get(name) || 0;
        completionMap.set(name, count + 1);
    }

    // 참가자 리스트를 순회하면서 완주하지 못한 사람을 찾음
    for (const name of participant) {
        const count = completionMap.get(name) || 0;
        if (count === 0) {
            return name; // 완주하지 못한 사람을 찾으면 즉시 반환
        }
        completionMap.set(name, count - 1); // 완주자 맵에서 참가자의 횟수를 감소
    }
}