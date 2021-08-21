
const arr = [396285104, 573261094, 759641832, 819230764, 364801279];

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

function miniMaxSum(array) {
    const swap = (arr, left, right) => {
        const swapValue = arr[left];
        arr[left] = arr[right];
        arr[right] = swapValue;
    }

    const partition = (arr, left, right) => {
        const pivot = arr[Math.floor((left + right) / 2)];
        let pointLeft = left;
        let pointRight = right;

        while (pointLeft <= pointRight) {
            while(arr[pointLeft] < pivot) {
                pointLeft++;
            }

            while(arr[pointRight] > pivot) {
                pointRight--;
            }

            if (pointLeft <= pointRight) {
                swap(arr, pointLeft, pointRight);
    
                pointLeft++;
                pointRight--;
            }
        }

        return pointLeft;
    }

    const quickSort = (arr, left, right) => {
        if (arr <= 1) {
            return arr;
        }

        const index = partition(arr, left, right);

        console.log("left", arr);

        if (index -1 > left) {
            quickSort(arr, left, index - 1);
        }

        
        if (index < right) {
            quickSort(arr, index, right);
        }

        console.log("right", arr);

        return arr;
    }
    
    const ordered = quickSort(array, 0, array.length - 1);

    console.log(ordered);
    
    const [max, min] = ordered.reduce((acc, a, index) => {
        if (index === 0) {
            return [a, 0];
        }
        
        if (index === ordered.length - 1) {
            return [acc[0], acc[1] + a];
        }

        return [acc[0] + a, acc[1] + a];
    }, [0, 0]);
    
    console.log(`${max} ${min}`);
}

// expect 2093989309 2548418794
miniMaxSum(arr);