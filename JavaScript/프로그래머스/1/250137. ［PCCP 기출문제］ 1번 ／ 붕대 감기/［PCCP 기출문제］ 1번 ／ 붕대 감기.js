const 붕대감기 = ([count, heal, bonus], i) => {
    if (count === i) {
        return [heal + bonus, 0]
    }
    return [heal, i] 
}

function solution(bandage, health, attacks) {
    const maxHealth = health
    const attackMap = new Map(attacks)
    const time = attacks[attacks.length - 1][0]
    
    let success = 0
    for (let i = 1; i <= time; i++) {
        if (attackMap.has(i)) {
            success = 0
            health -= attackMap.get(i)
        } else {
            success++
            const [회복량, 스택] = 붕대감기(bandage, success)
            health += 회복량
            success = 스택
        }
        if (health > maxHealth) {
            health = maxHealth
        }
        if (health <= 0) {
            health = -1
            break
        }
    }
    return health;
}