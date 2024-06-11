// 数组去重
export function distinct1(arr, key) {
  var newObj = {}
  var newArr = []
  for (var i = 0; i < arr.length; i++) {
    var item = arr[i]
    if (!newObj[item[key]]) {
      newObj[item[key]] = newArr.push(item)
    }
  }
  return newArr
}
