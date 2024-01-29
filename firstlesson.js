function nextInLine (arr, item){
    arr.push(item);
    return arr.shift();
}
var testArr=[1,2,3,4,5];
console.log("Before: " + JSON.stringify(testArr));
console.log(nextInLine(testArr, 6))
console.log("After: " + JSON.stringify(testArr));


function firstBoolean (x){
    if (x===8){
        return true
    }else{
        return false
    }
}
console.log(firstBoolean(5))

function repetitiveIfs (val){
    val < 5 - return "tiny";
    val < 10 - return "small";
    val < 15 - return "mid";
    val < 20 - return "large";
    val >= 5 - return "XXL";
}
console.log(repetitiveIfs(18));