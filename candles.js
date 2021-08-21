function birthdayCakeCandles(candles) {
    const max = Math.max(...candles);
    console.log(max);
    return candles.filter(c => c === max).length;
}

const input = [3, 2, 1, 3];

console.log(birthdayCakeCandles(input));