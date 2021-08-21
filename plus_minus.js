function plusMinus(arr) {
    const [pos, neg] = arr.reduce((acc, a) => {
        const p = acc[0];
        const n = acc[1];
        return [a > 0 ? p + 1: p, a < 0 ? n + 1: n]
    }, [0, 0]);
    
    const len = arr.length;

    const applyPrecision = (number) => parseFloat(number).toPrecision(6);
    
    console.log(applyPrecision(pos/len));
    console.log(applyPrecision(neg/len));
    console.log(applyPrecision(Math.abs(pos - neg), len));
}

const input = [1, -1, 0, 1, 0];

plusMinus(input);