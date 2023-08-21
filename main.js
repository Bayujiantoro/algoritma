



function urutAngka(arr) {
    
    for (let i = 0; i < arr.length - 1; i++) { 
        if (arr[i] < arr[i + 1]) {
            let element = arr[i];
            arr[i] = arr[i + 1];
            arr[i + 1] = element;
        }
        
    }
    return arr;
}

let arr = [1, 3, 2, 9, 5 ,6 ,8 ];

console.log(urutAngka(arr));
console.log(arr.length)