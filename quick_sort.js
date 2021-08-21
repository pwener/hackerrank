const arr = [5, 2, 1, 4, 3];

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

        swap(arr, pointLeft, pointRight);

        pointLeft++;
        pointRight--;
    }

    return pointLeft;
}

const quickSort = (arr, left, right) => {
    if (arr <= 1) {
        return arr;
    }

    const index = partition(arr, left, right);

    if (index -1 > left) {
        quickSort(arr, left, index - 1);
    }

    if (index < right) {
        quickSort(arr, index, right);
    }

    return arr;
}

console.log(quickSort(arr, 0, arr.length - 1).join().replace(",", " "));