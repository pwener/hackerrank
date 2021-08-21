/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

function breakingRecords(scores) {
    let record = scores[0];
    let min = scores[0];
   return scores.reduce((acc, next) => {
       console.log('evaluate:', next);
       let [best, worst] = acc; 

       if (next > record) {
           best++;
           record = next;
       }
       
       if (next < min) {
           worst++;
           min = next;
       }
       
       return [best, worst]
   }, [0, 0])
}

const arr = [3, 4, 21, 36, 10, 28, 35, 5, 24, 42];
const arr2 = [10, 5, 20, 20, 4, 5, 2, 25, 1];

console.log(breakingRecords(arr));