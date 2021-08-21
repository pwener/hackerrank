/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

function timeConversion(s) {
    const period = s.slice(s.length - 2, s.length);
    const time = s.slice(0, s.length - 2);
    const hour = parseInt(s.slice(0, 2));
    
    if (period === "AM") {
        return hour === 12 ? time.replace('12', '00') : time;
    }
    
    const conversion =  hour !== 12 ? hour + 12 : 12;
    
    return `${conversion}`.concat(s.slice(2, s.length - 2))
}


console.log(timeConversion('12:45:54PM'));